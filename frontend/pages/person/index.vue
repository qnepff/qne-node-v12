<template>
    <div class="page-wrapper">
        <div class="container">
            <h1>FHIR Person Resource</h1>
            <h2>A Versatile Standard for Person Data</h2>

            <div class="description-section">
                <p>While originally designed for healthcare, the FHIR Person resource is a robust and flexible standard that can be adapted for various industries and use cases beyond medicine.</p>
            </div>

            <div class="advantages-section">
                <h2>Advantages of Using FHIR Person Resource Across Industries</h2>
                <ul>
                    <li v-for="advantage in advantages" :key="advantage.title">
                        <strong>{{ advantage.title }}:</strong> {{ advantage.description }}
                    </li>
                </ul>
            </div>

            <div class="use-cases-section">
                <h2>Potential Use Cases Outside Healthcare</h2>
                <ul>
                    <li v-for="useCase in useCases" :key="useCase">{{ useCase }}</li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    setup() {
        const advantages = [
            { title: 'Comprehensive Data Model', description: 'Covers a wide range of personal attributes applicable to many domains.' },
            { title: 'Interoperability', description: 'Designed for easy data exchange, beneficial for any industry requiring standardized person data.' },
            { title: 'Extensibility', description: 'Can be customized to include industry-specific attributes while maintaining a standard core.' },
            { title: 'Identity Management', description: 'Robust framework for managing identities across systems and organizations.' },
            { title: 'Privacy and Security', description: 'Built-in mechanisms for consent and data protection, crucial for handling personal information in any field.' },
            { title: 'Global Applicability', description: 'Designed to work across different countries and systems, supporting international operations.' },
            { title: 'Relationship Modeling', description: 'Capable of representing complex relationships between individuals and organizations.' },
        ];

        const useCases = [
            'Human Resources: Standardizing employee data across multinational corporations',
            'Education: Managing student information across educational institutions',
            'Government Services: Centralizing citizen data for various public services',
            'Financial Services: Standardizing customer information across banking and insurance sectors',
            'Hospitality: Managing guest profiles across hotel chains and travel services',
            'Retail: Creating unified customer profiles for personalized shopping experiences',
            'Non-profit Organizations: Managing volunteer and beneficiary information',
        ];

        return {
            advantages,
            useCases,
        };
    }
}
</script>

<style scoped>
.page-wrapper {
    min-height: 100vh;
    background-color: #f0f4f8;
    padding: 20px;
}

.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
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

ul {
    list-style-type: none;
    padding-left: 0;
    width: 100%;
    text-align: left;
}

li {
    margin-bottom: 15px;
    background-color: #f8f9fa;
    padding: 10px;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.field-name, strong {
    font-weight: bold;
    color: #3498db;
    margin-right: 5px;
}

.field-description {
    color: #34495e;
}

.advantages-section {
    margin-top: 30px;
    background-color: #e6ffe6;
    border-radius: 8px;
    padding: 20px;
    width: 100%;
}

.advantages-section h2 {
    color: #2e8b57;
}

.description-section, .advantages-section, .use-cases-section {
    margin-top: 30px;
    background-color: #f8f9fa;
    border-radius: 8px;
    padding: 20px;
    width: 100%;
    text-align: left;
}

.advantages-section h2, .use-cases-section h2 {
    color: #2e8b57;
}

.advantages-section ul, .use-cases-section ul {
    padding-left: 20px;
}

.advantages-section li, .use-cases-section li {
    margin-bottom: 10px;
    background-color: white;
}

strong {
    color: #3498db;
}
</style>