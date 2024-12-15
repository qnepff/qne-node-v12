<template>
  <ul :class="{ 'submenu': isSubmenu }">
    <li v-for="item in items" :key="item.path" class="menu-item">
      <div class="menu-item-content">
        <NuxtLink 
          v-if="!item.children || !item.children.length" 
          :to="item.path" 
          @click="$emit('closeMenu')"
          class="menu-link"
        >
          {{ item.name }}
        </NuxtLink>
        <div 
          v-else 
          @mouseenter="openSubMenu(item)"
          @mouseleave="closeSubMenu(item)"
          class="submenu-wrapper"
        >
          <button class="submenu-toggle">
            {{ item.name }}
            <span class="toggle-icon">▶</span>
          </button>
          <MenuItems 
            v-if="item.isOpen"
            :items="item.children" 
            :isSubmenu="true"
            @closeMenu="$emit('closeMenu')"
          />
        </div>
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
interface MenuItem {
  name: string;
  path: string;
  children?: MenuItem[];
  isOpen?: boolean;
}

const props = defineProps<{
  items: MenuItem[];
  isSubmenu?: boolean;
}>()

const emit = defineEmits(['closeMenu'])

const openSubMenu = (item: MenuItem) => {
  item.isOpen = true;
}

const closeSubMenu = (item: MenuItem) => {
  item.isOpen = false;
}
</script>

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.menu-item {
  margin-bottom: 0.25rem;
  border-bottom: 1px solid #adb5bd; /* Darker border */
}

.menu-item:last-child {
  border-bottom: none;
}

.menu-item-content {
  display: flex;
  align-items: center;
}

.menu-link, .submenu-toggle {
  text-decoration: none;
  color: #495057;
  display: block;
  padding: 0.75rem 1rem;
  transition: background-color 0.3s ease, color 0.3s ease;
  width: 100%;
  text-align: left;
  background: none;
  border: none;
  font-size: inherit;
  cursor: pointer;
}

.menu-link:hover, .submenu-toggle:hover {
  background-color: #e9ecef;
  color: #0056b3;
}

.submenu-wrapper {
  position: relative;
  width: 100%;
}

.submenu {
  position: absolute;
  top: 0;
  left: 100%;
  background-color: #f8f9fa;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  border: 2px solid #6c757d; /* Darker border */
  border-radius: 4px;
  min-width: 180px;
  z-index: 1000;
  padding: 0.5rem 0;
}

.toggle-icon {
  float: right;
  font-size: 0.8em;
  margin-left: 0.5rem;
}
</style>
