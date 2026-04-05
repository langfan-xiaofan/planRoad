<template>
  <div class="h-full flex flex-col overflow-hidden p-2 relative">
    <!-- 无目标时的引导提示 -->
    <transition name="el-fade-in-linear">
      <div v-if="!currentTargetJob" class="absolute inset-0 flex flex-col items-center justify-center bg-white/50 backdrop-blur-sm z-10 p-6">
        <!-- 引导提示图标 -->
        <div class="flex flex-col items-center mb-10 animate-fade-in-up">
          <div class="w-16 h-16 bg-[#F7EECD] rounded-full flex items-center justify-center mb-4 shadow-sm border-2 border-[#C2D68F]">
            <el-icon class="text-3xl text-[#8A9E58]"><Guide /></el-icon>
          </div>
          <!-- 引导提示内容 -->
          <div class="bg-white/90 backdrop-blur-md px-6 py-4 rounded-2xl rounded-tl-none border border-[#EFDCE2] shadow-sm max-w-lg text-center">
            <p class="text-[#4b5563] leading-relaxed text-sm">
              <strong class="text-[#8A9E58] text-base">你的能力雷达我已经吃透啦！✨</strong><br/>
              综合你当前的技能栈，我为你精选了下面这几个最具潜力的发展方向。<br/>
              你想先看看自己和哪个岗位的契合度？
            </p>
          </div>
        </div>
        <!-- 推荐岗位列表 -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 w-full max-w-5xl">
          <div 
            v-for="(job, index) in recommendedJobs" 
            :key="index"
            @click="selectRecommendedJob(job.name)"
            class="group cursor-pointer bg-white rounded-2xl p-6 border border-gray-100 shadow-sm hover:shadow-md hover:-translate-y-1 hover:border-[#C2D68F] transition-all duration-300 relative overflow-hidden flex flex-col"
          >
            <div class="absolute top-0 right-0 bg-[#F7EECD] text-[#8A9E58] text-xs font-bold px-3 py-1 rounded-bl-xl z-10">
              契合度: {{ job.matchScore }}
            </div>
            <!-- 岗位名称 -->
            <h3 class="text-lg font-extrabold text-brand-dark mt-2 mb-2 group-hover:text-[#8A9E58] transition-colors">
              {{ job.name }}
            </h3>
            <div class="flex items-center gap-2 mb-3">
              <el-tag size="small" type="info" effect="plain" class="border-none bg-gray-50">{{ job.type }}</el-tag>
            </div>
            <p class="text-xs text-gray-500 leading-relaxed flex-1">
              <span class="text-[#EFDCE2] font-bold">推荐理由：</span>{{ job.reason }}
            </p>
            <!-- 岗位类型 -->
            <div class="mt-4 flex justify-end opacity-0 group-hover:opacity-100 transition-opacity">
              <el-button circle size="small" class="!bg-[#8A9E58] !text-white !border-none">
                <el-icon><Right /></el-icon>
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </transition>
    <!-- 有目标时的图表 -->
    <!-- 图表标题 -->
    <transition name="el-fade-in-linear">
      <div v-if="currentTargetJob" class="h-full flex flex-col gap-4">
        
        <div class="shrink-0 bg-white rounded-3xl p-5 shadow-sm border border-brand-cream flex justify-between items-center">
          <div>
            <h2 class="text-xl font-extrabold text-brand-dark flex items-center gap-2">
              <el-icon class="text-brand-pink"><Position /></el-icon>
              人岗匹配度与差距分析
            </h2>
            <!-- 当前目标岗位 -->
            <div class="flex items-center gap-3 mt-1">
              <p class="text-sm text-gray-500">当前目标岗位：<span class="font-bold text-[#8A9E58]">{{ currentTargetJob }}</span></p>
              <el-button size="small" type="primary" link @click="clearTarget" class="!text-gray-400 hover:!text-[#F87171]">
                <el-icon class="mr-1"><Close /></el-icon> 重新选择
              </el-button>
            </div>
          </div>
          <el-button round class="!bg-brand-green !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/path')">
            生成职业路径规划 <el-icon class="ml-1"><Right /></el-icon>
          </el-button>
        </div>
        <!-- 图表内容 -->
        <div class="flex-1 grid grid-cols-1 md:grid-cols-2 grid-rows-2 gap-4 min-h-0">
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
            <div class="shrink-0 mb-2">
              <h3 class="font-extrabold text-brand-dark text-base">双基线能力雷达对比图</h3>
              <p class="text-xs text-gray-400">对比您的能力与岗位要求，未覆盖区域即为“技能缺口”</p>
            </div>
            <div ref="radarGapRef" class="flex-1 w-full min-h-0"></div>
          </div>
          <!-- 核心技能缺口瀑布图 -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
            <div class="shrink-0 mb-2">
              <h3 class="font-extrabold text-brand-dark text-base flex items-center gap-1">
                核心技能缺口瀑布图 <el-icon class="text-brand-green"><Aim /></el-icon>
              </h3>
              <p class="text-xs text-gray-400">对比目标岗位，找出当前亟待补齐的最大能力缺口（点击绿色区域生成路径）</p>
            </div>
            <div ref="scatterRoiRef" class="flex-1 w-full min-h-0"></div>
          </div>
          <!-- 同岗候选人竞争力分布 -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
            <div class="shrink-0 mb-2">
              <h3 class="font-extrabold text-brand-dark text-base">同岗候选人竞争力分布</h3>
              <p class="text-xs text-gray-400">红线为您当前所处的百分位，箱体代表市场主流竞争者的水平区间</p>
            </div>
            <div ref="boxplotRef" class="flex-1 w-full min-h-0"></div>
          </div>
          <!-- 匹配得分权重拆解树 -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
            <div class="shrink-0 mb-2">
              <h3 class="font-extrabold text-brand-dark text-base">匹配得分权重拆解树</h3>
              <p class="text-xs text-gray-400">面积越大权重越高。深色模块代表“表现不佳且高权重”的关键瓶颈</p>
            </div>
            <div ref="treemapRef" class="flex-1 w-full min-h-0"></div>
          </div>
        </div>

      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick, watch } from 'vue'
