<template>
  <header class="fixed top-0 left-0 right-0 bg-white shadow-md z-50 px-4">
    <div class="flex justify-between items-center h-full max-w-7xl mx-auto">
      <div class="flex items-center gap-4">
        <NuxtLink to="/" class="text-lg text-gray-700 hover:text-blue-700 transition-colors">
          <Icon icon="fa-solid:home" />
        </NuxtLink>
        <TopMenu />
        <QuickButton />
      </div>
      <div class="flex items-center gap-4">
        <button @click="openHelpDialog" class="text-lg text-gray-700 hover:text-blue-700 transition-colors">
          <Icon icon="fa-solid:question" />
        </button>
        <button class="text-lg text-gray-700 hover:text-blue-700 transition-colors">
          <Icon icon="fa-solid:sign-out-alt" />
        </button>
      </div>
    </div>

    <HelpDialog :show="showHelpDialog" @close="closeHelpDialog" />
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import TopMenu from './TopMenu.vue'
import QuickButton from './QuickButton.vue'
import HelpDialog from './HelpDialog.vue'

const showHelpDialog = ref(false)

const openHelpDialog = () => {
  showHelpDialog.value = true
}

const closeHelpDialog = () => {
  showHelpDialog.value = false
}
</script>

<style scoped>
header {
  height: 48px; /* Reduced from 64px to 48px */
}

.menu-item {
  text-transform: capitalize;
}
</style>
