<template>
  <div class="flex min-h-screen items-center justify-center">
    <div class="text-center">
      <h1 class="text-4xl font-bold mb-4">404 - Page Not Found</h1>
      <p class="mb-4">The page you are looking for does not exist.</p>
      <NuxtLink to="/" class="text-blue-500 hover:text-blue-700">
        Return to Home
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
// Set the response status code to 404
if (process.server) {
  setResponseStatus(404)
}
</script> 