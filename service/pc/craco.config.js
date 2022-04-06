// https://ant.design/docs/react/use-with-create-react-app-cn
// https://github.com/gsoft-inc/craco
const CracoLessPlugin = require('craco-less');

module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:8051',
                changeOrigin: true,
                pathRewrite: {
                    '^/api': '/api'
                }
            }
        }
    },
    plugins: [
        {
            plugin: CracoLessPlugin,
            options: {
                lessLoaderOptions: {
                    lessOptions: {
                        modifyVars: { '@primary-color': '#1DA57A' },
                        javascriptEnabled: true,
                    },
                },
            },
        },
    ],
};
