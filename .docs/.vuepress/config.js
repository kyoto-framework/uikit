
module.exports = {
    title: 'kyoto-uikit',
    description: 'Wake up, Samurai. We have a project to build',

    base: '/',

    themeConfig: {
        navbar: [
            {
                text: 'Tailwind UI',
                link: '/tailwindui'
            },
            {
                text: 'Material',
                link: '/material'
            }
        ],
    },

    plugins: [
        ['@vuepress/plugin-search', {
            locales: {
                '/': {
                    placeholder: 'Search'
                }
            }
        }]
    ],
}