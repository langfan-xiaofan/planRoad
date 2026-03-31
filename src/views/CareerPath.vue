<template>
  <div class="h-full flex flex-col gap-4 overflow-hidden p-2">
    
    <div class="shrink-0 bg-white rounded-3xl p-5 shadow-sm border border-brand-cream flex justify-between items-center">
      <div>
        <h2 class="text-xl font-extrabold text-brand-dark flex items-center gap-2">
          <el-icon class="text-brand-green"><Connection /></el-icon>
          定制化职业成长路径规划
        </h2>
        <p class="text-sm text-gray-500 mt-1">基于高 ROI 技能缺口，AI 已为您生成最优学习路径与时间表。</p>
      </div>
      <el-button round class="!bg-brand-green !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/report')">
        生成最终发展报告 <el-icon class="ml-1"><Right /></el-icon>
      </el-button>
    </div>

    <div class="flex-1 grid grid-cols-1 md:grid-cols-2 grid-rows-2 gap-4 min-h-0">
      
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">技能转化与成长链路 (Sankey)</h3>
          <p class="text-xs text-gray-400">现有技能储备 &rarr; 推荐学习项目 &rarr; 达成目标能力</p>
        </div>
        <div ref="sankeyRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">中短期里程碑甘特图</h3>
          <p class="text-xs text-gray-400">时间轴排期：任务的颜色对应不同的技能维度</p>
        </div>
        <div ref="ganttRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">关联岗位转型力导向图</h3>
          <p class="text-xs text-gray-400">中心为当前目标，连线粗细代表技能复用率与转型难度 (可拖拽节点)</p>
        </div>
        <div ref="graphRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">预期能力值增长趋势</h3>
          <p class="text-xs text-gray-400">执行规划后，各维度能力的季度预期累积量</p>
        </div>
        <div ref="stackedBarRef" class="flex-1 w-full min-h-0"></div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { Connection, Right, MagicStick } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()

// DOM 引用
const sankeyRef = ref(null)
const ganttRef = ref(null)
const graphRef = ref(null)
const stackedBarRef = ref(null)

let charts = []
let sankeyChart = null
let ganttChart = null

// 映射字典 
const skillMap = {
  '微前端 (Qiankun)': { sankey: '微前端Qiankun项目', ganttIndex: 1 },
  'Pinia 状态管理': { sankey: 'Vue3工程化实战', ganttIndex: 3 },
  'Vite 构建优化': { sankey: 'Vue3工程化实战', ganttIndex: 3 },
  'Node.js 性能监控': { sankey: 'Node.js性能监控', ganttIndex: 2 }
}

