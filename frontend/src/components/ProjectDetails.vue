<template>
    <div class="board grid grid-cols-10 p-4">
        <div class="board__objectives col-span-3 flex flex-col gap-2">
            <objective-title-card 
                :projectId="project.ID"
                @objective-created="getProject"
            />
            <objective-card 
                v-for="obj in project.Objectives" 
                :objective="obj"  
                :key="obj.ID"
                @objective-edited="getProject"
                @objective-deleted="getProject"
            />
        </div>
        <div class="board__goals col-span-7 px-4 flex flex-col gap-2">
            <details-summary-card :project="project"/>
            <div class="grid grid-cols-2 gap-2">
                <goal-card v-for="i in 4" />
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

const route = useRoute();

const objectives = ref([
    {name:"kiruu", description:"also kiruuuu", progress:40},
    {name:"objective1", description:"description for objective1", progress:20},
    {name:"objective2", description:"description for objective2", progress:50},
    {name:"objective3", description:"description for objective3", progress:80},
    {name:"objective4", description:"description for objective4", progress:100},
]);

const project =  computed(()=>{
    return projectStore.currentProject
})


const projectStore = useProjectStore();

const getProject = async ()=>{
    const projectId = route.params.id;
    await projectStore.getOneProject(projectId)
}


const progress = ref(60);

onMounted(async ()=>{
    await getProject()
})
</script>
