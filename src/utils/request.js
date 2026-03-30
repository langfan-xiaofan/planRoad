import axios from 'axios'

export const http = axios.create({
  baseURL: '/api',
  timeout: 15000,
})

export const sleep = (ms) => new Promise((r) => setTimeout(r, ms))

export const jobOptions = [
  { label: '高级前端工程师', value: 'senior_fe' },
  { label: '全栈开发工程师', value: 'fullstack' },
  { label: '技术经理', value: 'tech_lead' },
  { label: 'UI 设计师', value: 'ui_designer' },
]

export const buildCapabilityProfileMock = (parsed) => {
  const radarDetails = [
    { name: '专业技能', value: 90 },
    { name: '逻辑思维', value: 85 },
    { name: '工程架构', value: 70 },
    { name: '沟通协作', value: 80 },
    { name: '管理潜质', value: 60 },
    { name: '创新学习', value: 88 },
  ]

  return {
    summary:
      '你具备扎实的工程化与业务落地能力，建议重点补齐架构与服务端视野，形成“可交付 + 可演进”的能力闭环。',
    radarDetails,
    strengths: ['Vue3 生态', '工程化实践', '业务拆解与落地'],
    weaknesses: ['大型架构经验', '服务端与部署', '跨团队协作方法论'],
    evidence: {
      '专业技能': ['熟练使用 Vue3/TS', '能独立完成复杂业务模块'],
      '逻辑思维': ['项目拆分清晰', '问题定位效率高'],
      '工程架构': ['组件抽象能力较强', '缺少大型系统演进经验'],
      '沟通协作': ['能与产品/后端对齐需求', '缺少跨团队影响力案例'],
      '管理潜质': ['有带教意愿', '缺少明确的管理实践'],
      '创新学习': ['学习速度快', '能快速落地新方案'],
    },
    parsed,
  }
}

const mockMatchResults = {
  senior_fe: {
    score: 88,
    comment: '你的前端基础非常扎实，与该岗位的核心需求高度契合。建议在架构设计方面积累经验。',
    tips: ['参与开源项目以提升架构视野', '深入学习性能优化方案', '补充 Node.js 服务端经验'],
    strengths: ['Vue3 精通', '工程化能力', 'CSS 还原度'],
    weaknesses: ['大型架构经验', '服务端渲染'],
    radarData: [85, 90, 70, 60, 80, 75],
  },
  fullstack: {
    score: 72,
    comment: '你具备全栈开发的潜力，但后端深度略有不足。',
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
    comment: '你的技能树主要集中在逻辑实现，与设计岗位的技能重合度较低。',
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

export const buildCareerPathMock = async ({ jobKey, matchResult }) => {
  await sleep(250)
  const opt = jobOptions.find((j) => j.value === jobKey)
  const label = opt?.label || '目标岗位'

  const focus = (matchResult?.weaknesses || []).slice(0, 2)
  const focusTasks = focus.length
    ? focus.map((w) => `补齐：${w}`)
    : ['完善项目复盘', '补齐短板能力', '形成可复用方法论']

  const variants = {
    senior_fe: [
      {
        time: '0-6 月',
        role: '高级前端工程师',
        title: '工程能力体系化',
        desc: '把零散经验沉淀为可复用规范：组件抽象、性能优化、质量体系。',
        tasks: [...focusTasks, '建立组件与规范库', '主导一次性能专项'],
      },
      {
        time: '6-12 月',
        role: '前端技术负责人',
        title: '架构视野与协作升级',
        desc: '通过架构设计与跨团队协作，提升交付效率与系统可演进性。',
        tasks: ['完成一次架构设计评审', '推动工程化改造落地', '带教 1-2 名同学'],
      },
      {
        time: '1-2 年',
        role: '架构师预备',
        title: '全局优化与系统演进',
        desc: '补齐服务端与部署视野，用系统思维解决业务问题。',
        tasks: ['掌握 Node/Nest 基础', '上线一次灰度发布', '完善监控与告警'],
      },
    ],
    fullstack: [
      {
        time: '0-6 月',
        role: '全栈开发工程师',
        title: '补齐后端基本功',
        desc: '围绕数据库、API 设计、鉴权与部署建立完整后端闭环。',
        tasks: ['系统学习数据库设计', '实现鉴权与权限控制', '完成一次容器化部署'],
      },
      {
        time: '6-12 月',
        role: '全栈核心开发',
        title: '高质量交付与性能优化',
        desc: '从功能实现走向可维护与可扩展，打造性能与质量优势。',
        tasks: ['优化接口性能', '完善测试与CI', '沉淀一套脚手架/模板'],
      },
      {
        time: '1-2 年',
        role: '技术负责人',
        title: '端到端负责业务模块',
        desc: '能独立负责从需求到上线的完整链路，并推动团队效率提升。',
        tasks: ['负责一条业务线', '制定技术规范', '输出复盘与方法论'],
      },
    ],
    tech_lead: [
      {
        time: '0-6 月',
        role: '资深工程师',
        title: '管理与影响力起步',
        desc: '从带教与小项目开始，培养沟通、协作与决策能力。',
        tasks: ['担任 Mentor', '主持一次技术评审', '学习项目管理方法'],
      },
      {
        time: '6-12 月',
        role: '技术经理',
        title: '团队协作与交付效率',
        desc: '建立规范与流程，提升团队整体交付质量。',
        tasks: ['建立研发规范', '推动质量体系落地', '跨部门对齐目标'],
      },
      {
        time: '1-2 年',
        role: '部门负责人预备',
        title: '战略规划与资源协调',
        desc: '从短期交付走向中长期规划，形成技术与业务协同能力。',
        tasks: ['制定技术路线', '人才梯队建设', '关键项目推进'],
      },
    ],
    ui_designer: [
      {
        time: '0-3 月',
        role: '设计入门',
        title: '设计基础与工具',
        desc: '建立色彩、排版、布局基础，熟练掌握主流设计工具。',
        tasks: ['系统学习排版与色彩', '熟练使用 Figma', '完成 3 个练习作品'],
      },
      {
        time: '3-6 月',
        role: 'UI 设计师',
        title: '组件与规范',
        desc: '用组件化思维建设设计系统，与前端协作提升一致性。',
        tasks: ['搭建设计组件库', '输出设计规范', '做一次可用性测试'],
      },
      {
        time: '6-12 月',
        role: '体验设计师',
        title: '体验与增长',
        desc: '关注用户体验与业务指标，用设计驱动增长。',
        tasks: ['梳理用户路径', '优化关键转化流程', '沉淀作品集'],
      },
    ],
  }

  return {
    title: `面向「${label}」的阶段性成长路线`,
    steps: variants[jobKey] || variants.senior_fe,
  }
}