// ==================== 1. 初始化所有图表 ====================
const initCharts = () => {
  // --- Sankey 图表 ---
  if (sankeyRef.value) {
    sankeyChart = echarts.init(sankeyRef.value)
    sankeyChart.setOption({
      tooltip: { trigger: 'item', triggerOn: 'mousemove' },
      series: [{
        type: 'sankey',
        left: '5%', right: '5%', top: '10%', bottom: '5%',
        nodeAlign: 'justify',
        data: [
          { name: 'JavaScript基础', itemStyle: { color: '#e5e7eb' } },
          { name: 'Vue2经验', itemStyle: { color: '#e5e7eb' } },
          { name: 'Vue3工程化实战', itemStyle: { color: '#8A9E58' } },
          { name: '微前端Qiankun项目', itemStyle: { color: '#F7EECD' } },
          { name: 'Node.js性能监控', itemStyle: { color: '#EFDCE2' } },
          { name: '高级前端架构', itemStyle: { color: '#4b5563' } },
          { name: '全栈思维', itemStyle: { color: '#4b5563' } }
        ],
        links: [
          { source: 'JavaScript基础', target: 'Node.js性能监控', value: 3 },
          { source: 'JavaScript基础', target: 'Vue3工程化实战', value: 2 },
          { source: 'Vue2经验', target: 'Vue3工程化实战', value: 4 },
          { source: 'Vue2经验', target: '微前端Qiankun项目', value: 2 },
          { source: 'Vue3工程化实战', target: '高级前端架构', value: 6 },
          { source: '微前端Qiankun项目', target: '高级前端架构', value: 3 },
          { source: 'Node.js性能监控', target: '全栈思维', value: 4 }
        ],
        lineStyle: { color: 'source', curveness: 0.5, opacity: 0.3 },
        label: { fontSize: 11, color: '#374151', fontWeight: 'bold' },
        emphasis: { 
          focus: 'adjacency', 
          itemStyle: { color: '#EFDCE2', borderColor: '#C2D68F', borderWidth: 2, shadowBlur: 15, shadowColor: 'rgba(194, 214, 143, 0.8)' },
          lineStyle: { color: '#C2D68F', opacity: 0.8 }
        }
      }]
    })
    charts.push(sankeyChart)
  }

  // --- Gantt 图表 ---
  if (ganttRef.value) {
    ganttChart = echarts.init(ganttRef.value)
    ganttChart.setOption({
      grid: { left: '3%', right: '4%', bottom: '10%', top: '10%', containLabel: true },
      xAxis: { type: 'value', name: '计划周期 (周)', nameLocation: 'middle', nameGap: 20, splitLine: { lineStyle: { type: 'dashed' } } },
      yAxis: { type: 'category', data: ['算法打卡', '微前端搭建', 'Node服务开发', 'Vue3源码剖析'], axisTick: { show: false }, axisLine: { show: false } },
      tooltip: { formatter: (params) => { if (params.seriesIndex === 1) return `${params.name}<br/>持续: ${params.value} 周` } },
      series: [
        { name: '占位', type: 'bar', stack: '总量', itemStyle: { borderColor: 'transparent', color: 'transparent' }, data: [1, 5, 2, 0] },
        { 
          name: '持续时间', type: 'bar', stack: '总量', barWidth: '50%',
          itemStyle: { borderRadius: 4, color: (params) => ['#e5e7eb', '#F7EECD', '#EFDCE2', '#8A9E58'][params.dataIndex] },
          emphasis: { itemStyle: { shadowBlur: 10, shadowColor: 'rgba(194, 214, 143, 0.8)', borderColor: '#8A9E58', borderWidth: 2 } },
          data: [10, 4, 3, 5] 
        }
      ]
    })
    charts.push(ganttChart)
    // 点击甘特图柱子，推迟任务节点 2 周
    ganttChart.on('click', (params) => {
      if (params.seriesIndex === 1) { 
        const currentOption = ganttChart.getOption()
        const startTimes = currentOption.series[0].data

        startTimes[params.dataIndex] += 2 
        ganttChart.setOption({ series: [{ data: startTimes }, {}] })

        sessionStorage.setItem('plan_delayed', 'true')

        ElMessage({
          message: '⚠️ 任务节点已推迟！系统已将延迟效应同步至【发展报告】的薪资预测模型中。',
          type: 'warning',
          duration: 3000,
          customClass: '!bg-[#EFDCE2]/90 !text-[#F87171] !border-[#EFDCE2] font-bold shadow-lg'
        })
      }
    })
  }

  // --- Graph 图表 ---
  if (graphRef.value) {
    const graphChart = echarts.init(graphRef.value)
    graphChart.setOption({
      tooltip: {},
      series: [{
        type: 'graph', layout: 'force', roam: true, label: { show: true, position: 'bottom', fontWeight: 'bold' },
        force: { repulsion: 250, edgeLength: 80 },
        data: [
          { name: '高级前端', symbolSize: 50, itemStyle: { color: '#8A9E58' } },
          { name: '全栈开发', symbolSize: 35, itemStyle: { color: '#EFDCE2' } },
          { name: '前端架构师', symbolSize: 40, itemStyle: { color: '#F7EECD' } },
          { name: 'Node研发', symbolSize: 30, itemStyle: { color: '#e5e7eb' } },
          { name: 'Web3前端', symbolSize: 25, itemStyle: { color: '#e5e7eb' } }
        ],
        links: [
          { source: '高级前端', target: '前端架构师', value: 80, lineStyle: { width: 4 } },
          { source: '高级前端', target: '全栈开发', value: 60, lineStyle: { width: 3 } },
          { source: '高级前端', target: 'Node研发', value: 45, lineStyle: { width: 2 } },
          { source: '高级前端', target: 'Web3前端', value: 30, lineStyle: { width: 1 } },
          { source: '全栈开发', target: 'Node研发', value: 70, lineStyle: { width: 3 } }
        ]
      }]
    })
    charts.push(graphChart)
  }
  
  // --- Stacked Bar 图表 ---
  if (stackedBarRef.value) {
    const stackedBarChart = echarts.init(stackedBarRef.value)
    stackedBarChart.setOption({
      tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
      legend: { bottom: 0, icon: 'roundRect', itemWidth: 12 },
      grid: { left: '3%', right: '4%', bottom: '15%', top: '10%', containLabel: true },
      xAxis: { type: 'category', data: ['Q1', 'Q2', 'Q3', 'Q4'] },
      yAxis: { type: 'value', splitLine: { lineStyle: { type: 'dashed' } } },
      series: [
        { name: '硬技能增长', type: 'bar', stack: 'total', barWidth: '40%', itemStyle: { color: '#8A9E58' }, data: [15, 30, 45, 60] },
        { name: '项目经验累积', type: 'bar', stack: 'total', itemStyle: { color: '#EFDCE2' }, data: [10, 25, 40, 50] },
        { name: '软实力提升', type: 'bar', stack: 'total', itemStyle: { color: '#F7EECD' }, data: [5, 10, 15, 25] }
      ]
    })
    charts.push(stackedBarChart)
  }
}

