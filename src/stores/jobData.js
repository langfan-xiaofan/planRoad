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
      name: file?.name || 'resume',
      size: file?.size || 0,
      type: file?.type || '',
      lastModified: file?.lastModified || Date.now(),
    })

    const parsed = {
      name: '张三',
      education: '本科',
      major: '计算机科学与技术',
      targetJobKey: 'senior_fe',
      skills: ['Vue3', 'JavaScript', 'TypeScript', 'Node.js', 'CSS3'],
      skillScores: {  
        '专业技能': 85,      // 对应 Vue 中的 scores['专业技能']
        '逻辑思维': 78,      // 对应 Vue 中的 scores['逻辑思维']
        '工程架构': 65,      // 对应 Vue 中的 scores['工程架构']
        '沟通协作': 80,      // 对应 Vue 中的 scores['沟通协作']
        '管理潜质': 55,      // 对应 Vue 中的 scores['管理潜质']
        '创新学习': 90       // 对应 Vue 中的 scores['创新学习'] },
    },
  }

    const progressSteps = [8, 22, 35, 52, 68, 82, 94, 100]
    for (const p of progressSteps) {
      onProgress?.(p)
      await new Promise((r) => setTimeout(r, 220))
    }

    setResumeParsed(parsed)
    setTargetJobKey(parsed.targetJobKey)
    const profile = buildCapabilityProfileMock(parsed)
    setCapabilityProfile(profile)

// 注意：如果 buildCapabilityProfileMock 内部逻辑依赖旧的英文 Key，它可能会覆盖或出错。
    // 最安全的做法是：如果该函数不可控，我们直接使用 parsed 中的数据，或者确保它返回的数据也包含上述中文 Key。
    // 这里为了保险，我们假设它只是做转换，如果它转换错了，我们优先保证 resumeParsed 是对的。
    try {
      const profile = buildCapabilityProfileMock(parsed)
      // 如果 profile 里的 skillScores 也是错的，我们可以手动修正它，或者忽略它直接用 parsed 的
      if (profile && profile.skillScores) {
         // 可选：如果 profile 数据也是错的，可以强制把正确的塞进去，或者只依赖 parsed
         // 这里暂时保留原逻辑，但你要确保 request.js 里的 buildCapabilityProfileMock 也能处理中文 Key
         setCapabilityProfile(profile)
      }
    } catch (e) {
      console.warn('Capability profile build skipped or failed, using raw parsed data', e)
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
