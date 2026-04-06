<template>
  <!-- 解析简历页面 -->
  <div class="h-full flex flex-col p-4 overflow-hidden relative">
    <!-- 解析简历表单 -->
    <transition name="el-fade-in">
      <div v-if="!jobStore.hasResumeProfile" class="absolute inset-0 z-10 flex flex-col items-center justify-center p-6 overflow-y-auto">
        <div class="w-full max-w-2xl bg-white/80 backdrop-blur-md rounded-3xl p-8 border border-[#EFDCE2]/60 shadow-sm flex flex-col items-center">
          <!-- 上传简历图标 -->
          <div class="w-16 h-16 bg-[#F7EECD] rounded-full flex items-center justify-center mb-6 shadow-sm border border-[#C2D68F]">
            <el-icon class="text-3xl text-[#8A9E58]"><DocumentAdd /></el-icon>
          </div>
          <!-- 上传简历标题 -->
          <h2 class="text-2xl font-extrabold text-brand-dark mb-2">上传简历开始解析</h2>
          <p class="text-sm text-gray-400 mb-8 text-center">AI 将自动提取核心技能，为你生成专属的六维能力模型与发展建议</p>
          <!-- 上传简历文件区域 --> 
          <div class="w-full mb-6">
            <el-upload
              class="resume-uploader"
              drag
              action="#"
              :auto-upload="false"
              :show-file-list="true"
              :limit="1"
              :on-change="handleFileChange"
              :on-remove="handleFileRemove"
              accept=".pdf,.doc,.docx"
            >
              <el-icon class="el-icon--upload text-[#C2D68F]"><upload-filled /></el-icon>
              <div class="el-upload__text text-gray-500">
                将文件拖到此处，或 <em class="el-upload__text text-[#8A9E58] font-bold not-italic">点击上传</em>
              </div>
              <template #tip>
                <div class="el-upload__tip text-center text-xs text-gray-400 mt-2">
                  支持 PDF, DOC, DOCX 格式文件，大小不超过 10MB
                </div>
              </template>
            </el-upload>
          </div>
          <!-- 个人介绍输入框 --> 
          <div class="w-full mb-8">
            <h3 class="text-sm font-bold text-gray-700 mb-2">个人介绍 (选填)</h3>
            <el-input
              v-model="personalIntro"
              type="textarea"
              :rows="4"
              placeholder="可以在此补充简历中未体现的个人亮点、职业期望或技能说明..."
              class="custom-intro-input"
              resize="none"
            />
          </div>
          <!-- 解析按钮按钮 -->
          <div class="w-full flex justify-center">
            <el-button 
              round 
              class="!bg-[#8A9E58] !text-white !border-none shadow-md hover:opacity-90 font-bold px-10 py-5 text-base w-full max-w-sm" 
              :loading="isParsing"
              :disabled="!selectedFile && !personalIntro.trim()"
              @click="startParse"
            >
              {{ isParsing ? 'AI 正在深度解析...' : '开始智能解析' }}
            </el-button>
          </div>

        </div>
      </div>
    </transition>
    <!-- 解析简历报告区域 -->
    <div v-if="jobStore.hasResumeProfile" class="flex-1 flex flex-col min-h-0 bg-white/40 backdrop-blur-md rounded-3xl p-6 border border-brand-pink/30 shadow-sm overflow-y-auto">
      <div class="flex items-center justify-between mb-6 shrink-0">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 bg-white rounded-2xl shadow-sm flex items-center justify-center border border-brand-pink/30">
            <el-icon class="text-2xl text-brand-green"><Document /></el-icon>
          </div>
          <div>
            <h2 class="text-xl font-extrabold text-brand-dark leading-tight">简历深度解析报告</h2>
            <p class="text-sm text-gray-500">AI 已完成对您的多维能力评估</p>
          </div>
        </div>
        <!-- 重新上传按钮 -->
        <el-button round class="!bg-brand-pink !text-gray-700 !border-none hover:bg-brand-pink/80 font-bold px-5" @click="jobStore.resetAll">
          <el-icon class="mr-1"><RefreshRight /></el-icon> 重新上传
        </el-button>
      </div>
      <!-- 解析简历报告内容区域 -->
      <div class="flex-1 min-h-0">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 h-full">
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-brand-pink/20 flex flex-col relative">
            <div class="shrink-0 mb-2">
              <h3 class="font-extrabold text-brand-dark text-base">六维能力模型</h3>
              <p class="text-xs text-gray-400">基于您过往经历的客观量化评分</p>
            </div>
            <div ref="radarRef" class="flex-1 w-full min-h-0"></div>
          </div>
          <!-- 能力总结与发展建议区域 -->
          <div class="bg-white rounded-3xl p-6 shadow-sm border border-brand-pink/20 flex flex-col">
            <div class="shrink-0 mb-4">
              <h3 class="font-extrabold text-brand-dark text-base mb-1">能力总结与发展建议</h3>
              <el-tag effect="light" type="success" round size="small" class="bg-brand-cream/50 text-brand-green border-none font-bold">
                已提炼核心优势
              </el-tag>
            </div>
            <!-- 核心优势区域 -->
            <div class="flex-1 overflow-y-auto pr-2 custom-scrollbar">
              <div class="space-y-4">
                <div class="bg-gray-50/80 rounded-2xl p-4 border border-gray-100">
                  <h4 class="text-sm font-bold text-gray-700 mb-2 flex items-center gap-2">
                    <el-icon class="text-brand-pink"><StarFilled /></el-icon> 核心优势
                  </h4>
                  <ul class="text-sm text-gray-600 space-y-1.5 list-disc pl-5 marker:text-brand-green">
                    <li>具备扎实的 <strong>前端工程化</strong> 基础，在组件库建设和构建工具有较深积累。</li>
                    <li>沟通与业务理解能力突出，能在团队中起到关键桥梁作用。</li>
                  </ul>
                </div>
                <!-- 潜在瓶颈区域 -->
                <div class="bg-[#F7EECD]/20 rounded-2xl p-4 border border-[#EFDCE2]/30">
                  <h4 class="text-sm font-bold text-[#8A9E58] mb-2 flex items-center gap-2">
                    <el-icon><WarningFilled /></el-icon> 潜在瓶颈
                  </h4>
                  <p class="text-sm text-gray-600 leading-relaxed">
                    在<strong>架构设计</strong>与<strong>底层性能调优</strong>方面经验略显不足，若期望晋升高级岗，需加强微前端等复杂场景的实战经验。
                  </p>
                </div>
                <!-- 下一步行动建议区域 -->
                <div class="bg-brand-green/5 rounded-2xl p-4 border border-brand-green/20">
                  <h4 class="text-sm font-bold text-brand-dark mb-2">下一步行动建议</h4>
                  <p class="text-sm text-gray-600 leading-relaxed">
                    建议前往 <span class="text-brand-pink font-bold">人岗匹配</span> 页面，选择您的目标岗位，查看具体的技能缺口，并生成量身定制的学习路径。
                  </p>
                </div>
              </div>
            </div>
            <!-- 前往人岗匹配按钮 -->
            <div class="shrink-0 mt-5 flex justify-end">
              <el-button round class="!bg-brand-green !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/match')">
                前往人岗匹配 <el-icon class="ml-1"><Right /></el-icon>
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { DocumentAdd, UploadFilled, Document, RefreshRight, StarFilled, WarningFilled, Right } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import { useJobDataStore } from '@/stores/jobData'