import { Position, Right, Aim, Guide, Close } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
// import { useUserStore } from '@/stores/user'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()

// ==================== 1. 状态与数据源 ====================
const currentTargetJob = ref(route.query.target || '') 

// AI 推荐岗位数据
const recommendedJobs = ref([
  {
    name: '高级前端开发工程师',
    type: '技术研发类',
    matchScore: '85% 极高',
    reason: '您的 Vue3 和工程化能力表现亮眼，非常契合目前主流互联网大厂的前端基建需求。'
  },
  {
    name: '全栈开发工程师',
    type: '综合技术类',
    matchScore: '72% 较高',
    reason: '您具备一定的 Node.js 基础，若能进一步补齐后端性能监控知识，薪资上限将大幅提升。'
  },
  {
    name: '产品经理 (技术向)',
    type: '产品/业务类',
    matchScore: '68% 值得尝试',
    reason: '您的业务沟通能力评分优异，技术转产品是一个能极大发挥您软技能优势的路径。'
  }
])

// ==================== 2. DOM 引用与图表实例 ====================
const radarGapRef = ref(null)
const scatterRoiRef = ref(null)
const boxplotRef = ref(null)
const treemapRef = ref(null)

let radarChart = null
let gapChart = null
let boxChart = null
let treeChart = null
let charts = []

// ==================== 3. 交互逻辑 ====================
const selectRecommendedJob = (jobName) => {
  router.push({ query: { target: jobName } })
}

const clearTarget = () => {
  router.push({ query: {} })
}

// 销毁图表实例（用于从“有目标”切回“无目标”时）
const disposeCharts = () => {
  charts.forEach(c => c.dispose())
  charts = []
  radarChart = gapChart = boxChart = treeChart = null
}

