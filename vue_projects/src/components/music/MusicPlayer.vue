<template>
  <div
    v-if="musicStore.isPlayerVisible && musicStore.currentMusic"
    class="fixed bottom-0 left-0 right-0 bg-slate-900/95 backdrop-blur-xl border-t border-slate-800 shadow-2xl shadow-black/40 z-40"
  >
    <div class="container mx-auto px-4 py-3">
      <div class="flex items-center gap-4">
        <div class="flex items-center gap-3 flex-shrink-0 w-64 group">
          <div class="relative">
            <img
              :src="musicStore.currentMusic.cover || '/default-cover.jpg'"
              :alt="musicStore.currentMusic.title"
              class="w-14 h-14 rounded-2xl object-cover shadow-lg group-hover:scale-105 transition-transform"
            />
            <div
              v-if="musicStore.isPlaying"
              class="absolute inset-0 rounded-2xl bg-black/30 flex items-center justify-center"
            >
              <div class="w-2 h-2 bg-white rounded-full animate-pulse" />
            </div>
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold text-slate-100 truncate group-hover:text-emerald-300 transition-colors">
              {{ musicStore.currentMusic.title }}
            </p>
            <p class="text-xs text-slate-400 truncate">
              {{ musicStore.currentMusic.artist || '未知艺术家' }}
            </p>
          </div>
        </div>

        <div class="flex-1 flex flex-col items-center gap-2">
          <div class="flex items-center gap-4">
            <el-button
              :icon="PreviousIcon"
              circle
              size="small"
              @click="musicStore.prev()"
              :disabled="musicStore.currentIndex <= 0"
            />
            <el-button
              :icon="musicStore.isPlaying ? PauseIcon : PlayIcon"
              circle
              type="primary"
              @click="togglePlay"
            />
            <el-button
              :icon="NextIcon"
              circle
              size="small"
              @click="musicStore.next()"
              :disabled="musicStore.currentIndex >= musicStore.playlist.length - 1"
            />
          </div>
          <div class="w-full flex items-center gap-2">
            <span class="text-xs text-slate-400 w-10 text-right">
              {{ formatTime(musicStore.currentTime) }}
            </span>
            <el-slider
              v-model="musicStore.progress"
              :max="100"
              :show-tooltip="false"
              @change="handleProgressChange"
              class="flex-1"
            />
            <span class="text-xs text-slate-400 w-10">
              {{ formatTime(musicStore.duration) }}
            </span>
          </div>
        </div>

        <div class="flex items-center gap-2 w-32">
          <svg class="w-5 h-5 text-slate-400" fill="currentColor" viewBox="0 0 24 24">
            <path
              d="M3 9v6h4l5 5V4L7 9H3zm13.5 3c0-1.77-1.02-3.29-2.5-4.03v8.05c1.48-.73 2.5-2.25 2.5-4.02zM14 3.23v2.06c2.89.86 5 3.54 5 6.71s-2.11 5.85-5 6.71v2.06c4.01-.91 7-4.49 7-8.77s-2.99-7.86-7-8.77z"
            />
          </svg>
          <el-slider
            v-model="musicStore.volume"
            :max="1"
            :step="0.01"
            :format-tooltip="(val) => Math.round(val * 100) + '%'"
            @change="handleVolumeChange"
            class="flex-1"
          />
        </div>

        <div class="flex items-center gap-2">
          <el-button :icon="ListIcon" circle @click="showPlaylist = !showPlaylist" />
          <el-button :icon="CloseIcon" circle @click="musicStore.closePlayer()" />
        </div>
      </div>
    </div>

    <transition name="playlist-slide">
      <div
        v-if="showPlaylist"
        class="fixed bottom-24 right-6 w-80 max-h-[60vh] bg-slate-900/95 border border-slate-800 rounded-2xl shadow-2xl shadow-black/40 overflow-hidden z-50"
      >
        <div class="flex items-center justify-between px-4 py-3 border-b border-slate-800">
          <span class="text-sm text-slate-200">播放列表</span>
          <el-button text :icon="CloseIcon" @click="showPlaylist = false" />
        </div>
        <div class="max-h-[50vh] overflow-y-auto p-3 space-y-2">
          <div
            v-for="(music, index) in musicStore.playlist"
            :key="music.id"
            @click="musicStore.play(music, index)"
            :class="[
              'flex items-center gap-3 p-3 rounded-2xl cursor-pointer transition-all',
              musicStore.currentIndex === index
                ? 'bg-emerald-500/10 border border-emerald-400'
                : 'border border-transparent hover:border-slate-700'
            ]"
          >
            <img
              :src="music.cover || '/default-cover.jpg'"
              :alt="music.title"
              class="w-12 h-12 rounded-xl object-cover"
            />
            <div class="flex-1 min-w-0">
              <p class="text-sm text-slate-100 truncate">{{ music.title }}</p>
              <p class="text-xs text-slate-400 truncate">
                {{ music.artist || '未知艺术家' }}
              </p>
            </div>
            <el-icon v-if="musicStore.currentIndex === index" class="text-emerald-400 text-lg">
              <VideoPlayIcon />
            </el-icon>
          </div>
          <div v-if="musicStore.playlist.length === 0" class="text-center py-8 text-slate-500">
            播放列表为空
          </div>
        </div>
      </div>
    </transition>

    <audio
      ref="audioRef"
      :src="musicStore.currentMusic?.url"
      @loadedmetadata="handleLoadedMetadata"
      @timeupdate="handleTimeUpdate"
      @ended="handleEnded"
      @volumechange="handleVolumeChange"
    />
  </div>