const router = useRouter()
const jobStore = useJobDataStore()

// 空状态下的表单数据
const selectedFile = ref(null)
const personalIntro = ref('')
const isParsing = ref(false)

// 图表引用
const radarRef = ref(null)
let chartInstance = null

// --- 文件选择与解析逻辑 ---
const handleFileChange = (uploadFile, uploadFiles) => {
  // 只保存最新选择的文件
  selectedFile.value = uploadFile.raw
}

const handleFileRemove = () => {
  selectedFile.value = null
}

const startParse = async () => {
  if (!selectedFile.value && !personalIntro.value.trim()) return

  isParsing.value = true
  try {
    // 调用 store 的方法，后续可修改该方法接受文本和文件
    await jobStore.parseResumeMock(selectedFile.value, { 
      onProgress: (p) => {
        // 展示进度条在这里做

        console.log(`解析进度: ${p}%`)
      }
    })
    
    // 渲染完成后初始化图表
    nextTick(() => {
      initChart()
    })
  } catch (error) {
    console.error('解析失败', error)
  } finally {
    isParsing.value = false
  }
}

// --- ECharts 图表逻辑 ---
const initChart = () => {
  if (!radarRef.value) return
  if (chartInstance) chartInstance.dispose()

  chartInstance = echarts.init(radarRef.value)
  const option = {
    tooltip: { trigger: 'item' },
    radar: {
      indicator: [
        { name: '前端工程化', max: 100 },
        { name: '全栈与Node', max: 100 },
        { name: '架构与性能', max: 100 },
        { name: '业务沟通', max: 100 },
        { name: '管理潜质', max: 100 },
        { name: '创新认知', max: 100 }
      ],
      radius: '60%',
      center: ['50%', '50%'],
      splitNumber: 4,
      axisName: { color: '#4b5563', fontSize: 12, fontWeight: 'bold' },
      splitArea: { areaStyle: { color: ['#f9fafb', '#ffffff'] } },
      axisLine: { lineStyle: { color: '#e5e7eb' } },
      splitLine: { lineStyle: { color: '#e5e7eb' } }
    },
    series: [{
      type: 'radar',
      data: [{
        value: [85, 78, 65, 80, 55, 90],
        name: '能力量化得分',
        itemStyle: { color: '#8A9E58' },
        areaStyle: { color: 'rgba(138, 158, 88, 0.4)' },
        lineStyle: { width: 2 }
      }]
    }]
  }
  chartInstance.setOption(option)
}

