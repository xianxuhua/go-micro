// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    ssr: false,
    vite: {
        vue: {
            reactivityTransform: true
        }
    },
    typescript: {
        strict: true
    }
})
