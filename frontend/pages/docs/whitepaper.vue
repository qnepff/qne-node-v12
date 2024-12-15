<script setup lang="ts">
import { ref } from 'vue'
import { useHead } from '@unhead/vue' // Changed from '#imports'

// SEO metadata
useHead({
  title: 'QNE Community Whitepaper - Detailed Version',
  meta: [
    { name: 'description', content: 'Comprehensive whitepaper on the QNE Community, its technology, and its impact on secure, efficient communication and data sharing.' }
  ]
})

const sections = ref([
  {
    title: 'Executive Summary',
    content: `The QNE Community represents a paradigm shift in digital interaction, offering unparalleled security, efficiency, and user control.<br><br>
    This whitepaper details the revolutionary platform that leverages QuickNEasy software to create a secure, efficient, and user-centric environment for communication and data sharing. It explores the core technologies, unique identification system, and governance model that make the QNE Community a groundbreaking solution for digital privacy and seamless data exchange.<br><br>
    Join us in shaping the future of secure, user-controlled digital interactions.`
  },
  {
    title: 'Introduction to QNE Community',
    content: `The QNE Community is built on the principle that once data is entered into the internet, it should never need to be entered again.<br><br>
    Our mission is to empower individuals and organizations with control over their data while simplifying digital interactions. The QNE Community uses QuickNEasy software to enable organizations and individuals to live in an environment that is easy to use, easy to manage, and easy to maintain.<br><br>
    We are committed to keeping data out of the reach of bad actors who might abuse it, ensuring a secure and trustworthy digital ecosystem.`
  },
  {
    title: 'Core Technology',
    content: `QuickNEasy software forms the backbone of the QNE Community.<br><br>
    It utilizes advanced peer-to-peer (P2P) communication, the Noise Protocol Framework for encryption, and Apache Avro for data serialization. This combination ensures secure, direct communication between peers without relying on centralized servers. The software is installed on devices located in a person's or organization's premises, with the option for remote hosting in data centers.<br><br>
    Additionally, the QNE Community leverages standardized data formats, including the UN register of Product types, to ensure consistency and interoperability across the platform.`
  },
  {
    title: 'Data Standards and Interoperability',
    content: `The QNE Community is committed to using and promoting global data standards to ensure maximum interoperability and efficiency. Key standards incorporated into the platform include:`,
    listItems: [
      { number: 1, title: 'FHIR (Fast Healthcare Interoperability Resources)', content: 'Used for medical data handling and as a basis for identity verification through the FHIR Patient resource.' },
      { number: 2, title: 'United Nations Standard Products and Services Code (UNSPSC)', content: 'Adopted to standardize product and service information across the community, facilitating seamless e-commerce and inventory management.' },
      { number: 3, title: 'Apache Avro', content: 'Employed for efficient data serialization, ensuring consistent data representation across the platform.' },
      { number: 4, title: 'Ongoing Standards Integration', content: 'As the QNE Community grows, many other industry-specific and global standards will be incorporated to meet the evolving needs of our diverse user base and ensure broad interoperability.' }
    ],
    conclusion: `By adhering to these international standards and remaining open to adopting new ones, the QNE Community not only enhances its internal consistency but also positions itself for easy integration with external systems and future expansions. This approach significantly reduces data entry redundancy and improves overall data quality and reliability across the community.`
  },
  {
    title: 'Data Control and Privacy',
    content: `At the heart of the QNE Community is the principle of user-controlled data sharing.<br><br>
    Members decide who can access their data, for how long, and under what conditions. This granular control ensures privacy while enabling efficient information exchange when needed. The community enforces strict rules on data handling, including mandatory deletion of data upon request.<br><br>
    This approach empowers users to maintain control over their personal information in the digital realm.`
  },
  {
    title: 'Unique Identification System',
    content: `The QNE Community employs an innovative identification system that balances ease of use with the ability to accommodate billions of users:`,
    listItems: [
      { number: 1, title: 'Segment-Based Naming', content: 'Users are identified by a segment number and a chosen name, such as "1-alice" or "23-bob".' },
      { number: 2, title: 'Base 36 Numbering', content: 'The system uses base 36 numbers (0-9 and A-Z) for segment identification, allowing for over 76 billion unique peer IDs.' },
      { number: 3, title: 'Easy-to-Remember', content: 'This format creates easy-to-remember identifiers while ensuring global uniqueness.' },
      { number: 4, title: 'Scalability', content: 'As the community grows, new segments can be added without disrupting existing identifiers.' },
      { number: 5, title: 'Local Uniqueness', content: 'Within each segment, users can choose any available name, allowing for personalization.' }
    ],
    conclusion: `This system enables efficient peer-to-peer connections across the entire QNE Community.<br><br>
    It provides a scalable solution that can grow with the community while maintaining simplicity for users. The combination of segment-based naming and base 36 numbering ensures a vast address space for future expansion.<br><br>
    Ultimately, this identification system forms the backbone of the QNE Community's interconnected network.`
  },
  {
    title: 'Identity Verification and Security',
    content: `The QNE Community implements a robust, multi-layered approach to identity verification and security:`,
    listItems: [
      { number: 1, title: 'FHIR Patient Resource', content: 'The foundation of identity validation is the FHIR (Fast Healthcare Interoperability Resources) Patient resource, which provides a standardized format for storing comprehensive identity information.' },
      { number: 2, title: 'Multi-Tiered Verification', content: 'Identity is verified through trusted entities such as medical practitioners, financial institutions, and government agencies.' },
      { number: 3, title: 'Annual Photo Verification', content: 'Members submit a yearly head-and-shoulders photo to ensure ongoing identity confirmation.' },
      { number: 4, title: 'Granular Access Control', content: 'Users have full control over which aspects of their identity information are shared and with whom.' },
      { number: 5, title: 'Noise Protocol Framework', content: 'All communications are secured using the Noise Protocol Framework, which provides strong encryption with perfect forward secrecy.' },
      { number: 6, title: 'Decentralized Architecture', content: 'The peer-to-peer nature of the system eliminates central points of failure or attack.' }
    ],
    conclusion: `This comprehensive approach ensures a high level of trust within the community.<br><br>
    By combining standardized data formats, multi-tiered verification, and advanced encryption, the QNE Community creates a secure environment for digital interactions. The decentralized architecture further enhances security by eliminating single points of failure.<br><br>
    User privacy and control over personal data remain paramount in this secure ecosystem.`
  },
  {
    title: 'Communication Protocols',
    content: `The QNE Community employs two distinct protocols to ensure secure, direct, and tamper-proof communication between peers.<br><br>
    Web access is handled via HTTPS, providing secure browsing and data transfer through traditional web interfaces. For peer-to-peer communication, including messaging and general data transfer, the community utilizes UDP secured by the Noise Protocol Framework. This combination guarantees end-to-end encryption and eliminates any possibility of third-party interference or man-in-the-middle attacks.<br><br>
    <div class="crucial-info-green">
      <strong>Very important:</strong> As long as peers maintain control of their passwords, it is cryptographically guaranteed that you are connecting directly with the intended peer, without any intermediaries able to intercept or alter the communication.
    </div><br>
    This dual-protocol approach not only ensures both accessibility and unparalleled security across all types of digital interactions within the community but also provides a level of communication integrity that is unmatched in conventional systems. The direct peer-to-peer nature of these connections, coupled with state-of-the-art encryption, makes the QNE Community a fortress of digital privacy and security.`
  },
  {
    title: 'Medical Data Handling',
    content: `The QNE Community bases its operations on the FHIR (Fast Healthcare Interoperability Resources) standard for handling medical data, with the FHIR Patient resource playing a crucial role in identity management.`,
    listItems: [
      { number: 1, title: 'Standardized Identity Verification', content: 'The FHIR Patient resource provides a comprehensive, standardized way to store and verify identity information, enhancing the reliability of user authentication.' },
      { number: 2, title: 'Interoperability', content: 'By using FHIR standards, the QNE Community ensures compatibility with existing health information exchange systems, facilitating seamless and secure sharing of health-related information.' },
      { number: 3, title: 'Granular Access Control', content: 'The FHIR standard allows for fine-grained control over data access, aligning perfectly with the QNE Community\'s principle of user-controlled data sharing.' },
      { number: 4, title: 'Enhanced Security', content: 'The structured nature of FHIR resources makes it easier to implement robust security measures and audit trails for all data access and modifications.' },
      { number: 5, title: 'Future-Proofing', content: 'As the healthcare industry increasingly adopts FHIR standards, the QNE Community is well-positioned for future integrations and expansions in health data management.' }
    ],
    conclusion: `This approach ensures secure and standardized handling of medical data.<br><br>
    By leveraging FHIR standards, the QNE Community not only enhances interoperability with existing healthcare systems but also provides a solid foundation for identity verification. The granular access control and enhanced security measures align perfectly with the community's principles of user data control and privacy.<br><br>
    This standardized approach to medical data handling positions the QNE Community at the forefront of secure and efficient health information exchange.`
  },
  {
    title: 'Community Governance',
    content: `The QNE Community is self-governing, with members playing a crucial role in defining and enforcing community rules.<br><br>
    This democratic approach ensures that the platform evolves to meet the needs of its users while maintaining high standards of conduct and data protection. Members participate in decision-making processes, propose and vote on new features, and help shape the community's policies. Non-compliance with community rules can result in exclusion from the community, ensuring that all members adhere to the agreed-upon standards.<br><br>
    This governance model fosters a sense of ownership and responsibility among community members, driving the platform's continuous improvement and adaptation to user needs.<br><br>
    <div class="crucial-info-green">
      <strong>Community Growth and Integration:</strong> A large community will have members inside organizations as well as customers of the organization. Members inside the organizations are able to introduce the community and its benefits and help to implement the community system. This dual membership structure facilitates organic growth and seamless integration of the QNE Community across various sectors.
    </div>`
  },
  {
    title: 'Membership and Fees',
    content: `Members pay an annual fee of USD30 plus VAT (if required in their country of origin).<br><br>
    This fee covers the cost of running the community, providing the initial QuickNEasy software, and funding ongoing development. The community reinvests much of its income into contract developers for software maintenance and new feature development based on member requests. This model ensures that the platform remains financially sustainable while continuously evolving to meet user needs.<br><br>
    The low annual fee makes the QNE Community accessible to a wide range of users, promoting inclusivity and diversity within the community.`
  },
  {
    title: 'Future Development',
    content: `The QNE Community is committed to continuous improvement and innovation.<br><br>
    Member feedback drives the development roadmap, with contracted developers and volunteers working to implement new features and enhance existing ones. This collaborative approach ensures that the platform remains cutting-edge and user-focused. Future developments may include special interest groups, chat rooms, and other community-driven features, all designed to enhance the user experience and expand the capabilities of the QNE Community.<br><br>
    By staying adaptable and responsive to user needs, the QNE Community aims to remain at the forefront of secure, efficient digital communication and data sharing.`
  },
  {
    title: 'Case Studies',
    content: ``,
    subSections: [
      {
        title: 'Streamlined Hotel Booking',
        content: `A QNE Community member, 87-bob, books a room at 34-HotelABC. Using their QNE identities, the entire booking process is completed without manual data entry. Upon arrival, 87-bob's identity is quickly verified using their QNE profile image, demonstrating the efficiency and security of the system.`
      },
      {
        title: 'Secure Medical Information Sharing',
        content: `A patient needs to share medical records with a specialist. Through the QNE Community, the patient grants temporary, specific access to their medical data. The specialist receives only the necessary information, which is automatically revoked after the consultation, ensuring data privacy and eliminating the need for physical document transfers.`
      },
      {
        title: 'Simplified E-commerce',
        content: `An online retailer, 56-EcoStore, integrates with the QNE Community. Customers can now make purchases using their QNE identity, with shipping and payment details automatically populated. Post-purchase, customers can easily manage what information the store retains, enhancing privacy control.`
      }
    ]
  },
  {
    title: 'Conclusion',
    content: `The QNE Community represents a paradigm shift in digital interaction, offering unparalleled security, efficiency, and user control. By joining the QNE Community, individuals and organizations can participate in shaping a future where data privacy and seamless communication go hand in hand. The community's commitment to user-driven development and democratic governance ensures that it will continue to evolve to meet the changing needs of its members.`
  }
])
</script>

