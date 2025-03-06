<template>
    <div class="board py-4">
        <div class="board__notes grid grid-cols-12 gap-2">
            <div class="hide-scrollbar col-span-4 pr-2 overflow-scroll h-[94vh]">
                <notes-title-card class="mb-3"/>
                <note-card 
                    v-for="note in sortedNotes" 
                    :key="note.ID" 
                    :note="note" 
                    :activeId="activeNote.ID"
                    @new-activation="handleNewActivation"
                />
            </div>
            <div class="board__notes__current col-span-8 px-8" v-if="activeNote!=null">
                <h3 class="font-medium text-lg">{{ activeNote.Title }}</h3>
                <div v-html="activeNote.Content"></div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import NoteCard from './NoteCard.vue';
import NotesTitleCard from './NotesTitleCard.vue';
import { useNotesStore } from '../stores/notes';

const noteStore = useNotesStore()

const route = useRoute()

const activeNote = ref(null)

const handleNewActivation = (data)=>{
    activeNote.value = data
}

const sortedNotes = computed(()=>{
    const notes = noteStore.notes
    if(notes.length>0){
        const sorted = notes.sort((a,b)=>{return a.ID > b.ID ? -1 : 1})
        activeNote.value = sorted[0] 
        return sorted
    }else{
        activeNote.value = null
        return []
    }
})

onMounted(async()=>{
    if(route.name=="ProjectNotes"){
        console.log(route.params)
        await noteStore.filterNotesBy({Name:'project_id' ,Value: route.params.projectId})
    }else if(route.name=="ObjectiveNotes"){
        await noteStore.filterNotesBy({Name:'objective_id' ,Value: route.params.objectiveId})
    }else if(route.name=="GoalNotes"){
        await noteStore.filterNotesBy({Name:'goal_id' ,Value: route.params.goalId})
    }else{
        await noteStore.getNotes()
    }
})

</script>

<style scoped>
.hide-scrollbar::-webkit-scrollbar {
    display: none; /* Safari and Chrome */
}

.hide-scrollbar {
    -ms-overflow-style: none;  /* IE and Edge */
    scrollbar-width: none;  /* Firefox */
}
</style>