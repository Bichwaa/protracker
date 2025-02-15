import { defineStore } from 'pinia';
import {ref} from 'vue';
import {useProjectController } from '../APIs/project-controller.js'

export const useProjectStore = defineStore('project', () => {
  
  let projects =ref([])
  
  async function getProjects() {
    const {getProjects: gp} = useProjectController()
    projects.value =  await gp()
  }

  return { projects, getProjects }
})

