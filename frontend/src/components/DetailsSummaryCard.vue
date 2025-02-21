<template>
    <div class="flex flex-col py-4 px-8 bg-gloom rounded-lg">
        <div class="flex mb-3">
            <h2 class="capitalize text-black font-black text-3xl">{{ project.Name }}</h2>
        </div>
        <div class="grid grid-cols-5">
            <div class="col-span-4 flex gap-2">
                <div class="flex flex-col justify-center">
                    <span class="font-semibold my-1">Tracking objectives</span>
                    <div class="flex flex-col gap-2 h-36 overflow-scroll hide-scrollbar">
                        <div class="flex items-center gap-2" v-for="o in objectives">
                            <Circle/>
                            <span>{{ o.Name }}: {{ o.Progress }}%</span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-span-1 flex flex-col justify-between">
                <div class="flex flex-col">
                    <span class="font-semibold text-5xl text-black text-right">{{avgProgress}}%</span>
                    <span class="capitalize text-black text-lg font-medium text-right">Completed</span>
                </div>
            </div>
        </div>
        <div class="mt-5 flex justify-between">
            <span class="font-medium">Overseeer: {{ project.Overseer }}</span>
            <div class="text-right font-medium text-lg">{{ grade }}</div>
        </div>
    </div>
</template>

<script>
const emptyDate = new Date()
</script>

<script setup>
import {computed} from 'vue';
import Circle from './svg/Circle.vue';

const props = defineProps({
    project:{
        type:Object,
        default:{
            ID:0,
            CreatedAt: emptyDate,
            UpdatedAt: emptyDate,
            DeletedAt:emptyDate,
            Name:"null",
            Tags:"null",
            BeginDate:emptyDate,
            EstimatedEndDate:emptyDate,
            Progress:0,
            Description:"null",
            Objectives:[],
            Notes:[]
}
    }
})

const objectives = computed(()=>{
    if(props.project.Objectives!=null){
     return props.project.Objectives
    }else{
        return []
    }
})

const avgProgress = computed(()=>{
    if(props.project.Objectives==null){
     return 0
    }else if(props.project.Objectives.length==0){
        return 0
    }else{
        const total =  props.project.Objectives.reduce((acc, obj) => {
            return acc + obj.Progress;
        }, 0);
        console.log("ttal", total)
        return total/props.project.Objectives.length
    }
})

const grade = computed(() => {
    const progress = avgProgress.value;
    if (progress >= 90) {
        return 'A+';
    } else if (progress >= 80) {
        return 'A';
    } else if (progress >= 70) {
        return 'B';
    } else if (progress >= 60) {
        return 'C';
    } else if (progress >= 50) {
        return 'D';
    } else {
        return 'F';
    }
})
</script>

<style scoped>
   .hide-scrollbar::-webkit-scrollbar {
       display: none; /* Safari and Chrome */
   }

   .hide-scrollbar {
       -ms-overflow-style: none;  /* IE and Edge */
       scrollbar-width: none;  /* Firefox */
   }
</style>