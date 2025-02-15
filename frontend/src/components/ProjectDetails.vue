<template>
    <router-link to="/" class="a">
    <img src="../assets/images/home.png" alt="">
    <span>HOME</span>
</router-link>
    <div class="board">
        <div class="board__objectives">
            <add-button class="btn-align" label="Add Objective"/>
            <objective-card v-for="obj in objectives" objective="obj"  :key="obj.name"/>
        </div>
        <div class="board__goals">
            <goal-card/>
        </div>
    </div>
</template>

<script>
import AddButton from './AddButton.vue';
import ObjectiveCard from './ObjectiveCard.vue';
import GoalCard from './GoalCard.vue';

export default {
    components:{
        AddButton,
        ObjectiveCard,
        GoalCard
    },
    data: function(){
        return{
            objectives : [{name:"kiruu", description:"also kiruuuu", progress:40}],
            progress: 60
        }
    }
}
</script>

<style scoped>
    a{
        display:flex;
        flex-direction: row-reverse;
        padding:.5rem 2rem;
        text-decoration: none;
        align-items: center;
        color:#b3ddc8;
        margin-left:1.5rem;
    }

    .board{
        display: flex;
    }

    .btn-container{
        display:block;
    }
    .btn-align{
        margin-left:2rem;
    }

    .board__objectives{
        flex-basis: 25%;
        display: flex;
        flex-direction: column;
        align-content: center;
        /* align-items: center; */
        min-height:95vh;
        border-right: solid .5px rgba(181, 201, 191, .05);
        border-radius: 15px;
    }

    .board__goals{
        display:grid;
        grid-template-columns: 1fr 1fr;
        width:100%;
    }
</style>