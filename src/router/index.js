import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../layout/Index.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Auth.vue') //认证页面包含登录注册等
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/Dashboard.vue')
      },
      {
        path: 'profile',
        name: 'ProfileBuilder',
        component: () => import('../views/ProfileBuilder.vue')
      },
      {
        path: 'jobs',
        name: 'JobGraph',
        component: () => import('../views/JobGraph.vue')
      },
      {
        path: 'report',
        name: 'CareerReport',
        component: () => import('../views/CareerReport.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

//路由守卫示例，后续在此接入Token验证
//router.beforeEach((to,from,next)=>{...})

export default router