<template>
  <div class="pt-12"> <!-- Add padding-top here -->
    <div class="whitepaper-container">
      <main>
        <header class="mb-12 text-center">
          <h1 class="text-5xl font-bold mb-4">QNE Community Whitepaper</h1>
          <p class="text-2xl text-gray-600">Revolutionizing Secure Communication and Data Sharing</p>
        </header>
        
        <div v-for="(section, index) in sections" :key="index" class="mb-16">
          <h2 class="text-3xl font-semibold mb-6 text-blue-800">{{ section.title }}</h2>
          <div class="mb-6 leading-relaxed text-xl" v-html="section.content"></div>
          
          <ol v-if="section.listItems" class="list-none pl-0 mb-6">
            <li v-for="item in section.listItems" :key="item.number" class="mb-4 flex items-start">
              <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 text-blue-800 flex items-center justify-center mr-4 mt-1">{{ item.number }}</span>
              <div>
                <h3 class="text-xl font-semibold text-blue-700 mb-1">{{ item.title }}</h3>
                <p class="text-lg">{{ item.content }}</p>
              </div>
            </li>
          </ol>
          
          <div v-if="section.conclusion" class="mb-6 leading-relaxed text-xl" v-html="section.conclusion"></div>
          
          <div v-if="section.subSections" class="ml-8">
            <div v-for="(subSection, subIndex) in section.subSections" :key="subIndex" class="mb-8">
              <h3 class="text-2xl font-medium mb-4 text-blue-700">{{ subSection.title }}</h3>
              <div class="text-lg" v-html="subSection.content"></div>
            </div>
          </div>
        </div>
      </main>
      
      <footer class="mt-16 pt-8 border-t border-gray-300 text-center text-gray-600">
        <p class="text-lg">© {{ new Date().getFullYear() }} QNE Community. All rights reserved.</p>
      </footer>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Merriweather:wght@300;400;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600;700&display=swap');

.whitepaper-container {
  font-family: 'Merriweather', serif;
  max-width: 64rem;
  margin: 0 auto;
  padding: 3rem 2rem;
  background-color: white;
  color: #1f2937;
}

h1, h2, h3 {
  font-family: 'Open Sans', sans-serif;
}

p {
  font-size: 1.25rem;
  line-height: 1.8;
}

h2 {
  color: #1e40af;
}

h3 {
  color: #1e3a8a;
}

/* Add this new class definition */
.crucial-info-green {
  background-color: #d1fae5;
  border-left: 4px solid #10b981;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 0.25rem;
}

.crucial-info-green strong {
  color: #047857;
}
</style>