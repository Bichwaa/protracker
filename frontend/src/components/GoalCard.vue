<template>
    <div class="project-card bg-gloom rounded-lg p-4"  @click="handleClickOutside">
        <div class="flex justify-between relative">
            <h3 class="project-card__title capitalize font-semibold">
                <router-link :to="{name:'Details', params:{id:1}}" class="font-semibold">{{"jina kama lipo"}}</router-link>
            </h3>

            <Dialog 
                v-if="showDelete"  
                @cancelled="showDelete=false"
                @deleted="deleteProject" 
                class="absolute top-0"/>

            <div class="relative" ref="menuContainer">
                <Bars3BottomRightIcon class="text-black w-6 h-6 cursor-pointer" @click="popupMenu= !popupMenu"/>

                <div class="popout-menu absolute right-5 top-0 shadow-xl rounded-md w-44 bg-gloom p-2"  v-if="popupMenu" >
                    <ul class="flex flex-col">
                        <li class="cursor-pointer my-1 p-2">
                            <router-link :to="{name:'Details', params:{id:1}}"  class="flex gap-4 items-center ">
                                <EyeIcon class="w-3 h-3 text-black"/>
                                <span class="text-sm">Goal Details</span>
                            </router-link>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 p-2">
                            <BookOpenIcon class="w-3 h-3"/>
                            <span class="text-sm">Goal Notes</span>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 p-2" @click="editProjectClicked">
                            <PencilIcon class="w-3 h-3  text-black"/>
                            <span class="text-sm text-black">Edit Goal</span>
                        </li>
                        <li class="cursor-pointer flex gap-4 items-center my-1 hover:bg-red-300 hover:rounded-md duration-300 p-2" @click="deleteSwitch">
                            <TrashIcon class="w-3 h-3   text-red-600 "/>
                            <span class="text-sm text-red-600 ">Delete Goal</span>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="project-card__notes w-24">
            <router-link to="/notes" class="flex gap-2 my-2 items-center">
                <BookOpenIcon class="text-royal w-4 h-4" />
                <span class="note-text font-medium text-royal">0 Notes</span>
            </router-link>
        </div>
        <div class="progress w-full">
            <span class="font-semibold">Progress:</span> <span class="font-semibold text-3xl">{{8}}%</span>
            <br>
            <progress max="100" :value="8" class="w-[80%]"></progress>
        </div>
        <span class="font-semibold">Tasks</span>
        <div class="flex flex-col gap-2 mt-2 h-36 overflow-scroll hide-scrollbar">
            <TaskCard v-for="i in 4" />
        </div>
    </div>
</template>

<script setup>
import {ref, computed} from 'vue';
import {useProjectController} from "../APIs/project-controller";
import {Bars3BottomRightIcon, PencilIcon, TrashIcon, BookOpenIcon,  EyeIcon } from '@heroicons/vue/24/solid';
import Dialog from "./Dialog.vue";
import TaskCard from './TaskCard.vue';

    const props = defineProps(['project'])

    const emit = defineEmits([
        "project-deleted", "edit-clicked",
        "tag-clicked",
    ])

    const showDelete = ref(false)
    
    const r =  ref("no!")
    const popupMenu = ref(false) 
    const menuContainer = ref(null);
    
    
    const tagClicked = (tag)=> {
            emit("tag-clicked",{tag});
        }
    
    const handleClickOutside = (event)=> {
      // Check if the click event occurred outside the menu container
      if (!menuContainer.value.contains(event.target)) {
        popupMenu.value = false;
      }
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
</script>

<style scoped>
progress {
  border-radius: 15px;
  /* height:20%; */
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

.hide-scrollbar::-webkit-scrollbar {
    display: none; /* Safari and Chrome */
}

.hide-scrollbar {
    -ms-overflow-style: none;  /* IE and Edge */
    scrollbar-width: none;  /* Firefox */
}
</style>