<template>
    <div class="task-card rounded-sm px-1">
        <div class="flex items-center justify-between">
            <div class="task-card__info flex items-center gap-1">
                <input type="checkbox" class="accent-royal"> 
                <span class="text-sm">A task named TASK</span>
            </div>
            <span class="text-xl cursor-pointer text-red-300 hover:text-red-500 duration-300" @click="()=>true">&times;</span>
        </div>
        
    </div>
</template>
