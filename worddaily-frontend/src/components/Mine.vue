<template>
  <div class="mine-container">
    <div class="user-info">
      <h2 @click="switchUser">{{ username }}</h2>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '用户1'
    };
  },
  methods: {
    switchUser() {
      // 切换用户的逻辑
    }
  }
};
</script>

<style scoped>
.mine-container {
  padding: 20px;
}

.user-info h2 {
  cursor: pointer;
}
</style>