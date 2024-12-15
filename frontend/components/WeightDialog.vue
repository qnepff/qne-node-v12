<template>
  <div class="weight-dialog">
    <h2>Weight Tracker</h2>
    <div class="unit-toggle">
      <span class="toggle-label">Select:</span>
      <ToggleSwitch
        v-model="isKg"
        label-on="kg"
        label-off="lbs"
        id="weight-unit-toggle"
        @update:modelValue="toggleUnit"
      />
    </div>
    <div v-if="!initialWeightSet" class="initial-weight-input">
      <label for="initial-weight">Enter your current weight:</label>
      <div class="input-group">
        <input
          id="initial-weight"
          ref="initialWeightInput"
          v-model.number="initialWeight"
          type="number"
          step="0.1"
          @blur="setInitialWeight"
          @keyup.enter="setInitialWeight"
        />
        <div class="custom-spinner">
          <button @click="incrementWeight" class="spinner-up">▲</button>
          <button @click="decrementWeight" class="spinner-down">▼</button>
        </div>
      </div>
    </div>
    <div v-else-if="!selectedWeight" class="weight-options" ref="weightOptionsRef">
      <label v-for="option in weightOptions" :key="option">
        <input
          type="radio"
          :value="option"
          v-model="selectedWeight"
          name="weight"
          @change="saveWeight"
        >
        {{ formatWeight(option) }} {{ unit }}
      </label>
    </div>
    <div v-else class="weight-history">
      <h3>Weight History</h3>
      <svg width="300" height="150" viewBox="0 0 300 150">
        <polyline
          :points="graphPoints"
          fill="none"
          stroke="#4CAF50"
          stroke-width="2"
        />
        <circle
          v-for="(point, index) in graphPoints.split(' ')"
          :key="index"
          :cx="point.split(',')[0]"
          :cy="point.split(',')[1]"
          r="3"
          fill="#4CAF50"
        />
      </svg>
      <button @click="backToRadioView">Back</button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted } from 'vue';
import ToggleSwitch from './ToggleSwitch.vue';

// This section defines reactive variables using Vue's ref function

// Stores the initial weight entered by the user
const initialWeight = ref(null);

// Tracks whether the initial weight has been set
const initialWeightSet = ref(false);

// Determines if the weight unit is kilograms (true) or pounds (false)
const isKg = ref(true);

// Stores the selected weight from the radio button options
const selectedWeight = ref(null);

// Reference to the DOM element containing weight options
const weightOptionsRef = ref(null);

// Reference to the input field for initial weight
const initialWeightInput = ref(null);

// Mock weight history data
const weightHistory = ref([80, 79.5, 79, 78.5, 79, 78, 77.5, 77, 76.5, 76]);

const unit = computed(() => isKg.value ? 'kg' : 'lbs');

const weightOptions = computed(() => {
  if (!initialWeight.value) return [];
  const baseWeight = Math.round(initialWeight.value * 2) / 2; // Round to nearest 0.5
  return Array.from({ length: 21 }, (_, i) => baseWeight - 5 + i * 0.5);
});

const graphPoints = computed(() => {
  const points = weightHistory.value.map((weight, index) => {
    const x = index * (300 / (weightHistory.value.length - 1));
    const y = 150 - (weight - Math.min(...weightHistory.value)) / (Math.max(...weightHistory.value) - Math.min(...weightHistory.value)) * 150;
    return `${x},${y}`;
  });
  return points.join(' ');
});

const setInitialWeight = () => {
  if (initialWeight.value) {

    // Convert weight to FHIR Observation format
    const fhirObservation = {
      resourceType: 'Observation',
      status: 'preliminary',
      category: [
        {
          coding: [
            {
              system: 'http://terminology.hl7.org/CodeSystem/observation-category',
              code: 'vital-signs',
              display: 'Vital Signs'
            }
          ]
        }
      ],
      code: {
        coding: [
          {
            system: 'http://loinc.org',
            code: '29463-7',
            display: 'Body weight'
          }
        ],
        text: 'Body weight'
      },
      valueQuantity: {
        value: initialWeight.value,
        unit: unit.value,
        system: 'http://unitsofmeasure.org',
        code: unit.value === 'kg' ? 'kg' : '[lb_av]'
      },
      effectiveDateTime: new Date().toISOString()
    };

    // Convert FHIR Observation to Avro format
    const avroSchema = {
      type: 'record',
      name: 'WeightObservation',
      fields: [
        { name: 'resourceType', type: 'string' },
        { name: 'status', type: 'string' },
        { name: 'category', type: { type: 'array', items: { type: 'record', name: 'Category', fields: [{ name: 'coding', type: { type: 'array', items: { type: 'record', name: 'Coding', fields: [{ name: 'system', type: 'string' }, { name: 'code', type: 'string' }, { name: 'display', type: 'string' }] } } }] } } },
        { name: 'code', type: { type: 'record', name: 'Code', fields: [{ name: 'coding', type: { type: 'array', items: 'Coding' } }, { name: 'text', type: 'string' }] } },
        { name: 'valueQuantity', type: { type: 'record', name: 'ValueQuantity', fields: [{ name: 'value', type: 'double' }, { name: 'unit', type: 'string' }, { name: 'system', type: 'string' }, { name: 'code', type: 'string' }] } },
        { name: 'effectiveDateTime', type: 'string' }
      ]
    };

    // In a real-world scenario, you would use an Avro library to serialize the data
    // For this example, we'll use a simple JSON.stringify as a placeholder
    const serializedData = JSON.stringify(fhirObservation);
    console.log('Serialized Avro-like data:', serializedData);
    const avroData = {
      // This line spreads all properties from the fhirObservation object into the avroData object.
      // It's a shorthand way to copy all properties from fhirObservation to avroData.
      // However, this approach doesn't actually convert the data to Avro format.
      // In a real-world scenario, you would use an Avro library to properly serialize the data.
      // This is just a placeholder to demonstrate the concept.
      ...fhirObservation
    };

    console.log('FHIR Observation:', fhirObservation);
    console.log('Avro Data:', avroData);

    // Here you would typically send the Avro data to your server or process it further


    
    initialWeightSet.value = true;
    selectedWeight.value = null;
    // Send initial weight to the server
    sendInitialWeightToServer(initialWeight.value);

    // This code uses Vue's nextTick function to ensure that the DOM has updated
    // before attempting to scroll to the selected weight option.
    // nextTick waits for the next DOM update cycle before executing the callback.
    nextTick(() => {
      // Once the DOM has updated, we call the scrollToSelectedWeight function.
      // This function will scroll the view to center on the currently selected weight option.
      scrollToSelectedWeight();
    });
  }
};

