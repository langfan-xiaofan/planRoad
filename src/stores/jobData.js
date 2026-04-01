import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import {
  buildCareerPathMock,
  buildCapabilityProfileMock,
  jobOptions,
  runJobMatchMock,
} from '../utils/request'

export const useJobDataStore = defineStore('jobData', () => {
  // --- 状态 (State) ---
  const resumeFile = ref(null)
  const resumeParsed = ref(null)
  const capabilityProfile = ref(null)
  const targetJobKey = ref('')
  const matchResult = ref(null)
  const careerPath = ref(null)

  const targetJobLabel = computed(() => {
    const opt = jobOptions.find((j) => j.value === targetJobKey.value)
    return opt?.label || ''
  })

  const hasResumeProfile = computed(() => Boolean(capabilityProfile.value && resumeParsed.value))
  const hasMatch = computed(() => Boolean(matchResult.value))
  const hasCareerPath = computed(() => Boolean(careerPath.value))

  const resetAll = () => {
    resumeFile.value = null
    resumeParsed.value = null
    capabilityProfile.value = null
    targetJobKey.value = ''
    matchResult.value = null
    careerPath.value = null
  }

  const setTargetJobKey = (key) => {
    targetJobKey.value = key
    matchResult.value = null
    careerPath.value = null
  }

  const setResumeFile = (file) => {
    resumeFile.value = file
  }

  const setResumeParsed = (parsed) => {
    resumeParsed.value = parsed
  }

  const setCapabilityProfile = (profile) => {
    capabilityProfile.value = profile
  }

  const parseResumeMock = async (file, { onProgress } = {}) => {
    resetAll()
    setResumeFile({
      name: file?.name || 'resume_linmumu.pdf',
      size: file?.size || 102400,
      type: file?.type || 'application/pdf',
      lastModified: file?.lastModified || Date.now(),
    })

    const parsed = {
      name: '林木木',
      education: '本科',
      major: '软件工程',
      targetJobKey: 'senior_fe',
      skills: ['Vue3 工程化', '微前端(Qiankun)', 'Node.js 中间层', '前端架构设计', '性能调优'],
      skillScores: {
        '前端工程化': 85,
        '全栈与Node': 78,
        '架构与性能': 65,
        '业务沟通': 80,
        '管理潜质': 55,
        '创新认知': 90
      }
    }

    const progressSteps = [8, 22, 35, 52, 68, 82, 94, 100]
    for (const p of progressSteps) {
      onProgress?.(p)
      await new Promise((r) => setTimeout(r, 220))
    }

    setResumeParsed(parsed)
    setTargetJobKey(parsed.targetJobKey)

    try {
      const profile = buildCapabilityProfileMock(parsed)
      if (profile && profile.skillScores) {
        setCapabilityProfile(profile)
      } else {
        // 兜底策略：如果 request 里的方法崩了，直接用 parsed 里的数据生成 Profile
        setCapabilityProfile({ summary: '具备扎实的前端基础，正在向高阶架构师转型。', skillScores: parsed.skillScores })
      }
    } catch (e) {
      console.warn('Capability profile build skipped or failed, using raw parsed data', e)
      setCapabilityProfile({ summary: '具备扎实的前端基础，正在向高阶架构师转型。', skillScores: parsed.skillScores })
    }

    return { parsed, profile: capabilityProfile.value }
  }

  const runMatchMock = async (jobKey) => {
    if (!hasResumeProfile.value) throw new Error('missing_profile')
    if (!jobKey) throw new Error('missing_target_job')

    setTargetJobKey(jobKey)
    const result = await runJobMatchMock({
      jobKey,
      resumeParsed: resumeParsed.value,
      capabilityProfile: capabilityProfile.value,
    })
    matchResult.value = result
    careerPath.value = null
    return result
  }

  const buildCareerPathFromMatchMock = async () => {
    if (!matchResult.value || !targetJobKey.value) throw new Error('missing_match')

    const plan = await buildCareerPathMock({
      jobKey: targetJobKey.value,
      resumeParsed: resumeParsed.value,
      matchResult: matchResult.value,
    })
    careerPath.value = plan
    return plan
  }

  return {
    resumeFile,
    resumeParsed,
    capabilityProfile,
    targetJobKey,
    targetJobLabel,
    matchResult,
    careerPath,
    hasResumeProfile,
    hasMatch,
    hasCareerPath,
    resetAll,
    setTargetJobKey,
    setResumeFile,
    setResumeParsed,
    setCapabilityProfile,
    parseResumeMock,
    runMatchMock,
    buildCareerPathFromMatchMock,
  }
})