<template>
    <div class="note-card">
        <h3 class="note-card__header">
            Lorem ipsum dolor sit amet.
        </h3>
        <div class="note-card__body">
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
        <div class="note-card__btns">
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
    .note-card{
        display: flex;
        flex-direction: column;
        width:12rem;
        max-height: 17rem;
        background-color:#2b2f33 ;
        margin:1rem;
        padding:1rem;
        border-radius:15px;
    }

    .note-card div{
        margin:.5rem 0;
    }

    .note-card__header{
        text-transform: uppercase;
    }

    .note-card__tags{
        font-size: .7rem;
        font-weight: 600;
    }

    .note-card__tags a{
        text-decoration: none;
        color:black;
        text-transform:uppercase;
        background-color:rgba(248, 248, 248, 0.541);
        margin-right:.5rem;
        padding: .2rem;
        border-radius:5px;
    }

    .note-card__body{
        overflow: hidden;
    }

    .note-card__btns{
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .note-card__btns a{
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