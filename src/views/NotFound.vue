<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
    <div class="text-center max-w-md">
      <!-- 404 图标 -->
      <div class="mb-8 relative inline-block">
        <div class="absolute inset-0 bg-teal-100 rounded-full blur-xl opacity-50 animate-pulse"></div>
        <el-icon :size="80" class="text-teal-500 relative z-10"><WarningFilled /></el-icon>
      </div>
      
      <h1 class="text-6xl font-extrabold text-gray-800 mb-4 tracking-tight">404</h1>
      <p class="text-xl text-gray-600 mb-2">哎呀，页面走丢了</p>
      <p class="text-gray-400 mb-8">您访问的页面不存在或已被移除</p>
      
      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <el-button type="primary" round size="large" @click="goHome" class="shadow-lg shadow-teal-200">
          返回首页
        </el-button>
        <el-button plain round size="large" @click="goBack">
          返回上一页
        </el-button>
      </div>
      
      <div class="mt-12 pt-8 border-t border-gray-200">
        <p class="text-xs text-gray-400">
          如果是演示模式中断，请尝试 <span class="text-teal-600 cursor-pointer hover:underline" @click="restartDemo">重新启动演示</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { WarningFilled } from '@element-plus/icons-vue'
import DemoTourManager from '@/utils/demoTour'

const router = useRouter()

const goHome = () => {
  router.push('/')
}

const goBack = () => {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push('/')
  }
}

const restartDemo = () => {
  const tour = new DemoTourManager()
  tour.start()
}
</script>

<style scoped>
/*添加一些简单的进入动画 */
.v-enter-active,
.v-leave-active {
  transition: opacity 0.3s ease;
}
.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>