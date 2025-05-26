import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },
    css: ['~/assets/css/tailwind.css'],
    postcss: {
        plugins: {
            cssnano: { preset: 'default' },
            autoprefixer: {},
        },
    },
    vite: {
        plugins: [
            tailwindcss(),
        ],
    },
})
