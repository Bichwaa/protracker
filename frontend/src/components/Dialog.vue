<template>
    <div class="">
        <dialog ref="confirmationDialog" open class="px-12 py-8 border border-1 border-[#490A47] rounded-md">
            <form>
                <p class="text-center">Delete this project?</p>
                <div class="flex gap-12 justify-between my-4">
                <button value="cancel" formmethod="dialog" class="p-3 bg-gray-500 rounded-md text-white font-medium" @click.prevent="deleteCancelled">Cancel</button>
                <button id="confirmBtn" value="default" class="p-3 bg-red-500 rounded-md text-white font-medium" @click.prevent="deleted">Delete</button>
                </div>
            </form>
        </dialog>
    </div>
    <output ref="outputBox"></output>
</template>

<script setup>
    import {ref} from 'vue';
    const confirmationDialog = ref(null)
    const emits = defineEmits(['cancelled',"delete"])

    const deleteCancelled= ()=>{
        emits("cancelled")
    }

    const deleted= ()=>{
        emits("deleted")
    }

</script>