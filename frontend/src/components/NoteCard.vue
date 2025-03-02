<template>
    <div 
        class="note-card rounded-lg p-4 mb-2"
        :class="activeId==note.ID?'active': 'bg-gloom'"
        >
        <h3 class="note-card__header font-medium cursor-pointer" @click="activate">
            Lorem ipsum dolor sit amet.
        </h3>
        <div class="note-card__body line-clamp-3 cursor-pointer" @click="activate">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. 
            Provident, assumenda temporibus incidunt adipisci reiciendis iure sit esse. 
            Quasi voluptates excepturi quod quo pariatur repellat possimus qui voluptatum 
            laborum aliquid quaerat maiores aperiam explicabo enim voluptatibus, obcaecati 
            cupiditate! Consequuntur, eligendi amet reprehenderit id corporis et totam quam
            earum tempore fugiat, quia aliquam asperiores sequi. Perferendis quae hic 
            voluptate iste asperiores ratione architecto. Similique quisquam ad voluptatum
            quasi, nihil quod eius consectetur placeat eum nulla aliquid rerum quidem est. 
            Fugiat ipsum dolore tenetur voluptatum earum! Minima natus corporis maiores
            tenetur temporibus ipsa doloribus nesciunt officiis! Rerum ullam hic vel, 
            perspiciatis vitae ab ratione, vero illo atque odit veritatis cumque quae 
            dolorum voluptas ea nihil eveniet exercitationem at dolorem maxime facere a? 
            Dolore repudiandae modi, temporibus animi quisquam assumenda reprehenderit. 
            Laborum eum eius temporibus quaerat voluptas facere, ullam sed optio quos ab 
            error incidunt harum officia animi aliquid sit. Amet omnis, fugit nihil 
            blanditiis earum maxime! Suscipit sunt fuga a. Architecto placeat ea 
            necessitatibus numquam excepturi, facilis dolores, incidunt officiis ab 
            doloribus fugiat sed nulla vel itaque! Ratione nam, quas quibusdam magnam.
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
