import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/admin'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/admin',
  name: 'Admin',
  component: () => import('@/view/login/index.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
