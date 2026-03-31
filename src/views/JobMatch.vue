<template>
  <div class="h-full flex flex-col gap-4 overflow-hidden p-2">
    
    <div class="shrink-0 bg-white rounded-3xl p-5 shadow-sm border border-brand-cream flex justify-between items-center">
      <div>
        <h2 class="text-xl font-extrabold text-brand-dark flex items-center gap-2">
          <el-icon class="text-brand-pink"><Position /></el-icon>
          人岗匹配度与差距分析
        </h2>
        <p class="text-sm text-gray-500 mt-1">当前目标岗位：<span class="font-bold text-brand-dark">高级前端开发工程师</span> | 核心缺口已标红预警</p>
      </div>
      <el-button round class="!bg-brand-green !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/path')">
        生成职业路径规划 <el-icon class="ml-1"><Right /></el-icon>
      </el-button>
    </div>

    <div class="flex-1 grid grid-cols-1 md:grid-cols-2 grid-rows-2 gap-4 min-h-0">
      
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">双基线能力雷达对比图</h3>
          <p class="text-xs text-gray-400">对比您的能力与岗位要求，未覆盖区域即为“技能缺口”</p>
        </div>
        <div ref="radarGapRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base flex items-center gap-1">
            核心技能缺口瀑布图 <el-icon class="text-brand-green"><Aim /></el-icon>
          </h3>
          <p class="text-xs text-gray-400">对比目标岗位，找出当前亟待补齐的最大能力缺口（点击绿色区域生成路径）</p>
        </div>
        <div ref="scatterRoiRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">同岗候选人竞争力分布</h3>
          <p class="text-xs text-gray-400">红线为您当前所处的百分位，箱体代表市场主流竞争者的水平区间</p>
        </div>
        <div ref="boxplotRef" class="flex-1 w-full min-h-0"></div>
      </div>

      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">匹配得分权重拆解树</h3>
          <p class="text-xs text-gray-400">面积越大权重越高。深色模块代表“表现不佳且高权重”的关键瓶颈</p>
        </div>
        <div ref="treemapRef" class="flex-1 w-full min-h-0"></div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { Position, Right,Aim } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import { useUserStore } from '@/stores/user'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()
// DOM 引用
const radarGapRef = ref(null)
const scatterRoiRef = ref(null)
const boxplotRef = ref(null)
const treemapRef = ref(null)

let charts = []

