<template>
  <div>
    <Header />
    <main class="pt-12"> <!-- Add padding-top here -->
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import Header from '~/components/Header.vue'
</script>