const handleResize = () => {
  chartInstance?.resize()
}

onMounted(() => {
  if (jobStore.hasResumeProfile) {
    nextTick(() => initChart())
  }
  window.addEventListener('resize', handleResize)
})

watch(() => jobStore.hasResumeProfile, (newVal) => {
  if (newVal) {
    nextTick(() => initChart())
  } else {
    // 重置表单
    selectedFile.value = null
    personalIntro.value = ''
  }
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  chartInstance?.dispose()
})
</script>

<style scoped>
/* 拖拽上传组件定制样式 */
.resume-uploader :deep(.el-upload-dragger) {
  background-color: #fafafa;
  border: 1.5px dashed #EFDCE2;
  border-radius: 16px;
  padding: 30px;
  transition: all 0.3s;
}
.resume-uploader :deep(.el-upload-dragger:hover) {
  border-color: #C2D68F;
  background-color: rgba(194, 214, 143, 0.05);
}
.resume-uploader :deep(.el-upload-dragger.is-dragover) {
  background-color: rgba(194, 214, 143, 0.1);
  border-color: #8A9E58;
}

/* 文本域定制样式 */
.custom-intro-input :deep(.el-textarea__inner) {
  border-radius: 12px;
  background-color: #f9fafb;
  border: 1px solid #e5e7eb;
  padding: 12px;
  transition: all 0.3s;
}
.custom-intro-input :deep(.el-textarea__inner:focus) {
  background-color: #ffffff;
  border-color: #C2D68F;
  box-shadow: 0 0 0 2px rgba(194, 214, 143, 0.2);
}

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #e5e7eb; border-radius: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: #d1d5db; }
</style>