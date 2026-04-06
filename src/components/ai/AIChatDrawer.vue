<template>
  <div 
    class="ai-fab group" 
    ref="fabRef"
    @mousedown.prevent="startDrag"
    @touchstart.passive="startDrag"
    :style="fabStyle"
    :class="{ 
      'has-unread': hasUnread, 
      'animate-bounce-slow': isIdle && hasUnread,
      'is-dragging': isDragging 
    }"
    title="AI 职业规划助手"
  >
    <div class="ai-fab-inner">
      <el-icon :size="28" class="text-white"><ChatDotRound /></el-icon>
      <span class="pulse-ring" v-if="isIdle"></span>
    </div>
    <span v-if="hasUnread" class="unread-badge"></span>
  </div>
 
  <el-drawer 
    v-model="isAiChatVisible" 
    direction="rtl" 
    size="420px"
    :with-header="false"
    class="ai-chat-drawer"
    @close="closeAiChat"
    destroy-on-close
  >
  
    <div class="chat-container flex flex-col h-full bg-[#F7EECD]/30"> 
      <div class="chat-header flex-shrink-0 bg-white/90 backdrop-blur-md border-b border-[#EFDCE2]/50 px-6 py-4 flex items-center gap-4 sticky top-0 z-10">
        <div class="ai-avatar-wrapper relative">
          <div class="w-10 h-10 rounded-full bg-[#C2D68F] flex items-center justify-center shadow-sm text-white">
            <el-icon :size="20"><ChatDotRound /></el-icon>
          </div>
          <span class="online-dot absolute -bottom-0.5 -right-0.5 w-3 h-3 bg-[#E0DC95] border-2 border-white rounded-full"></span>
        </div>
        <div class="header-info flex-1">
          <h3 class="text-sm font-bold text-gray-800 leading-tight">AI 职业规划助手</h3>
          <p class="text-xs text-[#8A9E58] font-medium mt-0.5">在线 · 随时为你服务</p>
        </div>
        <el-button text circle :icon="Close" class="text-gray-400 hover:text-[#C2D68F] hover:bg-[#F7EECD]" @click="closeAiChat" />
      </div>

      <div class="chat-messages flex-1 overflow-y-auto px-4 py-6 space-y-6 scroll-smooth" ref="messagesRef">
        <div v-for="(msg, idx) in messages" :key="idx" :class="['message-group flex gap-3 animate-fade-in-up', msg.role]">
          
          <div class="message-avatar flex-shrink-0 w-8 h-8 rounded-full bg-[#C2D68F] flex items-center justify-center text-white shadow-sm" v-if="msg.role === 'assistant'">
            <el-icon :size="16"><ChatDotRound /></el-icon>
          </div>
          
          <div class="message-content flex flex-col max-w-[85%]" :class="{'items-end': msg.role === 'user'}">
            <div 
              class="message-bubble px-4 py-3 text-sm leading-relaxed shadow-sm break-words"
              :class="msg.role === 'user' ? 'bg-[#EFDCE2] text-gray-800 rounded-2xl rounded-tr-sm' : 'bg-white text-gray-700 rounded-2xl rounded-tl-sm border border-[#EFDCE2]/50'"
            >
              <span v-html="formatMessage(msg.content)"></span>
            </div>
            <span class="text-[10px] text-gray-400 mt-1 px-1">{{ msg.time }}</span>
          </div>

          <div class="message-avatar flex-shrink-0 w-8 h-8 rounded-full bg-white border border-[#EFDCE2] flex items-center justify-center text-[#EFDCE2]" v-if="msg.role === 'user'">
            <el-icon :size="16"><User /></el-icon>
          </div>
        </div>
        
        <div v-if="isTyping" class="message-group flex gap-3 animate-fade-in-up">
          <div class="w-8 h-8 rounded-full bg-[#C2D68F] flex items-center justify-center text-white shadow-sm">
            <el-icon :size="16"><ChatDotRound /></el-icon>
          </div>
          <div class="bg-white border border-[#EFDCE2]/50 px-4 py-3 rounded-2xl rounded-tl-sm shadow-sm">
            <div class="typing-indicator flex gap-1.5">
              <span class="w-1.5 h-1.5 bg-[#C2D68F] rounded-full animate-bounce" style="animation-delay: 0s"></span>
              <span class="w-1.5 h-1.5 bg-[#C2D68F] rounded-full animate-bounce" style="animation-delay: 0.15s"></span>
              <span class="w-1.5 h-1.5 bg-[#C2D68F] rounded-full animate-bounce" style="animation-delay: 0.3s"></span>
            </div>
          </div>
        </div>
      </div>

      <div class="chat-input-area flex-shrink-0 bg-white border-t border-[#EFDCE2]/50 p-4">
        <div class="relative flex items-end gap-2 bg-[#F7EECD]/20 border border-[#EFDCE2] rounded-2xl p-1.5 focus-within:ring-2 focus-within:ring-[#EFDCE2] focus-within:border-[#C2D68F] transition-all">
          <el-input
            v-model="inputText"
            type="textarea"
            :rows="1"
            placeholder="输入你的问题..."
            @keyup.enter="sendMessage"
            :disabled="isTyping"
            class="custom-textarea"
            resize="none"
          />
          <el-button 
            color="#C2D68F"
            :icon="Promotion" 
            :loading="isTyping"
            :disabled="!inputText.trim()"
            circle
            class="!w-9 !h-9 !rounded-xl !m-0 flex-shrink-0 shadow-sm disabled:opacity-50 text-white"
            @click="sendMessage"
          />
        </div>
        <div class="text-center mt-2">
          <span class="text-[10px] text-[#8A9E58]/60">AI 生成内容仅供参考，建议结合实际情况</span>
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, computed, nextTick, watch, onBeforeUnmount } from 'vue'
import { ChatDotRound, Close, User, Promotion } from '@element-plus/icons-vue'
import { useAiChat } from './useAiChat' // 全局状态