</template>

<script setup>
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { useMusicStore } from '../../stores/music'
import {
  VideoPlay as PlayIcon,
  VideoPause as PauseIcon,
  ArrowLeft as PreviousIcon,
  ArrowRight as NextIcon,
  List as ListIcon,
  VideoPlay as VideoPlayIcon,
  Close as CloseIcon,
} from '@element-plus/icons-vue'

const musicStore = useMusicStore()
const audioRef = ref(null)
const showPlaylist = ref(false)

const togglePlay = () => {
  if (musicStore.isPlaying) {
    audioRef.value?.pause()
    musicStore.pause()
  } else {
    audioRef.value?.play()
    musicStore.play(musicStore.currentMusic, musicStore.currentIndex)
  }
}

const handleProgressChange = (val) => {
  if (audioRef.value) {
    const time = (val / 100) * musicStore.duration
    audioRef.value.currentTime = time
    musicStore.setCurrentTime(time)
  }
}

const handleVolumeChange = () => {
  if (audioRef.value) {
    audioRef.value.volume = musicStore.volume
  }
}

const handleLoadedMetadata = () => {
  if (audioRef.value) {
    musicStore.setDuration(audioRef.value.duration)
    audioRef.value.volume = musicStore.volume
  }
}

const handleTimeUpdate = () => {
  if (audioRef.value) {
    musicStore.setCurrentTime(audioRef.value.currentTime)
  }
}

const handleEnded = () => {
  musicStore.next()
}

const formatTime = (seconds) => {
  if (!seconds || isNaN(seconds)) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

watch(
  () => musicStore.isPlaying,
  (isPlaying) => {
    if (audioRef.value) {
      if (isPlaying) {
        audioRef.value.play()
      } else {
        audioRef.value.pause()
      }
    }
  }
)

watch(
  () => musicStore.currentMusic,
  () => {
    if (audioRef.value && musicStore.currentMusic) {
      audioRef.value.load()
      if (musicStore.isPlaying) {
        audioRef.value.play()
      }
    }
  }
)

watch(
  () => musicStore.isPlayerVisible,
  (visible) => {
    if (!visible) {
      showPlaylist.value = false
    }
  }
)

onMounted(() => {
  musicStore.loadMusics()
  if (audioRef.value) {
    audioRef.value.volume = musicStore.volume
  }
})

onUnmounted(() => {
  if (audioRef.value) {
    audioRef.value.pause()
  }
})
</script>

<style scoped>
.playlist-slide-enter-active,
.playlist-slide-leave-active {
  transition: transform 0.25s ease, opacity 0.25s ease;
}
.playlist-slide-enter-from,
.playlist-slide-leave-to {
  transform: translateY(20px);
  opacity: 0;
}
</style>
