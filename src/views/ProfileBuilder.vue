<!-- 简历构建器 -->
<template>
  <div class="h-[calc(100vh-120px)] flex gap-4 relative">
    <!-- 加载状态提示 -->
    <div v-if="isGenerating" class="absolute inset-0 bg-white/80 backdrop-blur-sm z-50 flex flex-col items-center justify-center rounded-xl transition-all duration-300">
      <el-progress type="dashboard" :percentage="loadingProgress" color="#84B4C3" :width="150">
        <template #default="{ percentage }">
          <span class="block text-3xl font-bold text-themePrimary">{{ percentage }}%</span>
        </template>
      </el-progress>
      <p class="mt-6 text-themePrimary font-bold text-lg tracking-widest animate-pulse">
        {{ loadingText }}
      </p>
    </div>
    <!-- 数据源录入区域 -->
    <div v-if="!hasData" class="w-3/12 bg-white rounded-xl shadow-sm border border-gray-100 flex flex-col h-full overflow-hidden shrink-0 transition-all duration-300">
      <div class="h-14 px-4 border-b border-gray-50 bg-themeBg/30 flex items-center gap-2 font-bold text-gray-700 shrink-0">
        <el-icon class="text-themePrimary"><Document /></el-icon> 数据源录入
      </div>
    
      <div class="p-4 flex-1 overflow-y-auto">
        <div class="text-sm text-gray-500 mb-2">1. 简历文件上传</div>
        <el-upload
          class="mb-6 w-full"
          drag
          action="https://mock-api.com/upload"
          multiple
          :show-file-list="true"
        >
          <el-icon class="el-icon--upload text-themePrimary"><upload-filled /></el-icon>
          <div class="el-upload__text">
            将文件拖到此处，或 <em class="text-themePrimary">点击上传</em>
          </div>
          <template #tip>
            <div class="text-xs text-gray-400 mt-2">支持 PDF/Word，AI将自动提取能力项</div>
          </template>
        </el-upload>
        <!-- 补充描述录入 -->
        <div class="text-sm text-gray-500 mb-2">2. 补充描述 (可选)</div>
        <el-input 
          v-model="rawInput" 
          type="textarea" 
          :rows="6" 
          placeholder="可在此处直接粘贴文本，或补充未写在简历中的项目经历、证书等信息..."
          class="text-sm"
        ></el-input>
      </div>
      <!-- 生成按钮区域 -->
      <div class="p-4 border-t border-gray-50 shrink-0">
        <el-button type="primary" class="w-full bg-themePrimary border-none h-10 text-base" @click="startGeneration">
          生成 AI 能力画像
        </el-button>
      </div>
    </div>
    <!-- 结构化数据核对区域 -->
    <div :class="hasData ? 'w-4/12' : 'w-4/12'" class="bg-white rounded-xl shadow-sm border border-gray-100 flex flex-col h-full overflow-hidden shrink-0 transition-all duration-300">
      <div class="h-14 px-4 border-b border-gray-50 bg-themeBg/30 flex items-center justify-between font-bold text-gray-700 shrink-0">
        <div class="flex items-center gap-2">
          <el-button v-if="hasData" @click="resetData" size="small" circle class="border-themePrimary text-themePrimary hover:bg-themePrimary hover:text-white transition-colors" title="重新录入数据">
            <el-icon><Back /></el-icon>
          </el-button>
          <el-icon class="text-themePrimary"><Edit /></el-icon> 结构化数据核对
        </div>
        <el-tag v-if="hasData" size="small" color="#C7E2E0" class="border-none text-themePrimary">支持手动微调</el-tag>
      </div>

      <div class="p-4 flex-1 overflow-y-auto bg-gray-50/50">
        <div v-if="!hasData" class="h-full flex flex-col items-center justify-center opacity-50">
          <el-icon :size="48" class="text-gray-300 mb-4"><List /></el-icon>
          <p class="text-sm text-gray-500">等待 AI 解析数据...</p>
        </div>
      
        <el-collapse v-else v-model="activeNames" class="border-none">
          <el-collapse-item name="1" class="mb-3 bg-white rounded-lg border border-gray-100 px-2 shadow-sm">
            <template #title><span class="font-bold text-gray-700">📌 核心专业技能</span></template>
            <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 10 }" v-model="parsedData.skills" class="mt-2"></el-input>
          </el-collapse-item>
          
          <el-collapse-item name="2" class="mb-3 bg-white rounded-lg border border-gray-100 px-2 shadow-sm">
            <template #title><span class="font-bold text-gray-700">💼 实习与工作经验</span></template>
            <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 10 }" v-model="parsedData.experience" class="mt-2"></el-input>
          </el-collapse-item>

          <el-collapse-item name="3" class="mb-3 bg-white rounded-lg border border-gray-100 px-2 shadow-sm">
            <template #title><span class="font-bold text-gray-700">🏆 证书与荣誉</span></template>
            <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 6 }" v-model="parsedData.certs" class="mt-2"></el-input>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>
    <!-- AI 综合评测报告区域 -->
    <div :class="hasData ? 'w-8/12' : 'w-5/12'" class="bg-white rounded-xl shadow-sm border border-gray-100 flex flex-col h-full overflow-hidden shrink-0 transition-all duration-300">
      <div class="h-14 px-4 border-b border-gray-50 bg-themeBg/30 flex items-center gap-2 font-bold text-gray-700 shrink-0">
        <el-icon class="text-themePrimary"><DataLine /></el-icon> AI 综合评测报告
      </div>
      <!-- 评测报告内容区域 -->
      <div class="flex-1 relative overflow-hidden bg-white">
        <div v-if="!hasData" class="absolute inset-0 flex flex-col items-center justify-center">
          <el-empty description="暂无画像数据，请先在左侧录入源文件" :image-size="160">
            <span class="text-xs text-themePrimary opacity-70">系统将在此生成多维度雷达图谱与能力评分</span>
          </el-empty>
        </div>

        <div v-else class="p-6 h-full flex flex-col">
          <div class="w-full flex gap-2 mb-4 shrink-0">
            <el-tag color="#FAEDE5" class="border-none text-gray-700">前端开发</el-tag>
            <el-tag color="#FAEDE5" class="border-none text-gray-700">工程化经验丰富</el-tag>
            <el-tag color="#EAC0C3" class="border-none text-white">高抗压</el-tag>
          </div>
          <!-- 能力雷达图谱区域 -->
          <div class="flex-1 flex gap-6 min-h-0">
            <div class="w-[55%] h-full bg-gray-50/50 rounded-xl border border-gray-100 flex items-center justify-center p-2 shadow-inner">
              <div ref="radarChartRef" class="w-full h-full min-h-[300px]"></div>
            </div>
            <!-- 能力评分区域 -->
            <div class="w-[45%] flex flex-col gap-4 h-full">
              
              <div class="flex gap-4 shrink-0">
                <div class="flex-1 bg-white border border-gray-100 rounded-xl p-4 flex flex-col items-center shadow-sm">
                  <el-progress type="dashboard" :percentage="85" color="#84B4C3" :width="110" :stroke-width="8">
                    <template #default="{ percentage }">
                      <span class="block text-2xl font-black text-themePrimary">{{ percentage }}</span>
                      <span class="text-xs text-gray-400 mt-1">系统评级: 优秀</span>
                    </template>
                  </el-progress>
                  <div class="text-sm font-bold text-gray-700 mt-2">画像完整度</div>
                </div>
              
                <div class="flex-1 bg-white border border-gray-100 rounded-xl p-4 flex flex-col items-center shadow-sm">
                  <el-progress type="dashboard" :percentage="92" color="#EAC0C3" :width="110" :stroke-width="8">
                    <template #default="{ percentage }">
                      <span class="block text-2xl font-black text-themeAccent">{{ percentage }}</span>
                      <span class="text-xs text-gray-400 mt-1">击败 89% 同学</span>
                    </template>
                  </el-progress>
                  <div class="text-sm font-bold text-gray-700 mt-2">职场竞争力</div>
                </div>
              </div>
              <!-- AI 深度诊断区域 -->
             <div class="flex-1 bg-themeBg/40 border border-themeSecondary rounded-xl p-5 flex flex-col shadow-sm min-h-0">
                <div class="text-themePrimary font-bold mb-3 flex items-center gap-2 shrink-0">
                  <el-icon><Monitor /></el-icon> AI 深度诊断
                </div>
                <div class="flex-1 overflow-y-auto text-gray-700 text-sm leading-relaxed space-y-2 pr-2">
                  <p>根据提取的简历特征，你的<span class="text-themePrimary font-bold">学习能力</span>与<span class="text-themePrimary font-bold">专业技能</span>表现极为突出，技术栈（Vue3/React/Vite）高度符合当前一二线大厂的前端工程化标准。</p>
                  <p><strong>短板提示：</strong>在<span class="text-themeAccent font-bold">实习与实践经历</span>维度略显单薄。当前市场对候选人的业务落地能力要求较高。</p>
                  <p><strong>行动建议：</strong>建议在接下来的时间里，利用自身学习能力强的优势，主导或参与 1-2 个具有实际复杂业务场景（如：低代码平台、复杂大屏可视化）的开源协作项目，以弥补实战经验的不足。</p>
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onBeforeUnmount } from 'vue'
import { UploadFilled, Document, Edit, DataLine, List, Monitor, Back } from '@element-plus/icons-vue'
import * as echarts from 'echarts'

