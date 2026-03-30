import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', {
  // 1. State: 定义数据源，并直接从 localStorage 初始化（实现持久化，刷新不丢失）
  state: () => ({
    token: localStorage.getItem('token') || '',
    // 用户详细信息（包含 username, email 等）
    userInfo: JSON.parse(localStorage.getItem('userInfo') || '{}'),
    // AI 对话收集的规划画像数据
    onboardingData: JSON.parse(localStorage.getItem('onboardingData') || '{}'),
  }),

  // 2. Getters: 类似于组件的 computed，用于派生出所需的状态
  getters: {
    // 是否已登录：只要有 token 就认为已登录
    isLoggedIn: (state) => !!state.token,
    // 获取用户名，优先读取 userInfo 中的 username，没有则返回 '访客'
    userName: (state) => state.userInfo.userInfo?.username || '访客',
    // 是否完成了初始化的 AI 对话画像（Home.vue 用于判断显示哪个视图）
    hasCompletedOnboarding: (state) => Object.keys(state.onboardingData).length > 0,
  },

  // 3. Actions: 类似于组件的 methods，用于修改 State 和执行业务逻辑
  actions: {
    // 保存token
    setToken(token) {
      this.token = token
      // 持久化到 localStorage
      localStorage.setItem('token', token)
    },
    // 保存用户信息
    setUserInfo(info) {
      this.userInfo = info
      // 持久化到 localStorage
      localStorage.setItem('userInfo', JSON.stringify(info))
    },
    // 核心：处理登录成功后的逻辑
    login(token, user) {
      this.setToken(token)
      this.setUserInfo(user)
    },
    // 核心：退出登录，清空所有状态和本地存储
    logout() {
      this.setToken('')
      this.userInfo = {}
      this.onboardingData = {}
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      localStorage.removeItem('onboardingData')
    },
    // 保存 AI 问卷数据 (供 AIChatDrawer.vue 使用)
    setOnboardingData(data) {
      this.onboardingData = data
      // 持久化到 localStorage
      localStorage.setItem('onboardingData', JSON.stringify(data))
    },
    // 获取 AI 问卷数据 (供 Home.vue 使用)
    getOnboardingData() {
      return this.onboardingData
    },
  }
}
)
