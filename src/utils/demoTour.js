import { ElMessage, ElNotification } from 'element-plus'

const DEMO_DATA = {
  user: { name: '林木木 (演示账号)', target: '高级前端架构师' },
  resume: {
    skills: ['Vue3', '微前端', 'Node.js', '前端架构设计', '性能调优'],
    score: 88
  },
  match: {
    job: 'senior_fe',
    label: '高级前端架构师',
    score: 85,
    comment: '您的技术栈与行业顶尖需求高度重合，特别是在工程化与创新认知领域表现卓越。'
  }
}

class DemoTourManager {
  constructor(router) {
    this.isRunning = false
    this.abortController = null
    this.router = router
  }

  async wait(ms) {
    return new Promise(r => setTimeout(r, ms))
  }

  async goToStep(path) {
    if (this.router.currentRoute.value.path === path) return
    this.router.push(path)
    await this.wait(1000) // 基础跳转缓冲时间
  }

  // 简历解析
  async step1_Resume() {
    await this.goToStep('/resume')

    ElNotification({
      title: '📄 步骤 1: 全息简历解析',
      message: `AI 正在提取【${DEMO_DATA.user.name}】的技能经验，构建六维能力指纹...`,
      type: 'info',
      duration: 3500
    })

    window.dispatchEvent(new CustomEvent('demo-trigger-parse'))

    // 模拟解析时间
    await this.wait(4000)
    ElNotification.closeAll()
  }

  // 人岗匹配
  async step2_Match() {
    await this.goToStep('/match')

    ElNotification({
      title: '🎯 步骤 2: 核心缺口探测',
      message: `正在将能力指纹与【${DEMO_DATA.match.label}】进行双基线对齐...`,
      type: 'warning',
      duration: 3500
    })

    window.dispatchEvent(new CustomEvent('demo-trigger-match', {
      detail: { job: DEMO_DATA.match.job }
    }))

    await this.wait(4000)
    ElNotification.closeAll()
  }

  // 职业路径
  async step3_Path() {
    await this.goToStep('/path')

    ElNotification({
      title: '🗺️ 步骤 3: 动态路径推演',
      message: '正在根据缺口瀑布图，为您编排最优的技能转化链路与甘特图排期...',
      type: 'success',
      duration: 3500
    })

    window.dispatchEvent(new CustomEvent('demo-trigger-path'))

    await this.wait(4000)
    ElNotification.closeAll()
  }

  //  最终报告
  async step4_Report() {
    await this.goToStep('/report')
    await this.wait(2000) // 等待 Echarts 图表渲染完毕

    ElNotification({
      title: '🎉 推演完成',
      message: '这就是为您量身定制的【全景发展与风险预测报告】，系统已生成三年期薪资预测。',
      type: 'success',
      duration: 0, // 最后一个提示框保持常亮
      position: 'bottom-right',
      onClick: () => {
        ElNotification({
          title: '立即体验',
          message: '注册账号，生成属于您的专属报告',
          type: 'success',
          duration: 5000,
          onClick: () => window.location.href = '/auth/register'
        })
      }
    })

    await this.wait(5000)
    this.isRunning = false
  }

  // 启动演示
  async start() {
    sessionStorage.setItem('is_demo_mode', 'true')
    if (this.isRunning) {
      ElMessage.warning('演示正在进行中...')
      return
    }

    this.isRunning = true
    ElMessage.info('🎬 AI 演示模式已启动 | 请勿手动操作')

    try {
      await this.step1_Resume()
      if (!this.isRunning) return

      await this.step2_Match()
      if (!this.isRunning) return

      await this.step3_Path()
      if (!this.isRunning) return

      await this.step4_Report()
    } catch (error) {
      console.error('Demo Tour Error:', error)
      ElMessage.error('演示过程中断，请检查控制台报错')
      this.isRunning = false
    }
  }

  stop() {
    this.isRunning = false
    ElNotification.closeAll()
    ElMessage.info('演示已终止')
  }
}

export default DemoTourManager