import {CreateGoal, FetchGoal, FetchGoals, UpdateGoal, DeleteGoal} from '../../wailsjs/go/main/App'

export function useGoalController(){

    async function createGoal(projectData){
        await CreateGoal(projectData);
    }

    async function getGoals() { 
       return await FetchGoals()
    }

    async function deleteGoal(projectId) {
        return DeleteGoal(projectId)
    }

    async function updateGoal(updatedData) {
        return UpdateGoal(updatedData)
    }

    async function getOneGoal(projectId) {
        return FetchGoal(projectId)
    }

    return {createGoal, getGoals, deleteGoal, updateGoal, getOneGoal}
}
