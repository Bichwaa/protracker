<template>
    <div class="flex justify-between px-4 py-2">
        <div class="home-wrapper">
            <router-link to="/"> 
            <HomeIcon  class="h-6 w-6 text-black"/>
            </router-link>
        </div>
        <div class="theme-wrapper">
            <router-link to="/" v-if="theme=='dark'"> 
            <SunIcon  class="h-6 w-6 text-black"/>
            </router-link>
            <router-link to="/" v-else> 
            <MoonIcon  class="h-6 w-6 text-black"/>
            </router-link>
        </div>
    </div>
</template>

<script setup>
import {ref} from "vue";
import { HomeIcon, SunIcon, MoonIcon } from '@heroicons/vue/24/solid';

const theme =  ref("light")
</script>