const { isAiChatVisible, closeAiChat, openAiChat: globalOpen } = useAiChat()

const messages = ref([])
const isTyping = ref(false)
const inputText = ref('')
const hasUnread = ref(false)
const messagesRef = ref(null)
const isIdle = ref(true)

// ==================== 拖拽引擎 ====================
const fabRef = ref(null)
const isDragging = ref(false)
const translateX = ref(0)
const translateY = ref(0)
let startX = 0, startY = 0, initialTx = 0, initialTy = 0, hasMoved = false
let minTx = 0, maxTx = 0, minTy = 0, maxTy = 0

const fabStyle = computed(() => ({
  transform: `translate3d(${translateX.value}px, ${translateY.value}px, 0)`,
  transition: isDragging.value ? 'none' : 'transform 0.2s cubic-bezier(0.16, 1, 0.3, 1)'
}))

const startDrag = (e) => {
  if (e.type === 'mousedown' && e.button !== 0) return
  hasMoved = false
  const el = fabRef.value
  if (!el) return
  const rect = el.getBoundingClientRect()
  const baseLeft = rect.left - translateX.value
  const baseTop = rect.top - translateY.value
  minTx = -baseLeft
  maxTx = window.innerWidth - baseLeft - rect.width
  minTy = -baseTop
  maxTy = window.innerHeight - baseTop - rect.height

  const isTouch = e.type.startsWith('touch')
  startX = isTouch ? e.touches[0].clientX : e.clientX
  startY = isTouch ? e.touches[0].clientY : e.clientY
  initialTx = translateX.value
  initialTy = translateY.value

  if (isTouch) {
    window.addEventListener('touchmove', doDrag, { passive: false })
    window.addEventListener('touchend', stopDrag)
  } else {
    window.addEventListener('mousemove', doDrag)
    window.addEventListener('mouseup', stopDrag)
  }
}

const doDrag = (e) => {
  const isTouch = e.type.startsWith('touch')
  const clientX = isTouch ? e.touches[0].clientX : e.clientX
  const clientY = isTouch ? e.touches[0].clientY : e.clientY
  const dx = clientX - startX
  const dy = clientY - startY

  if (!isDragging.value && (Math.abs(dx) > 3 || Math.abs(dy) > 3)) {
    isDragging.value = true
    hasMoved = true
  }

  if (isDragging.value) {
    if (e.cancelable) e.preventDefault() 
    let newTx = initialTx + dx
    let newTy = initialTy + dy
    translateX.value = Math.max(minTx, Math.min(newTx, maxTx))
    translateY.value = Math.max(minTy, Math.min(newTy, maxTy))
  }
}

const stopDrag = () => {
  window.removeEventListener('mousemove', doDrag)
  window.removeEventListener('mouseup', stopDrag)
  window.removeEventListener('touchmove', doDrag)
  window.removeEventListener('touchend', stopDrag)

  if (!hasMoved) {
    openDrawer()
  }
  setTimeout(() => { isDragging.value = false }, 0)
}

onBeforeUnmount(() => {
  window.removeEventListener('mousemove', doDrag)
  window.removeEventListener('mouseup', stopDrag)
  window.removeEventListener('touchmove', doDrag)
  window.removeEventListener('touchend', stopDrag)
})
// ====================================================

