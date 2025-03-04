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
            <div class="board__notes__current col-span-8 px-8" v-if="activeNote">
                <h3 class="font-medium text-lg">{{ activeNote.Title }}</h3>
                <div v-html="activeNote.Content"></div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import NoteCard from './NoteCard.vue';
import NotesTitleCard from './NotesTitleCard.vue';
import { useNotesStore } from '../stores/notes';

const noteStore = useNotesStore()

const PLACEHOLDER_TEXT = `Lorem ipsum dolor sit amet consectetur adipisicing elit. 
                    Provident, assumenda temporibus incidunt adipisci reiciendis iure sit esse. 
                    Quasi voluptates excepturi quod quo pariatur repellat possimus qui voluptatum 
                    laborum aliquid quaerat maiores aperiam explicabo enim voluptatibus, obcaecati 
                    cupiditate! Consequuntur, eligendi amet reprehenderit id corporis et totam quam
                    earum tempore fugiat, quia aliquam asperiores sequi. Perferendis quae hic 
                    voluptate iste asperiores ratione architecto. Similique quisquam ad voluptatum
                    quasi, nihil quod eius consectetur placeat eum nulla aliquid rerum quidem est. 
                    Fugiat ipsum dolore tenetur voluptatum earum! Minima natus corporis maiores
                    tenetur temporibus ipsa doloribus nesciunt officiis! Rerum ullam hic vel, 
                    perspiciatis vitae ab ratione, vero illo atque odit veritatis cumque quae 
                    dolorum voluptas ea nihil eveniet exercitationem at dolorem maxime facere a? 
                    Dolore repudiandae modi, temporibus animi quisquam assumenda reprehenderit. 
                    Laborum eum eius temporibus quaerat voluptas facere, ullam sed optio quos ab 
                    error incidunt harum officia animi aliquid sit. Amet omnis, fugit nihil 
                    blanditiis earum maxime! Suscipit sunt fuga a. Architecto placeat ea 
                    necessitatibus numquam excepturi, facilis dolores, incidunt officiis ab 
                    doloribus fugiat sed nulla vel itaque! Ratione nam, quas quibusdam magnam.`

const activeNote = ref(null)

const handleNewActivation = (data)=>{
    activeNote.value = data
}

// const notes=ref([
//                 {ID:1, Title:"this is a note title", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["one", "three"]},
//                 {ID:2, Title:"this is a note title 2", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["two", "three"]},
//                 {ID:3, Title:"this is a tripple note title", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["one", "three"]},
//                 {ID:4, Title:"this is a title 4 a note", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["one", "three"]},
//                 {ID:5, Title:"this is a galaxy note5 title", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["one", "three"]},
//                 {ID:6, Title:"Six is a note title", Content:PLACEHOLDER_TEXT.slice(0, Math.floor(Math.random()*PLACEHOLDER_TEXT.length)), tags:["one", "three"]},
//             ])


const sortedNotes = computed(()=>{
    const notes = noteStore.notes
    if(notes.length>0){
        const sorted = notes.sort((a,b)=>{return a.ID > b.ID ? -1 : 1})
        if(!activeNote.value){
            activeNote.value = sorted[0] 
        }
        return sorted
    }else{
        return []
    }
})

onMounted(async()=>{
 await noteStore.getNotes()
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