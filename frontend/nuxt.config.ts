import { defineNuxtConfig } from 'nuxt/config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['~/assets/css/tailwind.css'],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  ssr: true,
  nitro: {
    preset: 'static',
    prerender: {
      crawlLinks: true,
      routes: [
        '/',
        '/communication/video-chat'
      ]
    }
  },

  app: {
    baseURL: '/',
    buildAssetsDir: '/_nuxt/',
    head: {
      link: [
        { rel: 'stylesheet', href: '/_nuxt/entry.css' }
      ]
    }
  },

  // Configure Vite within Nuxt
  vite: {
    server: {
      fs: {
        allow: [
          // Add the paths that need to be accessible by Vite
          '/home/paulf/node_modules/@vue/devtools-api',
          '/home/paulf/Documents/vscode/current/qne-frontend/qne-frontend',
          '/home/paulf/node_modules/nuxt',
          '/home/paulf/node_modules/@nuxt/devtools',
          '/home/paulf/node_modules/vite/dist/client'
        ]
      }
    }
  },
})
