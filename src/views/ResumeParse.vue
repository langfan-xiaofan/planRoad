<template>
  <div class="h-full flex flex-col gap-3 overflow-hidden p-2">
    <!-- 解析报告标题 -->
    <div class="shrink-0 bg-white rounded-3xl p-5 shadow-sm border border-brand-cream flex justify-between items-center">
      <div>
        <h2 class="text-xl font-extrabold text-brand-dark flex items-center gap-2">
          <el-icon class="text-brand-green"><Document /></el-icon>
          AI 简历深度解析报告
        </h2>
        <p class="text-sm text-gray-500 mt-1">已为你提取核心能力特征，生成多维结构化能力画像（基准线）。</p>
      </div>
      <el-button round class="!bg-brand-green !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/match')">
        确认画像并前往匹配 <el-icon class="ml-1"><Right /></el-icon>
      </el-button>
    </div>
<!-- 个人信息展示 -->
    <div class="shrink-0 bg-white rounded-2xl px-6 py-4 shadow-sm border border-brand-cream flex flex-wrap gap-8 items-center relative overflow-hidden">
      <div class="absolute right-0 top-0 w-32 h-32 bg-brand-green/5 rounded-full blur-2xl pointer-events-none"></div>
      
      <div class="flex items-center gap-4 border-r border-gray-100 pr-8 z-10">
        <div class="w-11 h-11 rounded-full bg-brand-pink text-white flex items-center justify-center font-bold text-lg shadow-inner">
          {{ userInfo.name.charAt(0) }}
        </div>
        <div class="flex flex-col">
          <h3 class="text-lg font-extrabold text-brand-dark flex items-center gap-1.5 leading-tight">
            {{ userInfo.name }}
            <el-icon v-if="userInfo.gender === 'male'" class="text-blue-400 text-sm"><Male /></el-icon>
            <el-icon v-else class="text-pink-400 text-sm"><Female /></el-icon>
          </h3>
          <span class="text-[10px] text-brand-green bg-brand-green/10 px-1.5 py-0.5 rounded mt-1 w-max">当前解析对象</span>
        </div>
      </div>

      <div class="flex items-center gap-8 flex-1 z-10">
        <div class="flex items-center gap-2 text-sm text-gray-600">
          <div class="w-7 h-7 rounded-full bg-gray-50 flex items-center justify-center text-gray-400"><el-icon><Calendar /></el-icon></div>
          <span class="font-bold text-gray-700">{{ userInfo.age }} <span class="text-xs font-normal text-gray-400">岁</span></span>
        </div>
        <!-- 位置信息 -->
        <div class="flex items-center gap-2 text-sm text-gray-600">
          <div class="w-7 h-7 rounded-full bg-gray-50 flex items-center justify-center text-gray-400"><el-icon><Location /></el-icon></div>
          <span class="font-bold text-gray-700">{{ userInfo.location }}</span>
        </div>
        <!-- 工作经验 -->
        <div class="flex items-center gap-2 text-sm">
          <div class="w-7 h-7 rounded-full flex items-center justify-center"
               :class="userInfo.isStudent ? 'bg-brand-yellow/10 text-brand-yellow-dark' : 'bg-brand-pink/10 text-brand-pink'">
            <el-icon v-if="userInfo.isStudent"><Reading /></el-icon>
            <el-icon v-else><Suitcase /></el-icon>
          </div>
          <div class="flex flex-col leading-tight">
            <span class="text-xs text-gray-400">{{ userInfo.isStudent ? '当前状态' : '工作经验' }}</span>
            <span class="font-bold" :class="userInfo.isStudent ? 'text-brand-yellow-dark' : 'text-brand-dark'">
              {{ userInfo.isStudent ? '在读学生 (未毕业)' : userInfo.experience }}
            </span>
          </div>
        </div>
      </div>
    </div>
<!-- 能力雷达图 -->
    <div class="flex-1 grid grid-cols-1 md:grid-cols-2 grid-rows-2 gap-4 min-h-0">
      
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">六维能力雷达图</h3>
          <p class="text-xs text-gray-400">实线为您的得分，虚线为同岗位行业平均分</p>
        </div>
        <div ref="radarChartRef" class="flex-1 w-full min-h-0"></div>
      </div>
<!-- 人岗初始匹配度预演 -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">人岗初始匹配度预演</h3>
          <p class="text-xs text-gray-400">基于当前解析结果与默认目标岗位（前端开发）的契合度预估</p>
        </div>
        <div ref="gaugeChartRef" class="flex-1 w-full min-h-0"></div>
      </div>
