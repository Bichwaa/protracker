<template>
  <div class="grid grid-cols-[90px_repeat(11,_1fr)]">
    <div class="col-span-1 mx-4">
      <AppBar/>
    </div>
    <div class="col-span-11 h-[100vh] overflow-scroll">
      <router-view/>
    </div>
  </div>
</template>

<script setup>
import AppBar from './components/AppBar.vue';
</script>