// ==================== 2. 处理跨页面穿透联动逻辑 (心跳呼吸版) ====================
const handleLinkage = async () => {
  await nextTick() 
  const targetSkill = route.query.focusSkill

  if (targetSkill && skillMap[targetSkill]) {
    const mapped = skillMap[targetSkill]

    // 1. 弹出话术提示 (层级调高防止被覆盖)
    setTimeout(() => {
      ElMessage({
        message: '我们的系统不仅能发现差距，还能直接生成填补差距的具体路径 ✨',
        type: 'success',
        icon: MagicStick,
        duration: 4000,
        customClass: '!bg-[#F7EECD] !text-[#8A9E58] !border-[#C2D68F] shadow-lg font-bold z-[9999]' 
      })
    }, 300)

    // 2. 视觉锚点闭环 - 心跳呼吸闪烁特效
    let blinkCount = 0
    // 使用 setInterval 让高亮状态每隔 400ms 切换一次 (一明一暗)
    const blinkTimer = setInterval(() => {
      const actionType = blinkCount % 2 === 0 ? 'highlight' : 'downplay'
      
      // 桑基图与甘特图同步闪烁
      if (sankeyChart) sankeyChart.dispatchAction({ type: actionType, name: mapped.sankey })
      if (ganttChart) ganttChart.dispatchAction({ type: actionType, seriesIndex: 1, dataIndex: mapped.ganttIndex })
      
      blinkCount++
      
      // 闪烁 3 次（即执行 6 轮切换）后停止，并最终保持在常亮高亮状态
      if (blinkCount > 5) {
        clearInterval(blinkTimer)
        if (sankeyChart) {
          sankeyChart.dispatchAction({ type: 'highlight', name: mapped.sankey })
          sankeyChart.dispatchAction({ type: 'showTip', name: mapped.sankey })
        }
        if (ganttChart) {
          ganttChart.dispatchAction({ type: 'highlight', seriesIndex: 1, dataIndex: mapped.ganttIndex })
          ganttChart.dispatchAction({ type: 'showTip', seriesIndex: 1, dataIndex: mapped.ganttIndex })
        }
      }
    }, 400) // 呼吸频率：400毫秒
  }
}

// ==================== 3. 生命周期挂载 ====================
onMounted(() => {
  nextTick(() => {
    initCharts() 
    handleLinkage() 
    window.addEventListener('resize', () => charts.forEach(c => c.resize()))
  })
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', () => charts.forEach(c => c.resize()))
  charts.forEach(c => c.dispose())
})
</script>