<template>
  <span>{{ displayValue }}</span>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'

const props = defineProps({
  value: {
    type: Number,
    required: true,
  },
  duration: {
    type: Number,
    default: 1000,
  },
  decimals: {
    type: Number,
    default: 0,
  },
})

const displayValue = ref(0)

const animate = (start, end) => {
  const startTime = performance.now()
  const difference = end - start

  const step = (currentTime) => {
    const elapsed = currentTime - startTime
    const progress = Math.min(elapsed / props.duration, 1)
    
    // 使用缓动函数
    const easeOutQuart = 1 - Math.pow(1 - progress, 4)
    displayValue.value = Math.floor(start + difference * easeOutQuart)
    
    if (progress < 1) {
      requestAnimationFrame(step)
    } else {
      displayValue.value = end
    }
  }
  
  requestAnimationFrame(step)
}

watch(() => props.value, (newValue) => {
  const start = displayValue.value
  animate(start, newValue)
}, { immediate: true })

onMounted(() => {
  displayValue.value = props.value
})
</script>

