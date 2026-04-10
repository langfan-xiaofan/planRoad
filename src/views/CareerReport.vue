<template>
  <div class="flex flex-col items-center h-full overflow-y-auto -mr-6 pr-6 pb-4">
    <div class="w-full max-w-4xl bg-white p-4 rounded-xl shadow-sm mb-6 flex items-center justify-between">
      <div class="flex items-center gap-4 w-1/2">
        <span class="text-sm text-gray-600 shrink-0">技能权重调节:</span>
        <el-slider v-model="skillWeight" class="w-full" :show-tooltip="false" />
      </div>
      <div class="flex gap-2">
        <el-button plain class="border-themePrimary text-themePrimary">AI 智能润色</el-button>
        <el-button type="primary" class="bg-themePrimary border-none">一键导出 PDF</el-button>
      </div>
    </div>

    <div class="w-full max-w-4xl bg-white p-12 rounded-sm shadow-md min-h-[1000px] mb-10">
      <h1 class="text-3xl font-bold text-center text-gray-800 mb-8 pb-4 border-b-2 border-themePrimary">
        个人职业生涯发展规划报告
      </h1>

      <section class="mb-8">
        <h2 class="text-xl font-bold text-themePrimary mb-4">1. 人岗匹配度分析</h2>
        <p class="text-gray-600 mb-4">综合对比目标岗位要求与个人能力，当前综合匹配度为 <strong class="text-themeAccent text-xl">82%</strong>。</p>
        <div class="bg-gray-50 h-40 flex items-center justify-center text-gray-400 rounded">
          匹配度图表占位（后续在此渲染对比数据）
        </div>
      </section>

      <section class="mb-8">
        <h2 class="text-xl font-bold text-themePrimary mb-4">2. 职业目标与发展路径</h2>
        <div class="text-gray-700 leading-relaxed text-sm">
          <strong>短期目标：</strong> 夯实前端开发基础，熟练掌握现代化框架构建流程。<br>
          <strong>中期目标：</strong> 能够独立负责中型项目的前端架构设计，向全栈方向拓展。<br>
          <strong>发展路径推荐：</strong> 初级开发工程师 -> 核心业务骨干 -> 前端架构师 / 研发组长
        </div>
      </section>

      <section>
        <h2 class="text-xl font-bold text-themePrimary mb-4">3. 行动计划时间线</h2>
        <el-timeline>
          <el-timeline-item timestamp="第 1-3 个月" placement="top" color="#84B4C3">
            <el-card>
              <h4 class="font-bold text-gray-700">基础巩固与考证期</h4>
              <p class="text-sm text-gray-500 mt-2">系统性学习缺失的核心专业知识，完成相关职业资格认证报名与备考。</p>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="第 4-6 个月" placement="top" color="#C7E2E0">
            <el-card>
              <h4 class="font-bold text-gray-700">项目实战与技能转化</h4>
              <p class="text-sm text-gray-500 mt-2">参与开源项目或校内双创项目，将理论知识转化为实际开发能力。</p>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="第 7-12 个月" placement="top" color="#EAC0C3">
            <el-card>
              <h4 class="font-bold text-gray-700">实习验证与方向调整</h4>
              <p class="text-sm text-gray-500 mt-2">投递意向岗位实习，在真实职场环境中检验规划合理性并动态调整。</p>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const skillWeight = ref(50)
</script>