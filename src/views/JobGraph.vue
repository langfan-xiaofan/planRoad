<!-- 岗位关联关系图谱 -->
<template>
  <div class="flex gap-6 h-full min-h-0">
    <div class="w-1/3 bg-white p-4 rounded-xl shadow-sm flex flex-col">
      <el-input placeholder="搜索岗位，例如：Java开发" class="mb-4" clearable></el-input>
      <el-scrollbar class="flex-1">
        <div v-for="i in 10" :key="i" class="p-3 border-b border-gray-100 hover:bg-themeBg cursor-pointer transition-colors">
          <div class="font-bold text-gray-800">高级前端工程师</div>
          <div class="text-sm text-gray-500 mt-1">薪资：15k-25k | 科技行业</div>
        </div>
      </el-scrollbar>
    </div>
<!-- 岗位核心要求分析与关联关系图谱切换 -->
    <div class="w-2/3 bg-white p-4 rounded-xl shadow-sm flex flex-col">
      <el-tabs v-model="activeTab" @tab-change="handleTabChange" class="flex-1 flex flex-col h-full">
        <!-- 岗位核心要求分析 -->
        <el-tab-pane label="岗位核心要求分析" name="profile" class="h-full flex flex-col">
          <div class="p-2 flex-1 flex flex-col h-full">
            <div class="flex justify-between items-end mb-4 shrink-0">
              <h3 class="text-xl font-bold text-themePrimary">岗位能力需求矩阵</h3>
              <div class="text-xs text-gray-400">区块面积代表该项能力在真实企业JD中的权重占比</div>
            </div>
            <!-- 标签云 -->
            <div class="flex flex-wrap gap-2 mb-4 shrink-0">
              <el-tag color="#FAEDE5" class="text-themePrimary border-themePrimary">Vue3精通</el-tag>
              <el-tag color="#FAEDE5" class="text-themePrimary border-themePrimary">工程化构建</el-tag>
              <el-tag color="#EAC0C3" class="text-white border-none">抗压能力强</el-tag>
            </div>
            <div ref="treemapChartRef" class="w-full flex-1 min-h-0"></div>
          </div>
        </el-tab-pane>
        <!-- 换岗关联图谱 -->
        <el-tab-pane label="换岗关联图谱" name="graph" class="h-full flex flex-col">
          <div class="p-2 flex-1 flex flex-col h-full">
            <h3 class="text-xl font-bold text-themePrimary mb-2 shrink-0">职业血缘关系网络</h3>
            <div ref="relationGraphRef" class="w-full flex-1 min-h-0"></div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'

const activeTab = ref('profile')
const treemapChartRef = ref(null)
const relationGraphRef = ref(null)

// 声明全局实例变量，用于生命周期管理与图表重绘
let treemapInstance = null
let relationInstance = null
let resizeObserver = null

onMounted(() => {
  nextTick(() => {
    // 默认优先初始化首屏展示的图表
    initTreemapChart()
    setupResizeObserver()
  })
})

// 监听Tabs切换，解决Echarts在隐藏容器中初始化导致宽度为0的经典Bug
const handleTabChange = (tabName) => {
  nextTick(() => {
    if (tabName === 'profile') {
      if (!treemapInstance) {
        initTreemapChart()
      } else {
        treemapInstance.resize()
      }
    } else if (tabName === 'graph') {
      if (!relationInstance) {
        initRelationGraph()
      } else {
        relationInstance.resize()
      }
    }
  })
}

// 设置全局容器尺寸监听器
const setupResizeObserver = () => {
  resizeObserver = new ResizeObserver(() => {
    if (treemapInstance && activeTab.value === 'profile') {
      treemapInstance.resize()
    }
    if (relationInstance && activeTab.value === 'graph') {
      relationInstance.resize()
    }
  })
  
  if (treemapChartRef.value) resizeObserver.observe(treemapChartRef.value)
  if (relationGraphRef.value) resizeObserver.observe(relationGraphRef.value)
}