<!-- 核心技能权重图谱 -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">核心技能权重图谱</h3>
          <p class="text-xs text-gray-400">词条大小反映该技能在您过往经历中的权重及贡献率</p>
        </div>
        <div ref="wordCloudRef" class="flex-1 w-full min-h-0"></div>
      </div>
<!-- 职业履历与项目时间轴 -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col h-full">
        <div class="shrink-0 mb-4">
          <h3 class="font-extrabold text-brand-dark text-base">职业履历与项目时间轴</h3>
          <p class="text-xs text-gray-400">AI已自动识别经历连续性与潜在空窗期</p>
        </div>
<!-- 有项目经历 -->
        <div v-if="hasExperiences" class="flex-1 overflow-y-auto pr-4 custom-scrollbar">
          <el-timeline>
            <el-timeline-item center timestamp="2023.09 - 至今" placement="top" color="#8A9E58">
              <div class="bg-brand-green/5 p-3 rounded-lg border border-brand-green/20">
                <h4 class="font-bold text-gray-800 text-sm">某科技公司 - 前端开发实习生</h4>
                <p class="text-xs text-gray-500 mt-1">独立负责 Vue3 后台管理系统模块开发...</p>
              </div>
            </el-timeline-item>
            
            <el-timeline-item center timestamp="2023.05 - 2023.08" placement="top" color="#F59E0B">
              <div class="bg-brand-yellow/10 p-3 rounded-lg border border-brand-yellow/30 flex items-start gap-2">
                <el-icon class="text-brand-yellow mt-0.5"><Warning /></el-icon>
                <div>
                  <h4 class="font-bold text-brand-yellow-dark text-sm">空窗期 (约3个月)</h4>
                  <p class="text-xs text-gray-500 mt-1">AI 提示：可能在后续面试中被问及，建议准备合理解释。</p>
                </div>
              </div>
            </el-timeline-item>

            <el-timeline-item center timestamp="2021.09 - 2023.05" placement="top" color="#8A9E58">
              <div class="bg-gray-50 p-3 rounded-lg border border-gray-100">
                <h4 class="font-bold text-gray-800 text-sm">校内创新创业项目 - 核心成员</h4>
                <p class="text-xs text-gray-500 mt-1">参与“互联网+”大赛，负责前端小程序界面搭建。</p>
              </div>
            </el-timeline-item>
          </el-timeline>
        </div>
<!-- 无项目经历 -->
        <div v-else class="flex-1 flex flex-col items-start justify-start bg-gray-50/80 rounded-2xl border border-dashed border-gray-200 p-5 text-left group transition-colors hover:bg-brand-green/5 hover:border-brand-green/30">
          
          <div class="flex items-center gap-2 mb-2">
            <div class="w-7 h-7 bg-brand-green/10 rounded-full flex items-center justify-center shrink-0 group-hover:scale-110 transition-transform">
              <el-icon class="text-brand-green text-sm"><MagicStick /></el-icon>
            </div>
            <h4 class="text-sm font-extrabold text-brand-dark">全新起点，无限可能 ✨</h4>
          </div>
          
          <p class="text-xs text-gray-500 leading-relaxed pl-9">
            AI 未提取到过往项目经历，这对于学生而言非常正常。<br/>
            在后续环节中，系统将为您重点规划 <span class="text-brand-green font-bold bg-brand-green/10 px-1 rounded">高含金量练手项目</span>，助您完成从 0 到 1 的破局。
          </p>
          
        </div>

      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
// 引用Element Plus 图标
import { 
  Document, Right, Warning, 
  User, Location, Calendar, Suitcase, Reading, Male, Female 
  , MagicStick
} from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import 'echarts-wordcloud'

// ==================== 个人基本信息数据 ====================
// 在真实场景下，这部分数据由你上传简历后AI解析返回
const userInfo = ref({
  name: '林木木',
  gender: 'female', // 'male' 或 'female'
  age: 21,
  location: '浙江省杭州市',
  isStudent: true, // true 表示在读，false 表示已工作
  experience: '应届生' // 如果 isStudent 为 false，这里可以是 '3年' 等
})

// ==================== 图表 DOM 引用 ====================
const radarChartRef = ref(null)
const gaugeChartRef = ref(null)
const wordCloudRef = ref(null)
let charts = []

