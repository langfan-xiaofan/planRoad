<!-- 核心布局与全局AI助手 (SaaS模式布局) -->
<template>
  <el-container class="h-screen bg-themeBg overflow-hidden">
    
    <el-header class="bg-white shadow-sm flex items-center justify-between px-6 h-16 border-b border-gray-100 z-50 relative">
      <div class="flex items-center gap-6 flex-1">
        
        <div class="h-16 flex items-center justify-center font-bold text-xl text-themePrimary border-r border-gray-100 pr-6 gap-2">
          <div 
            class="w-8 h-8 bg-themePrimary" 
            style="-webkit-mask: url('/src/assets/logo.svg') no-repeat center / contain; mask: url('/src/assets/logo.svg') no-repeat center / contain;"
          ></div>
          <span>AI生涯规划助手</span>
        </div>
        <!-- 顶部导航菜单 -->
        <el-menu
          :default-active="activeMenu"
          router
          mode="horizontal"
          class="border-none flex-1 mt-0 h-16 !bg-transparent"
        >
          <el-menu-item index="/dashboard" class="h-16 flex items-center justify-center text-sm">
            <span>工作台总览</span>
          </el-menu-item>
          <el-menu-item index="/profile" class="h-16 flex items-center justify-center text-sm">
            <span>能力画像评测</span>
          </el-menu-item>
          <el-menu-item index="/jobs" class="h-16 flex items-center justify-center text-sm">
            <span>岗位图谱探索</span>
          </el-menu-item>
          <el-menu-item index="/report" class="h-16 flex items-center justify-center text-sm">
            <span>AI生涯报告</span>
          </el-menu-item>
        </el-menu>

      </div> 
      <!-- 用户信息与退出登录按钮 -->
      <div class="flex items-center gap-4 border-l border-gray-100 pl-6 h-16">
        <el-avatar size="small" src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"></el-avatar>
        <span class="text-sm text-gray-700">admin</span>
        <el-link :underline="false" class="text-themeAccent text-sm" @click="handleLogout">退出</el-link>
      </div>
      
    </el-header>
    <!-- 主内容区域 -->
    <el-container class="relative flex-1 overflow-hidden">
      <el-main class="p-6 overflow-hidden flex-1 h-full flex flex-col">
      <router-view></router-view>
      </el-main>
      <!-- 全局AI助手按钮 -->
      <div class="fixed bottom-8 right-8 z-50">
        <button @click="aiDrawerVisible = true" class="w-14 h-14 rounded-full bg-themePrimary text-white shadow-lg flex items-center justify-center hover:bg-themeSecondary transition-colors duration-300">
          AI
        </button>
      </div>
      <!-- 全局AI助手抽屉内容 -->
      <el-drawer v-model="aiDrawerVisible" title="全局AI规划师" size="400px" direction="rtl">
        <div class="flex flex-col h-full">
          <el-scrollbar ref="scrollbarRef" class="flex-1 p-4 bg-gray-50 rounded-lg">
            <!--  whitespace-pre-wrap leading-relaxed 让后端传回来的换行符 \n 正确显示 -->
            <div v-for="(msg, index) in chatHistory" :key="index" :class="['mb-4 p-3 rounded-lg max-w-[85%] whitespace-pre-wrap leading-relaxed', msg.role === 'user' ? 'bg-themeSecondary text-gray-800 self-end ml-auto' : 'bg-white border border-gray-200 text-gray-700']">
              {{ msg.content }}
            </div>
          </el-scrollbar>
          
          <div class="mt-4 flex gap-2 shrink-0">
            <el-input 
              v-model="chatInput" 
              placeholder="向AI提问，例如：什么是UI设计师？" 
              @keyup.enter.native="sendMessage"
              :disabled="isAiTyping"
            ></el-input>
            <el-button 
              type="primary" 
              class="bg-themePrimary border-none w-20" 
              @click="sendMessage"
              :loading="isAiTyping"
            >
              {{ isAiTyping ? '思考中' : '发送' }}
            </el-button>
          </div>
        </div>
      </el-drawer>
    </el-container>
    
  </el-container>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const activeMenu = computed(() => route.path)

