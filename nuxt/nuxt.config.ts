// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },

  runtimeConfig: {
    // The private keys which are only available within server-side
    redisHost: 'localhost',
    redisPort: '56379',
    apiBase: 'http://localhost:58080',
    public: {
      apiBase: 'http://localhost:58080'
    }
  },

  nitro: {
    plugins: ['plugins/logs.ts'],
  },

  routeRules: {
    "/cached": {
      swr: false,
      cache: {
        maxAge: 60,
        base: "redis"
      }
    },
    "/api/nitro": {
      cache: {
        maxAge: 60,
        base: "redis",
        swr: false
      }
    },
  },

  modules: ["@nuxt/ui", '@pinia/nuxt'],
})