const initCharts = () => {
  // 1. 六维能力雷达图
  // 1. 六维能力雷达图
  if (radarChartRef.value) {
    const radarChart = echarts.init(radarChartRef.value)
    radarChart.setOption({
      tooltip: { trigger: 'item' },
      
      // 1：将图例改为垂直排列，并固定在图表容器的右侧垂直居中位置
      legend: { 
        orient: 'vertical', 
        right: '30%', 
        top: 'center', 
        data: ['我的画像', '行业平均'], 
        icon: 'circle', 
        itemWidth: 10,
        textStyle: { fontSize: 11 }
      },
      
      radar: {
        indicator: [
          { name: '技术储备', max: 100 }, { name: '沟通协作', max: 100 },
          { name: '项目管理', max: 100 }, { name: '创新思维', max: 100 },
          { name: '执行落地', max: 100 }, { name: '持续学习', max: 100 }
        ],
        radius: '45%',
        
        // 2：中心点为 ['42%', '55%']
        // 42% (X轴向左移)：给右侧的图例让出绝对空间，防止重叠
        // 55% (Y轴向下移)：把上方被切掉的“技术储备”等文字拉回到可视区域内
        center: ['42%', '55%'], 
        
        splitNumber: 4,
        axisName: { color: '#4b5563', fontSize: 10, fontWeight: 'bold' },
        splitArea: { areaStyle: { color: ['#f9fafb', '#ffffff'] } }
      },
      series: [{
        type: 'radar',
        data: [
          { value: [85, 70, 60, 75, 88, 92], name: '我的画像', itemStyle: { color: '#8A9E58' }, areaStyle: { color: 'rgba(138, 158, 88, 0.3)' }, lineStyle: { width: 2 } },
          { value: [75, 80, 70, 65, 80, 85], name: '行业平均', itemStyle: { color: '#9ca3af' }, lineStyle: { type: 'dashed', width: 2 }, areaStyle: { color: 'transparent' } }
        ]
      }]
    })
    charts.push(radarChart)
  }

  // 2. 人岗匹配度预演
  if (gaugeChartRef.value) {
    const gaugeChart = echarts.init(gaugeChartRef.value)
    gaugeChart.setOption({
      series: [{
        type: 'gauge', startAngle: 180, endAngle: 0, center: ['50%', '75%'], radius: '95%', min: 0, max: 100,
        axisLine: { lineStyle: { width: 12, color: [[0.6, '#F87171'], [0.8, '#FBBF24'], [1, '#8A9E58']] } },
        pointer: { icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z', length: '50%', width: 6, offsetCenter: [0, '-10%'], itemStyle: { color: 'auto' } },
        axisTick: { length: 10, lineStyle: { color: 'auto', width: 2 } },
        splitLine: { length: 15, lineStyle: { color: 'auto', width: 3 } },
        axisLabel: { color: '#464646', fontSize: 10, distance: -35 },
        detail: { fontSize: 30, offsetCenter: [0, '-10%'], valueAnimation: true, formatter: '{value}%', color: 'auto', fontWeight: 'bolder' },
        data: [{ value: 68 }]
      }]
    })
    charts.push(gaugeChart)
  }

  // 3. 技能权重词云
  if (wordCloudRef.value) {
    const wcChart = echarts.init(wordCloudRef.value)
    wcChart.setOption({
      tooltip: { show: true },
      series: [{
        type: 'wordCloud', shape: 'circle', width: '100%', height: '100%', sizeRange: [12, 35], rotationRange: [-45, 45], rotationStep: 45, gridSize: 6,
        textStyle: { fontWeight: 'bold', color: () => ['#8A9E58', '#EFDCE2', '#4b5563', '#F7EECD', '#6EE7B7'][Math.floor(Math.random() * 5)] },
        data: [
          { name: 'Vue3', value: 100 }, { name: 'JavaScript', value: 85 }, { name: '快速学习', value: 70 },
          { name: '团队协作', value: 60 }, { name: 'Node.js', value: 55 }, { name: 'TailwindCSS', value: 50 },
          { name: '问题解决', value: 45 }, { name: 'Git', value: 40 }, { name: 'ECharts', value: 35 }
        ]
      }]
    })
    charts.push(wcChart)
  }
}

onMounted(() => {
  nextTick(() => {
    initCharts()
    window.addEventListener('resize', () => charts.forEach(c => c.resize()))
  })
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', () => charts.forEach(c => c.resize()))
  charts.forEach(c => c.dispose())
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background-color: rgba(194, 214, 143, 0.5); border-radius: 4px; }
.custom-scrollbar:hover::-webkit-scrollbar-thumb { background-color: rgba(194, 214, 143, 0.8); }
</style>