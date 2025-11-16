<template>
  <div class="flex items-center justify-center py-12">
    <div class="relative w-16 h-16">
      <div class="absolute inset-0 border-4 border-primary-200 dark:border-primary-800 rounded-full"></div>
      <div class="absolute inset-0 border-4 border-transparent border-t-primary-600 rounded-full animate-spin"></div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>

