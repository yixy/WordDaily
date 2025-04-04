import { createRouter, createWebHistory } from 'vue-router';
import Punch from '@/components/Punch.vue';
import Diction from '@/components/Diction.vue';
import Mine from '@/components/Mine.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/punch',
      component: Punch
    },
    {
      path: '/diction',
      component: Diction
    },
    {
      path: '/mine',
      component: Mine
    },
    {
      path: '/',
      redirect: '/punch'
    }
  ]
});

export default router;