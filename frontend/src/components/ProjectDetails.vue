<template>
    <div class="board grid grid-cols-10 py-4">
        <div class="board__objectives col-span-3 flex flex-col gap-2">
            <objective-title-card 
                :projectId="project.ID"
                @objective-created="getProject"
            />
            <div>
                <objective-filter
                    @sort-by-created="sortsortObjectivesByCreated"
                    @sort-by-updated="sortsortObjectivesByUpdated"
                    @sort-by-progress="sortObjectivesByProgress"
                />
            </div>
            <objective-card 
                v-for="obj in sortedObjectives" 
                :objective="obj"  
                :activeId="activeObjectiveId"
                :key="obj.ID"
                @objective-edited="getProject"
                @objective-deleted="getProject"
                @new-activation="handleNewActivation"
            />
        </div>
        <div class="board__goals col-span-7 px-4 flex flex-col gap-2">
            <details-summary-card :project="project"/>
            <div class="grid grid-cols-2 gap-2">
                <goal-card v-for="i in currentObjectiveGoals" :goal="i" @refresh="getProject"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import {ref, onMounted, computed, watch} from 'vue';
import { useRoute } from 'vue-router';
import ObjectiveCard from './ObjectiveCard.vue';
import ObjectiveTitleCard from './ObjectiveTitleCard.vue';
import DetailsSummaryCard from './DetailsSummaryCard.vue';
import GoalCard from './GoalCard.vue';
import {useProjectStore} from "../stores/project.js";
import ObjectiveFilter from './ObjectiveFilter.vue';

const route = useRoute();

const project =  computed(()=>{
    return projectStore.currentProject
})

const activeObjectiveId = ref(0)

const projectStore = useProjectStore();


const handleNewActivation = (data)=>{
    //data is the objective ID
activeObjectiveId.value = data
}

const getProject = async ()=>{
    await projectStore.getOneProject(route.params.id)
}

const sortedObjectives = ref([])

const sortsortObjectivesByCreated =()=>{
    sortedObjectives.value = [
        ...sortedObjectives.value.sort((a,b)=>{return a.CreatedAt > b.CreatedAt ? -1 : 1})
    ]
}


const sortsortObjectivesByUpdated = ()=>{
    sortedObjectives.value = [
        ...sortedObjectives.value.sort((a,b)=>{ return a.UpdatedAt > b.UpdatedAt ? -1 : 1 })
    ]
}

const sortObjectivesByProgress = ()=>{
    sortedObjectives.value = [
        ...sortedObjectives.value.sort((a,b)=>{return a.Progress > b.Progress ? -1 : 1})

    ]
}



const currentObjectiveGoals = computed(() => {
    if (project.value.Objectives && project.value.Objectives.length > 0) {
        const obje = project.value.Objectives.find((obj) => obj.ID == activeObjectiveId.value);
        if (obje && obje.Goals != null) {
            return obje.Goals;
        } else {
            return [];
        }
    } else {
        return [];
    }
});

// Watch for changes in project and update activeObjectiveId if necessary
watch(project, (newProject) => {
    if (newProject.Objectives && newProject.Objectives.length > 0) {
        const obje = newProject.Objectives
        .sort((a,b)=>{return a.ID > b.ID ? -1 : 1})
        .find((obj) => obj.ID == activeObjectiveId.value);
        if (!obje) {
            activeObjectiveId.value = newProject.Objectives[0].ID; // Set to the first objective if current is not found
        }
    }
});

onMounted(async ()=>{
    await getProject()
    if(sortedObjectives.value.length==0){
        sortedObjectives.value = project.value.Objectives.sort((a, b) => { return a.ID > b.ID ? -1 : 1 })
        activeObjectiveId.value = sortedObjectives.value[0].ID;
    }
})
</script>

