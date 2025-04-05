import { createRouter, createWebHistory } from "vue-router";
import Punch from "@/components/Punch.vue";
import Diction from "@/components/Diction.vue";
import Mine from "@/components/Mine.vue";
import Login from "@/components/Login.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/punch",
      component: Punch,
      meta: { requiresAuth: true },
    },
    {
      path: "/diction",
      component: Diction,
      meta: { requiresAuth: true },
    },
    {
      path: "/mine",
      component: Mine,
      meta: { requiresAuth: true },
    },
    {
      path: "/login",
      component: Login,
    },
    {
      path: "/:pathMatch(.*)*",
      redirect: "/punch",
    },
  ],
});

// 全局前置守卫
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !isUserLoggedIn()) {
    next("/login");
  } else {
    next();
  }
});

// 模拟用户登录验证
function isUserLoggedIn() {
  // 这里可以根据实际需求实现用户登录验证逻辑
  // 例如检查localStorage或sessionStorage中的token
  if (sessionStorage.getItem("isLoggedIn")) {
    return true;
  }
  return false;
}

export default router;
