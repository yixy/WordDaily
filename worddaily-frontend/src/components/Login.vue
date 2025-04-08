<template>
  <div class="login-container">
    <div class="login-form">
      <h2>登录</h2>
      <div class="input-group">
        <label for="username">用户名</label>
        <input
          type="text"
          id="username"
          v-model="username"
          placeholder="请输入用户名"
        />
      </div>
      <div class="input-group">
        <label for="password">密码</label>
        <input
          type="password"
          id="password"
          v-model="password"
          placeholder="请输入密码"
        />
      </div>
      <button class="login-button" @click="login">登录</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    async login() {
      if (this.username && this.password) {
        try {
          console.log("Sending login request...");
          const response = await this.$http.post("/api/login", {
            username: this.username,
            password: this.password,
          });
          console.log("Login response received:", response);
          if (response.data.success) {
            sessionStorage.setItem("isLoggedIn", "true");
            sessionStorage.setItem("authToken", response.data.authToken);
            this.$router.push("/punch");
          }
        } catch (error) {
          console.error("Login request failed:", error);
          alert("登录请求失败，请稍后重试：" + error.message);
        }
      } else {
        alert("请输入用户名和密码");
      }
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-form {
  width: 300px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.input-group {
  margin-bottom: 15px;
}

.input-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.input-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.login-button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.login-button:hover {
  background-color: #0056b3;
}
</style>
