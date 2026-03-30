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

  // 2. 核心数据组装
  const reportData = computed(() => {
    if (!canGenerate.value) return null

    // --- 基础信息 ---
    const targetLabel =
      jobOptions.find((o) => o.value === job.targetJobKey)?.label || job.targetJobLabel || '目标岗位'

    const matchScore = Number(job.matchResult.value?.score ?? 0)
    const userName = job.resumeParsed.value?.name || '匿名用户'

    // --- 雷达图数据 (确保顺序和数值安全) ---
    const rawRadar = job.matchResult.value?.radarData || [0, 0, 0, 0, 0, 0]
    const radarDetails = [
      { name: '专业技能', value: rawRadar[0] },
      { name: '逻辑思维', value: rawRadar[1] },
      { name: '沟通协作', value: rawRadar[2] }, // 调整顺序对应你的 UI
      { name: '工程架构', value: rawRadar[3] },
      { name: '管理潜质', value: rawRadar[4] },
      { name: '创新学习', value: rawRadar[5] },
    ].map((i) => ({
      ...i,
      value: Math.max(0, Math.min(100, Number(i.value) || 0))
    }))

    // --- 核心指标计算 (新增) ---

    // 1. 市场身价估算 (基于分数简单模拟)
    let salaryRange = '15-25 万'
    if (matchScore >= 90) salaryRange = '45-60 万'
    else if (matchScore >= 80) salaryRange = '35-45 万'
    else if (matchScore >= 70) salaryRange = '25-35 万'
    else if (matchScore >= 60) salaryRange = '18-25 万'

    // 2. 爆发潜力
    const potentialLevel = matchScore >= 90 ? 'S' : matchScore >= 80 ? 'A+' : matchScore >= 70 ? 'A' : 'B'
    const potentialPercent = matchScore >= 85 ? 95 : matchScore >= 70 ? 80 : 60

    // 3. 行业风险 (模拟算法：分数越高风险越低，但高分段竞争大风险略升)
    let riskScore = 50
    let riskComment = '行业整体平稳，需关注技术迭代。'
    if (matchScore < 60) {
      riskScore = 75
      riskComment = '基础岗位替代风险较高，建议尽快提升核心技能。'
    } else if (matchScore >= 85) {
      riskScore = 30
      riskComment = '高端人才稀缺，抗风险能力强，但需警惕技术方向偏差。'
    } else {
      riskScore = 45
      riskComment = '中坚力量，需向架构或管理转型以规避中年危机。'
    }

    // 4. 简短总结 (取长总结的前30个字)
    const fullSummary = job.capabilityProfile.value?.summary || job.matchResult.value?.comment || '您的综合竞争力良好，具备较大的发展潜力。'
    const summaryShort = fullSummary.length > 30 ? fullSummary.substring(0, 30) + '...' : fullSummary

    // 5. 桑基图数据转换 (将 steps 转换为 source->target 结构)
    const rawSteps = job.careerPath.value?.steps || []
    const sankeyPath = []

    if (rawSteps.length > 0) {
      // 第一个节点：当前 -> 第一步
      sankeyPath.push({
        source: `当前：${targetLabel}`,
        target: `阶段一：${rawSteps[0].role || '初级进阶'}`,
        value: 100
      })

      // 中间节点流转
      for (let i = 0; i < rawSteps.length - 1; i++) {
        sankeyPath.push({
          source: `阶段${i + 1}：${rawSteps[i].role || '发展中'}`,
          target: `阶段${i + 2}：${rawSteps[i + 1].role || '资深专家'}`,
          value: 80 - (i * 10) // 模拟留存率/概率递减
        })
      }

      // 最后一个节点 -> 终局
      const lastStep = rawSteps[rawSteps.length - 1]
      sankeyPath.push({
        source: `阶段${rawSteps.length}：${lastStep.role || '专家'}`,
        target: '终局：行业领袖/CTO',
        value: 50
      })
    }

    // 6. 建议列表格式化
    const rawTips = job.matchResult.value?.tips || ['持续学习新技术', '参与开源项目', '提升沟通能力']
    const tips = rawTips.map((t, idx) => {
      const titles = ['技能提升', '项目实践', '软性实力', '行业视野', '人脉积累']
      return {
        title: titles[idx % titles.length],
        content: t
      }
    })

    // --- 返回最终对象 ---
    return {
      userName,
      targetJob: targetLabel,
      overallScore: matchScore,
      summary: fullSummary,
      summaryShort, // 新增
      matchRate: matchScore,
      competitiveness: matchScore >= 80 ? '前 15%' : matchScore >= 60 ? '前 35%' : '前 60%',
      growthSpeed: matchScore >= 80 ? '快' : matchScore >= 60 ? '中' : '稳',

      // 新增字段
      salaryRange,
      potentialLevel,
      potentialPercent,
      riskScore,
      riskComment,

      radarDetails,
      path: sankeyPath, // 替换为桑基图格式
      tips,
    }
  })

  const filename = computed(() => {
    const name = reportData.value?.userName || '职业发展'
    const jobName = reportData.value?.targetJob || '报告'
    // 去除文件名中的特殊字符
    const safeName = name.replace(/[\/\\:*?"<>|]/g, '')
    const safeJob = jobName.replace(/[\/\\:*?"<>|]/g, '')
    return `${safeName}-${safeJob}-规划报告.pdf`
  })

  return {
    canGenerate,
    reportData,
    filename,
  }
})