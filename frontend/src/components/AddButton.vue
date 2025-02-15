<template>
    <div class="buttons" @mouseover="showLabel" @mouseout="hideLabel">
       <div class="label">{{label||"Add"}}</div> <button class="buttons__add"></button> 
    </div>
</template>

<script>
export default {
    props:["label"],
    methods:{
        showLabel(){
            document.querySelector(".buttons .label").classList.add("grow")
        },
        hideLabel(){
            document.querySelector(".buttons .label").classList.remove("grow")
        }
    }
}
</script>

<style scoped>
    .buttons{
        position: relative;
        z-index: 1;
        float:right;
        right:1rem;
        margin: 1rem 1rem;
        padding:0;
    }
    .buttons__add{
    width: 2.5rem;
    height: 2.5rem;
    color: #1a1a1a;
    background-color:#2dbd75b0;
    border-radius: 50%;
    border: none;
    cursor:pointer;
    transition: all 1s;
    }

    .buttons__add:hover{
      background-color: #2dbd7560;
    }

    .buttons__add:after{
        content: "+";
        font-size: 2.5em;
        line-height: 1.1em;
        color:white;
    }

    .label{
        position: relative;
        bottom:.5rem;
        display: inline;
        background-color:#212528;
        padding: .4rem;
        border-radius:5px;
        overflow: hidden;
        color:#212528;
        transition: all 2s;
    }

    .grow{
        width:auto;
        height:auto;
        background-color:rgba(248, 248, 248, 0.541);
        color:white;
    }
</style>