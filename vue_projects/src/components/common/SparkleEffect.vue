<template>
  <div class="sparkle-container absolute inset-0 pointer-events-none overflow-hidden">
    <div
      v-for="i in count"
      :key="i"
      class="sparkle"
      :style="getSparkleStyle(i)"
    ></div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  count: {
    type: Number,
    default: 20,
  },
})

const getSparkleStyle = (index) => {
  const delay = Math.random() * 2
  const duration = 2 + Math.random() * 2
  const left = Math.random() * 100
  const top = Math.random() * 100
  const size = 4 + Math.random() * 4
  
  return {
    left: `${left}%`,
    top: `${top}%`,
    width: `${size}px`,
    height: `${size}px`,
    animationDelay: `${delay}s`,
    animationDuration: `${duration}s`,
  }
}
</script>

<style scoped>
.sparkle-container {
  overflow: hidden;
}

.sparkle {
  position: absolute;
  background: radial-gradient(circle, #fff 0%, transparent 70%);
  border-radius: 50%;
  opacity: 0;
  animation: sparkle-animation linear infinite;
  pointer-events: none;
}

@keyframes sparkle-animation {
  0% {
    opacity: 0;
    transform: scale(0) rotate(0deg);
  }
  50% {
    opacity: 1;
    transform: scale(1) rotate(180deg);
  }
  100% {
    opacity: 0;
    transform: scale(0) rotate(360deg);
  }
}
</style>

