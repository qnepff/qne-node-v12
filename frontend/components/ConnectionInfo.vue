<template>
  <div class="fixed top-2 right-2 bg-gray-800 text-white px-3 py-1 rounded-full text-sm z-50">
    Protocol: {{ protocol }}
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const protocol = ref('Loading...')

onMounted(async () => {
  try {
    // First try direct request to current origin
    const response = await fetch('/api/protocol', {
      headers: {
        'Accept': '*/*',
        'Cache-Control': 'no-cache'
      }
    })
    console.log('Protocol response headers:', response.headers)
    const proto = response.headers.get('X-Protocol')
    if (proto) {
      protocol.value = proto
    } else {
      // Fallback to checking if HTTP/3 is supported
      const httpVersion = response.headers.get('alt-svc')
      protocol.value = httpVersion?.includes('h3') ? 'HTTP/3' : 'HTTP/2'
    }
  } catch (error) {
    protocol.value = 'Error: ' + error.message
    console.error('Error checking protocol:', error)
  }
})
</script>
