<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold text-center mb-8">WebRTC Video Chat</h1>
    
    <div class="flex flex-col md:flex-row justify-center gap-6 mb-8">
      <div class="flex-1">
        <video
          ref="localVideoRef"
          :srcObject="localStream"
          class="w-full aspect-video bg-gray-800 rounded-lg"
          autoplay
          playsinline
          muted
        />
        <p class="text-center mt-2 text-gray-600">Local Video</p>
      </div>
      <div class="flex-1">
        <video
          ref="remoteVideoRef"
          :srcObject="remoteStream"
          class="w-full aspect-video bg-gray-800 rounded-lg"
          autoplay
          playsinline
        />
        <p class="text-center mt-2 text-gray-600">Remote Video</p>
      </div>
    </div>

    <div v-if="error" class="mb-6 p-4 bg-red-100 text-red-700 rounded-lg text-center">
      {{ error }}
    </div>

    <div class="flex justify-center gap-4">
      <button
        @click="init"
        :disabled="isCallEnabled"
        class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        Start Camera
      </button>
      <button
        @click="startCall"
        :disabled="!isCallEnabled"
        class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        Start Call
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount, watch } from 'vue'
import { useWebRTC } from '~/composables/communication/useWebRTC.ts'

const { localStream, remoteStream, error, isCallEnabled, init, startCall, cleanup } = useWebRTC()

// Create refs for the video elements
const localVideoRef = ref<HTMLVideoElement | null>(null)
const remoteVideoRef = ref<HTMLVideoElement | null>(null)

// Watch for stream changes and update video elements
watch(localStream, (stream) => {
  if (localVideoRef.value && stream) {
    localVideoRef.value.srcObject = stream
  }
})

watch(remoteStream, (stream) => {
  if (remoteVideoRef.value && stream) {
    remoteVideoRef.value.srcObject = stream
  }
})

// Clean up resources when component is destroyed
onBeforeUnmount(() => {
  cleanup()
})
</script>
