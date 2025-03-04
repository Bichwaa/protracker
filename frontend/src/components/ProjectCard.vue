<template>
    <div class="project-card bg-gloom rounded-lg p-4"  @click="handleClickOutside">
        <div class="flex justify-between relative">
            <h3 class="project-card__title capitalize font-semibold">
                <router-link :to="{name:'Details', params:{id:project.ID}}" class="font-semibold">{{project.Name}}</router-link>
            </h3>

            <Dialog 
                v-if="showDelete"  
                @cancelled="showDelete=false"
                @deleted="deleteProject"
            />

            <div class="relative" ref="menuContainer">
                <Bars3BottomRightIcon class="text-black w-6 h-6 cursor-pointer" @click="popupMenu= !popupMenu"/>

                <div class="popout-menu absolute right-5 top-0 shadow-xl rounded-md w-44 bg-gloom p-2"  v-if="popupMenu" >
                    <ul class="flex flex-col">
                        <li class="cursor-pointer my-1 p-2">
                            <router-link :to="{name:'Details', params:{id:project.ID}}"  class="flex gap-4 items-center ">
                                <EyeIcon class="w-3 h-3 text-black"/>
                                <span class="text-sm">Project Details</span>
                            </router-link>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 p-2">
                            <BookOpenIcon class="w-4 h-4"/>
                            <span class="text-sm" @click="openProjectNoteForm">Add Note</span>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 p-2" @click="editProjectClicked">
                            <PencilIcon class="w-3 h-3  text-black"/>
                            <span class="text-sm text-black">Edit Project</span>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 hover:bg-red-300 hover:rounded-md duration-300 p-2" @click="deleteSwitch">
                            <TrashIcon class="w-3 h-3   text-red-600 "/>
                            <span class="text-sm text-red-600 ">Delete Project</span>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
        <div v-if="showNoteForm">
            <NoteForm 
                @modal-off="closeProjectNoteForm"
                :projectId="project.ID"
                @note-created="closeProjectNoteForm"
            />
        </div>
        <div class="project-card__notes w-24">
            <router-link :to="{ name: 'ProjectNotes', params: { projectId: project.ID }}" class="flex gap-2 my-2 items-center">
                <BookOpenIcon class="text-royal w-4 h-4" />
                <span class="note-text font-medium text-royal">
                    {{ project.Notes.length }} Note{{ project.Notes.length!=1?'s':'' }}
                </span>
            </router-link>
        </div>
        <div class="progress w-full">
            <span class="font-semibold">Progress:</span> <span class="font-semibold text-3xl">{{project.Progress}}%</span>
            <br>
            <progress max="100" v-bind:value="project.Progress" class="w-[80%]"></progress>
        </div>
        <div class="time"><span class="font-semibold">Age:</span> <span> {{timeLapsed.days}} Days | {{timeLapsed.hours}} Hours</span></div>
        <div class="tags my-4">
           <span> <a class="text-sm bg-gray-200 mx-1 rounded-xl px-2 py-1 cursor-pointer" v-for="tag in tagList" @click="tagClicked(tag)">{{tag}}</a></span>
        </div>
        <div class="descrption flex flex-col">
            <span class="font-semibold">Description:</span>
            <div>
                {{project.Description}}
            </div>
        </div>
    </div>
</template>

<script setup>
import {ref, computed} from 'vue';
import {useProjectController} from "../APIs/project-controller";
import {Bars3BottomRightIcon, PencilIcon, TrashIcon, BookOpenIcon,  EyeIcon } from '@heroicons/vue/24/solid';
import Dialog from "./Dialog.vue";
import NoteForm from '../components/NoteForm.vue'

    const props = defineProps(['project'])

    const emit = defineEmits([
        "project-deleted", "edit-clicked",
        "tag-clicked",
    ])

    const showDelete = ref(false)
    const showNoteForm = ref(false)
    
    const r =  ref("no!")
    const popupMenu = ref(false) 
    const menuContainer = ref(null);
    
    
    const tagClicked = (tag)=> {
            emit("tag-clicked",{tag});
        }
    
    const handleClickOutside = (event)=> {
      // Check if the click event occurred outside the menu container
      if (menuContainer.value !=null && !menuContainer.value.contains(event.target)) {
        popupMenu.value = false;
    }
    }

    const openProjectNoteForm = ()=>{
        showNoteForm.value = true
    }

    const closeProjectNoteForm = ()=>{
        showNoteForm.value = false;
    }

    const editProjectClicked = ()=>{
        emit("edit-clicked")
        popupMenu.value = false;
    }

    const deleteProject = async ()=>{
        const {deleteProject} = useProjectController();
        try {
            await deleteProject(props.project.ID);
            showDelete.value = false;
            emit('project-deleted')
        }catch (e){
            throw e
        }
    }

    const deleteSwitch = ()=>{
        showDelete.value = true;
        popupMenu.value = false;
    }

    const tagList = computed(()=>{
        const tagString = props.project.Tags;
        if(tagString && tagString.length>0){
            return tagString.split(',')
        }else {
            return []
        }
    })

    const timeLapsed = computed(() => {
        // Parse the input date string into a Date object
        const inputDate = new Date(props.project.CreatedAt);
        
        // Get the current date and time
        const currentDate = new Date();
        
        // Calculate the difference in milliseconds
        const differenceInMilliseconds = currentDate - inputDate;
        
        // Convert milliseconds to seconds
        const differenceInSeconds = differenceInMilliseconds / 1000;
        
        // Convert seconds to hours
        const differenceInHours = differenceInSeconds / 3600;
        
        // Calculate the number of full days
        const days = Math.floor(differenceInHours / 24);
        
        // Calculate the remaining hours
        const hours = Math.floor(differenceInHours % 24);
        
        return { days, hours };
    })

</script>


<style scoped>
progress {
  border-radius: 15px;
  /* height:60%; */
}
progress::-webkit-progress-bar {
  /* style rules */
  border-radius: 15px;
  height: 60%;
  /* background-color: #490A47; */
  
}
progress::-webkit-progress-value {
  /* style rules */
  border-radius: 15px;
  background-color: #490A47;
  
}
progress::-moz-progress-bar {
  /* style rules */
  border-radius: 15px;
  /* height: 60%; */
  background-color: #490A47;
}
</style>