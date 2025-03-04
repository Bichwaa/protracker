<template>
    <div 
        class="note-card rounded-lg p-4 mb-2"
        :class="activeId==note.ID?'active': 'bg-gloom'"
        >
        <h3 class="note-card__header font-medium cursor-pointer" @click="activate">
            {{ note.Title }}
        </h3>
        <div class="note-card__body line-clamp-3 cursor-pointer" @click="activate" v-html="note.Content">
        </div>
        <div class="note-card__tags">
            TAGS:
            <span> 
                <a class="text-sm bg-gray-200 mx-1 rounded-xl px-2 py-1 cursor-pointer" v-for="tag in tagList">{{tag}}</a>
            </span>
        </div>
    </div>
</template>

<script setup>

import { computed } from 'vue'

const props = defineProps({
    note:{
        type: Object,
        default: {}
    },
    activeId:{
        type:Number,
        default:0
    }
})

const emit = defineEmits(['new-activation'])

const activate = ()=>{
 emit('new-activation', props.note)
}

const tagList = computed(()=>{
        const tagString = props.note.Tags;
        if(tagString && tagString.length>0){
            return tagString.split(',')
        }else {
            return []
        }
    })

</script>

<style scoped>
.active{
    background-color: rgba(73, 10, 71, 0.29);
}
</style>
