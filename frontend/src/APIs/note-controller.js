import { CreateNote, FetchNote, FetchNotes, UpdateNote, DeleteNote } from '../../wailsjs/go/main/App.js'

export function useNoteController() {

    async function createNote(noteData) {
        await CreateNote(noteData);
    }

    async function getNotes() { 
        return await FetchNotes();
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

    return { createNote, getNotes, deleteNote, updateNote, getOneNote }
}