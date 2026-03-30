<template>
  <div class="min-h-screen bg-gradient-to-br from-teal-50 via-white to-blue-50 flex flex-col relative overflow-hidden">

    <!-- 背景装饰光斑 -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <!-- 左上角光斑-->
      <div class="absolute -top-40 -right-40 w-96 h-96 bg-teal-100/40 rounded-full mix-blend-multiply filter blur-3xl opacity-70 animate-blob"></div>
      <!-- 右下角光斑-->
      <div class="absolute -bottom-40 -left-40 w-96 h-96 bg-blue-100/40 rounded-full mix-blend-multiply filter blur-3xl opacity-70 animate-blob animation-delay-2000"></div>
      <!-- 新增第三个光斑  -->
      <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-purple-100/30 rounded-full mix-blend-multiply filter blur-3xl opacity-50 animate-blob animation-delay-4000"></div>
    </div>

    <!-- 顶部Logo区域 - 整体上移 -->
    <div class="relative z-10 pt-8 pb-4">
      <div class="text-center">
        <div class="inline-flex items-center justify-center w-14 h-14 rounded-2xl bg-white shadow-lg mb-2">
          <el-icon :size="28" class="text-teal-600"><UserFilled /></el-icon>
        </div>  
        <h1 class="text-xl font-bold text-gray-800">CareerAgent</h1>
        <p class="text-gray-500 text-xs mt-0.5">AI 驱动的职业规划平台</p>
      </div>
    </div>

    <!-- 主内容卡片容器 -->
    <div class="relative flex-1 flex items-start justify-center px-4 pb-4 z-10 pt-4">
      <div class="w-full max-w-md">
        <!-- 卡片区域  -->
        <div class="bg-white/80 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/60 p-6 md:p-8 transition-all duration-300 relative z-10 min-h-[400px]">

          <!-- 关闭按钮  -->
          <div class="absolute top-3 right-3 cursor-pointer text-gray-400 hover:text-teal-600 transition-colors z-20" @click="handleClose">
            <el-icon :size="18"><Close /></el-icon>
          </div>

          <!-- 登录、注册表单渲染在此 - 使用router-view渲染子路由 -->
          <router-view></router-view>

        </div>

        <!-- 底部区域  -->
        <div class="text-center mt-4 text-xs text-ref-mid-blue/60">
          <p>© 2026 CareerAgent. All rights reserved.</p>
          <!-- 添加政策链接 -->
          <div class="mt-1 space-x-3 text-gray-500">
            <a href="#" class="hover:text-teal-600 transition">隐私政策</a>
            <a href="#" class="hover:text-teal-600 transition">服务条款</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Close, UserFilled } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleClose = () => {
  router.push('/')
}
</script>

<style scoped>
@keyframes blob {
  0%, 100% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
}
.animate-blob {
  animation: blob 7s infinite;
}
.animation-delay-2000 {
  animation-delay: 2s;
}
.animation-delay-4000 {
  animation-delay: 4s;
}
</style>
