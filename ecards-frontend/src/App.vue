<template>
  <div id="app">
    <div class="status-bar"></div>
    <div class="content">
      <router-view />
    </div>
    <div class="nav-bar" v-if="$route.path !== '/login'">
      <router-link to="/punch">打卡</router-link>
      <router-link to="/diction">词库</router-link>
      <router-link to="/mine">我的</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App'
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  max-width: 375px;
  margin: 0 auto;
  border-radius: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.status-bar {
  height: 20px;
  background-color: #f8f9fa;
}

.content {
  height: calc(100vh - 80px);
  overflow-y: auto;
}

.nav-bar {
  display: flex;
  justify-content: space-around;
  padding: 10px;
  background-color: #f8f9fa;
  border-top: 1px solid #ddd;
}

.nav-bar a {
  text-decoration: none;
  color: #2c3e50;
}

.nav-bar a.router-link-exact-active {
  color: #007bff;
}
</style>