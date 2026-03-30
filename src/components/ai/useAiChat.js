import { ref } from 'vue'

// 在模块作用域外定义 ref，使其成为全局单例状态
const isAiChatVisible = ref(false)

export function useAiChat() {
  const toggleAiChat = () => {
    isAiChatVisible.value = !isAiChatVisible.value
  }

  const openAiChat = () => {
    isAiChatVisible.value = true
  }

  const closeAiChat = () => {
    isAiChatVisible.value = false
  }

  return {
    isAiChatVisible,
    toggleAiChat,
    openAiChat,
    closeAiChat
  }
}