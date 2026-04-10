<template>
  <!-- 模拟基于 AI 推荐的热门岗位数据 -->
  <div class="space-y-4 h-full overflow-y-auto -mr-6 pr-6 pb-4">
    <div class="grid grid-cols-4 gap-4">
      <el-card shadow="hover" class="border-none rounded-lg" body-class="p-4">
        <div class="flex justify-between items-center">
          <div class="text-gray-500 text-sm">已收录岗位</div>
          <el-icon class="text-themePrimary" :size="20"><Suitcase /></el-icon>
        </div>
        <div class="text-3xl font-bold text-gray-800 mt-2">10,245</div>
        <div class="text-xs text-green-500 mt-2 flex items-center">
          <el-icon><Top /></el-icon> <span>较上周增加 4%</span>
        </div>
      </el-card>
      <!-- 行业平均薪资 -->
      <el-card shadow="hover" class="border-none rounded-lg" body-class="p-4">
        <div class="flex justify-between items-center">
          <div class="text-gray-500 text-sm">行业平均薪资</div>
          <el-icon class="text-themePrimary" :size="20"><Money /></el-icon>
        </div>
        <div class="text-3xl font-bold text-gray-800 mt-2">12.5k</div>
        <div class="text-xs text-green-500 mt-2 flex items-center">
          <el-icon><Top /></el-icon> <span>环比上涨 2%</span>
        </div>
      </el-card>
      <!-- 热门换岗路径 -->
      <el-card shadow="hover" class="border-none rounded-lg" body-class="p-4">
        <div class="flex justify-between items-center">
          <div class="text-gray-500 text-sm">热门换岗路径</div>
          <el-icon class="text-themeAccent" :size="20"><Connection /></el-icon>
        </div>
        <div class="text-3xl font-bold text-gray-800 mt-2">2,315</div>
        <div class="text-xs text-red-400 mt-2 flex items-center">
          <el-icon><Bottom /></el-icon> <span>变动率 1.2%</span>
        </div>
      </el-card>
      <!-- 解析技能标签 -->
      <el-card shadow="hover" class="border-none rounded-lg" body-class="p-4">
        <div class="flex justify-between items-center">
          <div class="text-gray-500 text-sm">解析技能标签</div>
          <el-icon class="text-themePrimary" :size="20"><CollectionTag /></el-icon>
        </div>
        <div class="text-3xl font-bold text-gray-800 mt-2">1,284</div>
        <div class="text-xs text-green-500 mt-2 flex items-center">
          <el-icon><Top /></el-icon><span>特征库持续扩充中</span>
        </div>
      </el-card>
    </div>
    <!-- 近半年热门行业岗位需求趋势 -->
    <el-card shadow="hover" class="border-none rounded-lg">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="font-bold text-gray-700">近半年热门行业岗位需求趋势</span>
          <el-select v-model="trendTime" size="small" class="w-32">
            <el-option label="近半年" value="6m" />
            <el-option label="近一年" value="1y" />
          </el-select>
        </div>
      </template>
      <div ref="trendChartRef" class="h-[280px] w-full"></div>
    </el-card>
    <!-- 企业核心技能需求 Top5 -->
    <div class="grid grid-cols-3 gap-4">
      <el-card shadow="hover" class="border-none rounded-lg h-full">
        <template #header><span class="font-bold text-gray-700">企业核心技能需求 Top5</span></template>
        <div class="space-y-4 pt-2">
          <div v-for="(item, index) in skills" :key="index">
            <div class="flex justify-between text-sm mb-1">
              <span class="text-gray-600">{{ item.name }}</span>
              <span class="text-gray-400">{{ item.count }} 岗位要求</span>
            </div>
            <el-progress :percentage="item.percent" :color="item.color" :show-text="false" />
          </div>
        </div>
      </el-card>
      <!-- 库内岗位行业分布 -->
      <el-card shadow="hover" class="border-none rounded-lg h-full">
        <template #header><span class="font-bold text-gray-700">库内岗位行业分布</span></template>
        <div ref="pieChartRef" class="h-[200px] w-full"></div>
      </el-card>
      <!-- 我的规划进度 -->
      <el-card shadow="hover" class="border-none rounded-lg h-full bg-themeBg/50">
        <template #header><span class="font-bold text-gray-700">我的规划进度</span></template>
        <div class="flex flex-col items-center justify-center h-[180px]">
          <el-progress type="dashboard" :percentage="30" color="#84B4C3" :width="120">
            <template #default="{ percentage }">
              <span class="block text-2xl font-bold text-themePrimary">{{ percentage }}%</span>
              <span class="text-xs text-gray-500">画像完成度</span>
            </template>
          </el-progress>
          <p class="text-xs text-gray-500 mt-2 text-center px-4">
            你还有部分能力数据未完善，会影响 AI 推荐的精准度。
          </p>
          <el-button type="primary" size="small" class="bg-themePrimary border-none mt-3" @click="$router.push('/profile')">
            去完善画像
          </el-button>
        </div>
      </el-card>
    </div>
    <!-- 新手友好：近期热门岗位推荐 -->
    <div class="grid grid-cols-3 gap-4 pb-6">
      <el-card shadow="hover" class="col-span-2 border-none rounded-lg">
  <template #header>
    <div class="flex justify-between items-center">
      <span class="font-bold text-gray-700"> 🔥  新手友好：近期热门岗位推荐</span>
      <el-link :underline="false" class="text-themePrimary text-sm" @click="$router.push('/jobs')">
        查看更多 <el-icon><ArrowRight /></el-icon>
      </el-link>
    </div>
  </template>
  <div class="space-y-0">
    <div v-for="(job, index) in recommendedJobs" :key="index" class="flex justify-between items-center py-3 border-b border-gray-100 last:border-0 hover:bg-gray-50 transition px-2 cursor-pointer rounded">
      <div>
        <div class="font-bold text-gray-800 text-sm mb-1">{{ job.title }}</div>
        <div class="text-xs text-gray-500">{{ job.company }} · {{ job.exp }} · {{ job.edu }}</div>
      </div>
      <div class="text-right">
        <div class="text-themeAccent font-bold">{{ job.salary }}</div>
        <div class="flex gap-1 mt-1 justify-end">
          <el-tag 
            v-for="(skill, sIndex) in job.skills.slice(0, 3)" 
            :key="sIndex" 
            size="small" 
            :color="sIndex % 2 === 0 ? '#FAEDE5' : '#C7E2E0'" 
            :class="sIndex % 2 === 0 ? 'text-gray-600 border-none' : 'text-themePrimary border-none'"
          >
            {{ skill }}
          </el-tag>
        </div>
      </div>
    </div>
  </div>
