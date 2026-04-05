import { createRouter, createWebHistory } from 'vue-router'

// 1. 使用懒加载导入核心组件 (减少首屏加载时间)
const Home = () => import('../views/Home.vue')
const ResumeParse = () => import('../views/ResumeParse.vue')
const JobMatch = () => import('../views/JobMatch.vue')
const CareerPath = () => import('../views/CareerPath.vue')
const Report = () => import('../views/Report.vue')

// 2.认证组件
const Auth = () => import('../views/auth/Auth.vue')

// 3. 404 页面
const NotFound = () => import('../views/NotFound.vue')

// 4. 定义路由表
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: '首页 - AI 职业规划智能体',
      keepAlive: true,
      transition: 'fade'
    }
  },
  {
    path: '/resume',
    name: 'Resume',
    component: ResumeParse,
    meta: {
      title: '简历智能解析',
      requiresAuth: false,
      transition: 'slide-left'
    }
  },
  {
    path: '/match',
    name: 'Match',
    component: JobMatch,
    meta: {
      title: '人岗匹配分析',
      requiresAuth: false,
      transition: 'slide-left'
    }
  },
  {
    path: '/path',
    name: 'Path',
    component: CareerPath,
    meta: {
      title: '职业发展路径',
      requiresAuth: false,
      transition: 'slide-left'
    }
  },
  {
    path: '/report',
    name: 'Report',
    component: Report,
    meta: {
      title: '完整发展报告',
      requiresAuth: false,
      transition: 'zoom-in'
    }
  },
  // ====== Auth 路由 ======
  {
    path: '/auth/login',
    name: 'Auth',
    component: Auth,
    meta: {
      title: '用户登录与注册',
      transition: 'fade'
    }
  },
  // 404 路由 (必须放在最后)
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue'),
    meta: { title: '404 - 页面不存在' }
  }
]

// 5. 创建 Router 实例 
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else if (to.hash) {
      return { el: to.hash, behavior: 'smooth' }
    } else {
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve({ top: 0 })
        }, 300)
      })
    }
  }
})

// 6. 全局前置守卫
router.beforeEach((to, from) => {
  // 设置页面标题
  document.title = to.meta.title || 'AI 职业规划智能体'

  // 这里预留未来的登录拦截逻辑
  // const token = localStorage.getItem('token')
  // const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  // if (requiresAuth && !token) {
  //   next('/auth/login')
  // } else {
  //   next()
  // }

  // 直接返回 true 代表放行
  return true
})

// 7. 导出路由
export default router