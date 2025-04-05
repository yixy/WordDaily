import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";

const app = createApp(App);

const http = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL || "/api",
});

http.interceptors.request.use(
  (config) => {
    const token = sessionStorage.getItem("authToken");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

http.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      sessionStorage.removeItem("isLoggedIn");
      sessionStorage.removeItem("authToken");
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

// 确保axios实例正确挂载到Vue实例上
app.config.globalProperties.$http = http;

app.use(router);
app.mount("#app");

// 添加全局错误处理
window.onerror = function (message, source, lineno, colno, error) {
  console.error("Global error:", message, source, lineno, colno, error);
};
