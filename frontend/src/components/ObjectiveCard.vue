<template>
    <div 
        class="p-4 rounded-lg" @click="handleClickOutside"
        :class="activeId==objective.ID?'active': 'bg-gloom'"
        >
        <div class="flex items-center justify-between relative" ref="menuContainer">
            <h4 class="font-semibold cursor-pointer" @click="activate">{{objective.Name || "OBJECTIVE NAME"}}</h4>
            <bars3-bottom-right-icon class="w-5 h-5 font-semibold stroke-2 cursor-pointer" @click="popupMenu= !popupMenu"/>
            <div 
                class="popout-menu absolute right-5 top-0 shadow-xl rounded-md w-44 bg-gloom p-2"  
                v-if="popupMenu"
                @click="activate"
            >
                <ul class="flex flex-col">
                    <li class="cursor-pointer flex gap-4 items-center my-1 p-2" @click="toggleGoalFormOn">
                        <PlusIcon class="w-3 h-3  text-black"/>
                        <span class="text-sm text-black">Add Goal</span>
                    </li>
                    <li class="cursor-pointer flex gap-4 items-center my-1 p-2">
                        <BookOpenIcon class="w-4 h-4"/>
                        <span class="text-sm" @click="openObjectiveNoteForm">Add Note</span>
                    </li>
                    <li class="cursor-pointer flex gap-4 items-center my-1 p-2" @click="editObjectiveClicked">
                        <PencilIcon class="w-3 h-3  text-black"/>
                        <span class="text-sm text-black">Edit Objective</span>
                    </li>
                    <li class="cursor-pointer flex gap-4 items-center my-1 hover:bg-red-300 hover:rounded-md duration-300 p-2" @click="deleteSwitch">
                        <TrashIcon class="w-3 h-3   text-red-600 "/>
                        <span class="text-sm text-red-600 ">Delete Objective</span>
                    </li>
                </ul>
            </div>
        </div>

        <div v-if="showNoteForm">
            <NoteForm 
                @modal-off="closeObjectiveNoteForm"
                :objectiveId="objective.ID"
                @note-created="closeObjectiveNoteForm"
            />
        </div>

        <div>
            <ObjectiveForm
            :edit="true"
            :objective="objective"
            v-if="ObjectiveFormOn" 
            @modal-off="toggleObjectiveFormOff"
            @objective-updated = "handleProjectUpdated"
            />
        </div>

        <div>
            <GoalForm
            :objectiveId="objective.ID"
            v-if="GoalFormOn" 
            @modal-off="toggleGoalFormOff"
            @goal-created="handleGoalCreated"
            />
        </div>

        <Dialog 
            v-if="showDelete"  
            entity="objective"
            @cancelled="showDelete=false"
            @deleted="deleteObjective" 
        />
        <div class="objective-card__notes w-24">
            <router-link to="/notes" class="flex gap-2 my-2 items-center">
                <BookOpenIcon class="text-royal w-4 h-4" />
                <span class="note-text font-medium text-royal">0 Notes</span>
            </router-link>
        </div>
        <div class="progress w-full">
            <span class="font-semibold py-1">Progress:</span> <span class="font-semibold text-2xl">{{objective.Progress}}%</span>
            <br>
            <progress max="100" class="w-[80%]" :value="objective.Progress"></progress>
        </div>
        <div class="objective-card__progress">
            <span class="font-medium">Overseer: {{objective.Overseer}}</span>
        </div>
        <div class="objective-card__descrption my-2">
            <span>Goals: {{ objective.Goals == null? 0 : objective.Goals.length }}</span>
        </div>
        <div class="descrption flex flex-col">
            <span class="font-medu">Description:</span>
            <div>
                {{objective.Description}}
            </div>
        </div>
    </div>
</template>

<script setup>
import {ref} from 'vue';
import {Bars3BottomRightIcon, PencilIcon, TrashIcon, BookOpenIcon, PlusIcon} from '@heroicons/vue/24/solid';
import ObjectiveForm from '../components/ObjectiveForm.vue'
import GoalForm from '../components/GoalForm.vue'
import { useObjectiveController } from "../APIs/objective-controller";
import Dialog from "./Dialog.vue";
import NoteForm from '../components/NoteForm.vue'

const props = defineProps({
    objective:{
        type: Object,
        default: {}
    },
    activeId:{
        type:Number,
        default:0
    }
})

const emit = defineEmits(['objective-edited','objective-deleted', 'new-activation'])
const showDelete = ref(false)
const popupMenu = ref(false) 
const menuContainer = ref(null);
const ObjectiveFormOn = ref(false)  
const GoalFormOn = ref(false)
const showNoteForm = ref(false)

const activate = ()=>{
 emit('new-activation', props.objective.ID)
}

const handleClickOutside = (event)=> {
      // Check if the click event occurred outside the menu container
      if (!menuContainer.value.contains(event.target)) {
        popupMenu.value = false;
      }
    }

const openObjectiveNoteForm = ()=>{
    showNoteForm.value = true
}

const closeObjectiveNoteForm = ()=>{
    showNoteForm.value = false;
}

const editObjectiveClicked = ()=>{
    toggleObjectiveFormOn()
    popupMenu.value = false;
}

const handleProjectUpdated = ()=>{
    emit("objective-edited")
    toggleObjectiveFormOff()
}

const deleteObjective = async ()=>{
    const {deleteObjective} = useObjectiveController();
    try {
        await deleteObjective(props.objective.ID);
        showDelete.value = false;
        emit('objective-deleted')
    }catch (e){
        throw e
    }
}

const toggleObjectiveFormOn = ()=>{
    ObjectiveFormOn.value = true
}

const toggleObjectiveFormOff = ()=>{
    ObjectiveFormOn.value = false
}

const toggleGoalFormOn = ()=>{
    GoalFormOn.value=true
}


const toggleGoalFormOff = ()=>{
    GoalFormOn.value=false
}


const deleteSwitch = ()=>{
    showDelete.value = true;
    popupMenu.value = false;
}

const handleGoalCreated = ()=>{
    toggleGoalFormOff();
    emit("objective-edited") //to reload project, also... it is edited
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

.active{
    background-color: rgba(73, 10, 71, 0.29);
}
</style>
