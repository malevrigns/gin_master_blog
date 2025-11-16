<template>
  <footer class="bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 mt-12">
    <div class="container mx-auto px-4 py-8">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div>
          <h3 class="text-lg font-semibold mb-4 text-gray-800 dark:text-gray-200">关于博客</h3>
          <p class="text-gray-600 dark:text-gray-400">
            一个基于 Go + Gin + Vue3 构建的现代化博客系统
          </p>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4 text-gray-800 dark:text-gray-200">快速链接</h3>
          <ul class="space-y-2">
            <li>
              <router-link
                to="/"
                class="text-gray-600 dark:text-gray-400 hover:text-primary-600 dark:hover:text-primary-400"
              >
                首页
              </router-link>
            </li>
            <li>
              <router-link
                to="/about"
                class="text-gray-600 dark:text-gray-400 hover:text-primary-600 dark:hover:text-primary-400"
              >
                关于
              </router-link>
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4 text-gray-800 dark:text-gray-200">联系方式</h3>
          <p class="text-gray-600 dark:text-gray-400">
            Email: your-email@example.com
          </p>
        </div>
      </div>
      <div class="mt-8 pt-8 border-t border-gray-200 dark:border-gray-700 text-center text-gray-600 dark:text-gray-400">
        <p>&copy; 2024 我的博客. All rights reserved.</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
</script>

