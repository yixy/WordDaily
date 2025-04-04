module.exports = {
    devServer: {
      proxy: {
        '/api': { // 代理前缀，可以根据需要修改
          target: 'http://localhost:8081', // 后端服务器地址
          changeOrigin: true, // 允许跨域
          pathRewrite: {
            '^/api': '', // 重写路径，去掉 `/api` 前缀
          },
        },
      },
    },
  };