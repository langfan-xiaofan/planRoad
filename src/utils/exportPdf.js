import html2canvas from 'html2canvas'
import { jsPDF } from 'jspdf'
import { defineStore } from 'pinia'
import { ElMessage } from 'element-plus'

export const useExportPdfStore = defineStore('exportPdf', {
  state: () => ({
    isExporting: false // 控制按钮的 loading 状态
  }),
  actions: {
    async exportToPdf(elementId, filename = 'CareerAgent_职业发展报告.pdf') {
      const element = document.getElementById(elementId)
      if (!element) {
        ElMessage.error('找不到需要导出的页面内容')
        return
      }

      this.isExporting = true
      const loadingMsg = ElMessage({
        message: '正在为您生成高清 PDF 报告，请稍候...',
        type: 'info',
        duration: 0 // 持续显示直到手动关闭
      })

      try {
        // 1. 将 HTML 节点转换为 Canvas (配置抗锯齿和高分辨率)
        const canvas = await html2canvas(element, {
          scale: 3, // 3倍缩放，保证文字和 ECharts 图表清晰度
          useCORS: true, // 允许加载跨域图片（如头像）
          backgroundColor: '#f9fafb', // 给点浅色底，防止 PDF 边缘刺眼
          logging: false
        })

        // 2. 将 Canvas 转换为图片数据
        const imgData = canvas.toDataURL('image/jpeg', 1.0)

        // 3. 计算 A4 纸张的比例，生成 PDF
        const pdf = new jsPDF('l', 'mm', 'a4')
        const pdfWidth = pdf.internal.pageSize.getWidth()
        // const pdfHeight = (canvas.height * pdfWidth) / canvas.width
        const pdfHeight = pdf.internal.pageSize.getHeight()

        // 先尝试以 A4 纸的宽度为基准铺满
        let renderWidth = pdfWidth
        let renderHeight = (canvas.height * pdfWidth) / canvas.width
        let offsetX = 0
        let offsetY = 0
        // 如果按宽度铺满后，发现高度超出了 A4 纸的高度，说明屏幕太方了
        // 那就改为以高度为基准铺满，宽度左右留白居中
        if (renderHeight > pdfHeight) {
          renderHeight = pdfHeight
          renderWidth = (canvas.width * pdfHeight) / canvas.height
          offsetX = (pdfWidth - renderWidth) / 2 // 水平居中
        } else {
          // 否则就是高度没占满，上下留白居中
          offsetY = (pdfHeight - renderHeight) / 2 // 垂直居中
        }

        // 4. 将图片精准贴入 PDF
        pdf.addImage(imgData, 'JPEG', offsetX, offsetY, renderWidth, renderHeight)
        // 触发下载
        pdf.save(filename)
        loadingMsg.close()
        ElMessage.success('🎉 PDF 报告导出成功！')
      } catch (error) {
        console.error('PDF 导出失败:', error)
        loadingMsg.close()
        ElMessage.error('导出失败，请检查控制台报错')
      } finally {
        this.isExporting = false
      }
    }
  }
})