// ==================== 4. 图表初始化与更新 ====================
const performMatching = (jobName) => {
  if (!jobName || !radarChart) return
  
  const newStandard = jobName.includes('前端') 
    ? [85, 80, 90, 75, 85, 80]
    : [80, 75, 85, 70, 85, 80] 

  radarChart.setOption({
    series: [{
      data: [
        {
          value: [85, 60, 40, 50, 70, 90],
          name: '我的当前能力',
          itemStyle: { color: '#8A9E58' },
          areaStyle: { color: 'rgba(138, 158, 88, 0.5)' },
          lineStyle: { width: 2 }
        },
        {
          value: newStandard,
          name: '岗位标准要求',
          itemStyle: { color: '#EFDCE2' },
          areaStyle: { color: 'rgba(239, 220, 226, 0.3)' },
          lineStyle: { width: 2, type: 'dashed' }
        }
      ]
    }]
  })
}
// 初始化图表
const initCharts = () => {
  if (radarGapRef.value && !radarChart) {
    radarChart = echarts.init(radarGapRef.value)
    radarChart.setOption({
      tooltip: { trigger: 'item' },
      legend: { orient: 'vertical', right: '30%', top: 'center', data: ['我的当前能力', '岗位标准要求'], icon: 'circle', itemWidth: 10, textStyle: { fontSize: 11 } },
      radar: {
        indicator: [
          { name: 'Vue3工程化', max: 100 }, { name: '算法与数据结构', max: 100 }, { name: '微前端架构', max: 100 },
          { name: 'Node.js中间层', max: 100 }, { name: '性能优化', max: 100 }, { name: '业务沟通', max: 100 }
        ],
        radius: '45%', center: ['40%', '55%'], splitNumber: 4,
        axisName: { color: '#4b5563', fontSize: 10, fontWeight: 'bold' },
        splitArea: { areaStyle: { color: ['#f9fafb', '#ffffff'] } }
      },
      series: [{ type: 'radar', data: [] }] 
    })
    charts.push(radarChart)
  }
  // 核心技能缺口瀑布图
  if (scatterRoiRef.value && !gapChart) { 
    gapChart = echarts.init(scatterRoiRef.value)
    const rawData = [
      { name: '微前端 (Qiankun)', percent: 30 }, { name: 'Vite 构建优化', percent: 45 },
      { name: 'Node.js 性能监控', percent: 50 }, { name: 'Pinia 状态管理', percent: 70 }
    ]
    const sortedData = rawData.map(item => ({ name: item.name, current: item.percent, gap: 100 - item.percent })).sort((a, b) => a.gap - b.gap)

    gapChart.setOption({
      grid: { top: '15%', bottom: '15%', left: '5%', right: '8%', containLabel: true },
      tooltip: {
        trigger: 'axis', axisPointer: { type: 'shadow' },
        formatter: (params) => {
          const current = params[0].value; const gap = params[1].value;
          return `<div style="font-weight:bold;margin-bottom:4px;">${params[0].name}</div><div style="color:#9ca3af;font-size:12px;">当前掌握: ${current}%</div><div style="color:#8A9E58;font-size:12px;font-weight:bold;">亟待补齐: ${gap}% (点击生成路径)</div>`
        }
      },
      legend: { bottom: 0, icon: 'circle', textStyle: { fontSize: 11 } },
      xAxis: { type: 'value', max: 100, splitLine: { lineStyle: { type: 'dashed' } } },
      yAxis: { type: 'category', data: sortedData.map(d => d.name), axisLabel: { fontWeight: 'bold', color: '#4b5563' }, axisLine: { show: false }, axisTick: { show: false } },
      series: [
        { name: '当前掌握', type: 'bar', stack: '总量', barWidth: '45%', itemStyle: { color: '#f3f4f6', borderRadius: [0, 0, 0, 0] }, data: sortedData.map(d => d.current) },
        { name: '核心缺口 (点击生成路径)', type: 'bar', stack: '总量', itemStyle: { color: '#8A9E58', borderRadius: [0, 4, 4, 0], shadowBlur: 5, shadowColor: 'rgba(138, 158, 88, 0.3)' }, emphasis: { itemStyle: { color: '#EFDCE2', borderColor: '#C2D68F', borderWidth: 2, shadowBlur: 10 } }, label: { show: true, position: 'inside', formatter: '+{c}%', color: '#fff', fontSize: 10 }, data: sortedData.map(d => d.gap) }
      ]
    })
    charts.push(gapChart)
// 核心技能缺口瀑布图点击事件
    gapChart.on('click', (params) => {
      if (params.seriesIndex === 1) {
        ElMessage({ message: `正在为您生成【${params.name}】的专属学习路径...`, type: 'success', duration: 1000, customClass: '!bg-[#F7EECD] !text-[#8A9E58] !border-[#C2D68F]' })
        setTimeout(() => { router.push({ path: '/path', query: { focusSkill: params.name, from: 'match' } }) }, 600)
      }
    })
  }
// 同岗候选人竞争力分布
  if (boxplotRef.value && !boxChart) {
    boxChart = echarts.init(boxplotRef.value)
    boxChart.setOption({
      grid: { top: '20%', bottom: '15%', left: '5%', right: '5%' },
      xAxis: { type: 'value', min: 0, max: 100, splitLine: { lineStyle: { type: 'dashed' } } },
      yAxis: { type: 'category', data: ['综合匹配度'], show: false },
      series: [
        { name: '行业分布', type: 'boxplot', data: [[40, 60, 72, 85, 98]], itemStyle: { color: '#f3f4f6', borderColor: '#9ca3af', borderWidth: 2 } },
        { type: 'scatter', data: [], markLine: { symbol: ['none', 'none'], label: { formatter: '您的位置\n68分', position: 'end', color: '#F87171', fontWeight: 'bold' }, lineStyle: { color: '#F87171', width: 2, type: 'solid' }, data: [{ xAxis: 68 }] } }
      ]
    })
    charts.push(boxChart)
  }
// 匹配得分权重拆解树
  if (treemapRef.value && !treeChart) {
    treeChart = echarts.init(treemapRef.value)
    treeChart.setOption({
      tooltip: { formatter: '{b}: 占比 {c}%' },
      series: [{
        type: 'treemap', width: '100%', height: '100%', roam: false, nodeClick: false, breadcrumb: { show: false }, itemStyle: { borderColor: '#fff', borderWidth: 2, gapWidth: 2 },
        data: [
          { name: '专业技能', value: 45, itemStyle: { color: '#EFDCE2' } }, { name: '项目经验', value: 25, itemStyle: { color: '#8A9E58' } },
          { name: '学历背景', value: 15, itemStyle: { color: '#F7EECD' } }, { name: '软素质', value: 10, itemStyle: { color: '#e5e7eb' } }, { name: '稳定性', value: 5, itemStyle: { color: '#e5e7eb' } }
        ]
      }]
    })
    charts.push(treeChart)
  }
}

// ==================== 5. 监听与生命周期 ====================
const handleResize = () => charts.forEach(c => c.resize())

// 监听路由变化
watch(
  () => route.query.target,
  (newTarget) => {
    currentTargetJob.value = newTarget || ''
    
    if (newTarget) {
      // 状态：变为“有目标” 等 DOM 渲染出来后再初始化图表
      nextTick(() => {
        if (charts.length === 0) {
          initCharts()
          window.addEventListener('resize', handleResize)
        }
        performMatching(newTarget)
      })
    } else {
      // 状态：变为“无目标” 清理图表实例
      disposeCharts()
      window.removeEventListener('resize', handleResize)
    }
  }
)
// 组件挂载时初始化图表
onMounted(() => {
  // 如果首次进入页面时就已经带了参数，直接初始化
  if (currentTargetJob.value) {
    nextTick(() => {
      initCharts()
      performMatching(currentTargetJob.value)
      window.addEventListener('resize', handleResize)
    })
  }
})
// 组件卸载时清理图表实例
onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  disposeCharts()
})
</script>

<style scoped>
/* 保证动画平滑过渡 */
.animate-fade-in-up {
  animation: fadeInUp 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>