/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          cream: '#F7EECD',  // 奶黄 (背景/侧边栏底色)
          pink: '#EFDCE2',   // 浅粉 (选中状态/次要点缀)
          yellow: '#E0DC95', // 青柠黄 (警告/辅助强调)
          green: '#C2D68F',  // 抹茶绿 (主操作/按钮/核心图表)

          // 补充深色系，用于文字和图标，保证极佳的可读性
          text: '#4B5563',   // 默认文字灰
          dark: '#2A342B',   // 极深灰绿 (标题)
          'green-dark': '#8A9E58', // 抹茶绿的加深版，用于Hover效果
        }
      },
      fontFamily: {
        sans: ['Inter', 'Microsoft YaHei', 'sans-serif'],
      }
    },
  },
  plugins: [],
}