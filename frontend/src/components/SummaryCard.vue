<template>
    <div class="flex flex-col p-4 bg-gloom rounded-lg">
        <div class="flex">
            <h2 class="capitalize text-black">Welcome,</h2>
        </div>
        <div class="grid grid-cols-5 place-items-center">
            <div class="col-span-4 flex gap-12">
                <div class="flex flex-col justify-center">
                    <span class="font-semibold text-3xl">{{ projects.length || 0 }}</span>
                    <span class="font-semibold text-xl">Project{{ projects.length==1?"":"s" }}</span>
                </div>
                <div>
                    <Pie :data="chartData" :options="chartOptions" class=""/>
                </div>
            </div>
            <div class="col-span-1 flex flex-col">
                <span class="capitalize text-black">average</span>
                <span class="capitalize text-black">progress</span>
                <span class="font-semibold text-4xl text-black">{{ avgProjProgress }}%</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import {ref, computed} from 'vue';
import { Pie, } from 'vue-chartjs';
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, ArcElement } from 'chart.js'

const props = defineProps(["projects"])

const projProgressArr = computed(()=>{
    if(props.projects){
        const counts = [0, 0, 0, 0, 0]; // Initialize counts for each grade
        props.projects.forEach(x => {
            if (x.Progress === 0) counts[0]++;
            else if (x.Progress < 60) counts[1]++;
            else if (x.Progress < 80) counts[2]++;
            else if (x.Progress < 100) counts[3]++;
            else counts[4]++;
        });
        return counts;
    }else{
        return [0,0,0,0,0]
    }
})

const avgProjProgress = computed(()=>{
    if(props.projects && props.projects.length > 0){
        const totalProgress = props.projects.reduce((acc, project) => acc + project.Progress, 0);
        return (totalProgress / props.projects.length).toFixed(); // Returns average rounded to 2 decimal places
    } else {
        return 0; // Return 0 if there are no projects
    }
})

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, ArcElement)

const chartData = ref({
        labels: [ 'No progress', '< 60% done', '> 60% done', '> 80% done', 'Finished' ],
        datasets: [{
        label: 'Perfomance',
        data: projProgressArr,
        backgroundColor: [
        '#FAFAFA',
        '#B8B5B5',
        '#917290',
        '#6D3E6C',
        '#490A47',
        ],
        borderColor:'transparent'
    }]
})

const chartOptions = {
    circumference:360,
    radius:"100%",
    aspectRatio:"2",
    layout: {
        padding: {
            top: 0,
            bottom:0,
            left:0,
            right:0
        }
    },
    plugins:{
        legend:{
            position:"left",
            strokeStyle:'red'
        }
    },
}
</script>