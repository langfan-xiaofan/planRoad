import axios from 'axios'

export const http = axios.create({
  baseURL: '/api',
  timeout: 15000,
})

export const sleep = (ms) => new Promise((r) => setTimeout(r, ms))

// ==================== 1. 目标岗位字典 ====================
export const jobOptions = [
  { label: '高级前端架构师', value: 'senior_fe' },
  { label: '全栈开发工程师', value: 'fullstack' },
  { label: '技术经理', value: 'tech_lead' },
  { label: 'UI 体验设计师', value: 'ui_designer' },
]

// ==================== 2. 模拟简历解析返回能力画像 ====================
export const buildCapabilityProfileMock = (parsed) => {
  // 与 Report.vue 的玫瑰图及双基线雷达图完全对齐
  const radarDetails = [
    { name: '前端工程化', value: 85 },
    { name: '全栈与Node', value: 78 },
    { name: '架构与性能', value: 65 },
    { name: '业务沟通', value: 80 },
    { name: '管理潜质', value: 55 },
    { name: '创新认知', value: 90 },
  ]

  return {
    summary: '具备扎实的前端基础，正在向高阶架构师转型。建议重点突破微前端体系与服务端性能监控，形成完整的工程化闭环。',
    radarDetails,
    strengths: ['Vue3 源码级理解', '工程化脚手架开发', '业务需求拆解'],
    weaknesses: ['大型微前端架构落地', 'Node.js 性能调优', '高并发场景处理'],
    evidence: {
      '前端工程化': ['主导过 Vite 构建优化，提升 40% 速度', '熟练封装企业级组件库'],
      '全栈与Node': ['能使用 Node 写中间层', '缺少底层性能排查经验'],
      '架构与性能': ['有组件抽象意识', '尚未独立操盘过十万级日活架构'],
      '业务沟通': ['需求对齐精准', '能有效协调产品与测试资源'],
      '管理潜质': ['负责新人带教', '缺乏正规的团队 KPI 考核管理经验'],
      '创新认知': ['对 Web3、AI 辅助编程等前沿技术敏感度高', '主动学习意愿极强'],
    },
    parsed,
  }
}

// ==================== 3. 模拟人岗匹配结果 ====================
const mockMatchResults = {
  senior_fe: {
    score: 85,
    comment: '你的前端工程化基础非常扎实，与架构师岗位的核心需求匹配度较高。亟待补齐架构设计与系统调优经验。',
    tips: ['深入微前端 Qiankun 源码与实战', '补充 Node.js 内存泄漏与性能监控经验', '体系化学习 Pinia 等状态管理的底层逻辑'],
    strengths: ['前端工程化', '业务沟通', '创新认知'],
    weaknesses: ['架构与性能', '管理潜质'],
    // 目标岗位的标准线 (雷达图)
    radarData: [80, 75, 85, 70, 85, 80],
    // 极简格式数据，配合核心技能缺口瀑布图
    coreGaps: [
      { name: '微前端 (Qiankun)', percent: 30 },
      { name: 'Vite 构建优化', percent: 45 },
      { name: 'Node.js 性能监控', percent: 50 },
      { name: 'Pinia 状态管理', percent: 70 }
    ]
  },
  // 其他职业兜底数据
  fullstack: {
    score: 72,
    comment: '具备全栈开发的潜力，但后端深度略有不足。',
    tips: ['系统学习数据库设计', '掌握 Docker 容器化部署', '尝试独立开发完整项目'],
    strengths: ['前端交互', '快速学习'],
    weaknesses: ['数据库优化', '高并发处理'],
    radarData: [80, 70, 60, 50, 40, 65],
  },
  tech_lead: {
    score: 65,
    comment: '技术能力达标，但管理经验和技术领导力需要加强。',
    tips: ['主动承担团队 Mentor 角色', '学习项目管理方法论', '提升技术选型决策力'],
    strengths: ['技术深度', '问题解决'],
    weaknesses: ['团队管理', '战略规划'],
    radarData: [85, 80, 70, 50, 30, 60],
  },
  ui_designer: {
    score: 45,
    comment: '你的技能树主要集中在逻辑实现，与设计岗位的重合度较低。',
    tips: ['系统学习色彩与排版理论', '熟练使用 Figma/Sketch', '积累作品集'],
    strengths: ['实现能力', '逻辑思维'],
    weaknesses: ['审美能力', '设计工具', '创意发散'],
    radarData: [40, 50, 60, 30, 20, 40],
  },
}

export const runJobMatchMock = async ({ jobKey }) => {
  await sleep(450)
  const base = mockMatchResults[jobKey]
  if (!base) throw new Error('unknown_job')

  return {
    jobKey,
    ...base,
  }
}

// ==================== 4. 模拟生成职业路径数据 ====================
export const buildCareerPathMock = async ({ jobKey, matchResult }) => {
  await sleep(250)
  const opt = jobOptions.find((j) => j.value === jobKey)
  const label = opt?.label || '目标岗位'

  // 数据结构完全对应 CareerPath.vue 的四个 ECharts 图表
  return {
    title: `面向「${label}」的定制化成长路线`,

    // 1. 技能转化桑基图数据 
    sankeyData: {
      nodes: [
        { name: 'JavaScript基础' }, { name: 'Vue2经验' }, { name: 'Vue3工程化实战' },
        { name: '微前端Qiankun项目' }, { name: 'Node.js性能监控' },
        { name: '高级前端架构' }, { name: '全栈思维' }
      ],
      links: [
        { source: 'JavaScript基础', target: 'Node.js性能监控', value: 3 },
        { source: 'JavaScript基础', target: 'Vue3工程化实战', value: 2 },
        { source: 'Vue2经验', target: 'Vue3工程化实战', value: 4 },
        { source: 'Vue2经验', target: '微前端Qiankun项目', value: 2 },
        { source: 'Vue3工程化实战', target: '高级前端架构', value: 6 },
        { source: '微前端Qiankun项目', target: '高级前端架构', value: 3 },
        { source: 'Node.js性能监控', target: '全栈思维', value: 4 }
      ]
    },

    // 2. 中短期里程碑甘特图数据 
    ganttData: {
      yAxisTasks: ['算法打卡', '微前端搭建', 'Node服务开发', 'Vue3源码剖析'],
      startOffset: [1, 5, 2, 0], // 起始点 (占位透明柱子)
      durations: [10, 4, 3, 5]   // 持续周期 (有色实际柱子)
    },

    // 3. 岗位转型力导向图数据 
    graphData: {
      nodes: [
        { name: '高级前端', symbolSize: 50 },
        { name: '全栈开发', symbolSize: 35 },
        { name: '前端架构师', symbolSize: 40 },
        { name: 'Node研发', symbolSize: 30 },
        { name: 'Web3前端', symbolSize: 25 }
      ],
      links: [
        { source: '高级前端', target: '前端架构师', value: 80 },
        { source: '高级前端', target: '全栈开发', value: 60 },
        { source: '高级前端', target: 'Node研发', value: 45 },
        { source: '高级前端', target: 'Web3前端', value: 30 },
        { source: '全栈开发', target: 'Node研发', value: 70 }
      ]
    },

    // 4. 预期能力值增长趋势数据 
    growthTrendData: {
      xAxis: ['Q1', 'Q2', 'Q3', 'Q4'],
      series: {
        hardSkills: [15, 30, 45, 60], // 硬技能增长
        projectExp: [10, 25, 40, 50], // 项目经验累积
        softSkills: [5, 10, 15, 25]   // 软实力提升
      }
    }
  }
}