<template>
  <div class="h-full flex flex-col gap-4 overflow-hidden p-2" id="report-container">
    <!--  报告标题  -->
    <div class="shrink-0 bg-white rounded-3xl p-5 shadow-sm border border-brand-cream flex justify-between items-center relative overflow-hidden">
      <div class="absolute -right-20 -top-20 w-64 h-64 bg-gradient-to-br from-brand-green/20 to-brand-pink/20 rounded-full blur-3xl"></div>
      <div class="z-10">
        <h2 class="text-xl font-extrabold text-brand-dark flex items-center gap-2">
          <el-icon class="text-brand-pink"><Trophy /></el-icon>
          AI 职业全景发展与风险预测报告
        </h2>
        <p class="text-sm text-gray-500 mt-1">
          根据您的执行计划，预计 <span class="text-brand-green font-bold bg-brand-green/10 px-1 rounded">12 个月</span> 后可达到 
          <span class="text-brand-dark font-bold">高级前端架构师</span> 标准。
        </p>
      </div>
      <!--  操作按钮  -->
      <div class="flex gap-3 z-10">
        <el-button round class="!text-brand-green !border-brand-green hover:!bg-brand-green/10 font-bold shadow-sm"
         :loading="pdfStore.isExporting"
         @click="pdfStore.exportToPdf('report-container')"
        >导出 PDF 报告</el-button>
        <el-button round class="!bg-brand-dark !text-white !border-none shadow-md hover:opacity-90 font-bold px-6" @click="$router.push('/')">
          完成规划，回到首页
        </el-button>
      </div>
    </div>
    <!--  报告内容  -->
    <div class="flex-1 grid grid-cols-1 md:grid-cols-2 grid-rows-2 gap-4 min-h-0">
      <!--  终极能力指纹 (Rose)  -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">终极能力指纹 (Rose)</h3>
          <p class="text-xs text-gray-400">综合前三页得分加权计算出的最终结构占比</p>
        </div>
        <div ref="roseChartRef" class="flex-1 w-full min-h-0"></div>
      </div>
      <!--  三年期薪资与能力复合增长预测  -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border flex flex-col relative h-full transition-all duration-700"
           :class="isDelayed ? 'border-red-200 ring-2 ring-red-50 shadow-red-100/50' : 'border-gray-100'">
        <div class="shrink-0 mb-5 flex justify-between items-start">
          <div>
            <h3 class="font-extrabold text-brand-dark text-base flex items-center gap-1">
              三年期薪资与能力复合增长预测
              <el-tooltip v-if="isDelayed" content="受任务推迟影响，预测模型已下调" placement="top">
                <el-icon class="text-red-500 animate-pulse"><Warning /></el-icon>
              </el-tooltip>
            </h3>
            <p class="text-xs transition-colors duration-500" :class="isDelayed ? 'text-red-500 font-bold' : 'text-gray-400'">
              {{ isDelayed ? '⚠️ 拖延导致复利流失，第3年预计月薪缩水 9,000 元' : '柱体为预期薪水(K)，折线为综合能力指数，虚线为未来预测' }}
            </p>
          </div>
        </div>
        <div ref="comboChartRef" class="flex-1 w-full min-h-0"></div>
      </div>
      <!--  技能贬值风险矩阵  -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2 flex justify-between items-start">
          <div>
            <h3 class="font-extrabold text-brand-dark text-base flex items-center gap-1">
              技能贬值风险矩阵 <el-icon class="text-brand-yellow"><Warning /></el-icon>
            </h3>
            <p class="text-xs text-gray-400">颜色越红代表该技能在未来被 AI 替代或淘汰的风险越高</p>
          </div>
        </div>
        <div ref="heatmapRef" class="flex-1 w-full min-h-0"></div>
      </div>
      <!--  学习向薪资转化的漏斗模型  -->
      <div class="bg-white rounded-3xl p-5 shadow-sm border border-gray-100 flex flex-col relative h-full">
        <div class="shrink-0 mb-2">
          <h3 class="font-extrabold text-brand-dark text-base">学习向薪资转化的漏斗模型</h3>
          <p class="text-xs text-gray-400">展示从“认知”到“变现”的转化率，评估规划可落地性</p>
        </div>
        <div ref="funnelRef" class="flex-1 w-full min-h-0"></div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { Trophy, Warning, InfoFilled } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
// 引入通知组件
import { ElNotification } from 'element-plus'
import { useExportPdfStore } from '@/stores/exportPdf'

