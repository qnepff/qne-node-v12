// This file defines a Nuxt plugin for file system operations

// Import the defineNuxtPlugin function from Nuxt
import { defineNuxtPlugin } from 'nuxt/app'

export default defineNuxtPlugin((nuxtApp) => {
  const $fetch = nuxtApp.$fetch

  const readFile = (fileName) => $fetch('/api/quick-n-easy', {
    method: 'POST',
    body: { action: 'read', fileName }
  })

  const writeFile = (fileName, content) => $fetch('/api/quick-n-easy', {
    method: 'POST',
    body: { action: 'write', fileName, content }
  })

  const deleteFile = (fileName) => $fetch('/api/quick-n-easy', {
    method: 'POST',
    body: { action: 'delete', fileName }
  })

  const listFiles = () => $fetch('/api/quick-n-easy', {
    method: 'POST',
    body: { action: 'list' }
  })

  return {
    provide: {
      quickNEasy: {
        readFile,
        writeFile,
        deleteFile,
        listFiles
      }
    }
  }
})