</el-card>
      <!-- 新手任务清单 -->
      <el-card shadow="hover" class="col-span-1 border-none rounded-lg">
        <template #header><span class="font-bold text-gray-700">新手任务清单</span></template>
        <div class="space-y-3">
          <div class="flex items-start gap-2">
            <el-checkbox v-model="todo1" />
            <span :class="{'line-through text-gray-400': todo1, 'text-sm text-gray-700': !todo1}">上传一份最新的个人简历解析</span>
          </div>
          <div class="flex items-start gap-2">
            <el-checkbox v-model="todo2" />
            <span :class="{'line-through text-gray-400': todo2, 'text-sm text-gray-700': !todo2}">查看一条适合自己的换岗路径</span>
          </div>
          <div class="flex items-start gap-2">
            <el-checkbox v-model="todo3" />
            <span :class="{'line-through text-gray-400': todo3, 'text-sm text-gray-700': !todo3}">生成第一份 AI 职业生涯报告</span>
          </div>
           <div class="flex items-start gap-2">
            <el-checkbox v-model="todo4" />
            <span :class="{'line-through text-gray-400': todo4, 'text-sm text-gray-700': !todo4}">向 AI 助手提问关于行业前景的问题</span>
          </div>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import * as echarts from 'echarts'
import { Suitcase, Money, Connection, UserFilled, Top, Bottom, ArrowRight, CollectionTag } from '@element-plus/icons-vue'

const trendTime = ref('6m')
const trendChartRef = ref(null)
const pieChartRef = ref(null)

// 任务清单状态
const todo1 = ref(true)
const todo2 = ref(false)
const todo3 = ref(false)
const todo4 = ref(false)

// 技能数据
const skills = [
  { name: 'JavaScript/TypeScript', count: '4,521', percent: 85, color: '#84B4C3' },
  { name: 'Vue.js / React', count: '3,890', percent: 70, color: '#C7E2E0' },
  { name: 'Java SpringBoot', count: '3,100', percent: 55, color: '#EAC0C3' },
  { name: 'Python 数据分析', count: '1,850', percent: 35, color: '#84B4C3' },
  { name: 'UI/UX 设计规范', count: '1,200', percent: 20, color: '#C7E2E0' }
]

