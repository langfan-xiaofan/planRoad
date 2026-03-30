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
    <div class="chat-container flex flex-col h-full bg-[#F7EECD]/30"> <div class="chat-header flex-shrink-0 bg-white/90 backdrop-blur-md border-b border-[#EFDCE2]/50 px-6 py-4 flex items-center gap-4 sticky top-0 z-10">
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

      <div v-if="showOptions" class="quick-options-panel flex-shrink-0 bg-white border-t border-[#EFDCE2]/50 p-4 transition-all duration-300">
        <div class="text-xs font-semibold text-gray-400 mb-3 uppercase tracking-wide">
          {{ isActionMode ? '选择下一步行动' : (currentQuestion?.optionsTitle || '请选择：') }}
        </div>
        
        <div class="options-grid grid grid-cols-2 gap-2.5" :class="{ 'grid-cols-1': currentLayout === 'vertical' }">
          <button 
            v-for="opt in displayOptions" 
            :key="opt.value"
            class="option-btn relative flex items-center justify-center gap-2 p-3 text-sm font-medium rounded-xl border transition-all duration-200 active:scale-95"
            :class="getOptionClass(opt)"
            @click="handleOptionClick(opt)"
          >
            <el-icon v-if="!isActionMode && getIconComponent(opt.icon)" class="text-base">
              <component :is="getIconComponent(opt.icon)" />
            </el-icon>
            <span>{{ opt.label }}</span>
            <el-icon v-if="!isActionMode && currentQuestion?.multiSelect && isOptionSelected(opt.value)" class="absolute top-1 right-1 text-[#C2D68F] w-4 h-4">
              <Check />
            </el-icon>
          </button>
        </div>

        <div v-if="!isActionMode && currentQuestion?.hasOther" class="other-input mt-3 flex gap-2 animate-fade-in">
          <el-input v-model="otherText" placeholder="或输入其他内容..." size="small" class="flex-1 custom-input-border" @keyup.enter="submitOther" />
          <el-button color="#C2D68F" size="small" class="!rounded-lg text-white" @click="submitOther">提交</el-button>
        </div>
      </div>

      <div v-else class="chat-input-area flex-shrink-0 bg-white border-t border-[#EFDCE2]/50 p-4">
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
import { 
  ChatDotRound, Close, User, Promotion, Check, 
  School,Medal, Suitcase, Search, Switch, 
  QuestionFilled, TrendCharts, Right, Reading, Clock, Timer 
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/user' 

// 引入全局状态
import { useAiChat } from './useAiChat'

const router = useRouter()
const userStore = useUserStore()

// 解构出全局变量和关闭方法
const { isAiChatVisible, closeAiChat, openAiChat: globalOpen } = useAiChat()

const messages = ref([])
const isTyping = ref(false)
const inputText = ref('')
const hasUnread = ref(false)
const messagesRef = ref(null)
const isIdle = ref(true)

// ==================== 终极拖拽引擎（ Translate3d 硬件加速版） ====================
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
// ========================================================================

const showOptions = ref(false)
const currentStep = ref(0)
const selectedOptions = ref({}) 
const otherText = ref('')
const collectedData = ref({})
const isActionMode = ref(false)
const actionOptions = ref([])

const iconMap = { School, Medal, Suitcase, Search, Switch, QuestionFilled, TrendCharts, Right, Reading, Clock, Timer }
const getIconComponent = (name) => iconMap[name] || null

const questions = [
  { id: 'identity', question: '你好！我是你的 AI 职业规划助手 🤖\n\n为了提供更精准的建议，先了解一下你的基本情况：\n\n**你目前是什么状态？**', optionsTitle: '选择当前身份', options: [ { value: 'student', label: '在校学生', icon: 'School' }, { value: 'freshman', label: '应届毕业生', icon: 'Medal' }, { value: 'employee', label: '职场人士', icon: 'Suitcase' }, { value: 'jobhunter', label: '待业/求职中', icon: 'Search' }, { value: 'career_change', label: '准备转行', icon: 'Switch' } ], optionsLayout: 'grid' },
  { id: 'experience', question: '明白了。请问你的工作或相关学习年限是？', optionsTitle: '选择年限', options: [ { value: '0', label: '在校/无经验' }, { value: '1', label: '1 年以内' }, { value: '2', label: '1-3 年' }, { value: '3', label: '3-5 年' }, { value: '4', label: '5-8 年' }, { value: '5', label: '8 年以上' } ], condition: (data) => data.identity !== 'student' },
  { id: 'field', question: '你目前所在的领域或专业方向是？', optionsTitle: '选择领域', options: [ { value: 'tech', label: '技术/开发' }, { value: 'product', label: '产品/运营' }, { value: 'design', label: '设计/创意' }, { value: 'marketing', label: '市场/销售' }, { value: 'finance', label: '金融/财务' }, { value: 'hr', label: '人力资源' }, { value: 'other', label: '其他' } ], hasOther: true, optionsLayout: 'grid' },
  { id: 'goal', question: '太好了！那你现阶段最想解决的核心问题是？', optionsTitle: '核心诉求', options: [ { value: 'explore', label: '迷茫探索', icon: 'QuestionFilled', desc: '不知道适合什么' }, { value: 'career_change', label: '转行规划', icon: 'Switch', desc: '想转行不知方向' }, { value: 'promotion', label: '晋升加薪', icon: 'TrendCharts', desc: '想升职/加薪' }, { value: 'job_change', label: '跳槽选择', icon: 'Right', desc: '想跳槽选机会' }, { value: 'skill_upgrade', label: '技能提升', icon: 'Reading', desc: '需要学习规划' } ], optionsLayout: 'vertical' },
  { id: 'target', question: '有没有心仪的具体岗位方向？', optionsTitle: '目标岗位', options: [ { value: 'has', label: '有明确目标' }, { value: 'vague', label: '有模糊方向' }, { value: 'none', label: '完全没头绪' } ], hasOther: true, otherPlaceholder: '例如：高级前端工程师、产品经理...' },
  { id: 'timeline', question: '你希望在多久内看到明显的成效或改变？', optionsTitle: '时间预期', options: [ { value: '1', label: '1 个月内', icon: 'Timer' }, { value: '3', label: '3 个月内', icon: 'Timer' }, { value: '6', label: '半年内', icon: 'Timer' }, { value: '12', label: '1 年内', icon: 'Timer' }, { value: 'long', label: '长期规划', icon: 'Clock' } ], optionsLayout: 'grid' },
  { id: 'priority', question: '最后，对你来说最重要的考量因素是什么？(可多选)', optionsTitle: '核心驱动力', options: [ { value: 'salary', label: '💰 薪资水平' }, { value: 'balance', label: '⚖️ 工作生活平衡' }, { value: 'growth', label: '📈 发展空间' }, { value: 'interest', label: '❤️ 兴趣匹配' }, { value: 'future', label: '🚀 行业前景' }, { value: 'stability', label: '🛡️ 稳定性' } ], optionsLayout: 'vertical', multiSelect: true }
]

const currentQuestion = computed(() => {
  if (currentStep.value >= questions.length) return null
  const q = questions[currentStep.value]
  if (q.condition && !q.condition(collectedData.value)) {
    currentStep.value++
    return currentQuestion.value
  }
  return q
})

const displayOptions = computed(() => isActionMode.value ? actionOptions.value : (currentQuestion.value?.options || []))
const currentLayout = computed(() => isActionMode.value ? 'vertical' : (currentQuestion.value?.optionsLayout || 'grid'))

const isOptionSelected = (value) => {
  const q = currentQuestion.value
  if (!q || !q.multiSelect) return false
  const selections = selectedOptions.value[q.id] || []
  return selections.includes(value)
}

const getOptionClass = (opt) => {
  if (isActionMode.value) return 'bg-white border-gray-100 text-gray-600 hover:border-[#C2D68F] hover:bg-[#F7EECD]/30'
  const isSelected = isOptionSelected(opt.value)
  // 选中的按钮颜色：青柠粉荔体系的微弱高亮
  return isSelected 
    ? 'bg-[#C2D68F]/10 border-[#C2D68F] text-[#8A9E58] shadow-sm ring-1 ring-[#C2D68F]/50' 
    : 'bg-white border-gray-100 text-gray-600 hover:border-[#C2D68F]/50 hover:bg-[#F7EECD]/30'
}

const openDrawer = () => {
  globalOpen() // 调用全局方法打开
  hasUnread.value = false
  isIdle.value = false
  if (messages.value.length === 0) startConversation()
}

// 监听全局状态，如果被外部打开，同样清空未读并初始化
watch(isAiChatVisible, (val) => {
  if (val) {
    hasUnread.value = false
    isIdle.value = false
    if (messages.value.length === 0) startConversation()
  } else {
    isIdle.value = true
  }
})

const startConversation = () => {
  currentStep.value = 0
  collectedData.value = {}
  selectedOptions.value = {}
  showOptions.value = false
  isActionMode.value = false
  addMessage('assistant', '👋 你好！我是你的 AI 职业规划助手。\n\n让我们花 2 分钟了解一下你的情况，我会为你提供**个性化**的规划建议。')
  setTimeout(() => askNextQuestion(), 800)
}

const askNextQuestion = () => {
  const q = currentQuestion.value
  if (!q) { generateSummary(); return }
  addMessage('assistant', q.question)
  showOptions.value = true
}

const addMessage = (role, content) => {
  messages.value.push({ role, content, time: formatTime(new Date()) })
  scrollToBottom()
}

const handleOptionClick = async (option) => {
  if (isActionMode.value) { executeAction(option.action || option.value); return }
  const q = currentQuestion.value
  if (!q) return
  if (q.multiSelect) {
    const currentSelections = selectedOptions.value[q.id] || []
    const index = currentSelections.indexOf(option.value)
    if (index > -1) currentSelections.splice(index, 1)
    else currentSelections.push(option.value)
    selectedOptions.value[q.id] = [...currentSelections]
    return 
  }
  selectedOptions.value[q.id] = option.value
  collectedData.value[q.id] = option.value
  addMessage('user', option.label)
  if (option.value === 'other' && q.hasOther) { showOptions.value = false; return }
  await nextQuestion()
}

const submitOther = async () => {
  if (!otherText.value.trim()) return
  const q = currentQuestion.value
  if (!q) return
  collectedData.value[q.id] = otherText.value.trim()
  addMessage('user', otherText.value.trim())
  otherText.value = ''
  await nextQuestion()
}

const nextQuestion = async () => {
  showOptions.value = false
  currentStep.value++
  await new Promise(resolve => setTimeout(resolve, 600))
  askNextQuestion()
}

const generateSummary = async () => {
  isTyping.value = true
  showOptions.value = false
  questions.forEach(q => { if (q.multiSelect && selectedOptions.value[q.id]) collectedData.value[q.id] = selectedOptions.value[q.id] })
  await new Promise(resolve => setTimeout(resolve, 1500))
  const summary = generateSummaryText(collectedData.value)
  addMessage('assistant', summary)
  isTyping.value = false
  userStore.setOnboardingData(collectedData.value)
  setTimeout(() => {
    addMessage('assistant', '✨ 基于你的画像，建议优先进行以下步骤：\n\n1️⃣ **能力测评**：量化你的核心竞争力\n2️⃣ **人岗匹配**：发现最适合你的机会\n3️⃣ **路径规划**：制定 actionable 的计划\n\n你想先从哪个开始？')
    showOptions.value = true
    isActionMode.value = true
    actionOptions.value = [
      { value: 'goResume', label: '📄 开始能力测评', action: 'goResume' },
      { value: 'goMatch', label: '🎯 查看人岗匹配', action: 'goMatch' },
      { value: 'goPath', label: '🗺️ 生成职业路径', action: 'goPath' }
    ]
  }, 500)
}

const executeAction = (action) => {
  closeAiChat()
  setTimeout(() => {
    try {
      if (action === 'goResume') router.push('/resume')
      else if (action === 'goMatch') router.push('/match')
      else if (action === 'goPath') router.push('/path')
      else router.push('/match') 
    } catch (e) {
      console.error('Navigation failed', e)
      ElMessage.warning('页面跳转失败，请手动导航')
    }
  }, 300)
}

const generateSummaryText = (data) => {
  const identityMap = { student: '在校学生', freshman: '应届生', employee: '职场人', jobhunter: '求职者', career_change: '转行者' }
  const goalMap = { explore: '方向探索', career_change: '转行', promotion: '晋升', job_change: '跳槽', skill_upgrade: '提升' }
  return `📋 **规划画像已生成**\n\n` +
    `• 身份：${identityMap[data.identity] || data.identity}\n` +
    `• 领域：${data.field || '待定'}\n` +
    `• 目标：${goalMap[data.goal] || data.goal}\n` +
    `• 周期：${data.timeline ? (data.timeline + '个月') : '长期'}\n\n` +
    `正在为你定制专属方案...`
}

const sendMessage = async () => {
  if (!inputText.value.trim() || isTyping.value) return
  const text = inputText.value.trim()
  addMessage('user', text)
  inputText.value = ''
  isTyping.value = true
  showOptions.value = false
  isActionMode.value = false 
  await new Promise(resolve => setTimeout(resolve, 1000))
  addMessage('assistant', '收到！正在分析...\n\n' + generateMockResponse(text))
  isTyping.value = false
}

const formatTime = (date) => `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
const formatMessage = (content) => content ? content.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>').replace(/\n/g, '<br>') : ''
const scrollToBottom = () => { nextTick(() => { if (messagesRef.value) messagesRef.value.scrollTop = messagesRef.value.scrollHeight }) }
const generateMockResponse = () => '这是一个模拟回复。接入真实 AI 后，我将为你提供深度分析。'
</script>

<style scoped>
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

.ai-fab.is-dragging {
  cursor: grabbing !important;
  opacity: 0.9;
}

/* 悬浮球配色重置：抹茶绿渐变 */
.ai-fab-inner {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #C2D68F 0%, #8A9E58 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10px 25px -5px rgba(194, 214, 143, 0.5), 0 8px 10px -6px rgba(194, 214, 143, 0.4);
  position: relative;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.3s;
}

.ai-fab:not(.is-dragging):hover .ai-fab-inner {
  transform: translateY(-4px) scale(1.05);
  box-shadow: 0 20px 25px -5px rgba(194, 214, 143, 0.6), 0 10px 10px -5px rgba(194, 214, 143, 0.4);
}

.ai-fab:active .ai-fab-inner {
  transform: scale(0.95);
}

.pulse-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: rgba(194, 214, 143, 0.5);
  animation: pulse-ring 2s infinite;
}

@keyframes pulse-ring {
  0% { transform: scale(1); opacity: 0.6; }
  100% { transform: scale(1.6); opacity: 0; }
}

/* 未读红点修改为浅粉色点缀 */
.unread-badge {
  position: absolute;
  top: 0;
  right: 0;
  width: 12px;
  height: 12px;
  background: #EFDCE2;
  border: 2px solid white;
  border-radius: 50%;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.animate-bounce-slow {
  animation: bounce 2s infinite;
}

.ai-chat-drawer :deep(.el-drawer__body) {
  padding: 0;
  overflow: hidden;
}

.animate-fade-in-up {
  animation: fadeInUp 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: translateY(0); }
}

.chat-messages::-webkit-scrollbar {
  width: 6px;
}
.chat-messages::-webkit-scrollbar-track {
  background: transparent;
}
.chat-messages::-webkit-scrollbar-thumb {
  background: #EFDCE2; /* 浅粉色滚动条 */
  border-radius: 3px;
}
.chat-messages::-webkit-scrollbar-thumb:hover {
  background: #C2D68F; /* 悬停抹茶绿 */
}

.custom-textarea :deep(.el-textarea__inner) {
  background: transparent;
  border: none;
  box-shadow: none;
  padding: 8px 12px;
  font-size: 14px;
  line-height: 1.5;
  max-height: 100px;
}
.custom-textarea :deep(.el-textarea__inner):focus {
  box-shadow: none;
}
.custom-input-border :deep(.el-input__wrapper) {
  box-shadow: 0 0 0 1px #EFDCE2 inset;
}
.custom-input-border :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #C2D68F inset;
}

.option-btn {
  transition: all 0.2s ease;
}
.option-btn:hover {
  transform: translateY(-1px);
}
.option-btn:active {
  transform: translateY(0);
}
</style>