<template>
  <div class="in-the-event-of">
    <h1>In the Event Of</h1>
    
    <div v-if="!isAuthenticated">
      <p>Please enter a password to view or edit your information.</p>
      <div class="password-input">
        <input 
          v-model="password" 
          type="password" 
          placeholder="Enter password"
        >
      </div>
      <button @click="authenticate" class="auth-btn">
        Access Information
      </button>
    </div>

    <div v-else>
      <p class="description">Use this page to list important information and items that will be needed by your loved ones in the event of your passing.</p>

      <div v-for="(section, index) in sections" :key="index" class="section">
        <h2>{{ section.title }}</h2>
        <div v-for="(item, itemIndex) in section.items" :key="itemIndex" class="item">
          <input v-model="item.content" :placeholder="item.placeholder">
          <button @click="removeItem(section, itemIndex)" class="remove-btn">Remove</button>
        </div>
        <button @click="addItem(section)" class="add-btn">Add Item</button>
      </div>

      <div class="actions">
        <button @click="saveInformation" class="save-btn">Save Information</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isAuthenticated = ref(false)
const password = ref('')

const authenticate = () => {
  if (password.value) {
    isAuthenticated.value = true
  } else {
    alert('Please enter a password')
  }
}

const sections = ref([
  {
    title: 'Important Documents',
    items: [
      { content: '', placeholder: 'e.g., Will location' },
      { content: '', placeholder: 'e.g., Birth certificate location' },
    ]
  },
  {
    title: 'Financial Information',
    items: [
      { content: '', placeholder: 'e.g., Bank account details' },
      { content: '', placeholder: 'e.g., Investment account information' },
    ]
  },
  {
    title: 'Digital Assets',
    items: [
      { content: '', placeholder: 'e.g., Email account credentials' },
      { content: '', placeholder: 'e.g., Social media account information' },
    ]
  },
  {
    title: 'Personal Wishes',
    items: [
      { content: '', placeholder: 'e.g., Funeral arrangements' },
      { content: '', placeholder: 'e.g., Organ donation preferences' },
    ]
  },
  {
    title: 'Important Contacts',
    items: [
      { content: '', placeholder: 'e.g., Lawyer contact information' },
      { content: '', placeholder: 'e.g., Executor of will' },
    ]
  },
])

const addItem = (section) => {
  section.items.push({ content: '', placeholder: 'Enter information' })
}

const removeItem = (section, index) => {
  section.items.splice(index, 1)
}

const saveInformation = () => {
  // This function will be implemented later to save the information to a database
  console.log('Saving information:', sections.value)
  alert('Information saved successfully (not actually saved to a database yet)')
}
</script>

<style scoped>
.in-the-event-of {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

h1 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.description {
  color: #34495e;
  margin-bottom: 2rem;
}

.password-input {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
}

.password-input input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #ced4da;
  border-radius: 4px;
}

.auth-btn, .save-btn {
  padding: 0.5rem 1rem;
  background-color: #2e5e20;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.auth-btn:hover, .save-btn:hover {
  background-color: #1e3e14;
}

.section {
  margin-bottom: 2rem;
  background-color: #f8f9fa;
  padding: 1rem;
  border-radius: 8px;
}

h2 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.item {
  display: flex;
  margin-bottom: 0.5rem;
}

.item input {
  flex-grow: 1;
  padding: 0.5rem;
  border: 1px solid #ced4da;
  border-radius: 4px;
}

.remove-btn, .add-btn {
  padding: 0.5rem 1rem;
  margin-left: 0.5rem;
  background-color: #2e5e20;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.remove-btn:hover, .add-btn:hover {
  background-color: #1e3e14;
}

.actions {
  margin-top: 2rem;
  text-align: right;
}

.save-btn {
  font-size: 1.1em;
  padding: 0.75rem 1.5rem;
}
</style>
