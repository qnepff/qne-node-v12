/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue",
    "./error.vue"
  ],
  safelist: [
    'bg-gray-50',
    'text-gray-900',
    'bg-gray-800',
    'text-white',
    'px-3',
    'py-1',
    'rounded-full',
    'text-sm',
    'z-50'
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
