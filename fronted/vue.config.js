const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,

})
module.exports={
  devServer: {
    port: 7235, // 自定义端口
    proxy: {//设置跨域请求
      '/apis': {   //只要是api开头的就会转到后端
        target: 'http://localhost:7234', // 后端监听的地址
        changeOrigin: true,
        pathRewrite: {
          '^/apis': ''
        }
      }
    }
  }
}