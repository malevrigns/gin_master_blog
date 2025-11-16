<template>
  <span class="typewriter-text">{{ displayText }}<span class="cursor">|</span></span>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  speed: {
    type: Number,
    default: 100,
  },
  delay: {
    type: Number,
    default: 0,
  },
})

const displayText = ref('')
let timeoutId = null
let index = 0

const type = () => {
  if (index < props.text.length) {
    displayText.value += props.text[index]
    index++
    timeoutId = setTimeout(type, props.speed)
  }
}

onMounted(() => {
  if (props.delay > 0) {
    timeoutId = setTimeout(type, props.delay)
  } else {
    type()
  }
})

onUnmounted(() => {
  if (timeoutId) {
    clearTimeout(timeoutId)
  }
})
</script>

<style scoped>
.typewriter-text {
  display: inline-block;
}

.cursor {
  animation: blink 1s infinite;
  color: currentColor;
}

@keyframes blink {
  0%, 50% {
    opacity: 1;
  }
  51%, 100% {
    opacity: 0;
  }
}
</style>

