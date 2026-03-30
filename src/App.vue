<template>
  <!-- 主应用布局 -->
  <div v-if="!isAuthPage" class="h-screen w-screen flex overflow-hidden bg-brand-cream/30 font-sans text-brand-text relative">
    <!-- 侧边栏 -->
    <aside :class="isSidebarCollapsed ? 'w-20' : 'w-64'" class="bg-brand-cream flex flex-col transition-all duration-300 ease-in-out z-20 shadow-xl shadow-brand-green/10 border-r border-white/50 shrink-0">
      <!-- 侧边栏标题 -->
      <div class="h-16 flex items-center justify-center cursor-pointer border-b border-brand-green/20 transition-all overflow-hidden" 
           :class="isSidebarCollapsed ? 'px-0' : 'px-6 gap-3'" 
           @click="$router.push('/')">
        <div class="w-9 h-9 shrink-0 rounded-xl bg-white flex items-center justify-center shadow-sm shadow-brand-green/40 text-brand-green">
          <el-icon :size="22"><Compass /></el-icon>
        </div>
        <span v-show="!isSidebarCollapsed" class="text-xl font-extrabold tracking-wider text-brand-dark whitespace-nowrap">CareerAgent</span>
      </div>
<!-- 导航栏 -->
      <nav class="flex-1 py-6 px-3 space-y-2 overflow-y-auto custom-scrollbar overflow-x-hidden">
        <router-link v-for="item in navItems" :key="item.path" :to="item.path"
          class="flex items-center rounded-2xl text-sm font-bold transition-all duration-300 group relative overflow-hidden"
          :class="[ 
            route.path === item.path ? 'bg-white text-brand-green shadow-sm' : 'text-gray-500 hover:text-brand-dark hover:bg-white/50',
            isSidebarCollapsed ? 'justify-center py-3 px-0' : 'gap-3 px-4 py-3'
          ]"
          :title="isSidebarCollapsed ? item.name : ''"
        >
          <div v-if="route.path === item.path" class="absolute inset-0 bg-brand-pink/20 z-0"></div>
          <div v-if="route.path === item.path && !isSidebarCollapsed" class="absolute left-2 top-1/2 -translate-y-1/2 w-1.5 h-1.5 bg-brand-pink rounded-full z-10"></div>
          
          <el-icon :size="20" class="z-10 transition-transform group-hover:scale-110 shrink-0" :class="route.path === item.path ? 'text-brand-green' : 'text-gray-400 group-hover:text-brand-green'">
            <component :is="item.icon" />
          </el-icon>
          <span v-show="!isSidebarCollapsed" class="z-10 whitespace-nowrap">{{ item.name }}</span>
        </router-link>
      </nav>
        <!-- 用户信息/登录/注册 -->
      <div class="p-3 border-t border-brand-green/20 overflow-hidden">
        <div class="flex items-center bg-white/60 rounded-2xl hover:bg-white transition-colors cursor-pointer shadow-sm" 
             :class="isSidebarCollapsed ? 'justify-center py-3 px-0' : 'gap-3 px-4 py-3'"
             @click="userStore.isLoggedIn ? (showAvatarDialog = true) : $router.push('/auth/login')"
             :title="isSidebarCollapsed ? (userStore.isLoggedIn ? userStore.userName : '登录/注册') : ''">
          <!-- 用户头像 -->
          <div class="w-9 h-9 shrink-0 rounded-full bg-[#EFDCE2] flex items-center justify-center text-white font-bold shadow-inner overflow-hidden border border-white">
            <img v-if="userStore.isLoggedIn && userStore.userInfo?.avatar" :src="userStore.userInfo.avatar" class="w-full h-full object-cover" alt="avatar" />
            <el-icon v-else class="text-gray-500"><User /></el-icon>
          </div>
          <!-- 用户信息 -->
          <div v-show="!isSidebarCollapsed" class="flex flex-col overflow-hidden">
            <span class="text-sm font-bold text-brand-dark whitespace-nowrap">
              {{ userStore.isLoggedIn ? userStore.userName : '游客 / 未登录' }}
            </span>
            <span class="text-xs mt-0.5 whitespace-nowrap transition-colors"
                  :class="userStore.isLoggedIn ? 'text-gray-400' : 'text-[#8A9E58]'">
              {{ userStore.isLoggedIn ? 'CareerAgent 规划中' : '点击登录开启规划' }}
            </span>
          </div>
          <!-- 退出登录按钮 -->
          <div v-if="userStore.isLoggedIn && !isSidebarCollapsed" 
               class="ml-auto w-6 h-6 rounded-full hover:bg-red-50 flex items-center justify-center text-gray-400 hover:text-red-400 transition-colors"
               @click.stop="userStore.logout()" title="退出登录">
            <el-icon><Close /></el-icon>
          </div>

        </div>
      </div>
    </aside>
    <!-- 主内容区域 -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden relative">
      <!-- 主内容区域标题 -->
      <header class="h-16 bg-white/80 backdrop-blur-md flex justify-between items-center px-6 shadow-sm shadow-brand-green/5 z-10 shrink-0">
        <div class="flex items-center gap-4">
          <div class="w-8 h-8 rounded-full bg-gray-50 flex items-center justify-center cursor-pointer hover:bg-brand-cream/50 transition-colors text-gray-500" @click="toggleSidebar">
            <el-icon :size="18"><component :is="isSidebarCollapsed ? Expand : Fold" /></el-icon>
          </div>
          <!-- 当前路由名称 -->
          <div class="text-lg font-extrabold text-brand-dark tracking-tight flex items-center gap-2">
            <span class="w-2.5 h-2.5 bg-brand-green rounded-tl-full rounded-br-full inline-block"></span>
            {{ currentRouteName }}
          </div>
        </div>
        <!-- 搜索框 -->
        <div class="flex items-center gap-5">
          <div class="relative hidden md:block group">
            <el-icon class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-300 group-focus-within:text-brand-green transition-colors"><Search /></el-icon>
            <input type="text" placeholder="探索你的职业..." class="pl-9 pr-4 py-1.5 bg-brand-cream/30 border border-brand-cream rounded-full text-sm focus:outline-none focus:ring-2 focus:ring-brand-green/50 focus:bg-white w-64 transition-all text-brand-dark placeholder-gray-400">
          </div>
          <!-- 通知图标 -->
          <div class="w-9 h-9 rounded-full bg-brand-cream/50 flex items-center justify-center cursor-pointer hover:bg-brand-pink hover:text-white transition-colors relative text-gray-500">
            <el-icon :size="18"><Bell /></el-icon>
            <span class="absolute top-1 right-1 w-2 h-2 bg-brand-yellow rounded-full border border-white"></span>
          </div>
        </div>
      </header>
      <!-- 主内容区域内容 -->
      <main class="flex-1 overflow-hidden relative p-4 lg:p-6 box-border">
        <router-view v-slot="{ Component, route }">
          <transition :name="route.meta.transition || 'fade'" mode="out-in" appear>
            <component :is="Component" :key="route.fullPath" />
          </transition>
        </router-view>
      </main>
    </div>
    <!-- AI聊天 -->
        <AIChatDrawer />
  </div>

  <div v-else>
    <router-view v-slot="{ Component, route }">
      <transition :name="route.meta.transition || 'fade'" mode="out-in" appear>
        <component :is="Component" :key="route.fullPath" />
      </transition>
    </router-view>
  </div>