const pdfStore = useExportPdfStore()

const roseChartRef = ref(null)
const comboChartRef = ref(null)
const heatmapRef = ref(null)
const funnelRef = ref(null)

let charts = []
// 响应式状态：是否发生了拖延
const isDelayed = ref(false)

const initCharts = () => {
  // ==================== 1. 南丁格尔玫瑰图 ====================
  if (roseChartRef.value) {
    const roseChart = echarts.init(roseChartRef.value)
    roseChart.setOption({
      tooltip: { trigger: 'item', formatter: '{b} : {c} ({d}%)' },
      legend: { orient: 'vertical', right: '5%', top: 'center', icon: 'circle', itemWidth: 10, textStyle: { fontSize: 11 } },
      series: [{
        name: '能力指纹', type: 'pie', radius: ['20%', '65%'], center: ['40%', '50%'], roseType: 'area', 
        itemStyle: { borderRadius: 4 }, label: { show: true, formatter: '{b}' },
        data: [
          { value: 40, name: '前端工程化', itemStyle: { color: '#8A9E58' } },
          { value: 33, name: '全栈/Node', itemStyle: { color: '#b5c58a' } },
          { value: 28, name: '架构与性能', itemStyle: { color: '#EFDCE2' } },
          { value: 22, name: '业务沟通', itemStyle: { color: '#F7EECD' } },
          { value: 18, name: '创新认知', itemStyle: { color: '#e5e7eb' } }
        ]
      }]
    })
    charts.push(roseChart)
  }

  // ==================== 2. 双轴预测混合图 (核心联动) ====================
  if (comboChartRef.value) {
    const comboChart = echarts.init(comboChartRef.value)
    
    // 理想数据 vs 拖延后的数据
    const idealSalary = [15, 22, 28, 35]
    const delayedSalary = [15, 18, 22, 26] // 薪资缩水
    
    const idealAbility = [65, { value: 78, lineStyle: { type: 'dashed' } }, { value: 88, lineStyle: { type: 'dashed' } }, { value: 95, lineStyle: { type: 'dashed' } }]
    const delayedAbility = [65, { value: 70, lineStyle: { type: 'dashed' } }, { value: 76, lineStyle: { type: 'dashed' } }, { value: 82, lineStyle: { type: 'dashed' } }] // 能力增速变缓

    // 动态生成 Series
    let seriesData = []
    if (isDelayed.value) {
      // 触发蝴蝶效应：显示原计划(虚化) vs 现计划(红色警示)
      seriesData = [
        { name: '原预估薪资(已错失)', type: 'bar', barWidth: '35%', itemStyle: { color: 'rgba(239, 220, 226, 0.4)', borderRadius: [4, 4, 0, 0] }, data: idealSalary },
        { name: '真实月薪(推迟后)', type: 'bar', barGap: '-100%', barWidth: '35%', itemStyle: { color: '#F87171', borderRadius: [4, 4, 0, 0] }, data: delayedSalary },
        { name: '原能力指数(已错失)', type: 'line', yAxisIndex: 1, smooth: true, itemStyle: { color: 'rgba(138, 158, 88, 0.3)' }, lineStyle: { width: 2, type: 'dashed' }, data: [65, 78, 88, 95] },
        { name: '真实能力指数', type: 'line', yAxisIndex: 1, smooth: true, itemStyle: { color: '#F87171' }, lineStyle: { width: 3 }, data: delayedAbility }
      ]
    } else {
      // 正常按计划执行的状态
      seriesData = [
        { name: '预估月薪', type: 'bar', barWidth: '35%', itemStyle: { color: '#EFDCE2', borderRadius: [4, 4, 0, 0] }, data: idealSalary },
        { name: '综合能力指数', type: 'line', yAxisIndex: 1, smooth: true, itemStyle: { color: '#8A9E58' }, lineStyle: { width: 3 }, data: idealAbility }
      ]
    }

    comboChart.setOption({
      tooltip: { trigger: 'axis', axisPointer: { type: 'cross' } },
      legend: { top: '0%', right: '15%', textStyle: { fontSize: 10 } },
      grid: { left: '5%', right: '5%', bottom: '10%', top: 50, containLabel: true },
      xAxis: [{ type: 'category', data: ['今年(历史)', '第1年(预测)', '第2年(预测)', '第3年(预测)'], axisPointer: { type: 'shadow' } }],
      yAxis: [
        { type: 'value', name: '月薪 (K)', min: 0, max: 40, splitLine: { lineStyle: { type: 'dashed' } }, axisLabel: { formatter: '{value} k' } },
        { type: 'value', name: '指数', min: 0, max: 100, splitLine: { show: false } }
      ],
      series: seriesData
    })
    charts.push(comboChart)
  }

  // ==================== 3. 行业风险热力矩阵 ====================
  if (heatmapRef.value) {
    const heatChart = echarts.init(heatmapRef.value)
    heatChart.setOption({
      tooltip: { position: 'top', formatter: (p) => `${['基础切图/UI', '常规管理后台', '跨端开发', '微前端架构', 'AI辅助工程化'][p.value[1]]}风险: ${p.value[2]}%` },
      grid: { left: '3%', right: '5%', bottom: '5%', top: '5%', containLabel: true },
      xAxis: { type: 'category', data: ['未来1年', '未来2年', '未来3年'], splitArea: { show: true } },
      yAxis: { type: 'category', data: ['基础切图/UI', '常规管理后台', '跨端开发', '微前端架构', 'AI辅助工程化'], splitArea: { show: true }, axisLabel: { fontSize: 11, color: '#4b5563' } },
      visualMap: { min: 0, max: 100, calculable: true, orient: 'horizontal', show: false, inRange: { color: ['#8A9E58', '#F7EECD', '#F87171'] } },
      series: [{ name: '替代风险', type: 'heatmap', data: [[0, 0, 80], [1, 0, 95], [2, 0, 100], [0, 1, 40], [1, 1, 60], [2, 1, 85], [0, 2, 20], [1, 2, 30], [2, 2, 50], [0, 3, 10], [1, 3, 15], [2, 3, 20], [0, 4, 5], [1, 4, 5], [2, 4, 10]], label: { show: true, formatter: '{@[2]}%', fontSize: 10 }, itemStyle: { borderColor: '#fff', borderWidth: 2 } }]
    })
    charts.push(heatChart)
  }

  // ==================== 4. 行动建议漏斗图 ====================
  if (funnelRef.value) {
    const funnelChart = echarts.init(funnelRef.value)
    funnelChart.setOption({
      tooltip: { trigger: 'item', formatter: '{b}转化率 : {c}%' },
      series: [{
        name: '计划转化漏斗', type: 'funnel', left: '10%', top: '5%', bottom: '5%', width: '80%', minSize: '20%', maxSize: '100%', sort: 'descending', gap: 2,
        label: { show: true, position: 'inside', formatter: '{b}', fontSize: 12, fontWeight: 'bold' },
        itemStyle: { borderColor: '#fff', borderWidth: 1 },
        data: [
          { value: 100, name: '推荐技能储备池', itemStyle: { color: '#e5e7eb' } },
          { value: 80, name: '已列入甘特图学习', itemStyle: { color: '#F7EECD' } },
          { value: 60, name: '实战项目落地应用', itemStyle: { color: '#EFDCE2' } },
          { value: 30, name: '产生实际薪资溢价', itemStyle: { color: '#8A9E58' } }
        ]
      }]
    })
    charts.push(funnelChart)
  }
}

