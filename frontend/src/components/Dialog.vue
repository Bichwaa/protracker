<template>
    <div class="">
        <dialog ref="confirmationDialog" open class=" border-[#490A47] rounded-md drop-shadow-lg bg-[#e2e2e2] z-50">
            <div class="flex w-full flex-row-reverse px-6 pt-3">
                <span class="text-2xl cursor-pointer hover:text-red-500 duration-300" @click="deleteCancelled">&times;</span>
            </div>
            <form class="py-8 px-12">
                <p class="text-center font-medium">Delete this {{entity}}?</p>
                <div class="flex gap-12 justify-between my-4">
                <button value="cancel" formmethod="dialog" class="px-3 py-1 bg-gray-500 rounded-md text-white text-sm font-medium" @click.prevent="deleteCancelled">Cancel</button>
                <button id="confirmBtn" value="default" class="px-3 py- rounded-md hover:bg-red-500 hover:text-white duration-300 text-red-500 text-sm font-medium" @click.prevent="deleted">Delete</button>
                </div>
            </form>
        </dialog>
    </div>
    <output ref="outputBox"></output>
</template>

<script setup>
    import {ref} from 'vue';

    const props = defineProps({
        entity:{
            type: String,
            default:"project"
        }
    })

    const confirmationDialog = ref(null)
    const emits = defineEmits(['cancelled',"deleted"])

    const deleteCancelled= ()=>{
        emits("cancelled")
    }

    const deleted= ()=>{
        emits("deleted")
    }

</script>