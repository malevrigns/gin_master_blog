<template>
  <button
    @click="$emit('click')"
    class="fixed bottom-24 right-6 w-14 h-14 bg-primary-600 hover:bg-primary-700 text-white rounded-full shadow-lg hover:shadow-xl transition-all duration-300 flex items-center justify-center z-40 group animate-float"
    :class="customClass"
  >
    <slot>
      <el-icon class="text-2xl">
        <component :is="icon" />
      </el-icon>
    </slot>
    <div class="absolute inset-0 rounded-full bg-white/20 scale-0 group-hover:scale-150 transition-transform duration-300"></div>
  </button>
</template>

<script setup>
defineProps({
  icon: {
    type: [String, Object],
    default: 'Plus',
  },
  customClass: {
    type: String,
    default: '',
  },
})

defineEmits(['click'])
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

