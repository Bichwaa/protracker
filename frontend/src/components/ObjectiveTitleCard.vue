<template>
    <div class="bg-gloom p-4 rounded-lg">
        <div class="flex items-center justify-between">
            <h4 class="font-medium">Objectives </h4> 
            <ArrowDownIcon class="h-4 w-4 font-medium" />
        </div>
        <div class="objective-card__progress py-2 text-gray-500 text-sm">
            <span>"The purpose of a system is what it does"</span>
        </div>
        <div class="objective-card__descrption">
            <button 
              class="mt-3 bg-[#490a47] hover:bg-black transition duration-300 disabled:bg-gray-200 text-white text-sm font-medium p-2 cursor-pointer grid place-content-center rounded-md"
              @click="toggleObjectiveFormOn"
            >
                <span class=""> + Add Objective </span>
            </button>
        </div>
        <div>
            <ObjectiveForm
            :edit="objectiveFormEditMode"
            :projectId="projectId"
            v-if="ObjectiveFormOn" 
            @modal-off="toggleProjectFormOff"
            @objective-created="handleObjectiveCreated"
            />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ArrowDownIcon } from '@heroicons/vue/24/solid';
import ObjectiveForm from '../components/ObjectiveForm.vue'

const props = defineProps(['projectId'])
const emit = defineEmits(['objective-created'])
const ObjectiveFormOn = ref(false);
const objectiveFormEditMode = ref(false);

const toggleObjectiveFormOn = ()=>{
  ObjectiveFormOn.value = true
}

const toggleProjectFormOff = ()=>{
  objectiveFormEditMode.value = false;
  ObjectiveFormOn.value = false
}

const handleObjectiveCreated = async ()=>{
  toggleProjectFormOff()
  emit("objective-created")
}

</script>