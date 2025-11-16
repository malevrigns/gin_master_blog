<template>
  <button
    @click="toggleMusicPlayer"
    class="fixed bottom-24 right-6 w-14 h-14 bg-primary-600 hover:bg-primary-700 text-white rounded-full shadow-lg hover:shadow-xl transition-all duration-300 flex items-center justify-center z-40 group animate-float"
    :class="{ 'bg-green-600 hover:bg-green-700': musicStore.currentMusic }"
  >
    <el-icon class="text-2xl">
      <component :is="musicStore.currentMusic ? VideoPlayIcon : HeadsetIcon" />
    </el-icon>
    <div class="absolute inset-0 rounded-full bg-white/20 scale-0 group-hover:scale-150 transition-transform duration-300"></div>
  </button>
</template>

<script setup>
import { onMounted } from 'vue'
import { VideoPlay as VideoPlayIcon, Headset as HeadsetIcon } from '@element-plus/icons-vue'
import { useMusicStore } from '../../stores/music'

const musicStore = useMusicStore()

const toggleMusicPlayer = async () => {
  if (musicStore.playlist.length === 0) {
    try {
      await musicStore.loadMusics()
    } catch (error) {
      console.error('Failed to load music:', error)
      return
    }
  }

  if (!musicStore.currentMusic && musicStore.playlist.length > 0) {
    musicStore.play(musicStore.playlist[0], 0)
    return
  }

  musicStore.togglePlayer()
}

onMounted(async () => {
  // 自动加载音乐列表
  if (musicStore.playlist.length === 0) {
    try {
      await musicStore.loadMusics()
    } catch (error) {
      console.error('Failed to load music:', error)
    }
  }
})
</script>

<style scoped>
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}
</style>

