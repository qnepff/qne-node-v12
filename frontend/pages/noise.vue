<template>
    <div>
      <h2>Noise Protocol Handshake</h2>
      <button @click="initiateNoiseHandshake">Initiate Handshake</button>
      <ul>
        <li v-for="log in logs" :key="log">{{ log }}</li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const logs = ref([]);
  
  // Function to call the server-side API that initiates the handshake
  async function initiateNoiseHandshake() {
    try {
      // Call the server-side API that handles the Noise handshake
      const result = await $fetch('/api/noise'); // or use axios/fetch
      logs.value = result.logs;
      console.log('API Response:', result);
    } catch (error) {
      console.error('Error initiating Noise handshake:', error);
    }
  }
  </script>
  