<template>
    <main class="max-w-4xl mx-auto px-8 py-12">
        <ConnectionInfo />
        <header>
            <h1 class="text-4xl font-bold text-center mb-2">QNE Community</h1>
            <h2 class="text-2xl text-center mb-4">A Better Internet Experience</h2>
            <p class="text-center mb-8">Join the visionary members shaping the future of online communities</p>
        </header>

        <section class="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-8">
            <h3 class="text-xl font-semibold mb-4">Become a Visionary Member</h3>
            <ul class="list-disc pl-6 space-y-2">
                <li>10,000 Visionary members will receive lifetime membership for one fee</li>
                <li>First members will receive priority choice of member names (up to 16 chars)</li>
                <li>Join the waiting list - Email visionary@qne.community</li>
                <li>Follow us on X @qnecommunity for updates</li>
            </ul>
        </section>

        <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <article v-for="feature in features" :key="feature.title" 
                     class="bg-gray-100 border border-gray-300 rounded-md p-6">
                <h3 class="text-xl font-semibold mb-2">{{ feature.title }}</h3>
                <p class="mb-4">{{ feature.description }}</p>
            </article>
        </section>

        <section class="mt-12 text-center">
            <h3 class="text-2xl font-semibold mb-4">Join the Community</h3>
            <p class="mb-2">Annual membership: USD30 (plus VAT/GST)</p>
            <p>Your membership fee directly improves the community experience</p>
        </section>
    </main>
</template>

<script setup>
import ConnectionInfo from '~/components/ConnectionInfo.vue'

const features = [
    {
        title: "AD-FREE Experience",
        description: "Enjoy social media and communication without intrusive advertising"
    },
    {
        title: "Secure Communication",
        description: "End-to-end encrypted mail, messaging, and chat between verified members"
    },
    {
        title: "Health Data Management",
        description: "Private health data management based on FHIR standards"
    },
    {
        title: "Community Control",
        description: "Members have a say in how the community functions, no venture capitalists"
    },
    {
        title: "Integrated Platform",
        description: "Access budgeting, shopping lists, and more on a single unified platform"
    },
    {
        title: "Member Privacy",
        description: "Your data remains yours - we don't share or sell your information"
    }
]
</script>
