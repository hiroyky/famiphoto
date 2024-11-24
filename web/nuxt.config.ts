// https://nuxt.com/docs/api/configuration/nuxt-config
import vuetify, {transformAssetUrls} from "vite-plugin-vuetify";

const runtimeConfig = {
  isDebug: process.env.IS_DEBUG === 'true',
  apiBaseUrl: process.env.API_BASE_URL,
  clientId: "famiphoto_web",
  clientSecret: process.env.CLIENT_SECRET,
  sessionSecret: process.env.SESSION_SECRET,
  public: {
    baseUrl: process.env.PUBLIC_BASE_URL,
    fallbackLocale: process.env.PUBLIC_FALLBACK_LOCALE ? process.env.PUBLIC_FALLBACK_LOCALE :'ja',
  }
}

export default defineNuxtConfig({
  app: {
    head: {
      charset: "utf-8",
      viewport: "width=device-width, initial-scale=1",
    }
  },
  runtimeConfig,
  serverHandlers: [
    { route: '/api/**', handler: '~/server/interfaces/http/app.ts' },
    { route: '/debug/**', handler: '~/server/interfaces/http/app.ts' }
  ],
  devtools: { enabled: true },
  experimental: {
    // GraphQLと通信する際にfalseにしないとエラーに
    renderJsonPayloads: false
  },
  build: {
    transpile: ['vuetify'],
  },
  modules: [
      '@pinia/nuxt',
    (_, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', (config) => {
        // @ts-expect-error
        config.plugins.push(vuetify({autoImport: true}))
      })
    },
      '@nuxtjs/i18n',
  ],
  vite: {
    vue: {
      template: {
        transformAssetUrls,
      }
    }
  },
  pinia: {
    storesDirs: ['./stores/**'],
  },
  i18n: {
    detectBrowserLanguage: {
      useCookie: true,
    },
    locales: [
        {code:'en', name: 'English'},
        {code:'ja', name:'日本語'}
    ],
    defaultLocale: 'en',
    strategy: 'no_prefix',
  }
})
