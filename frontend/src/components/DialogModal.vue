  <template>
  <div class="fixed inset-0 flex items-center justify-center z-50">
    <div class="fixed inset-0 bg-black opacity-50"></div>
    <div class="rounded-lg p-6 z-10 lg:min-w-[570px]">
    <div class="flex w-full justify-end">
      <!-- <span class="text-[#292a5e] cursor-pointer font-medium p-2" @click="closeClicked">x</span> -->
    </div>
      <!-- Modal content goes here -->
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['close-modal']) 

function closeClicked(){
  emit('close-modal')
}
</script>

<style scoped>
/* Add your custom styles here */
</style>
