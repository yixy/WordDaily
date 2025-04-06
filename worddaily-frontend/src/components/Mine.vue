<template>
  <div class="mine-container">
    <div class="user-info">
      <h2 @click="switchUser">{{ username }}</h2>
      <img
        :src="headshotUrl"
        alt="User Avatar"
        @click="confirmLogout"
        class="user-avatar"
      />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "用户名",
      headshotUrl: "",
    };
  },
  methods: {
    switchUser() {
      // 切换用户的逻辑
    },
    confirmLogout() {
      if (confirm("确定要退出吗？")) {
        this.$router.push("/login");
      }
    },
  },
  async created() {
    try {
      // 假设从后端获取用户信息
      const response = await this.$http.post(
        "/api/user",
        {
          username: "seven",
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      if (response.data.success) {
        const user = response.data.user;
        this.username = user.Username;
        this.headshotUrl = user.Headshot
          ? `data:image/png;base64,${user.Headshot}`
          : "";
      }
    } catch (error) {
      alert(error.message);
    }
  },
};
</script>

<style scoped>
.mine-container {
  padding: 20px;
}

.user-info h2 {
  cursor: pointer;
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  cursor: pointer;
}
</style>
