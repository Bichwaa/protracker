import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useNoteController } from '../APIs/note-controller.js'

export const useNotesStore = defineStore('notes', () => {
  
  let notesCache = ref([]);
  let notes = ref([]);
  let currentNote = ref({});

  async function getNotes() {
    const { getNotes: gn } = useNoteController();
    notes.value = await gn();
    notesCache.value = notes.value;
  }

  async function getOneNote(id) {
    const { getOneNote } = useNoteController();
    currentNote.value = await getOneNote(Number(id));
  }

  async function filterNotesByTag(tag) {
    notes.value = notesCache.value.filter(note => 
      note.Tags.toLowerCase().indexOf(tag.toLowerCase()) >= 0
    );
  }

  async function resetFilters() {
    notes.value = notesCache.value;
  }

  return { notes, getNotes, getOneNote, currentNote, filterNotesByTag, resetFilters }
})