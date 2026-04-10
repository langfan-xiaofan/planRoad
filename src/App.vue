<template>
  <router-view></router-view>
</template>

<script setup>
// App.vue 作为根容器，不需要写多余的逻辑
</script>

<style>
/* 确保根节点撑满全屏，防止背景色断层 */
html, body, #app {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
}
</style>