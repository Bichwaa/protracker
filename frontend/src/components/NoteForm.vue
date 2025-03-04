<template>
    <!-- The Modal -->
    <div id="myModal" class="modal">
        <!-- Modal content -->
        <div class="modal-content w-2/3 rounded-lg">
            <div class="w-full flex flex-row-reverse">
                <span class="text-2xl cursor-pointer hover:text-red-500 duration-300" @click="bye">&times;</span>
            </div>
            <form method="post" id="objective-form" @submit.prevent="submitForm" v-if="objectiveForm" class="flex flex-col">
                <div class="flex flex-col my-1">
                    <label for="Title">Title</label>
                    <input 
                        type="text" 
                        v-model="noteData.Title" 
                        name="Title" placeholder="Title" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Title = ''"
                        >
                    <span class="text-sm text-red-500" v-if="errors.Title" >{{ errors.Title }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="Tags">Tags</label>
                    <input 
                        ref="tagsInput"
                        type="text" 
                        v-model="noteData.Tags" 
                        name="Tags" placeholder="Tags" 
                        class="border border-[#490a47] rounded-sm p-1"
                        @focus="errors.Tags = ''"
                        >
                    <span class="text-sm text-red-500" v-if="errors.Tags" >{{ errors.Tags }}</span>
                </div>
                <div class="flex flex-col my-1">
                    <label for="description">Content</label>
                        <VueTailwindEditor @get-html="updateHtml" :width="inpWdt"/>
                    <span class="text-sm text-red-500" v-if="errors.Content" >{{ errors.Content }}</span>
                </div>
                <button type="submit" class="mt-3 bg-[#490a47] disabled:bg-gray-200 text-white font-medium p-3 cursor-pointer grid place-content-center rounded-md" >
                    <span v-if="!loading"> {{ edit ?  'Update Note' : 'Create Note' }} </span>
                    <Loader v-else  class="h-5 w-5"/>
                </button>
            </form>
        </div>
    </div> 
</template>

<script setup>
import { ref, defineEmits, onMounted, computed} from 'vue';
import { z } from "zod";
import { useNoteController } from "../APIs/note-controller";
import Loader from './Loader.vue';
import { VueTailwindEditor } from 'vue-tailwind-wysiwyg-editor';
import 'vue-tailwind-wysiwyg-editor/dist/style.css'

const props = defineProps({
    edit: {
        type: Boolean,
        default: false
    },
    projectId:{
        type: Number,
        default:0
    },
    objectiveId:{
        type: Number,
        default:0
    },
    goalId:{
        type: Number,
        default:0
    },
    title : {
        type: Object,
        default: {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Title:null,
            Tags:null,
            Content:null
        }
    }
});

const tagsInput = ref(null)

const inpWdt = computed(()=>{
    if(tagsInput.value!=null){
        return tagsInput.value.offsetWidth
    }else{
        return 0
    }
})

const html = ref("")

const updateHtml = (val) => {
  html.value = val;
};

const loading = ref(false)

const emit = defineEmits(['modal-off', 'note-created', 'objective-updated']);

const objectiveForm = ref(true);

const nullObjFields = {
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Title:null,
            Tags:null,
            Content:null
}

const noteData = ref(nullObjFields)
//validation Logic
const projectschema = z.object({
            ID:z.number().int().optional(),
            Title:z.string({message:"Note must have a title"}).nonempty(),
            Tags:z.string(),
            Content:z.string(),
})

const errors = ref({
            Title:"",
            Tags:"",
            Content:""
})

//End validation Logic

const bye = () => {
    emit("modal-off");
};

const submitForm = async () => {
    loading.value = true;
    const { createNote, updateNote } = useNoteController();
    let form = document.getElementById('objective-form');
        const formItems = [...form];
        //from array of form items filter out the submit value 
        //then grab name&value pairs for each of the remaining items.
        const keyVals = formItems.filter(i => i.type != "submit").map(item => [item.name, item.value]);
        //then turn them into an object
        const data = Object.fromEntries(keyVals);
        // console.log(data)
        data.Content = html.value;
        const validation = projectschema.safeParse(data)
        if (!validation.success) {
            console.log('Form data:', validation);
            // Handle validation errors
            validation.error.errors.forEach((err) => {
            errors.value[err.path[0]] = err.message;
            });
        } else {
            // Form is valid, proceed with submission
            loading.value = false;

        }
    if(!props.edit && validation.success){

        if(props.projectId!=0){
            data.ProjectID = props.projectId;
        } else if(props.objectiveId!=0){
            data.ObjectiveID = props.objectiveId;
        } else if(props.goalId!=0){
            data.GoalID = props.goalId;
        }
        await createNote(data);
        emit("note-created");
    }else if(props.edit && validation.success){
        data.ID = noteData.value.ID
        await updateNote(data);
        // console.log(data)
        emit("objective-updated");
    }
}

onMounted(()=>{
    noteData.value = nullObjFields
    if(props.edit){
        noteData.value = props.objective
        noteData.value.EstimatedEndDate = new Date(props.objective.EstimatedEndDate).toISOString().substring(0, 16)
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

