<template>
  <div class="page-wrapper">
    <div class="container">
      <h1>FHIR Patient Resource: Identity Verification for Global Use</h1>
      
      <div class="description-section">
        <p>While originally designed for healthcare, the FHIR Patient resource's identity verification mechanisms have broad applicability across various sectors. These mechanisms provide a robust framework for managing identity certainty, which is crucial not only in healthcare but in many global contexts where proving identity is essential.</p>
        <p>Notably, medical certification of identity often stands out as one of the most reliable methods. This is because patients typically interact with multiple healthcare providers over time, each interaction potentially raising the level of identity certainty.</p>
      </div>

      <div class="verification-section">
        <h2>Identity Verification Mechanisms</h2>
        <ul>
          <li v-for="mechanism in verificationMechanisms" :key="mechanism.name">
            <strong>{{ mechanism.name }}:</strong> {{ mechanism.description }}
          </li>
        </ul>
      </div>

      <div class="certainty-levels-section">
        <h2>Levels of Identity Certainty</h2>
        <p>These mechanisms can be used to define levels of certainty applicable across various sectors:</p>
        <ul>
          <li v-for="level in certaintyLevels" :key="level.name">
            <strong>{{ level.name }}:</strong> {{ level.description }}
          </li>
        </ul>
      </div>

      <div class="examples-section">
        <h2>Global Applications of Identity Verification</h2>
        <ul>
          <li v-for="example in globalExamples" :key="example.scenario">
            <strong>{{ example.scenario }}:</strong> 
            <span>{{ example.description }}</span>
          </li>
        </ul>
      </div>

      <div class="medical-certification-section">
        <h2>The Unique Value of Medical Certification</h2>
        <p>Medical certification of identity offers several advantages:</p>
        <ul>
          <li v-for="advantage in medicalAdvantages" :key="advantage">{{ advantage }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
// definePageMeta({
//   layout: 'qne'
// })  


export default {
  setup() {
    const verificationMechanisms = [
      { 
        name: "Identifier.type", 
        description: "Indicates the type and reliability of an identifier (e.g., official, temp, secondary)."
      },
      { 
        name: "Identifier.period", 
        description: "Specifies the time period for which an identifier is valid, useful for managing temporary or changing identifiers."
      },
      { 
        name: "Link element", 
        description: "Allows linking of records with a 'type' attribute to indicate the nature of the link (e.g., 'replaced-by', 'refer', 'seealso')."
      },
      { 
        name: "Active flag", 
        description: "Indicates whether a record is currently active, which can be used to mark uncertain or temporary records."
      },
    ];

    const certaintyLevels = [
      { 
        name: "Unverified", 
        description: "Initial record creation with minimal verification."
      },
      { 
        name: "Partially Verified", 
        description: "Some identifiers or information have been confirmed."
      },
      { 
        name: "Fully Verified", 
        description: "All critical identifiers and information have been confirmed."
      },
      { 
        name: "Legally Attested", 
        description: "Identity has been legally attested or notarized."
      },
    ];

    const globalExamples = [
      { 
        scenario: "Financial Services", 
        description: "Banks can use these mechanisms to track the verification status of customer identities for KYC compliance."
      },
      { 
        scenario: "Government Services", 
        description: "Agencies can manage citizen identities with varying levels of verification for different services."
      },
      { 
        scenario: "International Travel", 
        description: "Border control can use these mechanisms to track the reliability of travel documents and identity claims."
      },
      { 
        scenario: "Online Platforms", 
        description: "Social media or e-commerce sites can implement identity verification levels to enhance user trust and security."
      },
    ];

    const medicalAdvantages = [
      "Multiple touchpoints: Patients often interact with various healthcare providers, each verifying their identity.",
      "Rigorous verification: Medical contexts often require strict identity checks for patient safety and legal reasons.",
      "Longitudinal data: Health records typically span long periods, allowing for identity confirmation over time.",
      "Biometric data: Medical records often include unique identifiers like fingerprints or DNA profiles.",
      "Legal standing: Medical records are often considered legally valid forms of identification in many jurisdictions."
    ];

    return {
      verificationMechanisms,
      certaintyLevels,
      globalExamples,
      medicalAdvantages
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

.description-section, .verification-section, .certainty-levels-section, .examples-section, .medical-certification-section {
  margin-top: 20px;
  padding: 15px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.medical-certification-section {
  background-color: #e6ffe6;
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
</style>
