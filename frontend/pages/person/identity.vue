<template>
  <div class="page-wrapper">
    <div class="container">
      <h1>FHIR Person Resource: Identifiers</h1>
      
      <div class="description-section">
        <p>The FHIR Person resource has the capability to contain a list of identifiers. This feature allows for the representation of various forms of identification associated with an individual across different systems or contexts.</p>
      </div>

      <div class="identifier-section">
        <h2>Identifier Structure</h2>
        <p>Each identifier in FHIR typically includes:</p>
        <ul>
          <li><strong>Use:</strong> The purpose of the identifier (usual, official, temp, secondary, old)</li>
          <li><strong>Type:</strong> The type of identifier</li>
          <li><strong>System:</strong> The namespace for the identifier</li>
          <li><strong>Value:</strong> The actual identifier value</li>
        </ul>
      </div>

      <div class="examples-section">
        <h2>Examples of Identifiers</h2>
        <ul>
          <li v-for="example in identifierExamples" :key="example.type">
            <strong>{{ example.type }}:</strong> 
            <span>{{ example.description }}</span>
            <code v-if="example.example">{{ example.example }}</code>
          </li>
        </ul>
      </div>

      <div class="benefits-section">
        <h2>Benefits of Multiple Identifiers</h2>
        <ul>
          <li>Enables cross-system identification</li>
          <li>Supports various identification schemes</li>
          <li>Facilitates data integration from multiple sources</li>
          <li>Allows for tracking of historical identifiers</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  setup() {
    const identifierExamples = [
      { 
        type: "Driver's License", 
        description: "Official government-issued ID",
        example: '{"system": "urn:oid:2.16.840.1.113883.4.3.25", "value": "S99969696"}'
      },
      { 
        type: "Passport", 
        description: "International travel document",
        example: '{"system": "urn:oid:2.16.840.1.113883.4.330", "value": "X123456"}'
      },
      { 
        type: "Social Security Number", 
        description: "U.S. national identification number",
        example: '{"system": "http://hl7.org/fhir/sid/us-ssn", "value": "123-45-6789"}'
      },
      { 
        type: "Employee ID", 
        description: "Organization-specific identifier"
      },
      { 
        type: "Student ID", 
        description: "Educational institution identifier"
      },
      { 
        type: "Medical Record Number", 
        description: "Healthcare provider's patient identifier"
      },
      { 
        type: "Biometric Identifier", 
        description: "Such as fingerprint or retinal scan data"
      },
      { 
        type: "National ID", 
        description: "Country-specific national identification number"
      }
    ];

    return {
      identifierExamples
    };
  }
};
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  background-color: #f0f4f8;
  padding: 20px;
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

h1, h2 {
  color: #2c3e50;
}

.description-section, .identifier-section, .examples-section, .benefits-section {
  margin-top: 20px;
  padding: 15px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

ul {
  padding-left: 20px;
}

li {
  margin-bottom: 10px;
}

strong {
  color: #3498db;
}

code {
  display: block;
  margin-top: 5px;
  padding: 5px;
  background-color: #e9ecef;
  border-radius: 4px;
  font-size: 0.9em;
}
</style>
