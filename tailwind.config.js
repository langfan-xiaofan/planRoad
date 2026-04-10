/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        themePrimary: '#84B4C3', //深色主色调，用于按钮、侧边栏、强调文字
        themeSecondary: '#C7E2E0', //浅色辅色调，用于背景过渡、悬浮状态
        themeBg: '#FAEDE5', //主背景色，米白色带来温暖感
        themeAccent: '#EAC0C3', //强调色，用于警告、错误或高亮标签
      }
    },
  },
  plugins: [],
}