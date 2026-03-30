import { ElMessage, ElNotification } from 'element-plus'

const DEMO_DATA = {
  user: { name: '李四 (演示账号)', target: '高级前端架构师' },
  resume: {
    skills: ['Vue3', 'TypeScript', 'Node.js', '微前端架构', 'CI/CD'],
    score: 95
  },
  match: {
    job: 'senior_fe',
    label: '高级前端架构师',
    score: 94,
    comment: '您的技术栈与行业顶尖需求高度重合，特别是在微前端与工程化领域表现卓越。'
  }
}

class DemoTourManager {
  constructor(router) {
    this.isRunning = false
    this.abortController = null // 用于强制停止演示
    this.router = router // 接收 Vue Router 实例
  }

  // 工具：等待一定时间
  async wait(ms) {
    return new Promise(r => setTimeout(r, ms))
  }

  // 工具：安全跳转并等待渲染
  async goToStep(path) {
    if (this.router.currentRoute.value.path === path) return
    
    // 使用 router.push 进行页面跳转
    this.router.push(path)
    
    // 等待页面加载，不同页面需要不同的等待时间
    if (path === '/report') {
      await this.wait(4000)
    } else if (path === '/path') {
      await this.wait(3500)
    } else if (path === '/match') {
      await this.wait(3500)
    } else {
      await this.wait(3500)
    }
  }

  // 步骤 1: 简历解析
  async step1_Resume() {
    await this.goToStep('/resume')

    ElNotification({
      title: '📄 步骤 1: 智能简历解析',
      message: 'AI 正在提取技能、项目经历并构建六维能力画像...',
      type: 'info',
      duration: 0 // 手动关闭
    })

    // 触发解析事件
    window.dispatchEvent(new CustomEvent('demo-trigger-parse'))

    // 等待解析完成
    await this.wait(10000000)
    ElNotification.closeAll()
  }

  // 步骤 2: 人岗匹配
  async step2_Match() {
    await this.goToStep('/match')

    ElNotification({
      title: '🎯 步骤 2: 人岗深度匹配',
      message: `正在对比您与 "${DEMO_DATA.match.label}" 的差距...`,
      type: 'info',
      duration: 0
    })

    // 触发匹配事件
    window.dispatchEvent(new CustomEvent('demo-trigger-match', {
      detail: { job: DEMO_DATA.match.job }
    }))

    // 等待匹配完成
    await this.wait(4000)
    ElNotification.closeAll()
  }

  // 步骤 3: 职业路径
  async step3_Path() {
    await this.goToStep('/path')

    ElNotification({
      title: '🗺️ 步骤 3: 生成成长路径',
      message: '基于匹配结果，规划未来 3 年的进阶路线...',
      type: 'warning',
      duration: 0
    })

    // 触发路径生成事件
    window.dispatchEvent(new CustomEvent('demo-trigger-path'))

    // 等待路径生成
    await this.wait(4000)
    ElNotification.closeAll()
  }

  // 步骤 4: 最终报告
  async step4_Report() {
    await this.goToStep('/report')

    // 等待报告页完全渲染
    await this.wait(3000)

    ElNotification({
      title: '🎉 演示完成',
      message: '这就是为您生成的完整职业发展蓝图！包含能力画像、差距分析及行动路线。',
      type: 'success',
      duration: 0,
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

    // 增加发展报告页面的停留时间，让用户有足够时间查看内容
    await this.wait(5000)

    this.isRunning = false
  }

  // 启动演示
  async start() {
    // 先标记演示模式，触发路由守卫中的逻辑
    sessionStorage.setItem('is_demo_mode', 'true')
    if (this.isRunning) {
      ElMessage.warning('演示正在进行中...')
      return
    }

    this.isRunning = true
    ElMessage.info('🎬 演示模式启动 | 请勿手动操作')

    try {
      await this.step1_Resume()
      if (!this.isRunning) return // 检查是否被中断

      await this.step2_Match()
      if (!this.isRunning) return

      await this.step3_Path()
      if (!this.isRunning) return

      await this.step4_Report()
    } catch (error) {
      console.error('Demo Tour Error:', error)
      ElMessage.error('演示过程中断，请刷新页面重试')
      this.isRunning = false
    }
  }

  // 停止演示
  stop() {
    this.isRunning = false
    ElNotification.closeAll()
    ElMessage.info('演示已停止')
  }
}

export default DemoTourManager