import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// 引入 Element Plus 及全局样式
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 引入 Element Plus 的图标（如果后续侧边栏需要用到）
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// 引入 Tailwind CSS 及自定义全局样式
import './assets/style.css'

const app = createApp(App)

// 注册所有 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)
app.use(ElementPlus)

app.mount('#app')