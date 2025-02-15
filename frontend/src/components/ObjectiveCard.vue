<template>
    <div class="objective-card">
        <h4>{{objective.name || "OBJECTIVE NAME"}}</h4>
        <div class="objective-card__notes">
            <router-link to="/notes">
                <img src="../assets/images/notes.png"><span class="note-text">0 Notes</span>
            </router-link>
        </div>
        <div class="objective-card__progress">
            <span>Progress: {{objective.progress||"0%"}}</span>
        </div>
        <div class="objective-card__descrption">
            {{objective.description}}
        </div>
        <div class="objective-card__btns">
            <button class="btns__edit">EDIT</button>
            <button class="btns__delete">DELETE</button>
        </div>
    </div>
</template>

<script>
export default {
    props:["objective"]
}
</script>

<style scoped>
  .objective-card{
        background: #2b2f33;
        margin: 1rem 1rem; 
        border-radius: 15px;
        padding:1rem;
    }

    .objective-card > div{
        margin: 1rem 0;
    }

    .objective-card__notes a{
        display:flex;
        align-items: center;
        text-decoration:none;
        color: #b3ddc8;
    }

    .objective-card__btns{
        display:flex;
        justify-content: space-between;
    }

    .objective-card__btns button{
        border:none;
        padding:.5rem;
        height:1.8rem;
        width:3.5rem;
        border-radius:8px;
        text-align:center;
        cursor:pointer;
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