<template>
  <div class="global-search-container">
    <el-autocomplete
      v-model="searchQuery"
      :fetch-suggestions="querySearch"
      placeholder="想找哪份分析报告？或者直接问我..."
      @select="handleSelect"
      class="w-full"
    >
      <template #prefix>
        <el-icon class="search-icon"><Search /></el-icon>
      </template>

      <template #default="{ item }">
        <div class="flex items-center justify-between py-1 px-1">
          <div class="flex flex-col">
            <span class="text-sm font-medium text-brand-dark">{{ item.value }}</span>
            <span class="text-[10px] text-gray-400 opacity-80">{{ item.desc }}</span>
          </div>
          <el-tag 
            size="small" 
            :type="getTagType(item.type)" 
            effect="light" 
            class="rounded-full scale-90"
          >
            {{ item.typeLabel }}
          </el-tag>
        </div>
      </template>
    </el-autocomplete>
  </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'

const searchQuery = ref('')
const router = useRouter()

// 模拟数据库（仅用于演示）
const searchPool = [
  // 1. 历史记录 (History)
  { value: '2026-03-25 产品经理分析', desc: '包含六维雷达图', type: 'history', typeLabel: '历史', id: 'h1' },
  { value: '后端开发 匹配报告', desc: '匹配度 85%', type: 'history', typeLabel: '历史', id: 'h2' },
  
  // 2. 职业库 (Job Library)
  { value: '前端开发工程师', desc: '需求方：互联网/字节跳动等', type: 'job', typeLabel: '岗位', id: 'j1' },
  { value: 'UI/UX 设计师', desc: '需求方：创意设计/小红书等', type: 'job', typeLabel: '岗位', id: 'j2' },
  
  // 3. AI 语义指令 (Command)
  { value: '切换到求职模式', desc: '调整首页工作台布局', type: 'cmd', typeLabel: 'AI指令' },
  { value: '帮我看看大厂要求', desc: '开启 AI 深度咨询', type: 'cmd', typeLabel: 'AI指令' }
]

// 联想过滤逻辑
const querySearch = (queryString, cb) => {
  const results = queryString
    ? searchPool.filter(item => 
        item.value.toLowerCase().includes(queryString.toLowerCase()) ||
        item.typeLabel.includes(queryString)
      )
    : searchPool.slice(0, 5) // 为空时默认显示前5条
  cb(results)
}

// 标签颜色逻辑
const getTagType = (type) => {
  const map = { 'history': 'info', 'job': 'success', 'cmd': 'warning' }
  return map[type] || 'info'
}

// 点击后的路由处理逻辑
const handleSelect = (item) => {
  console.log('用户选择了:', item)
  
  if (item.type === 'history') {
    router.push({ name: 'ResumeAnalysis', query: { reportId: item.id } })
  } else if (item.type === 'job') {
    // 带参数跳转到人岗匹配页
    router.push({ path: '/match', query: { target: item.value } })
  } else if (item.type === 'cmd') {
    // 触发 AI 悬浮球逻辑
    alert(`AI 正在执行指令：${item.value}`)
  }
}
</script>

<style scoped>
:deep(.el-input__wrapper) {
  background-color: rgba(255, 255, 255, 0.4) !important;
  backdrop-filter: blur(4px);
  border-radius: 20px !important;
  border: 1px solid #EFDCE2 !important; 
  box-shadow: none !important;
}
:deep(.el-input__wrapper.is-focus) {
  border-color: #C2D68F !important; 
  background-color: rgba(255, 255, 255, 0.8) !important;
}
.search-icon {
  color: #8A9E58; 
}
</style>