<template>
  <div
    v-if="visible && announcement"
    class="bg-gradient-to-r from-primary-500 via-primary-600 to-primary-700 rounded-lg shadow-md p-4 mb-6 text-white relative overflow-hidden group hover:shadow-xl transition-all duration-300"
  >
    <!-- 背景动画 -->
    <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
    <button
      @click="close"
      class="absolute top-2 right-2 text-white hover:text-gray-200 transition-all hover:scale-110 z-10"
    >
      <el-icon><Close /></el-icon>
    </button>
    <div class="flex items-start space-x-3 relative z-10">
      <el-icon class="text-2xl animate-bounce"><Bell /></el-icon>
      <div class="flex-1">
        <h3 class="font-semibold mb-1">{{ announcement.title || '公告' }}</h3>
        <p class="text-sm opacity-90">{{ announcement.content }}</p>
        <a
          v-if="announcement.link"
          :href="announcement.link"
          class="text-sm underline mt-2 inline-block hover:opacity-80"
        >
          了解更多 →
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Bell, Close } from '@element-plus/icons-vue'

const visible = ref(true)
const announcement = ref({
  title: '欢迎来到我的博客',
  content: '这是一个基于 Go + Gin + Vue3 构建的现代化博客系统，参考了 Firefly 的设计理念，并添加了音乐播放功能。',
  link: '/about',
})

const close = () => {
  visible.value = false
  localStorage.setItem('announcement_closed', 'true')
}

onMounted(() => {
  const closed = localStorage.getItem('announcement_closed')
  if (closed === 'true') {
    visible.value = false
  }
})
</script>

