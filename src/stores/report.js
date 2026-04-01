import { defineStore } from 'pinia'
import { computed } from 'vue'
import { jobOptions } from '../utils/request'
import { useJobDataStore } from './jobData'

export const useReportStore = defineStore('report', () => {
  const job = useJobDataStore()

  // 1. 判断是否具备生成条件
  const canGenerate = computed(
    () => job.hasResumeProfile.value && job.hasMatch.value && job.hasCareerPath.value
  )

  // 2. 核心数据组装（与现有的 Report.vue 完美对齐）
  const reportData = computed(() => {
    // 基础信息 (自动回退到林木木与架构师，防止无数据时报错)
    if (!canGenerate.value) return null
    const userName = job.resumeParsed.value?.name || '林木木'
    const targetJob = job.targetJobLabel || '高级前端架构师'
    const matchScore = Number(job.matchResult.value?.score ?? 85)
    // --- 图表 1: 终极能力指纹 (Rose)  ---
    const roseData = [
      { value: 40, name: '前端工程化', itemStyle: { color: '#8A9E58' } },
      { value: 33, name: '全栈/Node', itemStyle: { color: '#b5c58a' } },
      { value: 28, name: '架构与性能', itemStyle: { color: '#EFDCE2' } },
      { value: 22, name: '业务沟通', itemStyle: { color: '#F7EECD' } },
      { value: 18, name: '创新认知', itemStyle: { color: '#e5e7eb' } }
    ].map((i) => ({
      ...i,
      value: Math.max(0, Math.min(100, Number(i.value) || 0))
    }))

    // --- 图表 2: 三年期薪资与能力复合增长预测 (Combo) ---
    const salaryGrowth = {
      idealSalary: [15, 22, 28, 35],
      delayedSalary: [15, 18, 22, 26], // 拖延后薪水缩水
      idealAbility: [
        65,
        { value: 78, lineStyle: { type: 'dashed' } },
        { value: 88, lineStyle: { type: 'dashed' } },
        { value: 95, lineStyle: { type: 'dashed' } }
      ],
      delayedAbility: [
        65,
        { value: 70, lineStyle: { type: 'dashed' } },
        { value: 76, lineStyle: { type: 'dashed' } },
        { value: 82, lineStyle: { type: 'dashed' } }
      ]
    }

    // --- 图表 3: 技能贬值风险矩阵 (Heatmap) ---
    const riskMatrix = {
      xAxis: ['未来1年', '未来2年', '未来3年'],
      yAxis: ['基础切图/UI', '常规管理后台', '跨端开发', '微前端架构', 'AI辅助工程化'],
      data: [
        [0, 0, 80], [1, 0, 95], [2, 0, 100], // 基础岗位高危
        [0, 1, 40], [1, 1, 60], [2, 1, 85],
        [0, 2, 20], [1, 2, 30], [2, 2, 50],
        [0, 3, 10], [1, 3, 15], [2, 3, 20],
        [0, 4, 5], [1, 4, 5], [2, 4, 10]   // AI辅助极低危
      ]
    }

    // --- 图表 4: 学习向薪资转化的漏斗模型 (Funnel) ---
    const funnelData = [
      { value: 100, name: '推荐技能储备池', itemStyle: { color: '#e5e7eb' } },
      { value: 80, name: '已列入甘特图学习', itemStyle: { color: '#F7EECD' } },
      { value: 60, name: '实战项目落地应用', itemStyle: { color: '#EFDCE2' } },
      { value: 30, name: '产生实际薪资溢价', itemStyle: { color: '#8A9E58' } }
    ]
    return {
      userName,
      targetJob,
      overallScore: matchScore,
      roseData,
      salaryGrowth,
      riskMatrix,
      funnelData
    }
  })
  //导出文件名将会自动带上真实数据，比如 "林木木-高级前端架构师-全景发展报告.pdf"
  const filename = computed(() => {
    const name = reportData.value?.userName || '林木木'
    const jobName = reportData.value?.targetJob || '高级前端架构师'
    const safeName = name.replace(/[\/\\:*?"<>|]/g, '')
    const safeJob = jobName.replace(/[\/\\:*?"<>|]/g, '')
    return `${safeName}-${safeJob}-全景发展报告.pdf`
  })

  return {
    canGenerate,
    reportData,
    filename,
  }
})