// 初始化矩形树图(替代原有的雷达图)
const initTreemapChart = () => {
  if (!treemapChartRef.value) return
  if (treemapInstance) treemapInstance.dispose()
  
  treemapInstance = echarts.init(treemapChartRef.value)
  
  // 模拟基于大模型解析出的岗位权重数据，后续替换为后端API对岗位JD的拆解数据
  const treemapData = [
    {
      name: '专业硬技能',
      itemStyle: { color: '#84B4C3' },
      children: [
        { name: '前端框架生态(Vue3/React)', value: 40 },
        { name: '工程化与构建(Vite/Webpack)', value: 25 },
        { name: 'TypeScript强类型开发', value: 15 },
        { name: 'Node.js服务端渲染', value: 10 }
      ]
    },
    {
      name: '通用软素质',
      itemStyle: { color: '#EAC0C3' },
      children: [
        { name: '复杂业务逻辑拆解能力', value: 20 },
        { name: '跨部门协同与沟通', value: 15 },
        { name: '高压环境抗压能力', value: 10 }
      ]
    },
    {
      name: '加分项',
      itemStyle: { color: '#C7E2E0' },
      children: [
        { name: '开源社区代码贡献', value: 10 },
        { name: '英语文档无障碍阅读', value: 5 }
      ]
    }
  ]

  treemapInstance.setOption({
    tooltip: {
      formatter: function (info) {
        const value = info.value
        const name = info.name
        return `<div style="font-weight:bold;color:#333;">${name}</div>需求权重占比: ${value}%`
      }
    },
    series: [{
      type: 'treemap',
      width: '100%',
      height: '100%',
      top: '0%', // 覆盖Echarts默认边距，强制图表顶部贴合容器，实现整体上移
      bottom: '0%', // 强制图表底部贴合容器，避免超出高度被截断显示不全
      left: '0%',
      right: '0%',
      roam: false,
      nodeClick: false,
      breadcrumb: { show: false },
      label: {
        show: true,
        formatter: '{b}\n\n权重:{c}%',
        fontSize: 13,
        fontWeight: 'bold',
        color: '#fff'
      },
      itemStyle: {
        borderColor: '#FAEDE5',
        borderWidth: 2,
        gapWidth: 2
      },
      data: treemapData
    }]
  })
}

// 初始化关联关系图谱
const initRelationGraph = () => {
  if (!relationGraphRef.value) return
  if (relationInstance) relationInstance.dispose()
  
  relationInstance = echarts.init(relationGraphRef.value)
  
  // 模拟满足任务要求的关联数据,后续替换为后端图数据库或关联算法生成的真实数据
  const graphData = [
    { name: '前端开发(当前)', symbolSize: 60, itemStyle: { color: '#84B4C3' } },
    { name: '全栈工程师', symbolSize: 40, itemStyle: { color: '#C7E2E0' } },
    { name: 'Node开发', symbolSize: 40, itemStyle: { color: '#C7E2E0' } },
    { name: 'UI设计', symbolSize: 40, itemStyle: { color: '#EAC0C3' } },
    { name: '产品经理', symbolSize: 40, itemStyle: { color: '#FAEDE5' }, label: { color: '#666' } }
  ]
  
  const graphLinks = [
    { source: '前端开发(当前)', target: '全栈工程师', label: { show:true, formatter: '补充后端知识', fontSize: 11 } },
    { source: '前端开发(当前)', target: 'Node开发', label: { show:true, formatter: '深化服务端运行环境', fontSize: 11 } },
    { source: '前端开发(当前)', target: 'UI设计', label: { show:true, formatter: '强化审美与交互', fontSize: 11 } },
    { source: 'UI设计', target: '前端开发(当前)', label: { show:true, formatter: '学习CSS框架', fontSize: 11 } },
    { source: '前端开发(当前)', target: '产品经理', label: { show:true, formatter: '提升业务统筹', fontSize: 11 } }
  ]

  relationInstance.setOption({
    tooltip: {},
    series: [{
      type: 'graph',
      layout: 'force',
      data: graphData,
      links: graphLinks,
      roam: true, // 保持开启漫游缩放和平移功能
      scaleLimit: { // 设定缩放界限阈值，防止图表过度缩放变形
        min: 0.6, // 限制最小缩放倍率为0.6倍，解决图表缩成一个点的问题
        max: 3    // 限制最大缩放倍率为3倍，防止无限放大超出容器边界
      },
      label: { show: true, position: 'right', fontWeight: 'bold' },
      force: { repulsion: 600, edgeLength: [120, 220] },
      lineStyle: { color: 'source', curveness: 0.2, width: 2 }
    }]
  })
}

// 组件卸载时销毁实例，防止内存泄漏
onBeforeUnmount(() => {
  if (resizeObserver) {
    resizeObserver.disconnect()
  }
  if (treemapInstance) treemapInstance.dispose()
  if (relationInstance) relationInstance.dispose()
})
</script>

<style scoped>
/* 穿透 Element Plus 组件，强行打通 Flex 高度继承链 */
:deep(.el-tabs) {
  display: flex;
  flex-direction: column;
}

:deep(.el-tabs__content) {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0; /* 防止内容撑破 flex 容器 */
}

:deep(.el-tab-pane) {
  flex: 1;
  display: flex;
  flex-direction: column;
}
</style>