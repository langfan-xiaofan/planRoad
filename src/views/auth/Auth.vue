<template>
  <!-- 登录注册页 -->
  <div class="min-h-screen w-full bg-[#F7EECD]/40 flex items-center justify-center relative overflow-hidden font-sans">
    <div class="absolute top-[-10%] left-[-10%] w-96 h-96 bg-[#EFDCE2]/60 rounded-full blur-3xl pointer-events-none"></div>
    <div class="absolute bottom-[-10%] right-[-10%] w-96 h-96 bg-[#C2D68F]/40 rounded-full blur-3xl pointer-events-none"></div>

    <div class="w-full max-w-md bg-white/90 backdrop-blur-xl rounded-3xl shadow-2xl shadow-[#C2D68F]/20 p-8 relative z-10 border border-white">
      <!-- 登录注册页标题 -->
      <div class="text-center mb-8">
        <div class="w-16 h-16 bg-gradient-to-br from-[#C2D68F] to-[#8A9E58] rounded-2xl mx-auto flex items-center justify-center shadow-md mb-4 text-white">
          <el-icon :size="32"><Compass /></el-icon>
        </div>
        <h2 class="text-2xl font-extrabold text-gray-800 tracking-wider">{{ viewTitles[currentView] }}</h2>
        <p class="text-sm text-gray-400 mt-2">{{ viewSubtitles[currentView] }}</p>
      </div>
      <!-- 登录注册页表单切换 -->
      <transition name="fade-slide" mode="out-in">
        <div v-if="currentView === 'login'" key="login">
          
          <div class="flex bg-[#F7EECD]/30 rounded-xl p-1 mb-6 relative">
            <div class="absolute inset-y-1 w-[calc(50%-4px)] bg-white rounded-lg shadow-sm transition-transform duration-300 ease-out"
                 :class="loginType === 'email' ? 'translate-x-[calc(100%+8px)]' : 'translate-x-0'"></div>
            <button class="flex-1 py-2 text-sm font-bold z-10 transition-colors"
                    :class="loginType === 'account' ? 'text-[#8A9E58]' : 'text-gray-400'"
                    @click="loginType = 'account'">账号密码登录</button>
            <button class="flex-1 py-2 text-sm font-bold z-10 transition-colors"
                    :class="loginType === 'email' ? 'text-[#8A9E58]' : 'text-gray-400'"
                    @click="loginType = 'email'">免密邮箱登录</button>
          </div>
          <!-- 账号密码登录表单 -->
          <el-form v-if="loginType === 'account'" :model="accountForm" :rules="accountRules" ref="accountFormRef" @keyup.enter="handleLogin">
            <el-form-item prop="username">
              <el-input v-model="accountForm.username" placeholder="请输入用户名" size="large" :prefix-icon="User" class="custom-input" />
            </el-form-item>
            <el-form-item prop="password">
              <el-input v-model="accountForm.password" type="password" placeholder="请输入密码" size="large" show-password :prefix-icon="Lock" class="custom-input" />
            </el-form-item>
            <div class="flex justify-between items-center mb-6 px-1">
              <el-checkbox v-model="accountForm.remember" class="custom-checkbox">记住我</el-checkbox>
              <span class="text-sm text-[#EFDCE2] hover:text-[#C2D68F] cursor-pointer font-bold transition-colors" @click="switchView('forgot')">忘记密码？</span>
            </div>
            <el-button color="#C2D68F" class="w-full !rounded-xl !h-11 text-white font-bold shadow-md shadow-[#C2D68F]/30 hover:opacity-90" :loading="loading" @click="handleLogin">
              登 录
            </el-button>
          </el-form>
          <!-- 免密邮箱登录表单 -->
          <el-form v-else :model="emailForm" :rules="emailRules" ref="emailFormRef" @keyup.enter="handleLogin">
            <el-form-item prop="email">
              <el-input v-model="emailForm.email" placeholder="请输入邮箱地址" size="large" :prefix-icon="Message" class="custom-input" />
            </el-form-item>
            <el-form-item prop="code">
              <div class="flex gap-3 w-full">
                <el-input v-model="emailForm.code" placeholder="请输入验证码" size="large" :prefix-icon="Key" class="custom-input flex-1" />
                <el-button class="!rounded-xl !h-11 !w-28 text-[#8A9E58] bg-[#F7EECD]/50 hover:bg-[#F7EECD] border-none" 
                           :disabled="countdown > 0" @click="sendCode('login', emailForm.email)">
                  {{ countdown > 0 ? `${countdown}s 后重试` : '获取验证码' }}
                </el-button>
              </div>
            </el-form-item>
            <el-button color="#C2D68F" class="w-full !rounded-xl !h-11 mt-2 text-white font-bold shadow-md shadow-[#C2D68F]/30 hover:opacity-90" :loading="loading" @click="handleLogin">
              验 证 并 登 录
            </el-button>
          </el-form>
        </div>
        <!-- 注册表单切换 -->
        <div v-else-if="currentView === 'register'" key="register">
          <el-form :model="registerForm" :rules="registerRules" ref="registerFormRef">
            <el-form-item prop="username">
              <el-input v-model="registerForm.username" placeholder="设置用户名" size="large" :prefix-icon="User" class="custom-input" />
            </el-form-item>
            <el-form-item prop="email">
              <el-input v-model="registerForm.email" placeholder="输入常用邮箱" size="large" :prefix-icon="Message" class="custom-input" />
            </el-form-item>
            <el-form-item prop="code">
              <div class="flex gap-3 w-full">
                <el-input v-model="registerForm.code" placeholder="邮箱验证码" size="large" :prefix-icon="Key" class="custom-input flex-1" />
                <el-button class="!rounded-xl !h-11 !w-28 text-[#8A9E58] bg-[#F7EECD]/50 hover:bg-[#F7EECD] border-none" 
                           :disabled="countdown > 0" @click="sendCode('register', registerForm.email)">
                  {{ countdown > 0 ? `${countdown}s 后重试` : '获取验证码' }}
                </el-button>
              </div>
            </el-form-item>
            <el-form-item prop="password">
              <el-input v-model="registerForm.password" type="password" placeholder="设置密码" size="large" show-password :prefix-icon="Lock" class="custom-input" />
            </el-form-item>
            <el-button color="#C2D68F" class="w-full !rounded-xl !h-11 mt-2 text-white font-bold shadow-md shadow-[#C2D68F]/30 hover:opacity-90" :loading="loading" @click="handleRegister">
              立 即 注 册
            </el-button>
          </el-form>
        </div>
        <!-- 忘记密码表单切换 -->
        <div v-else-if="currentView === 'forgot'" key="forgot">
          <el-form :model="forgotForm" :rules="forgotRules" ref="forgotFormRef">
            <el-form-item prop="email">
              <el-input v-model="forgotForm.email" placeholder="请输入绑定的邮箱" size="large" :prefix-icon="Message" class="custom-input" />
            </el-form-item>
            <el-form-item prop="code">
              <div class="flex gap-3 w-full">
                <el-input v-model="forgotForm.code" placeholder="邮箱验证码" size="large" :prefix-icon="Key" class="custom-input flex-1" />
                <el-button class="!rounded-xl !h-11 !w-28 text-[#8A9E58] bg-[#F7EECD]/50 hover:bg-[#F7EECD] border-none" 
                           :disabled="countdown > 0" @click="sendCode('reset-password', forgotForm.email)">
                  {{ countdown > 0 ? `${countdown}s` : '获取验证码' }}
                </el-button>
              </div>
            </el-form-item>
            <el-form-item prop="newPassword">
              <el-input v-model="forgotForm.newPassword" type="password" placeholder="输入新密码" size="large" show-password :prefix-icon="Lock" class="custom-input" />
            </el-form-item>
            <el-button color="#EFDCE2" class="w-full !rounded-xl !h-11 mt-2 !text-gray-700 font-bold shadow-md shadow-[#EFDCE2]/30 hover:!bg-[#e6cdd4]" :loading="loading" @click="handleResetPassword">
              重 置 密 码
            </el-button>
          </el-form>
        </div>
      </transition>
      <!-- 登录注册页切换 -->
      <div class="mt-8 text-center text-sm text-gray-500">
        <template v-if="currentView === 'login'">
          还没有账号？ 
          <span class="text-[#C2D68F] font-bold cursor-pointer hover:underline" @click="switchView('register')">立即注册</span>
        </template>
        <template v-else>
          想起来了？ 
          <span class="text-[#C2D68F] font-bold cursor-pointer hover:underline" @click="switchView('login')">返回登录</span>
        </template>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Compass, User, Lock, Message, Key } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user' 
 

