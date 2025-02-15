<template>
  <div class="home">
    <section class="home__topsection grid grid-cols-3 px-4 mt-4 gap-4">
      <div class="grid col-span-1">
        <InfoCard v-on:add-project-clicked="toggleProjectForm"/>
      </div>
      <div class="grid col-span-2">
        <SummaryCard />
      </div>
    </section>
    <section>
      <ProjectForm 
        :edit="projectFormEditMode"
        :project="projectToedit"
        v-if="ProjectFormOn" 
        @modal-off="toggleProjectForm"
        @project-created="handleProjectCreated"
        @project-updated = "handleProjectUpdated"
        />
    </section>
    <p v-if="loading">...loading</p>
    <section class="grid grid-cols-3 gap-4 p-4">
        <template v-if="projects">
        <ProjectCard
        class="col-span-1"
          v-for="project in projects" 
          :key="project.ID"
          :id="project.ID"
          v-bind:project="project"
          @project-deleted="getProjects"
          @edit-clicked="openEditForm(project)"
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

const toggleProjectForm = ()=>{
  ProjectFormOn.value = !ProjectFormOn.value
}

const openEditForm = (project)=>{
  projectFormEditMode.value = true;
  projectToedit.value = project;
  toggleProjectForm()
}


const handleProjectCreated = async ()=>{
  toggleProjectForm();
  await getProjects()
}

const handleProjectUpdated = async ()=>{
  toggleProjectForm();
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