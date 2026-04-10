<!-- 认证模块 (登录/注册/忘记密码) -->
 <template>
  <div class="min-h-screen bg-themeBg flex items-center justify-center">
    <div class="bg-white p-8 rounded-xl shadow-lg w-[400px]">
      <h2 class="text-2xl font-bold text-themePrimary text-center mb-6">职业规划智能体</h2>
      
      <el-tabs v-model="activeTab" class="w-full">
        <el-tab-pane label="登录" name="login">
          <el-radio-group v-model="loginType" class="mb-4 w-full flex justify-center">
            <el-radio-button label="account">账号密码</el-radio-button>
            <el-radio-button label="email">邮箱验证</el-radio-button>
          </el-radio-group>

          <el-form :model="loginForm">
            <el-form-item v-if="loginType==='account'">
              <el-input v-model="loginForm.account" placeholder="请输入账号/用户名"></el-input>
            </el-form-item>
            <el-form-item v-if="loginType==='account'">
              <el-input v-model="loginForm.password" type="password" placeholder="请输入密码" show-password></el-input>
            </el-form-item>
            
            <el-form-item v-if="loginType==='email'">
              <el-input v-model="loginForm.email" placeholder="请输入邮箱地址"></el-input>
            </el-form-item>
            <el-form-item v-if="loginType==='email'">
              <div class="flex w-full gap-2">
                <el-input v-model="loginForm.code" placeholder="验证码"></el-input>
                <el-button class="bg-themeSecondary text-themePrimary">获取验证码</el-button>
              </div>
            </el-form-item>

            <div class="flex justify-between items-center mb-4">
              <el-checkbox v-model="rememberMe">记住我</el-checkbox>
              <el-link :underline="false" class="text-themeAccent" @click="activeTab='forgot'">忘记密码?</el-link>
            </div>
            
            <el-button type="primary" class="w-full bg-themePrimary border-none hover:bg-themeSecondary hover:text-themePrimary" @click="handleLogin">登录</el-button>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="注册" name="register">
          <el-form :model="registerForm" class="flex flex-col items-center">
            <el-upload
              class="avatar-uploader mb-4"
              action="https://mock-api.com/upload" 
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
            >
              <img v-if="registerForm.avatar" :src="registerForm.avatar" class="w-20 h-20 rounded-full border-2 border-themePrimary" />
              <div v-else class="w-20 h-20 rounded-full border-2 border-dashed border-themePrimary flex items-center justify-center text-themePrimary hover:bg-themeSecondary">
                上传头像
              </div>
            </el-upload>
            
            <el-form-item class="w-full">
              <el-input v-model="registerForm.account" placeholder="设置用户名"></el-input>
            </el-form-item>
            <el-form-item class="w-full">
              <el-input v-model="registerForm.email" placeholder="设置绑定邮箱"></el-input>
            </el-form-item>
            <el-form-item class="w-full">
              <el-input v-model="registerForm.password" type="password" placeholder="设置密码" show-password></el-input>
            </el-form-item>
            <el-button type="primary" class="w-full bg-themePrimary border-none" @click="handleRegister">注册账号</el-button>
          </el-form>
        </el-tab-pane>

        <el-tab-pane label="找回密码" name="forgot">
          <el-form :model="forgotForm">
            <el-form-item>
              <el-input v-model="forgotForm.email" placeholder="请输入注册邮箱"></el-input>
            </el-form-item>
            <el-form-item>
              <div class="flex w-full gap-2">
                <el-input v-model="forgotForm.code" placeholder="邮箱验证码"></el-input>
                <el-button class="bg-themeSecondary text-themePrimary">获取验证码</el-button>
              </div>
            </el-form-item>
            <el-form-item>
              <el-input v-model="forgotForm.newPassword" type="password" placeholder="输入新密码" show-password></el-input>
            </el-form-item>
            <el-button type="primary" class="w-full bg-themePrimary border-none" @click="handleReset">重置密码</el-button>
            <div class="text-center mt-2">
              <el-link :underline="false" class="text-themePrimary" @click="activeTab='login'">返回登录</el-link>
            </div>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeTab = ref('login')
const loginType = ref('account')
const rememberMe = ref(false)

const loginForm = ref({ account: '', password: '', email: '', code: '' })
const registerForm = ref({ account: '', email: '', password: '', avatar: '' })
const forgotForm = ref({ email: '', code: '', newPassword: '' })

//后续在此处调用后端Login接口
const handleLogin = () => {
  //TODO:接入鉴权API,成功后写入Token并跳转
  router.push('/dashboard')
}

//后续在此处调用后端Register接口
const handleRegister = () => {}

//后续在此处调用后端重置密码接口
const handleReset = () => {}

//处理头像上传成功回调
const handleAvatarSuccess = (res, file) => {
  registerForm.value.avatar = URL.createObjectURL(file.raw)
}
</script>