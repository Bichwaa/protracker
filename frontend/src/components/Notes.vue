<template>
    <div class="board">
        <div class="board__path">
            <router-link to="/">
                <img src="../assets/images/home.png" alt="">
                <span>PROJECT </span>
            </router-link>
            <router-link to="/project-details">
                <span> > objective</span>
            </router-link>
            <router-link to="/project-details">
                <span> > goal</span>
            </router-link>
            <router-link to="/project-details">
                <span> > task</span>
            </router-link>
        </div>
        <div class="board__organizer">
        </div>
        <div class="board__notes">
            <div class="board__notes__list">
                <note-card v-for="note in notes" :key="note.id"/>
            </div>
            <div class="board__notes__current">
                <h3>Another boring ass header</h3>
                <p>
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
                </p>
            </div>
        </div>
    </div>
</template>

<script>
import NoteCard from './NoteCard.vue';
export default {
    data:function(){
        return{
            notes:[
                {id:1, title:"this is a note title", body:"lorem50", tags:["one", "three"]},
                {id:2, title:"this is a note title 2", body:"lorem50", tags:["two", "three"]},
                {id:3, title:"this is a tripple note title", body:"lorem50", tags:["one", "three"]},
                {id:4, title:"this is a title 4 a note", body:"lorem50", tags:["one", "three"]},
                {id:5, title:"this is a galaxy note5 title", body:"lorem50", tags:["one", "three"]},
                {id:6, title:"Six is a note title", body:"lorem50", tags:["one", "three"]},
            ]
        }
    },
    components:{
        NoteCard
    }
}
</script>

<style scoped>
    .board{
        display:flex;
        flex-direction: column;
    }

    .board__path{
        padding:1em 0;
    }

    .board__path a{
        color:#b3ddc8;
        text-decoration: none;
        text-transform: uppercase;
    }

    .board__path span{
        font-size:1.7rem;
        font-weight: 700;
    }

    .board__notes{
        display: grid;
        grid-template-columns: 60% 40%;
    }

    .board__notes__list{
        /* display:grid; */
        /* grid-template-columns: repeat(3, 1fr); */
        display:flex;
        flex-wrap:wrap;
    }

    .board__notes__current{
        border-left: solid .5px #b3ddc8;
        padding:1rem;
    }

    .board__notes__current h3{
        text-transform: uppercase;
    }
</style>