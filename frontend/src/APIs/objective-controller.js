import {CreateObjective, FetchObjective, FetchObjectives, UpdateObjective, DeleteObjective} from '../../wailsjs/go/main/App'

export function useObjectiveController(){

    async function createObjective(projectData){
        await CreateObjective(projectData);
    }

    async function getObjectives() { 
       return await FetchObjectives()
    }

    async function deleteObjective(projectId) {
        return DeleteObjective(projectId)
    }

    async function updateObjective(updatedData) {
        return UpdateObjective(updatedData)
    }

    async function getOneObjective(projectId) {
        return FetchObjective(projectId)
    }

    return {createObjective, getObjectives, deleteObjective, updateObjective, getOneObjective  }
}
