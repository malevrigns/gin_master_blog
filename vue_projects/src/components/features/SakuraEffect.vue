<template>
  <div v-if="enabled" class="sakura-container fixed inset-0 pointer-events-none z-0">
    <div
      v-for="i in count"
      :key="i"
      class="sakura"
      :style="getSakuraStyle(i)"
    ></div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useSettingsStore } from '../../stores/settings'

const settingsStore = useSettingsStore()

const enabled = computed(() => settingsStore.sakuraEnabled)
const count = computed(() => settingsStore.sakuraCount)

const getSakuraStyle = (index) => {
  const delay = Math.random() * 5
  const duration = 10 + Math.random() * 10
  const left = Math.random() * 100
  const animationDelay = Math.random() * 5
  
  return {
    left: `${left}%`,
    animationDelay: `${animationDelay}s`,
    animationDuration: `${duration}s`,
  }
}
</script>

<style scoped>
.sakura-container {
  overflow: hidden;
}

.sakura {
  position: absolute;
  width: 10px;
  height: 10px;
  background: radial-gradient(circle, #ffb7c5 0%, #ff9ec4 100%);
  border-radius: 50% 0 50% 0;
  opacity: 0.8;
  animation: fall linear infinite;
  pointer-events: none;
}

.sakura::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle, #ffb7c5 0%, #ff9ec4 100%);
  border-radius: 50% 0 50% 0;
  transform: rotate(45deg);
}

@keyframes fall {
  0% {
    transform: translateY(-100vh) rotate(0deg);
    opacity: 1;
  }
  100% {
    transform: translateY(100vh) rotate(360deg);
    opacity: 0;
  }
}
</style>