const router = useRouter()
const userStore = useUserStore()

// 视图与状态控制
const currentView = ref('login') 
const loginType = ref('account') 
const loading = ref(false)

// 文本配置
const viewTitles = { login: '欢迎回来', register: '开启职业探索', forgot: '找回密码' }
const viewSubtitles = { login: 'CareerAgent 智能职业规划', register: '只需几秒钟，创建你的专属档案', forgot: '通过验证邮箱重设您的密码' }

const switchView = (view) => {
  currentView.value = view
  // 切换视图时清除倒计时
  countdown.value = 0 
}

// ================= 表单数据与验证规则 =================
const accountFormRef = ref(null)
const accountForm = reactive({ username: '', password: '', remember: false })
const accountRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const emailFormRef = ref(null)
const emailForm = reactive({ email: '', code: '' })
const emailRules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }
  ],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
}

const registerFormRef = ref(null)
const registerForm = reactive({ username: '', email: '', password: '', code: '' })
const registerRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
  password: [{ required: true, message: '请设置密码', trigger: 'blur' }, { min: 6, message: '密码至少 6 位', trigger: 'blur' }],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
}

const forgotFormRef = ref(null)
const forgotForm = reactive({ email: '', code: '', newPassword: '' })
const forgotRules = {
  email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  newPassword: [{ required: true, message: '请设置新密码', trigger: 'blur' }, { min: 6, message: '密码至少 6 位', trigger: 'blur' }]
}

