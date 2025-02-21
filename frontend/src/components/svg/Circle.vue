<template>
    <div class="grid place-items-center">
        <svg width="17" height="17">
            <circle cx="10" cy="10" r="6" :fill="fillColor" />
        </svg>
    </div>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
    fillColor: {
        type: String,
        default: '#490A47'
    }
});
</script>