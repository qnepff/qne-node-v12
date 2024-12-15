import { useFetch } from 'nuxt/app'

export const useQuickNEasy = () => {
  return useFetch('/api/quick-n-easy')
}