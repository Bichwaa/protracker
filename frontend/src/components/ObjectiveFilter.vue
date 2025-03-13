<template>
    <div class="bg-gloom rounded-lg p-2 my-2">
        <div class="flex justify-between items-center" @click="toggleSettings">
            <FunnelIcon class="h-6 w-6 cursor-pointer"/>
            <span class="cursor-pointer">Sort
                 Objectives</span>
            <BarsArrowDownIcon class="h-6 w-6 cursor-pointer"/>
        </div>
        <div 
            class="settings grid grid-cols-4 my-2"
            :class="['settings', { 'expanded': isExpanded }]" :style="{ maxHeight: isExpanded ? '200px' : '0' }"
            >
            <div class="col-span-2 flex flex-col gap-2">
                <label class="flex items-center gap-2" @click="sortByCreated">
                    <input type="radio" name="filter" value="all" class="form-radio">
                    <span class="text-sm">Created</span>
                </label>
                <label class="flex items-center gap-2" @click="sortByUpdated">
                    <input type="radio" name="filter" value="active" class="form-radio">
                    <span class="text-sm">Updated</span>
                </label>
                <label class="flex items-center gap-2" @click="sortByProgress">
                    <input type="radio" name="filter" value="active" class="form-radio">
                    <span class="text-sm">Progress</span>
                </label>
            </div>

            <!-- <div class="flex flex-col gap-2">
                <label class="flex items-center gap-2">
                    <input type="radio" name="filter" value="all" class="form-radio">
                    <span class="text-sm">Ascending</span>
                </label>
                <label class="flex items-center gap-2">
                    <input type="radio" name="filter" value="active" class="form-radio">
                    <span class="text-sm">Descending</span>
                </label>
            </div> -->
        </div>
    </div>
</template>

<script setup>
import { FunnelIcon, BarsArrowDownIcon } from '@heroicons/vue/24/solid';
import { ref } from 'vue';

const emit = defineEmits([
    'sort-by-created',
    'sort-by-updated',
    'sort-by-progress'
])

const isExpanded = ref(false);

function toggleSettings() {
    isExpanded.value = !isExpanded.value;
}

function sortByCreated(){
    emit('sort-by-created')
}

function sortByUpdated(){
    emit('sort-by-updated')
}

function sortByProgress(){
    emit('sort-by-progress')
}
</script>

<style scoped>
.settings {
    transition: max-height 0.3s ease, opacity 0.3s ease;
    max-height: 0;
    overflow: hidden;
    opacity: 0;
}
.settings.expanded {
    max-height: 200px; /* Adjust this value based on your content */
    opacity: 1;
}
</style>