import {CreateTask, FetchTask, FetchTasks, UpdateTask, DeleteTask, MarkTaskAsCompleted} from '../../wailsjs/go/main/App.js'

export function useTaskController(){

    async function createTask(taskData){
        await CreateTask(taskData);
    }

    async function getTasks() { 
       return await FetchTasks();
    }

    async function deleteTask(taskId) {
        return DeleteTask(taskId);
    }

    async function updateTask(updatedData) {
        return UpdateTask(updatedData);
    }

    async function getOneTask(taskId) {
        return FetchTask(taskId);
    }

    async function markTaskAsCompleted(taskId) {
        return MarkTaskAsCompleted(taskId);
    }

    return {createTask, getTasks, deleteTask, updateTask, getOneTask, markTaskAsCompleted}
}
