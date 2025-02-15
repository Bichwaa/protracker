<template>
    <div class="task-card">
        <div class="task-card__notes">
            <router-link to="/notes">
                <img src="../assets/images/notes.png"><span class="note-text">0 Notes</span>
            </router-link>
        </div>
        <div class="task-card__info">
            <input type="checkbox" class="check" id=""> <span>A task named TASK</span>
        </div>

        <div class="task-card__btns">
            <a href="#">
                <img src="../assets/images/edit.png"><span class="edit">EDIT</span>
            </a>
            <a href="#">
                <img src="../assets/images/delete.png"><span class="delete">DELETE</span>
            </a>
        </div>
    </div>
</template>

<style scoped>
    .task-card{
        box-shadow: 3px 3px 8px rgba(0, 0, 0, 0.522);
        border-radius:5px;
        padding:.5rem;
    }
    .task-card div{
        margin:.7rem 0;
    }

    .task-card__notes a{
        display:flex;
        align-items: center;
        color:#b3ddc8;
        text-decoration: none;
    }

    .task-card__btns{
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .task-card__btns a{
        text-decoration: none;
        font-size:.8rem;
        display:flex;
        align-items: center;
        /* margin:0 1rem; */
    }

    .edit{
        color:rgba(218, 165, 32, 0.5);
        font-weight:700;
        margin:0 .5rem;
        transition: all 1s;
        
    }
    .edit:hover{
        color:rgba(218, 165, 32, 0.9);
        padding:.5px;
        background-color:rgba(218, 165, 32, 0.1);
    }

    .delete{
        color: rgba(208, 68, 36, 0.5);
        font-weight:700;
        margin:0 .5rem;
        transition: all 1s;
    }

    .delete:hover{
        color: rgba(208, 68, 36, 0.9);
        padding:.5px;
        background-color: rgba(208, 68, 36, 0.1);
    }
</style>