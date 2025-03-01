<template>
  <div class="home">
    <section class="home__topsection grid grid-cols-3 pr-4 mt-4 gap-4">
      <div class="grid col-span-1">
        <InfoCard v-on:add-project-clicked="toggleProjectFormOn"/>
      </div>
      <div class="grid col-span-2">
        <SummaryCard :projects="projects"/>
      </div>
    </section>
    <section>
      <ProjectForm 
        :edit="projectFormEditMode"
        :project="projectToedit"
        v-if="ProjectFormOn" 
        @modal-off="toggleProjectFormOff"
        @project-created="handleProjectCreated"
        @project-updated = "handleProjectUpdated"
        />
    </section>
    <section class="pr-4 pt-2" v-if="filters.show">
      <div class="flex p-1 items-center gap-2">
        <span class="text-sm font-medium">Filters:</span>
        <span> <a class="text-sm bg-gray-200 mx-1 rounded-xl px-4  cursor-pointer" v-for="tag in filters.filterTagList">{{tag}}</a></span>
        <span class="flex items-center cursor-pointer" @click="clearFilters">
          <span class="text-[1.35rem] text-red-400 mt-[.1rem]">&times;</span>
          <span class="text-sm text-red-400"> clear</span>
        </span>
      </div>
    </section>
    <section class="grid grid-cols-3 gap-4 pr-4 pt-4">
        <template v-if="projects">
        <ProjectCard
        class="col-span-1"
          v-for="project in projects" 
          :key="project.ID"
          :id="project.ID"
          v-bind:project="project"
          @project-deleted="getProjects"
          @edit-clicked="openEditForm(project)"
          @tag-clicked="filterProjects"
        />
        </template>
    </section>
  </div>
</template>

<script setup>
import {ref, onMounted,computed} from "vue";
import ProjectCard from './ProjectCard.vue';
import SummaryCard from './SummaryCard.vue';
import ProjectForm from './ProjectForm.vue';
import InfoCard from "./InfoCard.vue"
import {useProjectStore} from "../stores/project.js";


const loading = ref(false);
const ProjectFormOn = ref(false);
const projectFormEditMode = ref(false);
const projectToedit = ref({
            ID:null,
            CreatedAt: null,
            UpdatedAt: null,
            DeletedAt:null,
            Name:null,
            Tags:null,
            BeginDate:null,
            EstimatedEndDate:null,
            Progress:null,
            Description:null,
            Objectives:[],
            Notes:[]
        })

const filters = ref({
  show:false,
  filterTagList:[]
})

const projectStore = useProjectStore();

onMounted(()=>{
  getProjects()
});

const projects = computed(()=>{
  return projectStore.projects
})

const getProjects = async ()=>{
  const data = await projectStore.getProjects()
  // projectStore.projects = data
}

const filterProjects = async (data)=>{
  // data should be an object with the shape {tag: string}
  filters.value.filterTagList.push(data.tag)
  filters.value.show = true
  await projectStore.filterProjectsByTag(data.tag)
}

const clearFilters = async ()=>{
  // await getProjects()
  await projectStore.resetFilters()
  filters.value = {show:false,filterTagList:[]}
}

const toggleProjectFormOn = ()=>{
  ProjectFormOn.value = true
}

const toggleProjectFormOff = ()=>{
  projectFormEditMode.value = false;
  ProjectFormOn.value = false
}

const openEditForm = (project)=>{
  projectFormEditMode.value = true;
  projectToedit.value = project;
  toggleProjectFormOn()
}


const handleProjectCreated = async ()=>{
  toggleProjectFormOff()
  await getProjects()
}

const handleProjectUpdated = async ()=>{
  toggleProjectFormOff()
  await getProjects()
}

 
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

::backdrop {
  background-image: linear-gradient(
    45deg,
    magenta,
    rebeccapurple,
    dodgerblue,
    green
  );
  opacity: 0.75;
}
</style>