//AI助手状态控制
const aiDrawerVisible = ref(false)
const chatInput = ref('')
const chatHistory = ref([
  { role: 'ai', content: '你好！我是你的专属职业规划AI。在浏览系统时遇到任何岗位词汇或规划问题，都可以随时问我。' }
])

const scrollbarRef = ref(null) 
const isAiTyping = ref(false) // 控制按钮加载状态和并发拦截

// 自动滚动到底部的方法
const scrollToBottom = async () => {
  await nextTick() // 等待 Vue 完成 DOM 更新
  if (scrollbarRef.value) {
    // 设置一个极大的值，使原生滚动条触底
    scrollbarRef.value.setScrollTop(999999) 
  }
}

//  SSE 发送方法
const sendMessage = async () => {
  const userText = chatInput.value.trim()
  if (!userText || isAiTyping.value) return

  // 1. 将用户消息推入列表并清空输入框
  chatHistory.value.push({ role: 'user', content: userText })
  chatInput.value = ''
  isAiTyping.value = true

  // 2. 提前为 AI 创建一个空的消息气泡，后续将不断向里面追加文字
  chatHistory.value.push({ role: 'ai', content: '' })
  const aiMessageIndex = chatHistory.value.length - 1
  scrollToBottom()

  try {
    //构建表单数据
    const formData = new FormData()
    formData.append('message', userText) // 传入当前发送的单条文本消息
    formData.append('user_id', '0')      // 目前写死为0 后续替换为真实的登录用户ID
    // formData.append('file', fileObject) // 在这个组件加了上传文件的按钮(后面根据实际情况)

    //TODO: 待替换为后端的API地址
    const response = await fetch('http://localhost:8080/api/ai/chat', { 
      method: 'POST',
      headers: {
        // 'Content-Type': 'application/json',
        // 'Authorization': 'Bearer ' + localStorage.getItem('token') // 有鉴权的话解开这个注释（待改） 
      },
      // 把历史记录一起发给大模型 让它有上下文记忆
      // 
      body: formData
      })
    

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    // 3. 获取数据流并创建解码器
    const reader = response.body.getReader()
    const decoder = new TextDecoder('utf-8')
    let buffer = '' // 暂存被截断的不完整数据块

    // 4. 循环读取流数据
    while (true) {
      const { done, value } = await reader.read()
      if (done) break

      // 将二进制数据解码为字符串并拼接进 buffer
      buffer += decoder.decode(value, { stream: true })
      
      // 以 \n\n 结尾，但也可能单行返回，按换行符分割
      const lines = buffer.split('\n')
      
      // 弹出最后一行（可能是不完整的数据），留到下一次循环处理
      buffer = lines.pop() || ''

      for (const line of lines) {
        const trimmedLine = line.trim()
        if (!trimmedLine) continue

        // 拦截后端发送的结束标记，通常是 [DONE]
        if (trimmedLine === 'data: [DONE]') return 

        // 解析标准的 SSE 数据行: "data: {"content": "xxx"}"
        if (trimmedLine.startsWith('data:')) {
          const dataStr = trimmedLine.replace('data:', '').trim()
          try {
            const dataObj = JSON.parse(dataStr)
            
            // TODO: 根据后端实际返回的 JSON 结构提取字符串
            // 假设后端直接返回 { "content": "xxx" }
            const chunkText = dataObj.content || ''          //.content待改成后端返回的字段名
            
            // 将追加的文字拼接到最后一条 AI 消息中
            chatHistory.value[aiMessageIndex].content += chunkText
            
            // 每次追加文字都让滚动条向下滚一点
            scrollToBottom()
          } catch (e) {
            console.warn('JSON 解析失败，可能是数据块截断:', dataStr)
          }
        }
      }
    }
  } catch (error) {
    console.error('SSE 请求失败:', error)
    chatHistory.value[aiMessageIndex].content = '抱歉，网络连接异常或服务器未响应，请稍后再试。'
  } finally {
    // 无论成功失败，重置打字状态
    isAiTyping.value = false
  }
}

const handleLogout = () => {
  // TODO: 对接后端登出接口，清空Token并跳转
}
</script>