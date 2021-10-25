
module.exports = {
    title: 'kyoto-uikit',
    description: 'Wake up, Samurai. We have a project to build',

    base: '/',

    themeConfig: {
        navbar: [
            {
                text: 'General',
                link: '/general'
            },
            {
                text: 'Tailwind UI',
                link: '/tailwind-ui/'
            },
            {
                text: 'Material',
                link: '/material/'
            }
        ],
        sidebar: {
            '/tailwind-ui/': [
                {
                    text: 'Application UI',
                    link: '/tailwind-ui/application-ui/',
                    children: [
                        {
                            text: 'Navigation',
                            link: '/tailwind-ui/application-ui/navigation/',
                            children: [
                                {
                                    text: '· Navbar',
                                    link: '/tailwind-ui/application-ui/navigation/navbar'
                                },
                                {
                                    text: '· Sidebar',
                                    link: '/tailwind-ui/application-ui/navigation/sidebar'
                                },
                                {
                                    text: '· Breadcrumbs',
                                    link: '/tailwind-ui/application-ui/navigation/breadcrumbs'
                                },
                                {
                                    text: '· Pagination',
                                    link: '/tailwind-ui/application-ui/navigation/pagination'
                                }
                            ]
                        },
                        {
                            text: 'Heading',
                            link: '/tailwind-ui/application-ui/heading/',
                            children: [
                                {
                                    text: '· Page',
                                    link: '/tailwind-ui/application-ui/heading/page'
                                },
                            ]
                        },
                        {
                            text: 'Data',
                            link: '/tailwind-ui/application-ui/data/',
                            children: [
                                {
                                    text: '· Stats',
                                    link: '/tailwind-ui/application-ui/data/stats'
                                },
                            ]
                        },
                        {
                            text: 'Form',
                            link: '/tailwind-ui/application-ui/form/',
                            children: [
                                {
                                    text: '· Layout',
                                    link: '/tailwind-ui/application-ui/form/layout'
                                },
                            ]
                        },
                        {
                            text: 'List',
                            link: '/tailwind-ui/application-ui/list/',
                            children: [
                                {
                                    text: '· Table',
                                    link: '/tailwind-ui/application-ui/list/table'
                                },
                            ]
                        },
                        {
                            text: 'Overlay',
                            link: '/tailwind-ui/application-ui/overlay/',
                            children: [
                                {
                                    text: '· Modal',
                                    link: '/tailwind-ui/application-ui/overlay/modal'
                                },
                                {
                                    text: '· Side-over',
                                    link: '/tailwind-ui/application-ui/overlay/side-over'
                                },
                            ]
                        }
                    ]
                },
                {
                    text: 'Marketing',
                    link: '/tailwind-ui/marketing/'
                },
                {
                    text: 'E-Commerce',
                    link: '/tailwind-ui/e-commerce/'
                },
            ]
        }
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