<template>
    <!-- The Modal -->
    <div id="myModal" class="modal">
        <!-- Modal content -->
        <div class="modal-content w-2/3 rounded-lg">
            <div class="w-full flex flex-row-reverse">
                <span class="text-2xl cursor-pointer hover:text-red-500 duration-300" @click="bye">&times;</span>
            </div>
            <form method="post" id="goal-form" @submit.prevent="submitForm" v-if="goalFormShow" class="flex flex-col">
                <div class="flex flex-col my-1">
                    <label for="name">Name</label>
                    <input 
                        type="text" 
                        v-model="goalData.Name" 
                        name="Name" placeholder="Name" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Name = ''"
                        >
                    <span class="text-sm text-red-500" v-if="errors.Name" >{{ errors.Name }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="description">Description</label>
                    <textarea 
                        v-model="goalData.Description" 
                        name="Description" 
                        placeholder="Description" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Description = ''"
                        />
                    <span class="text-sm text-red-500" v-if="errors.Description" >{{ errors.Description }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="name">Deadline</label>
                    <input type="datetime-local" v-model="goalData.EstimatedEndDate" name="EstimatedEndDate" class="border border-[#490a47] rounded-sm p-1">
                    <span class="text-sm text-red-500" v-if="errors.EstimatedEndDate" >{{ errors.EstimatedEndDate }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="tags">Overseer:</label>
                    <input type="text" v-model="goalData.Overseer" name="Overseer" placeholder="Firstname Lastname" class="border border-[#490a47] rounded-sm p-1">
                    <span class="text-sm text-red-500" v-if="errors.Overseer" >{{ errors.Overseer }}</span>
                </div>
                <button type="submit" class="mt-3 bg-[#490a47] disabled:bg-gray-200 text-white font-medium p-3 cursor-pointer grid place-content-center rounded-md" >
                    <span v-if="!loading"> {{ edit ?  'Update Goal' : 'Create Goal' }} </span>
                    <Loader v-else  class="h-5 w-5"/>
                </button>
            </form>
        </div>
    </div> 
</template>

<script setup>
import { ref, defineEmits, onMounted} from 'vue';
import { z } from "zod";
import { useGoalController } from "../APIs/goal-controller";
import Loader from './Loader.vue';

const props = defineProps({
    edit: {
        type: Boolean,
        default: false
    },
    projectId:{
        type: Number,
        default:0
    },
    goal : {
        type: Object,
        default: {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Overseer:null,
            Goals:[],
            Notes:[]
        }
    }
});
const loading = ref(false)

const emit = defineEmits(['modal-off', 'goal-created', 'goal-updated']);

const goalFormShow = ref(true);

const nullGoalFields = {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Overseer:null,
            Goals:[],
            Notes:[]
}

const goalData = ref(nullGoalFields)
//validation Logic
const goalSchema = z.object({
            ID:z.number().int().optional(),
            Name:z.string({message:"Goal must have a name"}).nonempty(),
            EstimatedEndDate:z.string({message:'Expected date, received nothing'}).datetime().nonempty(), 
            Description:z.string(),
            Overseer:z.string(),
})

const errors = ref({
            Name:"",
            EstimatedEndDate:"",
            Description:""
})

//End validation Logic

const bye = () => {
    emit("modal-off");
};

const submitForm = async () => {
    loading.value = true;
    const { createGoal, updateGoal } = useGoalController();
    let form = document.getElementById('goal-form');
        const formItems = [...form];
        const keyVals = formItems.filter(i => i.type != "submit").map(item => [item.name, item.value]);
        const data = Object.fromEntries(keyVals);
        data.EstimatedEndDate = new Date(data.EstimatedEndDate).toJSON();
        const validation = goalSchema.safeParse(data)
        if (!validation.success) {
            validation.error.errors.forEach((err) => {
            errors.value[err.path[0]] = err.message;
            });
        } else {
            loading.value = false;
        }
    if(!props.edit && validation.success){
        data.ProjectID = props.projectId;
        await createGoal(data);
        emit("goal-created");
    }else if(props.edit && validation.success){
        data.ID = goalData.value.ID
        await updateGoal(data);
        emit("goal-updated");
    }
}

onMounted(()=>{
    goalData.value = nullGoalFields
    if(props.edit){
        goalData.value = props.goal
        goalData.value.EstimatedEndDate = new Date(props.goal.EstimatedEndDate).toISOString().substring(0, 16)
    }
})

</script>

<style scoped>
    /* The Modal (background) */
    .modal {
        display: block; /* Hidden by default */
        position: fixed; /* Stay in place */
        z-index: 1; /* Sit on top */
        left: 0;
        top: 0;
        width: 100%; /* Full width */
        height: 100%; /* Full height */
        overflow: auto; /* Enable scroll if needed */
        background-color: rgb(0,0,0); /* Fallback color */
        background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
    }

    /* Modal Content/Box */
    .modal-content {
        background-color: #fefefe;
        margin: 15% auto;
        padding: 20px;
        border: 1px solid #888;
    }

    /* The Close Button */
    .close {
        color: #aaa;
        float: right;
        font-size: 28px;
        font-weight: bold;
    }

    .close:hover,
    .close:focus {
        color: black;
        text-decoration: none;
        cursor: pointer;
    } 
</style>