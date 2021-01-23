export default {
    // Global page headers (https://go.nuxtjs.dev/config-head)
    head: {
        title: 'tsundoku-nuxt',
        meta: [
            { charset: 'utf-8' },
            { name: 'viewport', content: 'width=device-width, initial-scale=1' },
            { hid: 'description', name: 'description', content: '' }
        ],
        link: [
            { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
        ]
    },

    // Global CSS (https://go.nuxtjs.dev/config-css)
    css: [
    ],

    // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
    plugins: [
        '~/plugins/cookies-to-state',
        '~/plugins/axios'
    ],

    // Auto import components (https://go.nuxtjs.dev/config-components)
    components: true,

    // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
    buildModules: [
    // https://go.nuxtjs.dev/eslint
        '@nuxtjs/eslint-module'
    ],

    // Modules (https://go.nuxtjs.dev/config-modules)
    modules: [
    // https://go.nuxtjs.dev/bootstrap
        '@nuxtjs/bulma',
        // https://go.nuxtjs.dev/axios
        '@nuxtjs/axios',
        'cookie-universal-nuxt',
        '@nuxtjs/dotenv'
    ],

    // Axios module configuration (https://go.nuxtjs.dev/config-axios)
    axios: {
        baseURL: 'http://localhost:8080/'
    },

    // Build Configuration (https://go.nuxtjs.dev/config-build)
    build: {
        postcss: {
            preset: {
                features: {
                    customProperties: false
                }
            }
        },
        extend (config, ctx) {
            if (ctx.isDev && ctx.isClient) {
                config.devtool = 'inline-cheap-module-source-map';
            }
        }
    }
};