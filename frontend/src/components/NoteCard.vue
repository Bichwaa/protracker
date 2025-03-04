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
            TAGS: &nbsp;
            <a href="#">duh</a>
            <a href="#">uh!</a>
        </div>
    </div>
</template>

<script setup>

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

</script>

<style scoped>
.active{
    background-color: rgba(73, 10, 71, 0.29);
}
</style>
