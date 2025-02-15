<template>
    <div class="bg-gloom flex flex-col rounded-lg p-4">
        <div class="flex justify-between items-center">
            <div>
                <router-link to="/"> 
                    <UserCircleIcon  class="h-16 w-16 text-black"/>
                </router-link>
            </div>

            <button class="bg-royal text-white hover:bg-black transition duration-300 rounded-lg px-2 py-0 h-8" @click="addProjectClicked">+ Add Project</button>
        </div>
        <div class="info-wrapper flex flex-col gap-2">
            <span class="font-medium text-black">{{ "Bichwaa,"  }}</span>
            <span class="font-medium text-black">Last log date: {{ "14/05/2023" }}</span>
        </div>

        <div class="quote-wrapper flex flex-col py-2">
            <span class="font-medium text-black">{{ "'Rent will not pay itsef sir'. - A wise man"  }}</span>
        </div>
    </div>
</template>

<script setup>
import {ref} from "vue";
import { UserCircleIcon } from '@heroicons/vue/24/outline';

const emit = defineEmits(['add-project-clicked'])
const theme =  ref("light")

const addProjectClicked = ()=>{
    emit("add-project-clicked");
}
</script>