// ==================== 页面挂载与全局状态监听 ====================
onMounted(() => {
  nextTick(() => {
    // 检查是否在 CareerPath 中触发了推迟魔法
    if (sessionStorage.getItem('plan_delayed') === 'true') {
      isDelayed.value = true
      
      // 弹出话术
      setTimeout(() => {
        ElNotification({
          title: '洞察与预警',
          message: '规划的执行效率直接决定未来的职业高度，我们的系统能实时量化这种影响。',
          type: 'error',
          icon: InfoFilled,
          duration: 3000,
          customClass: 'shadow-2xl border-2 border-red-200 bg-red-50'
        })
      }, 400) // 延迟一点点，确保动画流畅
    }

    initCharts()
    window.addEventListener('resize', () => charts.forEach(c => c.resize()))
  })
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', () => charts.forEach(c => c.resize()))
  charts.forEach(c => c.dispose())
  
  // 离开页面时可以清除拖延标记，保证下次进来重新计算
  // sessionStorage.removeItem('plan_delayed') 
})
</script>

<style scoped>
@media print {
  body { background: white; }
  .print\:hidden { display: none !important; }
  .print\:shadow-none { box-shadow: none !important; }
  .print\:w-full { width: 100% !important; margin: 0 !important; padding: 0 !important; }
  * { -webkit-print-color-adjust: exact !important; print-color-adjust: exact !important; }
}
</style>