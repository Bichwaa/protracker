<template>
  <div 
    class="border-4 border-t-transparent border-gray-200 rounded-full animate-spin"
    ></div>
</template>

<script setup>
defineProps({
  width: {
    type: String,
    default: '10', // Default width (e.g., 16 for 4rem)
  },
  height: {
    type: String,
    default: '16', // Default height (e.g., 16 for 4rem)
  },
});
</script>

<style scoped>
/* Additional styles can be added here if needed */
</style>
