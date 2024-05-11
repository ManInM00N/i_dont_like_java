const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,

})
module.exports={
  devServer: {
    port: 7235,
    proxy: {
      '/apis': {
        target: 'http://127.0.0.1:7234',
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          '^/apis': ''
        }
      }
    }
  }
}