async function sendInitialWeightToServer(weight) {
      // Make an API call to send the initial weight to the server
      try {
        // Assuming you're using Nuxt 3 with the built-in $fetch utility
        const response = await $fetch('/api/weight', {
          method: 'POST',
          body: { weight, unit: unit.value },
        });

        if (!response.ok) {
          throw new Error('Failed to send weight to server');
        }
        console.log(`Successfully sent initial weight to server: ${weight} ${unit.value}`);
      } catch (error) {
        console.error('Error sending weight to server:', error);
        // Handle the error appropriately (e.g., show a user-friendly message)
      }
    }


const toggleUnit = () => {
  if (initialWeight.value) {
    if (!isKg.value) {
      initialWeight.value = Math.round(initialWeight.value * 2.20462 * 10) / 10;
      if (selectedWeight.value) {
        selectedWeight.value = Math.round(selectedWeight.value * 2.20462 * 2) / 2;
      }
    } else {
      initialWeight.value = Math.round(initialWeight.value / 2.20462 * 10) / 10;
      if (selectedWeight.value) {
        selectedWeight.value = Math.round(selectedWeight.value / 2.20462 * 2) / 2;
      }
    }
    nextTick(() => {
      scrollToSelectedWeight();
    });
  }
};

const formatWeight = (weight) => {
  return isKg.value ? weight.toFixed(1) : Math.round(weight);
};

const saveWeight = () => {
  console.log(`Weight saved: ${selectedWeight.value} ${unit.value}`);
  // Here you would typically save the weight to your data store
  weightHistory.value.push(selectedWeight.value);
  weightHistory.value = weightHistory.value.slice(-10); // Keep only last 10 entries
};

const backToRadioView = () => {
  selectedWeight.value = null;
  nextTick(() => {
    scrollToSelectedWeight();
  });
};

const scrollToSelectedWeight = () => {
  if (weightOptionsRef.value) {
    const selectedOption = weightOptionsRef.value.querySelector('input:checked');
    if (selectedOption) {
      selectedOption.scrollIntoView({ behavior: 'smooth', block: 'center' });
    }
  }
};

const incrementWeight = () => {
  if (initialWeight.value === null) initialWeight.value = 0;
  initialWeight.value = Math.round((initialWeight.value + 0.1) * 10) / 10;
};

const decrementWeight = () => {
  if (initialWeight.value === null) initialWeight.value = 0;
  initialWeight.value = Math.max(0, Math.round((initialWeight.value - 0.1) * 10) / 10);
};

onMounted(() => {
  if (initialWeightSet.value) {
    scrollToSelectedWeight();
  } else {
    // Set focus on the initial weight input
    initialWeightInput.value.focus();
  }
});

// Don't forget to define the emit
const emit = defineEmits(['close']);
</script>

<style scoped>
.weight-dialog {
  padding: 1rem;
  text-align: center;
}

h2 {
  margin-bottom: 1rem;
}

.unit-toggle {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 1rem;
}

.toggle-label {
  margin-right: 0.5rem;
}

.initial-weight-input {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 1rem;
}

.input-group {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 0.5rem;
}

input[type="number"] {
  width: 150px;
  height: 50px;
  padding: 0.75rem;
  font-size: 1.5rem;
  font-weight: bold;
  border: 1px solid #ced4da;
  border-radius: 4px;
  text-align: center;
}

/* Remove the default spinner styles in Firefox */
input[type="number"] {
  -moz-appearance: textfield;
}

/* Remove the default spinner styles in WebKit browsers */
input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

.custom-spinner {
  display: flex;
  flex-direction: column;
  margin-left: 10px;
}

.custom-spinner button {
  width: 30px;
  height: 30px;
  font-size: 1.2rem;
  background-color: #f0f0f0;
  border: 1px solid #ced4da;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.custom-spinner button:hover {
  background-color: #e0e0e0;
}

.spinner-up {
  border-top-right-radius: 4px;
  border-top-left-radius: 4px;
}

.spinner-down {
  border-bottom-right-radius: 4px;
  border-bottom-left-radius: 4px;
  margin-top: 1px;
}

.weight-options {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
  max-height: 300px;
  overflow-y: auto;
}

.weight-history {
  display: flex;
  flex-direction: column;
  align-items: center;
}

button {
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  font-size: 1rem;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>
