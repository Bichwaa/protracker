<template>
    <div class="bg-gloom p-4 rounded-lg">
        <div class="flex items-center justify-between">
            <h4 class="font-medium">Notes </h4> 
            <ArrowLeftIcon class="h-4 w-4 font-medium cursor-pointer" />
        </div>
        <!-- <div class="objective-card__descrption">
            <button 
              class="mt-3 bg-[#490a47] hover:bg-black transition duration-300 disabled:bg-gray-200 text-white text-sm font-medium p-2 cursor-pointer grid place-content-center rounded-md"
              @click="toggleObjectiveFormOn"
            >
                <span class=""> + Add Note </span>
            </button>
        </div>
        <div>
            <NoteForm
            :edit="objectiveFormEditMode"
            :projectId="projectId"
            v-if="ObjectiveFormOn" 
            @modal-off="toggleProjectFormOff"
            @objective-created="handleObjectiveCreated"
            />
        </div> -->
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ArrowLeftIcon } from '@heroicons/vue/24/solid';
// import NoteForm from '../components/NoteForm.vue'

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