// ================= 验证码倒计时逻辑 =================
const countdown = ref(0)
let timer = null

const startCountdown = () => {
  countdown.value = 60
  timer = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) clearInterval(timer)
  }, 1000)
}

// 发送验证码 API 调用
const sendCode = async (type, email) => {
  if (!email || !/^\S+@\S+\.\S+$/.test(email)) {
    ElMessage.warning('请先输入正确的邮箱地址')
    return
  }
  
  try {
    // 模拟 API 调用，(/api/auth/send-code)
    // await axios.post('/api/auth/send-code', { email, type })
    console.log(`Sending ${type} code to ${email}`)
    
    ElMessage.success('验证码发送成功，请注意查收')
    startCountdown()
  } catch (error) {
    ElMessage.error(error.response?.data?.message || '发送失败，请重试')
  }
}

// ================= 登录 API 调用 =================
const handleLogin = async () => {
  const formRef = loginType.value === 'account' ? accountFormRef : emailFormRef
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        let res;
        if (loginType.value === 'account') {
          // (/api/auth/login)
          // res = await axios.post('/api/auth/login', accountForm)
          console.log('Account Login Payload:', accountForm)
          res = { data: { token: 'mock_token_123', user: { id: 1, username: accountForm.username } } } // Mock response
        } else {
          // (/api/auth/email-login)
          // res = await axios.post('/api/auth/email-login', emailForm)
          console.log('Email Login Payload:', emailForm)
          res = { data: { token: 'mock_token_456', user: { id: 1, email: emailForm.email } } } // Mock response
        }

        // 保存登录状态 
        // userStore.userName = res.data.user.username || '用户'
        // userStore.isLoggedIn = true
        userStore.login(res.data.token, res.data.user)

        ElMessage.success('登录成功！')
        router.push('/') // 跳转到首页
      } catch (error) {
        ElMessage.error(error.response?.data?.message || '登录失败，请检查账号密码')
      } finally {
        loading.value = false
      }
    }
  })
}

// ================= 注册 API 调用 =================
const handleRegister = async () => {
  if (!registerFormRef.value) return
  await registerFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        // await axios.post('/api/auth/register', registerForm)
        console.log('Register Payload:', registerForm)
        ElMessage.success('注册成功，请登录')
        switchView('login')
      } catch (error) {
        ElMessage.error(error.response?.data?.message || '注册失败')
      } finally {
        loading.value = false
      }
    }
  })
}

// ================= 重置密码 API 调用 =================
const handleResetPassword = async () => {
  if (!forgotFormRef.value) return
  await forgotFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        // await axios.post('/api/auth/reset-password', forgotForm)
        console.log('Reset Password Payload:', forgotForm)
        ElMessage.success('密码重置成功，请使用新密码登录')
        switchView('login')
      } catch (error) {
        ElMessage.error(error.response?.data?.message || '重置失败')
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
/* 视图切换动画 */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(20px);
}

.custom-input :deep(.el-input__wrapper) {
  background-color: #f9fafb;
  box-shadow: 0 0 0 1px #f3f4f6 inset;
  border-radius: 12px;
  padding: 4px 12px;
  transition: all 0.3s;
}
.custom-input :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #C2D68F inset !important;
  background-color: #ffffff;
}

.custom-checkbox :deep(.el-checkbox__input.is-checked .el-checkbox__inner),
.custom-checkbox :deep(.el-checkbox__input.is-indeterminate .el-checkbox__inner) {
  background-color: #C2D68F;
  border-color: #C2D68F;
}
.custom-checkbox :deep(.el-checkbox__input.is-checked + .el-checkbox__label) {
  color: #8A9E58;
}
</style>