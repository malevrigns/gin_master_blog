<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-4xl mx-auto">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-8">
        <h1 class="text-4xl font-bold mb-6 text-gray-900 dark:text-gray-100">关于我</h1>
        <div class="prose dark:prose-invert max-w-none">
          <p class="text-lg text-gray-700 dark:text-gray-300 mb-4">
            欢迎来到我的博客！这里记录着我的技术学习、生活感悟和思考。
          </p>
          <p class="text-gray-700 dark:text-gray-300 mb-4">
            这个博客系统基于 Go + Gin + Vue3 构建，采用了现代化的技术栈，提供了完整的文章管理、评论系统、音乐播放等功能。
          </p>
          <h2 class="text-2xl font-bold mt-8 mb-4 text-gray-900 dark:text-gray-100">技术栈</h2>
          <ul class="list-disc list-inside text-gray-700 dark:text-gray-300 mb-4">
            <li>后端: Go + Gin + GORM</li>
            <li>前端: Vue3 + Vite + Tailwind CSS + Element Plus</li>
            <li>数据库: SQLite / MySQL</li>
            <li>认证: JWT</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.prose {
  @apply text-gray-700 dark:text-gray-300;
}
</style>

