<template>
  <div
    id="app"
    class="min-h-screen text-gray-900 dark:text-gray-100 transition-colors relative"
  >
    <!-- 背景与装饰层 -->
    <ParticleBackground />

    <!-- 主布局 -->
    <Header />
    <router-view v-slot="{ Component }">
      <transition name="page" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>

    <!-- 全局功能组件 -->
    <MusicPlayer />
    <MusicButton />
    <SakuraEffect />
    <SettingsPanel ref="settingsPanelRef" />
    <ScrollToTop />
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useThemeStore } from './stores/theme'
import { useSettingsStore } from './stores/settings'
import { useAuthStore } from './stores/auth'
import Header from './components/layout/Header.vue'
import Footer from './components/layout/Footer.vue'
import MusicPlayer from './components/music/MusicPlayer.vue'
import MusicButton from './components/music/MusicButton.vue'
import SakuraEffect from './components/features/SakuraEffect.vue'
import SettingsPanel from './components/features/SettingsPanel.vue'
import ParticleBackground from './components/common/ParticleBackground.vue'
import ScrollToTop from './components/common/ScrollToTop.vue'

const themeStore = useThemeStore()
const settingsStore = useSettingsStore()
const authStore = useAuthStore()
const settingsPanelRef = ref(null)

onMounted(() => {
  themeStore.initTheme()
  settingsStore.initSettings()
  authStore.initAuth()

  // 全局快捷键打开设置面板
  window.addEventListener('keydown', (e) => {
    if (e.ctrlKey && e.key === ',') {
      e.preventDefault()
      settingsPanelRef.value?.open()
    }
  })

  // 监听打开设置面板事件
  window.addEventListener('open-settings', () => {
    settingsPanelRef.value?.open()
  })
})
</script>

<style scoped>
#app {
  padding-bottom: 100px; /* 为音乐播放器留出空间 */
}
</style>

<style>
/* 页面过渡动画 */
.page-enter-active {
  animation: pageIn 0.4s ease-out;
}

.page-leave-active {
  animation: pageOut 0.3s ease-in;
}

@keyframes pageIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes pageOut {
  from {
    opacity: 1;
    transform: translateY(0);
  }
  to {
    opacity: 0;
    transform: translateY(-20px);
  }
}
</style>

