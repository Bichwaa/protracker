<template>
    <div class="task-card rounded-sm px-1">
        <div class="flex items-center justify-between">
            <div class="task-card__info flex items-center gap-1">
                <input 
                    type="checkbox" 
                    class="accent-royal" 
                    @change="handleTaskChecked"
                    :checked="task.Progress==100?true:false"
                > 
                <span class="text-sm cursor-pointer" @dblclick="updateTaskClicked">{{ task.Name }}</span>
            </div>
            <span class="text-xl cursor-pointer text-red-300 hover:text-red-500 duration-300" @click="deleteSwitch">&times;</span>
        </div>
        <div>
            <TaskForm
            v-if="TaskFormOn"
                :task="task"
                :edit="true"
                @task-updated="handleTaskUpdated"
                @modal-off="toggleTaskFormOff"
                 />
        </div>

        <div class="relative">
            <Dialog 
                v-if="showDelete"
                entity="Task"  
                @cancelled="showDelete=false"
                @deleted="deleteTask"
            />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import TaskForm from './TaskForm.vue';
import Dialog from "./Dialog.vue";
import { useTaskController } from '../APIs/task-controller.js';

const props = defineProps({
    task:{
        type: Object,
        default:{
            ID:0,
            Name:"A task named TASK",
            EstimatedEndDate: new Date(),
            Progress:100
        }
    }
});

const emit = defineEmits(["refresh"])

const TaskFormOn = ref(false)
const showDelete = ref(false)

const deleteSwitch = ()=>{
        showDelete.value = true;
    }

    const toggleTaskFormOn = ()=>{
        TaskFormOn.value=true
    }


    const toggleTaskFormOff = ()=>{
        console.log('task forming offf')
        TaskFormOn.value=false
    }

    const updateTaskClicked = ()=>{
        toggleTaskFormOn()
    }


    const handleTaskUpdated = ()=>{
        toggleTaskFormOff()
        emit("refresh")
    }

    const handleTaskChecked = async ()=>{
        console.log("thing got checked")
        const {markTaskAsCompleted} = useTaskController();
        try {
            await markTaskAsCompleted(props.task.ID);
            showDelete.value = false;
            emit('refresh')
        }catch (e){
            throw e
        }
    }

    const deleteTask = async ()=>{
        const {deleteTask} = useTaskController();
        try {
            await deleteTask(props.task.ID);
            showDelete.value = false;
            emit('refresh')
        }catch (e){
            throw e
        }
    }

</script>
