<template>
  <div class="top-menu" ref="menuRef">
    <button @click="toggleMenu" class="menu-toggle">
      <Icon icon="fa-solid:bars" />
    </button>
    <div v-if="isOpen" class="menu-dropdown">
      <nav>
        <MenuItems :items="menuItems" :isSubmenu="false" @closeMenu="closeMenu" />
      </nav>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import MenuItems from './MenuItems.vue'
import { useMenu } from '~/composables/useMenu'

const { menuItems } = useMenu()
const isOpen = ref(false)
const menuRef = ref(null)

const toggleMenu = () => {
  isOpen.value = !isOpen.value
}

const closeMenu = () => {
  isOpen.value = false
}

const handleClickOutside = (event) => {
  if (menuRef.value && !menuRef.value.contains(event.target)) {
    closeMenu()
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
.top-menu {
  position: relative;
}

.menu-toggle {
  background: none;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  padding: 5px;
  color: #495057;
  transition: color 0.3s ease;
}

.menu-toggle:hover {
  color: #0056b3;
}

.menu-dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #f8f9fa;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1), 0 1px 3px rgba(0, 0, 0, 0.08);
  border: 2px solid #6c757d;
  border-radius: 4px;
  z-index: 1000;
  min-width: 180px;
  padding: 0.5rem 0;
  margin-top: 0.5rem;
}

nav {
  padding: 0.5rem;
}
</style>
