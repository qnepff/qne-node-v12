<template>
  <div v-if="show" class="fixed inset-0 bg-transparent flex justify-center items-center z-50" @click="$emit('close')">
    <div class="bg-green-800 text-white p-8 rounded-lg shadow-lg max-w-lg w-full" @click.stop>
      <h2 class="text-2xl font-bold mb-4">Help: Quick Button Usage</h2>
      <div class="mb-4 text-lg leading-relaxed">
        <p>The Quick Button (bolt icon) provides quick access to frequently used features:</p>
        <ul class="list-disc pl-6">
          <li><strong class="text-green-300">Open/Close Menu:</strong> Click the bolt icon to toggle the quick menu.</li>
          <li><strong class="text-green-300">Access Features:</strong> Click on a menu item to open its corresponding dialog.</li>
          <li><strong class="text-green-300">Shopping:</strong> Manage your shopping lists.</li>
          <li><strong class="text-green-300">Todo:</strong> Keep track of your tasks.</li>
          <li><strong class="text-green-300">Event:</strong> Plan and manage your events.</li>
        </ul>
        <p>To close an open dialog, click the Quick Button again.</p>
      </div>
      <p class="text-center italic text-green-300 mt-4 font-bold">Click anywhere to close this help dialog</p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  show: boolean
}>()

defineEmits<{
  (e: 'close'): void
}>()
</script>