// 打开抽屉与初始化
const openDrawer = () => {
  globalOpen()
  hasUnread.value = false
  isIdle.value = false
  if (messages.value.length === 0) {
    addMessage('assistant', '👋 你好！我是你的 AI 职业规划助手。\n\n你可以随时向我提问关于简历优化、面试技巧或职业发展的问题。我能帮你做些什么？')
  }
}

watch(isAiChatVisible, (val) => {
  if (val) {
    hasUnread.value = false
    isIdle.value = false
    if (messages.value.length === 0) {
      addMessage('assistant', '👋 你好！我是你的 AI 职业规划助手。\n\n你可以随时向我提问关于简历优化、面试技巧或职业发展的问题。我能帮你做些什么？')
    }
  } else {
    isIdle.value = true
  }
})

const addMessage = (role, content) => {
  messages.value.push({ role, content, time: formatTime(new Date()) })
  scrollToBottom()
}

// 纯聊天发送逻辑 (待改   这里对接/api/chat 接口)
const sendMessage = async () => {
  if (!inputText.value.trim() || isTyping.value) return
  const text = inputText.value.trim()
  addMessage('user', text)
  inputText.value = ''
  isTyping.value = true
  
  // 模拟网络延迟
  await new Promise(resolve => setTimeout(resolve, 1000))
  addMessage('assistant', generateMockResponse(text))
  isTyping.value = false
}

const formatTime = (date) => `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
const formatMessage = (content) => content ? content.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>').replace(/\n/g, '<br>') : ''
const scrollToBottom = () => { nextTick(() => { if (messagesRef.value) messagesRef.value.scrollTop = messagesRef.value.scrollHeight }) }

// 模拟回复（待改）
const generateMockResponse = (text) => `针对“${text}”这个问题，我建议你：\n\n1. 结合实际项目经验来阐述。\n2. 突出你的核心竞争力。\n\n这只是个模拟回答，接入真实接口后会更智能哦！`
</script>

<style scoped>
/* 悬浮球样式保留你原有的极佳设计 */
.ai-fab {
  position: fixed;
  bottom: 2rem; 
  right: 2rem;  
  z-index: 9999; 
  cursor: grab; 
  touch-action: none; 
  user-select: none;
  -webkit-user-drag: none;
}
.ai-fab.is-dragging { cursor: grabbing !important; opacity: 0.9; }
.ai-fab-inner {
  width: 56px; height: 56px;
  background: linear-gradient(135deg, #C2D68F 0%, #8A9E58 100%);
  border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 10px 25px -5px rgba(194, 214, 143, 0.5), 0 8px 10px -6px rgba(194, 214, 143, 0.4);
  position: relative; transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.3s;
}
.ai-fab:not(.is-dragging):hover .ai-fab-inner { transform: translateY(-4px) scale(1.05); box-shadow: 0 20px 25px -5px rgba(194, 214, 143, 0.6), 0 10px 10px -5px rgba(194, 214, 143, 0.4); }
.ai-fab:active .ai-fab-inner { transform: scale(0.95); }
.pulse-ring { position: absolute; inset: 0; border-radius: 50%; background: rgba(194, 214, 143, 0.5); animation: pulse-ring 2s infinite; }
@keyframes pulse-ring { 0% { transform: scale(1); opacity: 0.6; } 100% { transform: scale(1.6); opacity: 0; } }
.unread-badge { position: absolute; top: 0; right: 0; width: 12px; height: 12px; background: #EFDCE2; border: 2px solid white; border-radius: 50%; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
.animate-bounce-slow { animation: bounce 2s infinite; }
.ai-chat-drawer :deep(.el-drawer__body) { padding: 0; overflow: hidden; }
.animate-fade-in-up { animation: fadeInUp 0.4s cubic-bezier(0.16, 1, 0.3, 1); }
@keyframes fadeInUp { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
.chat-messages::-webkit-scrollbar { width: 6px; }
.chat-messages::-webkit-scrollbar-track { background: transparent; }
.chat-messages::-webkit-scrollbar-thumb { background: #EFDCE2; border-radius: 3px; }
.chat-messages::-webkit-scrollbar-thumb:hover { background: #C2D68F; }
.custom-textarea :deep(.el-textarea__inner) { background: transparent; border: none; box-shadow: none; padding: 8px 12px; font-size: 14px; line-height: 1.5; max-height: 100px; }
.custom-textarea :deep(.el-textarea__inner):focus { box-shadow: none; }
</style>