const initCharts = () => {
  // ==================== 1. 双基线雷达图 ====================
  if (radarGapRef.value) {
    const radarChart = echarts.init(radarGapRef.value)
    radarChart.setOption({
      tooltip: { trigger: 'item' },
      
      // 将图例垂直排列，并固定在图表容器的右侧
      legend: { 
        orient: 'vertical', 
        right: '30%', 
        top: 'center', 
        data: ['我的当前能力', '岗位标准要求'], 
        icon: 'circle', 
        itemWidth: 10,
        textStyle: { fontSize: 11 }
      },
      
      radar: {
        indicator: [
          { name: 'Vue3工程化', max: 100 },
          { name: '算法与数据结构', max: 100 },
          { name: '微前端架构', max: 100 },
          { name: 'Node.js中间层', max: 100 },
          { name: '性能优化', max: 100 },
          { name: '业务沟通', max: 100 }
        ],
        
        // 缩小半径，调整中心点位置
        radius: '45%', // 强制缩小雷达图本体，给四周文字腾出空间
        center: ['40%', '55%'], // 40% (向左偏)：给右侧图例让出空间；55% (向下偏)：防止顶部文字被切
        
        splitNumber: 4,
        axisName: { color: '#4b5563', fontSize: 10, fontWeight: 'bold' },
        splitArea: { areaStyle: { color: ['#f9fafb', '#ffffff'] } }
      },
      series: [{
        type: 'radar',
        data: [
          {
            value: [85, 60, 40, 50, 70, 90],
            name: '我的当前能力',
            itemStyle: { color: '#8A9E58' },
            areaStyle: { color: 'rgba(138, 158, 88, 0.5)' },
            lineStyle: { width: 2 }
          },
          {
            value: [80, 75, 85, 70, 85, 80],
            name: '岗位标准要求',
            itemStyle: { color: '#EFDCE2' },
            areaStyle: { color: 'rgba(239, 220, 226, 0.3)' },
            lineStyle: { width: 2, type: 'dashed' }
          }
        ]
      }]
    })
    charts.push(radarChart)
  }

  // ==================== 2. 核心技能缺口图  ====================
  if (scatterRoiRef.value) { 
    const gapChart = echarts.init(scatterRoiRef.value)
    const rawData = [
      { name: '微前端 (Qiankun)', current: 30, target: 80 },
      { name: 'Vite 构建优化', current: 45, target: 85 },
      { name: 'Node.js 性能监控', current: 50, target: 80 },
      { name: 'Pinia 状态管理', current: 70, target: 90 }
    ]
    // 提取 X、Y 轴数据，按缺口大小降序排列，缺口最大的在最上面
    const sortedData = rawData.map(item => ({
      ...item,
      gap: item.target - item.current
    })).sort((a, b) => a.gap - b.gap) // ECharts 横向柱图默认是从下往上画的

    gapChart.setOption({
      grid: { top: '15%', bottom: '15%', left: '5%', right: '8%', containLabel: true },
      tooltip: {
        trigger: 'axis',
        axisPointer: { type: 'shadow' },
        formatter: (params) => {
          const current = params[0].value
          const gap = params[1].value
          return `
            <div style="font-weight:bold;margin-bottom:4px;">${params[0].name}</div>
            <div style="color:#9ca3af;font-size:12px;">当前掌握: ${current}%</div>
            <div style="color:#8A9E58;font-size:12px;font-weight:bold;">亟待补齐: ${gap}% (点击生成路径)</div>
          `
        }
      },
      legend: { bottom: 0, icon: 'circle', textStyle: { fontSize: 11 } },
      xAxis: { type: 'value', max: 100, splitLine: { lineStyle: { type: 'dashed' } } },
      yAxis: { 
        type: 'category', 
        data: sortedData.map(d => d.name),
        axisLabel: { fontWeight: 'bold', color: '#4b5563' },
        axisLine: { show: false },
        axisTick: { show: false }
      },
      series: [
        {
          name: '当前掌握',
          type: 'bar',
          stack: '总量',
          barWidth: '45%',
          itemStyle: { color: '#f3f4f6', borderRadius: [0, 0, 0, 0] },
          data: sortedData.map(d => d.current)
        },
        {
          name: '核心缺口 (点击生成路径)',
          type: 'bar',
          stack: '总量',
          itemStyle: { 
            color: '#8A9E58', 
            borderRadius: [0, 4, 4, 0],
            shadowBlur: 5,
            shadowColor: 'rgba(138, 158, 88, 0.3)'
          },
          // 悬停高亮特效，强调可点击
          emphasis: {
            itemStyle: {
              color: '#EFDCE2', 
              borderColor: '#C2D68F',
              borderWidth: 2,
              shadowBlur: 10
            }
          },
          label: {
            show: true,
            position: 'inside',
            formatter: '+{c}%',
            color: '#fff',
            fontSize: 10
          },
          data: sortedData.map(d => d.gap)
        }
      ]
    })
    charts.push(gapChart)

    // 绑定点击事件 (穿透逻辑无缝迁移)
    gapChart.on('click', (params) => {
      // 限制只有点击缺口柱子时才跳转 点击灰色“当前掌握”柱子不跳转
      if (params.seriesIndex === 1) {
        const clickedSkill = params.name
        
        // 添加一点加载延迟动画让跳转更自然
        ElMessage({
          message: `正在为您生成【${clickedSkill}】的专属学习路径，请稍候...`,
          type: 'success',
          duration: 1000,
          customClass: '!bg-[#F7EECD] !text-[#8A9E58] !border-[#C2D68F]'
        })
        // 延迟 600 毫秒后跳转，避免画面闪烁过快
        setTimeout(() => {
          router.push({ 
            path: '/path', 
            query: { focusSkill: clickedSkill, from: 'match' } 
          })
        }, 600)
      }
    })
  }
  // ==================== 3. 竞争力分布箱线图 ====================
  if (boxplotRef.value) {
    const boxChart = echarts.init(boxplotRef.value)
    boxChart.setOption({
      grid: { top: '20%', bottom: '15%', left: '5%', right: '5%' },
      xAxis: { type: 'value', min: 0, max: 100, splitLine: { lineStyle: { type: 'dashed' } } },
      yAxis: { type: 'category', data: ['综合匹配度'], show: false },
      series: [
        {
          name: '行业分布',
          type: 'boxplot',
          data: [
            // [最小值, 下四分位数(Q1), 中位数, 上四分位数(Q3), 最大值]
            [40, 60, 72, 85, 98]
          ],
          itemStyle: { color: '#f3f4f6', borderColor: '#9ca3af', borderWidth: 2 }
        },
        {
          // 用 markLine 标出当前用户的水平
          type: 'scatter',
          data: [],
          markLine: {
            symbol: ['none', 'none'],
            label: { formatter: '您的位置\n68分', position: 'end', color: '#F87171', fontWeight: 'bold' },
            lineStyle: { color: '#F87171', width: 2, type: 'solid' },
            data: [{ xAxis: 68 }]
          }
        }
      ]
    })
    charts.push(boxChart)
  }

  // ==================== 4. 匹配要素贡献树 (Treemap) ====================
  if (treemapRef.value) {
    const treeChart = echarts.init(treemapRef.value)
    treeChart.setOption({
      tooltip: { formatter: '{b}: 占比 {c}%' },
      series: [{
        type: 'treemap',
        width: '100%',
        height: '100%',
        roam: false,
        nodeClick: false,
        breadcrumb: { show: false },
        itemStyle: { borderColor: '#fff', borderWidth: 2, gapWidth: 2 },
        data: [
          // 颜色逻辑：表现越差且权重越高的，用深粉色/红色警示。表现好的用绿色或奶黄色。
          { name: '专业技能', value: 45, itemStyle: { color: '#EFDCE2' } }, // 权重最大，且有欠缺
          { name: '项目经验', value: 25, itemStyle: { color: '#8A9E58' } }, // 表现好
          { name: '学历背景', value: 15, itemStyle: { color: '#F7EECD' } },
          { name: '软素质', value: 10, itemStyle: { color: '#e5e7eb' } },
          { name: '稳定性', value: 5, itemStyle: { color: '#e5e7eb' } }
        ]
      }]
    })
    charts.push(treeChart)
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