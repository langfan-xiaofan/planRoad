<template>
  <div class="h-full flex items-center justify-center p-4">
    <div class="bg-white rounded-3xl p-10 shadow-sm border border-brand-cream/80 text-center max-w-md w-full relative overflow-hidden group">
      <div class="absolute -right-16 -top-16 w-48 h-48 bg-brand-pink/20 rounded-full blur-3xl transition-all duration-700 group-hover:bg-brand-pink/30"></div>
      <div class="absolute -left-16 -bottom-16 w-48 h-48 bg-brand-green/20 rounded-full blur-3xl transition-all duration-700 group-hover:bg-brand-green/30"></div>

      <div class="relative z-10 flex flex-col items-center">
        <div class="w-20 h-20 bg-brand-pink/10 rounded-full flex items-center justify-center mb-6 shadow-inner">
          <el-icon class="text-4xl text-brand-pink"><Warning /></el-icon>
        </div>
        
        <h1 class="text-7xl font-black text-brand-dark mb-2 tracking-tighter drop-shadow-sm">404</h1>
        <h2 class="text-xl font-extrabold text-brand-dark mb-3">哎呀，偏离了职业航线！</h2>
        
        <p class="text-sm text-gray-500 mb-8 leading-relaxed">
          您访问的页面就像那些石沉大海的简历一样，找不到了。<br>
          不如回到首页，让 AI 帮您重新规划航向吧。
        </p>
        
        <el-button
          round
          class="!bg-brand-green !text-white !border-none shadow-md shadow-brand-green/30 hover:-translate-y-0.5 transition-transform font-bold px-8 py-5 text-base"
          @click="$router.push('/')"
        >
          返回首页基地
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Warning } from '@element-plus/icons-vue'
</script>