// 模拟基于 AI 推荐的热门岗位数据
const recommendedJobs = ref([
  { 
    title: '初级前端开发工程师 (Vue方向)', 
    company: '网易(杭州)网络有限公司', 
    exp: '1-3年', edu: '本科', salary: '12k - 18k', 
    skills: ['Vue3', '前端工程化'] 
  },
  { 
    title: '数据分析师 (AI方向)', 
    company: '阿里大文娱', 
    exp: '应届生', edu: '硕士', salary: '15k - 25k', 
    skills: ['Python', 'SQL', '机器学习'] 
  },
  { 
    title: 'UI交互设计师', 
    company: '字节跳动', 
    exp: '1-3年', edu: '本科', salary: '14k - 22k', 
    skills: ['Figma', '用户体验设计'] 
  },
  { 
    title: '后端开发工程师 (Go)', 
    company: '腾讯科技', 
    exp: '1年以内', edu: '本科', salary: '16k - 25k', 
    skills: ['Go', '微服务', 'MySQL'] 
  }
])

onMounted(() => {
  nextTick(() => {
    initTrendChart()
    initPieChart()
    
    // 监听窗口缩放，重置图表大小
    window.addEventListener('resize', () => {
      echarts.getInstanceByDom(trendChartRef.value)?.resize()
      echarts.getInstanceByDom(pieChartRef.value)?.resize()
    })
  })
})
// 初始化趋势图表
const initTrendChart = () => {
  const chart = echarts.init(trendChartRef.value)
  chart.setOption({
    tooltip: { trigger: 'axis', axisPointer: { type: 'cross' } },
    legend: { data: ['人工智能', '软件开发'], top: 0, right: 0 },
    grid: { left: '3%', right: '4%', bottom: '3%', top: '15%', containLabel: true },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: ['1月', '2月', '3月', '4月', '5月', '6月'],
        axisLine: { lineStyle: { color: '#ccc' } }
      }
    ],
    yAxis: [{ type: 'value', splitLine: { lineStyle: { type: 'dashed', color: '#eee' } } }],
    series: [
      {
        name: '人工智能',
        type: 'line',
        smooth: true,
        lineStyle: { width: 0 },
        showSymbol: false,
        areaStyle: {
          opacity: 0.8,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#84B4C3' },
            { offset: 1, color: 'rgba(132, 180, 195, 0.1)' }
          ])
        },
        emphasis: { focus: 'series' },
        data: [120, 280, 150, 450, 220, 310]
      },
      {
        name: '软件开发',
        type: 'line',
        smooth: true,
        lineStyle: { width: 0 },
        showSymbol: false,
        areaStyle: {
          opacity: 0.6,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#C7E2E0' },
            { offset: 1, color: 'rgba(199, 226, 224, 0.1)' }
          ])
        },
        emphasis: { focus: 'series' },
        data: [220, 180, 350, 150, 380, 210]
      }
    ]
  })
}
// 初始化行业分布图表
const initPieChart = () => {
  const chart = echarts.init(pieChartRef.value)
// 模拟行业分布数据
  const rawData = [
    { value: 850, name: '软件开发' },
    { value: 420, name: '人工智能' },
    { value: 310, name: '数据分析' },
    { value: 250, name: '产品设计' },
    { value: 170, name: '运维测试' },
    { value: 120, name: '电子商务' },
    { value: 80, name: '游戏开发' },
    { value: 45, name: '区块链' },
    { value: 30, name: '物联网' }
  ]
  // 对原始数据进行降序排序
  rawData.sort((a, b) => b.value - a.value)

  // Top5降维算法：只保留前5个，其余合并为“其他领域”
  let processedData = []
  if (rawData.length > 5) {
    processedData = rawData.slice(0, 5)
    const othersValue = rawData.slice(5).reduce((sum, item) => sum + item.value, 0)
    processedData.push({ value: othersValue, name: '其他领域' })
  } else {
    processedData = rawData
  }

  chart.setOption({
    tooltip: { trigger: 'item' },
    legend: { show: false }, // 隐藏图例让图表更大
    color: ['#84B4C3', '#C7E2E0', '#EAC0C3', '#FAEDE5', '#d1d5db', '#f3f4f6'], // 最后一个颜色留给“其他”
    series: [
      {
        name: '行业分布',
        type: 'pie',
        radius: ['50%', '80%'], // 调整为环形图，更现代
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 4,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: { show: false, position: 'center' },
        emphasis: {
          label: { show: true, fontSize: 14, fontWeight: 'bold' }
        },
        labelLine: { show: false },
        color: ['#84B4C3', '#C7E2E0', '#EAC0C3', '#FAEDE5', '#d1d5db'], // 严格使用 P3 配色
        data: processedData
      }
    ]
  })
}
</script>