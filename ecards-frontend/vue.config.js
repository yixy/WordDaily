module.exports = {
  devServer: {
    proxy: {
      "/api": {
        // 代理前缀，可以根据需要修改
        target: "http://localhost:8081", // 后端服务器地址
        changeOrigin: true, // 允许跨域
        //默认情况下，http-proxy-middleware（Vue CLI 使用的代理工具）不会自动解析请求体（body）。因此，req.body 和 res.body 是 undefined。
        onProxyReq(proxyReq, req, res) {
          console.log(
            "[req Request]=============\n",
            req.method,
            req.url,
            req.headers,
            req.body,
            "\n[res Response]=============\n",
            res.statusCode,
            res.headers,
            res.body,
            "[Proxy Request]=============\n",
            proxyReq.method,
            proxyReq.url,
            proxyReq.headers,
            proxyReq.body
          );
        },
        onProxyRes(proxyRes, req, res) {
          console.log(
            "[Proxy Response]============\n",
            proxyRes.statusCode,
            proxyRes.url,
            proxyRes.headers,
            proxyRes.body,
            "[request]============\n",
            req.method,
            req.url,
            req.headers,
            req.body,
            "[response]============\n",
            res.statusCode,
            res.url,
            res.headers,
            res.body
          );
        },
        // pathRewrite: {
        //   "^/api": "", // 重写路径，去掉 `/api` 前缀
        // },
      },
    },
  },
};
