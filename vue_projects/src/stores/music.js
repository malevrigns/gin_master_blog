import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../api'

export const useMusicStore = defineStore('music', () => {
  const currentMusic = ref(null)
  const playlist = ref([])
  const currentIndex = ref(-1)
  const isPlaying = ref(false)
  const isPlayerVisible = ref(false)
  const volume = ref(0.8)
  const progress = ref(0)
  const duration = ref(0)
  const currentTime = ref(0)

  const loadMusics = async () => {
    try {
      const response = await api.get('/music')
      playlist.value = response.data.musics || []
      return response.data
    } catch (error) {
      console.error('Failed to load musics:', error)
      throw error
    }
  }

  const loadPlaylist = async (playlistId) => {
    try {
      const response = await api.get(`/music/playlists/${playlistId}`)
      playlist.value = response.data.musics || []
      return response.data
    } catch (error) {
      console.error('Failed to load playlist:', error)
      throw error
    }
  }

  const play = (music, index = null) => {
    if (index !== null) {
      currentIndex.value = index
    } else {
      currentIndex.value = playlist.value.findIndex(m => m.id === music.id)
    }
    currentMusic.value = playlist.value[currentIndex.value]
    isPlaying.value = true
    isPlayerVisible.value = true
  }

  const pause = () => {
    isPlaying.value = false
  }

  const next = () => {
    if (currentIndex.value < playlist.value.length - 1) {
      currentIndex.value++
      currentMusic.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
  }

  const prev = () => {
    if (currentIndex.value > 0) {
      currentIndex.value--
      currentMusic.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
  }

  const setVolume = (val) => {
    volume.value = val
  }

  const setProgress = (val) => {
    progress.value = val
    currentTime.value = (val / 100) * duration.value
  }

  const setCurrentTime = (time) => {
    currentTime.value = time
    if (duration.value > 0) {
      progress.value = (time / duration.value) * 100
    }
  }

  const setDuration = (dur) => {
    duration.value = dur
  }

  const openPlayer = () => {
    if (currentMusic.value) {
      isPlayerVisible.value = true
    }
  }

  const closePlayer = () => {
    isPlayerVisible.value = false
  }

  const togglePlayer = () => {
    if (!currentMusic.value) {
      return
    }
    isPlayerVisible.value = !isPlayerVisible.value
  }

  return {
    currentMusic,
    playlist,
    currentIndex,
    isPlaying,
    isPlayerVisible,
    volume,
    progress,
    duration,
    currentTime,
    loadMusics,
    loadPlaylist,
    play,
    pause,
    next,
    prev,
    setVolume,
    setProgress,
    setCurrentTime,
    setDuration,
    openPlayer,
    closePlayer,
    togglePlayer,
  }
})

