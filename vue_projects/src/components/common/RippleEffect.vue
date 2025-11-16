<template>
  <div
    class="ripple-container"
    @click="createRipple"
  >
    <slot></slot>
    <span
      v-for="(ripple, index) in ripples"
      :key="index"
      class="ripple"
      :style="{
        left: ripple.x + 'px',
        top: ripple.y + 'px',
      }"
    ></span>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const ripples = ref([])

const createRipple = (event) => {
  const container = event.currentTarget
  const rect = container.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top

  const ripple = { x, y }
  ripples.value.push(ripple)

  setTimeout(() => {
    ripples.value = ripples.value.filter(r => r !== ripple)
  }, 600)
}
</script>

<style scoped>
.ripple-container {
  position: relative;
  overflow: hidden;
}

.ripple {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.6);
  transform: scale(0);
  animation: ripple-animation 0.6s ease-out;
  pointer-events: none;
  width: 20px;
  height: 20px;
  margin-left: -10px;
  margin-top: -10px;
}

@keyframes ripple-animation {
  to {
    transform: scale(4);
    opacity: 0;
  }
}
</style>

