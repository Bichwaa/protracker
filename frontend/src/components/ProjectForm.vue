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
                    <input type="text" v-model="projectData.Name" name="name" placeholder="Name" class="border border-[#490a47] rounded-sm p-1">
                </div>
                <div class="flex flex-col my-1">
                    <label for="description">description</label>
                    <textarea v-model="projectData.Description" name="description" placeholder="description" class="border border-[#490a47] rounded-sm p-1"/>
                </div>
                <div class="flex flex-col my-1">
                    <label for="name">Deadline</label>
                    <input type="datetime-local" v-model="projectData.EstimatedEndDate" name="EstimatedEndDate" class="border border-[#490a47] rounded-sm p-1">
                </div>
                <div class="flex flex-col my-1">
                    <label for="tags">Tags:</label>
                    <input type="text" v-model="projectData.Tags" name="tags" placeholder="comma separated tags" class="border border-[#490a47] rounded-sm p-1">
                </div>
                <input type="submit" :value="edit ?  'update' : 'create'" class="mt-3 bg-[#490a47] text-white font-medium p-3 cursor-pointer" >
            </form>
        </div>
    </div> 
</template>

<script setup>
import { ref, defineEmits, onMounted} from 'vue';
import { useProjectController } from "../APIs/project-controller";

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

const emit = defineEmits(['modal-off', 'project-created', 'project-updated']);

const projectForm = ref(true);

const nullProjFields = {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            Tags:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Objectives:[],
            Notes:[]
}

const projectData = ref(nullProjFields)

const bye = () => {
    // projectData.value = nullProjFields
    emit("modal-off");
    console.log(projectData.value)
};

const submitForm = async () => {
    const { createProject, updateProject } = useProjectController();
    let form = document.getElementById('project-form');
        const formItems = [...form];
        //from array of form items filter out the submit value 
        //then grab name&value pairs for each of the remaining items.
        const keyVals = formItems.filter(i => i.type != "submit").map(item => [item.name, item.value]);
        //then turn them into an object
        const data = Object.fromEntries(keyVals);
        // console.log(data)
        data.EstimatedEndDate = new Date(data.EstimatedEndDate).toJSON()
    if(!props.edit){
        await createProject(data);
        emit("project-created");
    }else{
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

