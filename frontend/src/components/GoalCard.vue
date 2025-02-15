<template>
    <div class="goal-card">
        <div class="goal-card__btns">
            <button class="btns__edit">EDIT</button>
            <button class="btns__delete">DELETE</button>
        </div>
        <h3>GOAL NAME</h3>
        <div class="goal-card__notes">
            <router-link to="/notes">
                <img src="../assets/images/notes.png"><span class="note-text">0 Notes</span>
            </router-link>
        </div>
        <div class="goal-card__desc">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. 
            Labore impedit aliquid suscipit voluptatem, vitae adipisci excepturi 
            velit corrupti voluptates temporibus.
        </div>
        <div class="goal-card__tasks">
            <task-card/>
        </div>
    </div>
</template>

<script>
import TaskCard from './TaskCard.vue';
export default {
    components:{
        TaskCard
    }
}
</script>

<style scoped>
    .goal-card{
        background: #2b2f33;
        margin: 1rem 1rem; 
        border-radius: 15px;
        padding:1rem;
    }

    .goal-card > div{
        margin: 1rem 0;
    }
    .goal-card__btns{
        display:flex;
        justify-content: space-between;
    }

    .goal-card__btns button{
        border:none;
        padding:.5rem;
        height:1.8rem;
        width:3.5rem;
        border-radius:8px;
        text-align:center;
        cursor:pointer;
    }

    .goal-card__notes a{
        display:flex;
        align-items: center;
        text-decoration:none;
        color: #b3ddc8;
    }

    .btns__edit{
        color:rgba(255, 255, 255, 0.637);
        background-color:rgba(218, 165, 32, 0.2);
        transition: all .8s;
    }

    .btns__edit:hover{
        color:rgba(218, 165, 32);
        background-color:rgba(218, 165, 32, 0.1);
    }

    .btns__delete{
        color:rgba(255, 255, 255, 0.637);
        background-color: rgba(208, 68, 36, 0.2);
        transition: all .8s;
    }

    .btns__delete:hover{
        color:rgba(208, 68, 36);
        background-color: rgba(208, 68, 36, 0.1);
    }
</style>