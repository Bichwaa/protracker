<template>
    <div class="board grid grid-cols-10 py-4">
        <div class="board__objectives col-span-3 flex flex-col gap-2">
            <objective-title-card 
                :projectId="project.ID"
                @objective-created="getProject"
            />
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
import {ref, onMounted, computed} from 'vue';
import { useRoute } from 'vue-router';
import ObjectiveCard from './ObjectiveCard.vue';
import ObjectiveTitleCard from './ObjectiveTitleCard.vue';
import DetailsSummaryCard from './DetailsSummaryCard.vue';
import GoalCard from './GoalCard.vue';
import {useProjectStore} from "../stores/project.js";
import {useGoalController} from "../APIs/goal-controller.js";

const route = useRoute();

const project =  computed(()=>{
    return projectStore.currentProject
})

const activeObjectiveId = ref(0)

const projectStore = useProjectStore();

const {getGoals} = useGoalController()

const handleNewActivation = (data)=>{
    //data is the objective ID
activeObjectiveId.value = data
}

const getProject = async ()=>{
    const projectId = route.params.id;
    await projectStore.getOneProject(projectId)
}

const sortedObjectives = computed(()=>{
    if(project.value.Objectives!=null && project.value.Objectives.length>0){
        const sorted = project.value.Objectives.sort((a,b)=>{return a.ID > b.ID ? -1 : 1})
        if(activeObjectiveId.value==0){
            activeObjectiveId.value = sorted[0].ID 
        }
        return sorted
    }else{
        return []
    }
})

const currentObjectiveGoals = computed(()=>{
    if(project.value.Objectives && project.value.Objectives.length>0){
        const obje = project.value.Objectives.find((obj)=> obj.ID==activeObjectiveId.value);
        if(obje.Goals != null){
            return obje.Goals
        }else{
            return []
        }
    }else{
        return []
    }
})

onMounted(async ()=>{
    await getProject()
})
</script>