<!-- 头像上传弹窗 -->
  <el-dialog
      v-model="showAvatarDialog"
      title="更换头像"
      width="360px"
      class="avatar-dialog"
      destroy-on-close
      center
    >
      <div class="flex flex-col items-center justify-center py-4">
        <el-upload
          class="avatar-uploader"
          action="/api/user/avatar" 
          name="avatar"
          :headers="{ Authorization: `Bearer ${userStore.token}` }"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
          :on-error="handleAvatarError"
        >
          <img v-if="uploadPreviewUrl || userStore.userInfo?.avatar" 
               :src="uploadPreviewUrl || userStore.userInfo.avatar" 
               class="w-32 h-32 rounded-full object-cover shadow-md border-4 border-[#EFDCE2]/30" />
          
          <div v-else class="w-32 h-32 rounded-full border-2 border-dashed border-[#C2D68F] bg-[#F7EECD]/30 flex flex-col items-center justify-center text-[#8A9E58] hover:bg-[#C2D68F]/10 transition-colors">
            <el-icon :size="28" class="mb-1"><Plus /></el-icon>
            <span class="text-xs font-bold">点击上传</span>
          </div>
        </el-upload>
        <p class="text-xs text-gray-400 mt-6 text-center">支持 JPG/PNG 格式，图片大小不超过 2MB</p>
      </div>
    </el-dialog>
