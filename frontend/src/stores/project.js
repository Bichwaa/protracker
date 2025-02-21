import { defineStore } from 'pinia';
import {ref} from 'vue';
import {useProjectController } from '../APIs/project-controller.js'

export const useProjectStore = defineStore('project', () => {

  let projectCache = ref([])
  let filterTags = ref(new Set())
  let projects =ref([])
  let currentProject = ref({})
  
  async function getProjects() {
    const {getProjects: gp} = useProjectController()
    projects.value =  await gp()
    projectCache.value = projects.value
  }

  async function getOneProject(id) {
    const { getOneProject } = useProjectController()
    currentProject.value =  await getOneProject(Number(id))
  }

  async function filterProjectsByTag(tag) {
    if(filterTags.value.size<1){ // if projects haven't been filtered before
      filterTags.value.add(tag)
      projects.value = projects.value.filter(project=>{
      //tag is supposed to be a possible substring to project.Tags
      // we check if it is.
      return project.Tags.toLowerCase().indexOf(tag.toLowerCase()) >= 0
    })
  }else{ //if projects have been filtered before
    const preAddSize = filterTags.value.size
    filterTags.value.add(tag)
    const postAddSize = filterTags.value.size
    if(postAddSize>preAddSize){ //if the projects state hasn't already been filtered for this tag
      const newFilteredBatch = projectCache.value.filter(project=>{
        //tag is supposed to be a possible substring to project.Tags
        // we check if it is.
        return project.Tags.toLowerCase().indexOf(tag.trim().toLowerCase()) >= 0
      })
      //make sure all project elements are unique
      const currentProjectIDs = projects.value.map(x=>x.ID);
      newFilteredBatch.forEach(x=>{
        if(!currentProjectIDs.includes(x.ID)){
          projects.value.push(x)
        }
      })
    }
  }
}

async function resetFilters(){
  filterTags.value.clear()
  projects.value = projectCache.value;
}

  return { projects, getProjects, getOneProject, currentProject, filterProjectsByTag, resetFilters }
})