const rawInput = ref('')
const hasData = ref(false)
const isGenerating = ref(false)
const loadingProgress = ref(0)
const loadingText = ref('正在读取文件...')

const activeNames = ref(['1', '2', '3'])
const parsedData = ref({
  skills: '熟练掌握 Vue3、React 及其生态。\n熟悉 Vite 构建工具及前端工程化配置。\n掌握 JavaScript(ES6+) 及 TypeScript。',
  experience: '字节跳动 - 前端开发实习生 (2023.07 - 2024.01)\n负责内部创作者平台核心模块重构，提升页面加载速度 30%。',
  certs: 'CET-6 (520分)\n全国大学生计算机设计大赛 二等奖'
})

const radarChartRef = ref(null)

// 声明全局变量存储实例，用于销毁
let chartInstance = null
let resizeObserver = null

const startGeneration = () => {
  isGenerating.value = true
  loadingProgress.value = 0
  hasData.value = false
  loadingText.value = '正在解析非结构化文本...'

  const timer = setInterval(() => {
    loadingProgress.value += Math.floor(Math.random() * 15) + 5
    
    if (loadingProgress.value >= 30) loadingText.value = '正在提取专业技能特征...'
    if (loadingProgress.value >= 60) loadingText.value = '正在量化综合素质得分...'
    if (loadingProgress.value >= 85) loadingText.value = '正在构建多维雷达图谱...'

    if (loadingProgress.value >= 100) {
      loadingProgress.value = 100
      clearInterval(timer)
      
      // 延迟半秒关闭遮罩
      setTimeout(() => {
        isGenerating.value = false
        hasData.value = true
        
        // 【核心修复1】：等待 350ms，确保 Tailwind 的 duration-300 伸缩动画彻底执行完毕再画图
        setTimeout(() => {
          initRadarChart()
        }, 350)
        
      }, 500)
    }
  }, 400)
}

