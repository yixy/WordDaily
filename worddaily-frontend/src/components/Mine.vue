<template>
  <div class="mine-container">
    <div class="user-info">
      <h2 @click="switchUser">{{ username }}</h2>
      <img
        :src="headshotUrl"
        alt="User Avatar"
        @click="changeHeadshot"
        class="user-avatar"
      />
    </div>
    <input
      type="file"
      @change="handleFileUpload"
      ref="fileInput"
      style="display: none"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "用户名",
      headshotUrl: "",
      file: null,
    };
  },
  methods: {
    async switchUser() {
      try {
        // 切换用户的逻辑
        if (confirm("确定要退出吗？")) {
          const response = await this.$http.post("/api/logout");
        }
      } catch (error) {}
    },
    changeHeadshot() {
      this.$refs.fileInput.click();
    },
    handleFileUpload(event) {
      this.file = event.target.files[0];
      if (this.file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          const base64 = e.target.result.split(",")[1]; // 获取 base64 编码部分
          this.updateHeadshot(base64);
        };
        reader.readAsDataURL(this.file);
      }
    },
    async updateHeadshot(base64) {
      try {
        const response = await this.$http.post(
          "/api/usermod",
          {
            username: this.username,
            headshot: base64,
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        if (response.data.success) {
          this.headshotUrl = `data:image/png;base64,${base64}`;
          alert("头像更新成功");
        }
      } catch (error) {
        alert(error.message);
      }
    },
  },
  async created() {
    try {
      // 假设从后端获取用户信息
      const response = await this.$http.post(
        "/api/userquery",
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
          : ""; // 修改默认值为空字符串，避免显示错误图片
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
  width: 250px;
  height: 250px;
  border-radius: 50%;
  cursor: pointer;
}
</style>
