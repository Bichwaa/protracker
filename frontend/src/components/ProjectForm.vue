<template>
    <!-- The Modal -->
    <div id="myModal" class="modal">
        <!-- Modal content -->
        <div class="modal-content w-2/3 rounded-lg">
            <div class="w-full flex flex-row-reverse">
                <span class="text-2xl cursor-pointer hover:text-red-500 duration-300" @click="bye">&times;</span>
            </div>
            <form method="post" id="project-form" @submit.prevent="submitForm" v-if="projectForm" class="flex flex-col">
                <div class="flex flex-col my-1">
                    <label for="name">Name</label>
                    <input 
                        type="text" 
                        v-model="projectData.Name" 
                        name="Name" placeholder="Name" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Name = ''"
                        >
                    <span class="text-sm text-red-500" v-if="errors.Name" >{{ errors.Name }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="description">description</label>
                    <textarea 
                        v-model="projectData.Description" 
                        name="Description" 
                        placeholder="description" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Description = ''"
                        />
                    <span class="text-sm text-red-500" v-if="errors.Description" >{{ errors.Description }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="name">Deadline</label>
                    <input type="datetime-local" v-model="projectData.EstimatedEndDate" name="EstimatedEndDate" class="border border-[#490a47] rounded-sm p-1">
                    <span class="text-sm text-red-500" v-if="errors.EstimatedEndDate" >{{ errors.EstimatedEndDate }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="tags">Overseer:</label>
                    <input type="text" v-model="projectData.Overseer" name="Overseer" placeholder="Firstname Lastname" class="border border-[#490a47] rounded-sm p-1">
                    <span class="text-sm text-red-500" v-if="errors.Overseer" >{{ errors.Overseer }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="tags">Tags:</label>
                    <input type="text" v-model="projectData.Tags" name="Tags" placeholder="comma separated tags" class="border border-[#490a47] rounded-sm p-1">
                    <span class="text-sm text-red-500" v-if="errors.Tags" >{{ errors.Tags }}</span>
                </div>
                <button type="submit" class="mt-3 bg-[#490a47] disabled:bg-gray-200 text-white font-medium p-3 cursor-pointer grid place-content-center rounded-md" >
                    <span v-if="!loading"> {{ edit ?  'update' : 'create' }} </span>
                    <Loader v-else  class="h-5 w-5"/>
                </button>
            </form>
        </div>
    </div> 
</template>

<script setup>
import { ref, defineEmits, onMounted} from 'vue';
import { z } from "zod";
import { useProjectController } from "../APIs/project-controller";
import Loader from './Loader.vue';

const props = defineProps({
    edit: {
        type: Boolean,
        default: false
    },
    project : {
        type: Object,
        default: {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            Overseer: null,
            Tags:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Objectives:[],
            Notes:[]
        }
    }
});
const loading = ref(false)

const emit = defineEmits(['modal-off', 'project-created', 'project-updated']);

const projectForm = ref(true);

const nullProjFields = {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            Overseer: null,
            Tags:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Objectives:[],
            Notes:[]
}

const projectData = ref(nullProjFields)
//validation Logic
const projectschema = z.object({
            ID:z.number().int().optional(),
            Name:z.string({message:"project must have name"}).nonempty(),
            Tags:z.string(),
            Overseer: z.string(),
            EstimatedEndDate:z.string({message:'Expected date, received nothing'}).datetime().nonempty(), 
            Description:z.string(),
})

const errors = ref({
            Name:"",
            Tags:"",
            EstimatedEndDate:"",
            Description:"",
            Overseer:""
})

//End validation Logic

const bye = () => {
    emit("modal-off");
};

const submitForm = async () => {
    loading.value = true;
    const { createProject, updateProject } = useProjectController();
    let form = document.getElementById('project-form');
        const formItems = [...form];
        //from array of form items filter out the submit value 
        //then grab name&value pairs for each of the remaining items.
        const keyVals = formItems.filter(i => i.type != "submit").map(item => [item.name, item.value]);
        //then turn them into an object
        const data = Object.fromEntries(keyVals);
        // console.log(data)
        data.EstimatedEndDate = new Date(data.EstimatedEndDate).toJSON();

        const validation = projectschema.safeParse(data)
        if (!validation.success) {
            // Handle validation errors
            validation.error.errors.forEach((err) => {
            errors.value[err.path[0]] = err.message;
            });
        } else {
            // Form is valid, proceed with submission
            loading.value = false;

        }
    if(!props.edit && validation.success){
        await createProject(data);
        emit("project-created");
    }else if(props.edit && validation.success){
        data.ID = projectData.value.ID
        await updateProject(data);
        // console.log(data)
        emit("project-updated");
    }
}

onMounted(()=>{
    projectData.value = nullProjFields
    if(props.edit){
        projectData.value = props.project
        projectData.value.EstimatedEndDate = new Date(props.project.EstimatedEndDate).toISOString().substring(0, 16)
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

