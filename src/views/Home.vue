<template>
  <div class="h-full flex flex-col gap-4 overflow-hidden">
  <!--  首页内容  -->
    <transition name="el-fade-in" mode="out-in">
      <!--  已登录用户内容  -->
      <div v-if="hasData" key="with-data" class="shrink-0 bg-white rounded-3xl p-6 shadow-sm flex flex-col md:flex-row justify-between items-center gap-4 relative overflow-hidden">
        <div class="absolute -right-10 -top-10 w-40 h-40 bg-brand-pink/20 rounded-full blur-3xl"></div>
        <!--  用户信息  -->
        <div class="flex items-center gap-5 z-10">
          <div class="w-14 h-14 rounded-full bg-brand-cream flex items-center justify-center text-brand-green shrink-0 shadow-inner">
            <el-icon :size="28"><StarFilled /></el-icon>
          </div>
          <div>
            <h2 class="text-xl font-extrabold text-brand-dark">欢迎回来，{{ userName }}！ 👋 </h2>
            <p class="text-sm text-gray-500 mt-1">
              当前目标：<span class="font-bold text-brand-green bg-brand-green/10 px-2 py-0.5 rounded-md">{{ targetJob }}</span>
              <span v-if="daysSinceLastVisit > 0" class="text-xs text-brand-yellow ml-2 inline-flex items-center">
                <el-icon class="mr-1"><Clock /></el-icon> {{ daysSinceLastVisit }} 天未更新
              </span>
            </p>
          </div>
        </div>
        <!--  继续规划按钮  -->
        <el-button
          round
          class="!bg-brand-green !border-brand-green !text-white hover:!bg-brand-green-dark shadow-md shadow-brand-green/30 whitespace-nowrap z-10 font-bold px-6 h-10"
          @click="continuePlanning"
        >
          继续规划路径 <el-icon class="ml-1"><Right /></el-icon>
        </el-button>
      </div>
      <!--  未登录用户内容  -->
      <div v-else key="no-data" class="shrink-0 bg-gradient-to-r from-brand-cream to-white rounded-3xl p-6 shadow-sm border border-white flex flex-col md:flex-row justify-between items-center relative overflow-hidden">
        <div class="absolute left-0 bottom-0 w-32 h-32 bg-brand-yellow/20 rounded-full blur-2xl"></div>
        <!--  欢迎信息  -->
        <div class="z-10 flex items-center gap-5">
          <div class="w-14 h-14 rounded-full bg-white flex items-center justify-center text-brand-yellow shrink-0 shadow-sm">
            <el-icon :size="28"><Sunrise /></el-icon>
          </div>
          <div>
            <h2 class="text-2xl font-extrabold text-brand-dark mb-1">早上好，{{ userName }}！ ✨ </h2>
            <p class="text-gray-500 text-sm">你的职业探索之旅即将开始。基于AI大模型，量身定制从认知到入职的全链路规划。</p>
          </div>
        </div>
        <!--  按钮  -->
        <div class="flex gap-3 z-10 mt-4 md:mt-0">
          <el-button round class="!text-brand-green !border-brand-green hover:!bg-brand-green/10 font-bold" @click="startDemo">观看演示</el-button>
          <el-button round class="!bg-brand-green !text-white border-none hover:opacity-90 shadow-md shadow-brand-green/30 font-bold" @click="$router.push('/resume')">开始初次评估</el-button>
        </div>
      </div>
    </transition>
    <!--  仪表盘内容  -->
    <transition name="el-fade-in" mode="out-in">
      
      <div v-if="hasData" key="dashboard-view" class="flex-1 flex flex-col gap-4 min-h-0">
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 shrink-0">
          <!--  简历解析状态  -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-brand-cream/50 flex justify-between items-center relative overflow-hidden group hover:shadow-md transition-shadow cursor-pointer" @click="navigate('/report')">
            <div class="flex flex-col z-10">
              <span class="text-gray-400 text-xs font-bold tracking-wider uppercase mb-1">全息能力聚合</span>
              <div class="text-2xl font-extrabold text-brand-dark mb-1">3维 <span class="text-sm font-normal text-gray-400">对齐</span></div>
              <div class="text-xs text-brand-green flex items-center bg-brand-green/10 px-2 py-0.5 rounded-full w-max">
                <span class="w-1.5 h-1.5 bg-brand-green rounded-full mr-1 animate-pulse"></span> 增量推演中
              </div>
            </div>
            <div ref="resumeChartRef" class="w-20 h-20 z-10 -mr-2"></div>
          </div>
          <!--  高匹配岗位  -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border border-brand-cream/50 flex flex-col relative overflow-hidden group hover:shadow-md transition-shadow">
            <div class="flex justify-between items-start z-10">
              <div>
                <span class="text-gray-400 text-xs font-bold tracking-wider uppercase mb-1 block">高匹配岗位</span>
                <div class="text-3xl font-extrabold text-brand-dark leading-none">12 <span class="text-sm font-normal text-gray-400">个</span></div>
              </div>
              <div class="text-xs text-brand-pink bg-brand-pink/10 px-2 py-1 rounded-md font-bold">> 85%</div>
            </div>
            <div ref="matchChartRef" class="w-full h-10 mt-auto z-10"></div>
          </div>
          <!--  职场竞争力评分  -->
          <div class="bg-brand-green rounded-3xl p-5 shadow-md shadow-brand-green/20 text-white flex flex-col justify-between relative overflow-hidden">
            <span class="text-white/80 text-xs font-bold tracking-wider uppercase relative z-10">职场竞争力评分</span>
            <div class="flex items-end justify-between relative z-10 mt-2">
              <div class="text-4xl font-extrabold leading-none">88</div>
              <div class="text-xs text-white/90 bg-white/20 px-2 py-1 rounded-md backdrop-blur-sm">前 24%</div>
            </div>
            <div ref="scoreChartRef" class="absolute bottom-[-10px] right-[-10px] w-32 h-24 opacity-60"></div>
          </div>
          <!--  AI 建议行动  -->
          <div class="bg-white rounded-3xl p-5 shadow-sm border-2 border-brand-yellow/50 flex flex-col cursor-pointer hover:bg-brand-yellow/5 hover:border-brand-yellow transition-all group" @click="navigate('/path')">
            <span class="text-brand-yellow text-xs font-bold tracking-wider uppercase mb-2 flex items-center">
              <el-icon class="mr-1"><MagicStick /></el-icon> AI 建议行动
            </span>
            <div class="text-sm font-bold text-brand-dark leading-snug flex-1">
              建议在近期补充 <span class="text-brand-green bg-brand-green/10 px-1 rounded">Vue3 工程化</span> 相关开源项目经验。
            </div>
            <div class="mt-2 flex items-center text-xs font-bold text-brand-yellow group-hover:text-brand-yellow-dark transition-colors">
              立即查看专属路径 
              <el-icon class="ml-1 transform group-hover:translate-x-1 transition-transform"><Right /></el-icon>
            </div>
          </div>
        </div>
        <!--  功能模块  -->
        <div class="flex-1 grid grid-cols-1 md:grid-cols-3 gap-4 min-h-0">
          <div v-for="(feature, index) in features" :key="index"
               class="bg-white rounded-3xl p-6 border border-brand-cream/50 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 cursor-pointer flex flex-col relative overflow-hidden group"
               @click="navigate(feature.path)"
          >
            <div class="absolute inset-x-0 bottom-0 h-1 bg-gradient-to-r from-brand-pink to-brand-green opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
            <!--  图标  -->
            <div class="w-12 h-12 rounded-2xl bg-brand-cream/50 flex items-center justify-center text-brand-green mb-4 group-hover:bg-brand-green group-hover:text-white transition-colors">
              <el-icon :size="24"><component :is="feature.icon" /></el-icon>
            </div>
            <!--  标题  -->
            <h3 class="text-lg font-extrabold text-brand-dark mb-2">{{ feature.title }}</h3>
            <p class="text-gray-500 text-sm leading-relaxed flex-1">{{ feature.desc }}</p>
            <!--  按钮  -->
            <div class="mt-4 flex justify-end">
              <div class="w-8 h-8 rounded-full bg-gray-50 flex items-center justify-center text-gray-400 group-hover:bg-brand-cream group-hover:text-brand-green transition-colors">
                <el-icon><Right /></el-icon>
              </div>
            </div>
          </div>
        </div>
  
      </div>
      <!--  工作台内容  -->
      <div v-else key="workbench-view" class="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-4 min-h-0">
        
        <div class="lg:col-span-8 bg-white rounded-3xl border border-brand-cream/50 shadow-sm p-8 flex flex-col relative overflow-hidden">
          <div class="absolute right-0 top-0 w-64 h-64 bg-brand-green/5 rounded-bl-full pointer-events-none"></div>
          <!--  标题  -->
          <h3 class="text-xl font-extrabold text-brand-dark mb-2">新手起步工作台</h3>
          <p class="text-sm text-gray-500 mb-8">只需简单三步，AI 将为你生成全息职业路径推演图。</p>
          <!--  步骤  -->
          <div class="flex-1 flex flex-col gap-6 relative z-10 overflow-y-auto custom-scrollbar pr-2">
            
            <div class="flex gap-4 group cursor-pointer" @click="navigate('/resume')">
              <div class="flex flex-col items-center">
                <div class="w-10 h-10 rounded-full bg-brand-green text-white flex items-center justify-center font-bold shadow-md shadow-brand-green/30 group-hover:scale-110 transition-transform">1</div>
                <div class="w-0.5 h-full bg-brand-cream mt-2 group-hover:bg-brand-green/30 transition-colors"></div>
              </div>
              <div class="flex-1 bg-gray-50 rounded-2xl p-5 border border-gray-100 group-hover:border-brand-green/50 group-hover:bg-white transition-all pb-6">
                <h4 class="font-extrabold text-brand-dark flex items-center gap-2 mb-2"><el-icon class="text-brand-green"><Document /></el-icon> 上传简历解析</h4>
                <p class="text-sm text-gray-500">让 AI 快速了解你的过往经历、技能储备，生成六维能力雷达图基线。</p>
                <div class="mt-4"><el-button size="small" round class="!text-brand-green !border-brand-green bg-white">立即上传</el-button></div>
              </div>
            </div>
            <!--  步骤  -->
            <div class="flex gap-4 group opacity-70">
              <div class="flex flex-col items-center">
                <div class="w-10 h-10 rounded-full bg-brand-cream text-brand-dark flex items-center justify-center font-bold">2</div>
                <div class="w-0.5 h-full bg-brand-cream mt-2"></div>
              </div>
              <div class="flex-1 bg-gray-50 rounded-2xl p-5 border border-gray-100 pb-6">
                <h4 class="font-extrabold text-gray-700 flex items-center gap-2 mb-2"><el-icon><Position /></el-icon> 人岗匹配对比</h4>
                <p class="text-sm text-gray-500">选择你的理想职位，AI 将对比差距，计算各项技能的投资回报率 (ROI)。</p>
              </div>
            </div>
            
            <div class="flex gap-4 group opacity-70">
              <div class="flex flex-col items-center">
                <div class="w-10 h-10 rounded-full bg-brand-cream text-brand-dark flex items-center justify-center font-bold">3</div>
              </div>
              <div class="flex-1 bg-gray-50 rounded-2xl p-5 border border-gray-100">
                <h4 class="font-extrabold text-gray-700 flex items-center gap-2 mb-2"><el-icon><Connection /></el-icon> 获取路径报告</h4>
                <p class="text-sm text-gray-500">导出属于你的定制化职业路线图，包含短期、中长期里程碑任务。</p>
              </div>
            </div>
          </div>
        </div>
        <!--  工作台内容  -->
        <div class="lg:col-span-4 flex flex-col gap-4">
          
          <div class="flex-1 bg-brand-pink/10 rounded-3xl border border-brand-pink/30 p-6 flex flex-col items-center justify-center text-center relative overflow-hidden">
            <el-icon class="text-5xl text-brand-pink mb-4 opacity-80"><MagicStick /></el-icon>
            <h4 class="font-extrabold text-brand-dark mb-2">不知道从何开始？</h4>
            <p class="text-sm text-gray-500 mb-6">你可以随时召唤悬浮的 AI 规划助手，通过对话的方式让它帮你寻找职业方向。</p>
            <el-button @click="openAiChat" round class="!bg-brand-pink !text-white !border-none shadow-md hover:opacity-90">
              唤起 AI 助手 <el-icon class="ml-1"><ChatDotSquare /></el-icon>
            </el-button>
          </div>
          <!--  数据概览  -->
          <div class="shrink-0 bg-white rounded-3xl border border-brand-cream/50 shadow-sm p-6">
            <h4 class="text-sm font-extrabold text-gray-400 uppercase tracking-wider mb-4">平台实时数据</h4>
            <div class="flex justify-between items-center">
              <div class="text-center">
                <div class="text-2xl font-black text-brand-green">10w+</div>
                <div class="text-xs text-gray-500 mt-1">已服务用户</div>
              </div>
              <div class="w-px h-8 bg-brand-cream/80"></div>
              <div class="text-center">
                <div class="text-2xl font-black text-brand-dark">85%</div>
                <div class="text-xs text-gray-500 mt-1">平均匹配度</div>
              </div>
            </div>
          </div>
        </div>
  
      </div>
    </transition>
  </div>
  </template>
  
  <script setup>
  import { computed, ref, watch, nextTick, onBeforeUnmount } from 'vue'
  import { Document, TrendCharts, Connection, StarFilled, Right, Clock, Trophy, MagicStick, Position, Sunrise, ChatDotSquare } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/stores/user'
  import { useJobDataStore } from '@/stores/jobData'
  import DemoTourManager from '@/utils/demoTour'
  import * as echarts from 'echarts'
  // 引入全局状态
  import { useAiChat } from '../components/ai/useAiChat'


  // ==================== 全局状态与路由配置 ====================
  const router = useRouter()
  const userStore = useUserStore()
  const jobStore = useJobDataStore()
  let demoManager = null
  
  // ==================== 计算属性：视图逻辑判断 ====================
  // 获取当前用户名，若无则默认为访客
  const userName = computed(() => userStore.userName || '访客')
  
  // 核心判断逻辑：决定显示“新手引导(false)”还是“数据概览(true)”
  // 依赖于用户是否完成了简历解析，或是否走完了初始入职流程
  const hasData = computed(() => jobStore.hasResumeProfile || userStore.hasCompletedOnboarding)
  
  // 提取用户当前的求职目标
  const onboardingData = computed(() => userStore.getOnboardingData())
  const targetJob = computed(() => jobStore.targetJobLabel || onboardingData.value?.target || '高级前端架构师')
  
  // 模拟数据：距离上次访问的天数
  const daysSinceLastVisit = computed(() => 3)
  
  // ==================== ECharts 图表初始化与生命周期管理 ====================
  // 获取 DOM 容器的引用
  const resumeChartRef = ref(null)
  const matchChartRef = ref(null)
  const scoreChartRef = ref(null)
  let charts = [] // 统一存储图表实例以便管理销毁
  
  // 图表渲染主函数
  const initCharts = () => {
    // 1. 初始化全息聚合雷达图 (动态合成：简历 + 标准 + 增量)
    if (resumeChartRef.value) {
      const resumeChart = echarts.init(resumeChartRef.value)
      resumeChart.setOption({
        tooltip: {
          show: true,
          confine: true, // 防止提示框在小卡片里被截断
          backgroundColor: 'rgba(255, 255, 255, 0.95)',
          textStyle: { fontSize: 10, color: '#374151' },
          extraCssText: 'border-radius: 8px; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); padding: 6px 10px;',
          formatter: () => {
            return `
              <div style="font-weight:bold;margin-bottom:4px;">多维能力对齐</div>
              <div style="display:flex;align-items:center;gap:4px;"><span style="width:6px;height:6px;border-radius:50%;background:#e5e7eb;"></span> 简历原始底座</div>
              <div style="display:flex;align-items:center;gap:4px;"><span style="width:6px;height:6px;border-radius:50%;background:#EFDCE2;"></span> 岗位目标基线</div>
              <div style="display:flex;align-items:center;gap:4px;"><span style="width:6px;height:6px;border-radius:50%;background:#8A9E58;"></span> 规划后成长增量</div>
            `
          }
        },
        radar: {
          indicator: [
            { name: '技能', max: 100 }, { name: '经验', max: 100 }, { name: '架构', max: 100 },
            { name: '算法', max: 100 }, { name: '沟通', max: 100 }, { name: '视野', max: 100 }
          ],
          radius: '75%', // 撑满整个小容器
          center: ['50%', '50%'],
          axisName: { show: false }, // 微型图表隐藏文字标签
          splitNumber: 3,
          splitArea: { show: false },
          splitLine: { lineStyle: { color: '#f3f4f6', type: 'solid' } },
          axisLine: { show: false }
        },
        series: [{
          type: 'radar',
          symbol: 'none', // 隐藏拐点小圆圈，保持微型图表极简
          data: [
            {
              // Layer 1: 简历页原始能力 (灰色实心)
              value: [45, 55, 30, 60, 75, 40],
              name: '原始能力',
              itemStyle: { color: '#e5e7eb' },
              areaStyle: { color: '#f3f4f6' },
              lineStyle: { width: 1, color: '#d1d5db' },
              z: 1
            },
            {
              // Layer 2: 匹配页岗位标准 (粉色虚线外框，表示缺口)
              value: [85, 80, 85, 75, 80, 85],
              name: '岗位标准',
              itemStyle: { color: '#EFDCE2' },
              lineStyle: { width: 2, type: 'dashed', color: '#EFDCE2' },
              areaStyle: { color: 'transparent' }, // 透明，只显示外框
              z: 3
            },
            {
              // Layer 3: 路径页成长增量 (抹茶绿实心覆盖，完美填补差距)
              value: [85, 80, 85, 80, 85, 85],
              name: '成长增量',
              itemStyle: { color: '#8A9E58' },
              areaStyle: { color: 'rgba(138, 158, 88, 0.4)' },
              lineStyle: { width: 2, color: '#8A9E58' },
              z: 2
            }
          ]
        }]
      })
      charts.push(resumeChart)
    }
  
    // 2. 初始化高匹配岗位微型柱状图
    if (matchChartRef.value) {
      const matchChart = echarts.init(matchChartRef.value)
      matchChart.setOption({
        grid: { top: 5, bottom: 0, left: 0, right: 0 },
        xAxis: { type: 'category', show: false, data: ['1','2','3','4','5','6'] },
        yAxis: { type: 'value', show: false },
        tooltip: { show: false },
        series: [{
          type: 'bar',
          data: [65, 70, 85, 92, 88, 95], // 伪造的趋势数据
          itemStyle: { 
            color: '#EFDCE2',
            borderRadius: [2, 2, 0, 0] 
          },
          barWidth: '60%'
        }]
      })
      charts.push(matchChart)
    }
  
    // 3. 初始化职场竞争力半圆仪表盘
    if (scoreChartRef.value) {
      const scoreChart = echarts.init(scoreChartRef.value)
      scoreChart.setOption({
        series: [{
          type: 'gauge',
          startAngle: 180,
          endAngle: 0,
          center: ['50%', '75%'],
          radius: '100%',
          min: 0,
          max: 100,
          pointer: { show: false },
          progress: {
            show: true,
            overlap: false,
            roundCap: true,
            clip: false,
            itemStyle: { color: '#ffffff' }
          },
          axisLine: {
            lineStyle: { width: 6, color: [[1, 'rgba(255,255,255,0.2)']] }
          },
          splitLine: { show: false },
          axisTick: { show: false },
          axisLabel: { show: false },
          detail: { show: false },
          data: [{ value: 88 }] // 评分数据
        }]
      })
      charts.push(scoreChart)
    }
  }
  
  // 监听数据状态变化，确保图表仅在有数据视图渲染后初始化
  watch(hasData, async (newVal) => {
    if (newVal) {
      await nextTick() // 等待 Vue 完成 DOM 结构更新
      initCharts()     // DOM 就绪后执行 ECharts 挂载
    } else {
      // 切换到无数据视图时，及时销毁实例释放内存
      charts.forEach(chart => chart.dispose())
      charts = []
    }
  }, { immediate: true })
  
  // 监听窗口大小变化以重绘图表，保持响应式布局
  const resizeCharts = () => {
    charts.forEach(chart => chart.resize())
  }
  window.addEventListener('resize', resizeCharts)
  
  // 组件卸载前清除事件监听并销毁图表
  onBeforeUnmount(() => {
    window.removeEventListener('resize', resizeCharts)
    charts.forEach(chart => chart.dispose())
  })
  
  // ==================== 业务交互方法与静态配置 ====================
  // 触发页面漫游引导
  const startDemo = () => {
    if (!demoManager) demoManager = new DemoTourManager(router)
    demoManager.start()
  }
  
  // 顶部“继续规划”按钮的路由分发
  const continuePlanning = () => {
    if (hasData.value) router.push('/path')
    else router.push('/resume')
  }
  
  // 通用路由跳转方法
  const navigate = (path) => {
    router.push(path)
  }
  
  // 底部快捷入口模块的静态数据配置
  const features = [
    { title: '智能简历解析', desc: '上传简历，AI 自动拆解能力画像，挖掘潜藏优势。', icon: Document, path: '/resume' },
    { title: '精准人岗匹配', desc: '多维对比目标岗位的契合度与差距，ROI 收益一目了然。', icon: Position, path: '/match' },
    { title: '职业路径图谱', desc: '可视化呈现晋升路线，生成里程碑计划，清晰规划未来。', icon: Connection, path: '/path' },
  ]

  const { openAiChat } = useAiChat()
  </script>
  
  <style scoped>
  /* 视图切换的渐隐渐入过渡动画 */
  .el-fade-in-enter-active,
  .el-fade-in-leave-active {
    transition: opacity 0.3s cubic-bezier(0.16, 1, 0.3, 1), transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  .el-fade-in-enter-from,
  .el-fade-in-leave-to {
    opacity: 0;
    transform: translateY(10px) scale(0.99);
  }
  
  /* 左侧步骤流容器的自定义滚动条样式 */
  .custom-scrollbar::-webkit-scrollbar {
    width: 4px;
  }
  .custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
  }
  .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: rgba(194, 214, 143, 0.5);
    border-radius: 4px;
  }
  .custom-scrollbar:hover::-webkit-scrollbar-thumb {
    background-color: rgba(194, 214, 143, 0.8);
  }
  </style>