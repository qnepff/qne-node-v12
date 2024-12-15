<template>
  <div class="quick-button-container" ref="containerRef">
    <button @click="toggleDropdown" class="icon-button">
      <Icon icon="fa-solid:bolt" />
    </button>
    <div v-if="isOpen" class="quick-dropdown">
      <button 
        v-for="item in quickLinks" 
        :key="item.name"
        @click="openDialog(item.name)"
        class="quick-link"
      >
        {{ item.name }}
      </button>
    </div>

    <!-- Dialogs -->
    <div v-if="activeDialog" class="dialog-overlay" @click="closeDialog">
      <div class="dialog" @click.stop>
        <component :is="activeDialog" @close="closeDialog"></component>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import WeightDialog from './WeightDialog.vue'

const isOpen = ref(false)
const activeDialog = ref(null)
const containerRef = ref(null)

const quickLinks = [
  { name: 'Shopping', component: 'ShoppingDialog' },
  { name: 'Todo', component: 'TodoDialog' },
  { name: 'Event', component: 'EventDialog' },
  { name: 'Weight', component: WeightDialog },
]

const toggleDropdown = () => {
  if (activeDialog.value) {
    activeDialog.value = null
  } else {
    isOpen.value = !isOpen.value
  }
}

const openDialog = (dialogName) => {
  const dialog = quickLinks.find(link => link.name === dialogName)
  if (dialog) {
    activeDialog.value = dialog.component
    isOpen.value = false
  }
}

const closeDialog = () => {
  activeDialog.value = null
}

const handleClickOutside = (event) => {
  if (containerRef.value && !containerRef.value.contains(event.target)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.quick-button-container {
  position: relative;
}

.icon-button {
  background: none;
  border: none;
  font-size: 1.2rem;
  color: #495057;
  cursor: pointer;
  transition: color 0.3s ease;
}

.icon-button:hover {
  color: #0056b3;
}

.quick-dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #f8f9fa;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1), 0 1px 3px rgba(0, 0, 0, 0.08);
  border: 2px solid #6c757d;
  border-radius: 4px;
  z-index: 1000;
  min-width: 120px;
  padding: 0.5rem 0;
  margin-top: 0.5rem;
}

.quick-link {
  display: block;
  width: 100%;
  padding: 0.5rem 1rem;
  color: #495057;
  text-decoration: none;
  background: none;
  border: none;
  text-align: left;
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.quick-link:hover {
  background-color: #e9ecef;
  color: #0056b3;
}

.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.dialog {
  background-color: white;
  padding: 1rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  max-width: 90%;
  max-height: 90%;
  overflow-y: auto;
}

.dialog h2 {
  margin-top: 0;
  margin-bottom: 1rem;
}

.close-button {
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  background-color: #0056b3;
  color: #ffffff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.close-button:hover {
  background-color: #004494;
}
</style>
