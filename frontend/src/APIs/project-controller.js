import {CreateProject, FetchProject, FetchProjects, UpdateProject,DeleteProject} from '../../wailsjs/go/main/App'

export function useProjectController(){

    async function createProject(projectData){
        await CreateProject(projectData);
    }

    async function getProjects() { 
       return await FetchProjects()
    }

    async function deleteProject(projectId) {
        return DeleteProject(projectId)
    }

    async function updateProject(updatedData) {
        return UpdateProject(updatedData)
    }

    return {createProject, getProjects, deleteProject, updateProject}
}

// export async function getAllProjects(){
//     const store = useProjectStore()
//     await store.getProjects()
// }

