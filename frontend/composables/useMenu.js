import { ref, onMounted } from 'vue'

export function useMenu() {
  const menuItems = ref([])

  onMounted(async () => {
    const pages = import.meta.glob('/pages/**/*.vue')
    const routes = []

    for (const path in pages) {
      const name = path.replace('/pages/', '').replace('.vue', '').replace(/\/index$/, '')
      const parts = name.split('/')
      let currentLevel = routes

      // Skip top-level index.vue and empty names
      if (name === '' || parts[0] === '' || parts[0] === 'index') continue

      parts.forEach((part, index) => {
        const isLast = index === parts.length - 1
        const existingItem = currentLevel.find(item => item.name === part)

        if (existingItem) {
          if (isLast) {
            existingItem.path = `/${name}`
          } else {
            currentLevel = existingItem.children || (existingItem.children = [])
          }
        } else {
          const newItem = {
            name: part,
            path: isLast ? `/${name}` : '#',
            isOpen: false
          }
          if (!isLast) {
            newItem.children = []
          }
          currentLevel.push(newItem)
          if (!isLast) {
            currentLevel = newItem.children
          }
        }
      })
    }

    menuItems.value = routes
  })

  return { menuItems }
}