</template>

<script setup>
import { useUserStore } from '@/stores/user'
import { 
  Compass, DataBoard, Document, Position, MapLocation, 
  DataAnalysis, User, Search, Bell, Plus,Close,
   Fold, Expand // 引入折叠/展开图标
} from '@element-plus/icons-vue'
import {ElMessage} from 'element-plus'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
// 导入AI 抽屉
import AIChatDrawer from '@/components/ai/AIChatDrawer.vue'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()

// 控制侧边栏状态
const isSidebarCollapsed = ref(false)
const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value
  // 触发一次 window resize 事件，让图表重新自适应宽度
  setTimeout(() => window.dispatchEvent(new Event('resize')), 300)
}
// 是否为认证页面
const isAuthPage = computed(() => route.path.startsWith('/auth/'))
// 导航栏项
const navItems = [
  { name: '分析概览', path: '/', icon: DataBoard },
  { name: '简历解析', path: '/resume', icon: Document },
  { name: '人岗匹配', path: '/match', icon: Position },
  { name: '职业路径', path: '/path', icon: MapLocation },
  { name: '发展报告', path: '/report', icon: DataAnalysis },
]
// 当前路由名称
const currentRouteName = computed(() => {
  const current = navItems.find(item => item.path === route.path)
  return current ? current.name : '工作台'
})
// 头像上传弹窗状态
const showAvatarDialog = ref(false)
const uploadPreviewUrl = ref('')
// 上传前的校验逻辑（格式、大小等）
const beforeAvatarUpload = (rawFile) => {
  const isJPGOrPNG = rawFile.type === 'image/jpeg' || rawFile.type === 'image/png'|| rawFile.type === 'image/webp'
  if (!isJPGOrPNG) {
    ElMessage.error('请上传 JPG/PNG 格式的图片')
    return false
  }
  const isLt2MB = rawFile.size / 1024 / 1024 < 2
  if (!isLt2MB) {
    ElMessage.error('图片大小不能超过 2MB')
    return false
  }
  return true
}
// 上传成功后的逻辑处理
const handleAvatarSuccess = (response, updateFile) => {
  // 假设后端返回的数据结构是 { code: 200, data: { avatarUrl: '...' }, message: '...' }
  // 这里需要根据你 Apifox 实际的返回结构（比如 response.data.url）进行修改！
  const newAvatarUrl = response.data?.avatarUrl || response.data?.url || URL.createObjectURL(uploadFile.raw)
  
  // 更新本地预览图
  uploadPreviewUrl.value = newAvatarUrl
  
  // 更新全局 Store 中的用户信息，触发左下角头像实时刷新！
  userStore.userInfo.avatar = newAvatarUrl
  userStore.setUserInfo(userStore.userInfo) // 重新存入 localStorage 以持久化
  
  ElMessage.success('头像修改成功！')
  setTimeout(() => {
    showAvatarDialog.value = false // 延迟关闭弹窗，体验更平滑
  }, 800)
}
// 上传失败后的逻辑处理
const handleAvatarError = (err) => {
  ElMessage.error('头像上传失败，请检查网络连接或后端接口是否开启')
  console.error(err)
}
</script>

<style scoped>
/* 页面切换动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(5px);
}

/* 聊天框弹出动画 */
.chat-slide-enter-active,
.chat-slide-leave-active {
  transition: opacity 0.3s cubic-bezier(0.16, 1, 0.3, 1), transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  transform-origin: bottom right;
}
.chat-slide-enter-from,
.chat-slide-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(20px);
}

/* 柔和的滚动条 */
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: rgba(194, 214, 143, 0.5); 
  border-radius: 4px;
}
.custom-scrollbar:hover::-webkit-scrollbar-thumb {
  background-color: rgba(194, 214, 143, 0.8);
}
</style>