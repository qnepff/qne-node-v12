<template>
  <div class="fhir-page">
    <h1>Understanding FHIR</h1>

    <div class="explanation-block child">
      <h2>For Ages 5-10</h2>
      <p>Imagine you have a big box of colorful building blocks. Each block represents a piece of information about your health, like your name, age, or the medicine you take.</p>
      <p>FHIR is like a special set of rules that helps doctors, nurses, and computers share these blocks and build them into a picture of your health.</p>
      <p>Just like how you can share your blocks with friends to build together, FHIR helps different hospitals and doctors share your health information to take better care of you!</p>
      <p>The best part is, you can keep a copy of your own health blocks on your computer at home, so you always know what's in your health box!</p>
    </div>

    <div class="explanation-block teen">
      <h2>For Ages 13-17</h2>
      <p>FHIR (Fast Healthcare Interoperability Resources) is a standard for exchanging healthcare information electronically. Think of it like a universal language for health data.</p>
      <p>Just as social media platforms allow you to share photos and messages easily, FHIR allows healthcare systems to share patient information securely and efficiently.</p>
      <p>This means that if you go to a new doctor or hospital, they can quickly access your health history, making it easier to provide you with the best care possible.</p>
      <p>What's cool is that FHIR also allows you to keep your own personal health record on your computer or phone, giving you control over your health information.</p>
    </div>

    <div class="explanation-block adult">
      <h2>For Ages 18+</h2>
      <p>FHIR (Fast Healthcare Interoperability Resources) is a next-generation standards framework for exchanging electronic health information. It's designed to facilitate interoperability between legacy health care systems, to make it easy to provide health care information to health care providers and individuals on a wide variety of devices, and to allow third-party application developers to provide medical applications which can be easily integrated into existing systems.</p>
      <p>FHIR solutions are built from a set of modular components called "Resources". These resources can easily be assembled into working systems that solve real world clinical and administrative problems at a fraction of the price of existing alternatives.</p>
      <p>By enabling seamless data exchange, FHIR aims to improve patient care, reduce healthcare costs, and empower patients to take a more active role in their health management.</p>
      <p>Importantly, FHIR supports the concept of personal health records, allowing individuals to maintain their own comprehensive health data on personal devices. This feature promotes patient engagement and ensures that individuals have access to their complete health history, regardless of changes in healthcare providers.</p>
    </div>

    <div class="advantages-section">
      <h2>Advantages of FHIR Person Resource</h2>
      <p>The FHIR Person resource serves as a global standard for interoperability, offering several key advantages:</p>
      <ul>
        <li>
          <strong>Consistency:</strong> Provides a uniform way to represent individuals across different healthcare systems and applications.
        </li>
        <li>
          <strong>Flexibility:</strong> Can represent various roles an individual might have (patient, practitioner, related person) without duplicating information.
        </li>
        <li>
          <strong>Comprehensive:</strong> Includes a wide range of demographic data, making it suitable for diverse use cases.
        </li>
        <li>
          <strong>Linkage:</strong> Allows linking to other resources like Patient or Practitioner, facilitating a complete view of an individual's interactions with the healthcare system.
        </li>
        <li>
          <strong>Global Applicability:</strong> Designed to work across different countries and healthcare systems, supporting international interoperability.
        </li>
        <li>
          <strong>Extensibility:</strong> Can be extended to include additional data elements specific to certain implementations while maintaining core interoperability.
        </li>
        <li>
          <strong>Improved Data Quality:</strong> Standardization helps reduce errors and inconsistencies in person-related data across systems.
        </li>
        <li>
          <strong>Efficient Data Exchange:</strong> Streamlines the process of sharing person information between different healthcare providers and systems.
        </li>
        <li>
          <strong>Privacy and Security:</strong> Includes built-in mechanisms for managing consent and data protection, crucial for handling sensitive personal information.
        </li>
      </ul>
    </div>

    <div class="fun-fact">
      <h2>Fun Fact!</h2>
      <p>FHIR stands for "Fast Healthcare Interoperability Resources." It's pronounced like "fire" 🔥</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FHIRPage',
}
</script>

<style scoped>
.fhir-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  background-color: #f0f8ff;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
}

h1 {
  color: #4a90e2;
  text-align: center;
}

.explanation-block {
  margin-bottom: 30px;
  padding: 20px;
  border-radius: 8px;
}

.explanation-block h2 {
  margin-top: 0;
}

.explanation-block p {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 15px;
}

.child {
  background-color: #e6f3ff;
}

.teen {
  background-color: #e6ffe6;
}

.adult {
  background-color: #fff0e6;
}

.fun-fact {
  background-color: #ffe4b5;
  border-radius: 8px;
  padding: 15px;
  margin-top: 20px;
}

.fun-fact h2 {
  color: #ff7f50;
  margin-top: 0;
}

.fun-fact p {
  font-size: 16px;
  font-style: italic;
}

.advantages-section {
  background-color: #e6ffe6;
  border-radius: 8px;
  padding: 20px;
  margin-top: 30px;
}

.advantages-section h2 {
  color: #2e8b57;
  margin-top: 0;
}

.advantages-section ul {
  padding-left: 20px;
}

.advantages-section li {
  margin-bottom: 10px;
}

.advantages-section strong {
  color: #2e8b57;
}
</style>
