<template>
    <div class="flex flex-col gap-4 p-1 mt-4 h-[97vh] bg-gloom rounded-lg">
        <div 
            class="home-wrapper pl-3 py-1 rounded-md"
            :class="route.name=='Home'?'bg-[#61325f2e]':''"
            >
            <router-link to="/"> 
            <HomeIcon  
                class="h-6 w-6"
                :class="route.name=='Home'?'text-royal':'text-black'"
                />
            </router-link>
        </div>
        <div 
            v-if="route.name=='Details'"
            class="notes-wrapper pl-3 py-1 rounded-md bg-[#61325f2e] cursor-pointer"
            >
            <Squares2X2Icon  
                class="h-6 w-6 text-royal"
                />
        </div>
        <div 
            class="notes-wrapper pl-3 py-1 rounded-md"
            :class="route.name=='Notes'?'bg-[#61325f2e]':''"
            >
            <router-link to="/notes"> 
            <BookOpenIcon  
                class="h-6 w-6"
                :class="route.name=='Notes'?'text-royal':'text-black'"
                />
            </router-link>
        </div>
        <div class="theme-wrapper">
            <router-link to="/" v-if="theme=='dark'"> 
            <SunIcon  class="h-6 w-6 text-royal"/>
            </router-link>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { useRoute } from "vue-router";
import { HomeIcon, SunIcon, BookOpenIcon, Squares2X2Icon } from '@heroicons/vue/24/solid';

const theme = ref("light");
const route = useRoute();
console.log(route)
</script>