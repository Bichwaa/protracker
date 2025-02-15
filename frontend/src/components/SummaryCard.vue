<template>
    <div class="flex flex-col p-4 bg-gloom rounded-lg">
        <div class="flex">
            <h2 class="capitalize text-black">Welcome,</h2>
        </div>
        <div class="grid grid-cols-5 place-items-center">
            <div class="col-span-4 flex gap-12">
                <div class="flex flex-col justify-center">
                    <span class="font-semibold text-3xl">12</span>
                    <span class="font-semibold text-xl">Projects</span>
                </div>
                <div>
                    <Pie :data="chartData" :options="chartOptions" class=""/>
                </div>
            </div>
            <div class="col-span-1 flex flex-col">
                <span class="capitalize text-black">average</span>
                <span class="capitalize text-black">perfomance</span>
                <span class="font-semibold text-4xl text-black">83%</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import {ref} from 'vue';
import { Pie, } from 'vue-chartjs';
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, ArcElement } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, ArcElement)

const chartData = ref({
        labels: [ 'No progress', '< 60% done', '> 60% done', '> 80% done', 'Finished' ],
        datasets: [{
        label: 'Perfomance',
        data: [10,,60,32,75,90],
        backgroundColor: [
        '#FAFAFA',
        '#B8B5B5',
        '#917290',
        '#6D3E6C',
        '#490A47',
        ],
        // hoverOffset: 4
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
            position:"left"
        }
    },
}
</script>