const resetData = () => {
  hasData.value = false
  loadingProgress.value = 0
  if (chartInstance) {
    chartInstance.dispose()
    chartInstance = null
  }
}

const initRadarChart = () => {
  if(!radarChartRef.value) return
  
  // 防止重复初始化
  if (chartInstance) {
    chartInstance.dispose()
  }
  
  chartInstance = echarts.init(radarChartRef.value)
  
  // 【核心修复2】：使用 ResizeObserver 精确监听 div 容器的尺寸变化，取代有缺陷的 window.resize
  if (!resizeObserver) {
    resizeObserver = new ResizeObserver(() => {
      if (chartInstance) chartInstance.resize()
    })
    resizeObserver.observe(radarChartRef.value)
  }

  chartInstance.setOption({
    tooltip: {},
    radar: {
      // 【核心修复3】：缩小半径至 55%，并在中心点强制绝对居中，给外围文字留足空间防裁剪
      radius: '55%', 
      center: ['50%', '50%'],
      indicator: [
        { name: '专业技能', max: 100 },
        { name: '证书要求', max: 100 },
        { name: '创新能力', max: 100 },
        { name: '学习能力', max: 100 },
        { name: '抗压能力', max: 100 },
        { name: '沟通能力', max: 100 },
        { name: '实习能力', max: 100 }
      ],
      axisName: { color: '#84B4C3', fontSize: 13, fontWeight: 'bold' },
      splitArea: {
        areaStyle: {
          color: ['rgba(250, 237, 229, 0.2)', 'rgba(250, 237, 229, 0.4)', 'rgba(250, 237, 229, 0.6)', 'rgba(250, 237, 229, 0.8)', '#FAEDE5'],
          shadowColor: 'rgba(0, 0, 0, 0.05)',
          shadowBlur: 10
        }
      },
      axisLine: { lineStyle: { color: 'rgba(199, 226, 224, 0.5)' } },
      splitLine: { lineStyle: { color: 'rgba(199, 226, 224, 0.5)' } }
    },
    series: [{
      name: '能力画像',
      type: 'radar',
      data: [
        {
          value: [85, 70, 80, 95, 88, 75, 60],
          name: '当前评估',
          itemStyle: { color: '#84B4C3' },
          areaStyle: { color: 'rgba(132, 180, 195, 0.5)' },
          lineStyle: { width: 2 }
        }
      ]
    }]
  })
}

// 页面卸载时清理监听器，防止内存泄漏
onBeforeUnmount(() => {
  if (resizeObserver && radarChartRef.value) {
    resizeObserver.unobserve(radarChartRef.value)
    resizeObserver.disconnect()
  }
  if (chartInstance) {
    chartInstance.dispose()
  }
})
</script>

<style scoped>
:deep(.el-collapse-item__header) {
  border-bottom: none;
  height: 40px;
}
:deep(.el-collapse-item__wrap) {
  border-bottom: none;
}
</style>