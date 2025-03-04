import { CreateNote, FetchNote, FetchNotes,FilterNotes, UpdateNote, DeleteNote } from '../../wailsjs/go/main/App.js'

export function useNoteController() {

    async function createNote(noteData) {
        await CreateNote(noteData);
    }

    async function getNotes() { 
        return await FetchNotes();
    }

    async function filterNotes(param){
        return await FilterNotes(param)
    }

    async function deleteNote(noteId) {
        return DeleteNote(noteId);
    }

    async function updateNote(updatedData) {
        return UpdateNote(updatedData);
    }

    async function getOneNote(noteId) {
        return FetchNote(noteId);
    }

    return { createNote, getNotes, deleteNote, updateNote, getOneNote, filterNotes }
}