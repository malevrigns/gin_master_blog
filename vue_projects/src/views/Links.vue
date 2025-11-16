<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-6xl mx-auto">
      <h1 class="text-4xl font-bold mb-8 text-gray-900 dark:text-gray-100">友情链接</h1>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="link in links"
          :key="link.id"
          class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 hover:shadow-xl transition-all duration-300 border border-transparent hover:border-primary-500/50 group"
        >
          <a
            :href="link.url"
            target="_blank"
            rel="noopener noreferrer"
            class="block"
          >
            <div class="flex items-center space-x-4 mb-4">
              <img
                v-if="link.logo"
                :src="link.logo"
                :alt="link.name"
                class="w-16 h-16 rounded-lg object-cover border-2 border-gray-200 dark:border-gray-700 group-hover:border-primary-500 transition-colors"
              />
              <div v-else class="w-16 h-16 rounded-lg bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center text-white text-2xl font-bold">
                {{ link.name[0] }}
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors truncate">
                  {{ link.name }}
                </h3>
                <p class="text-sm text-gray-500 dark:text-gray-400 truncate">
                  {{ link.desc || '暂无描述' }}
                </p>
              </div>
            </div>
            <div class="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
              <span class="truncate flex-1">{{ link.url }}</span>
              <el-icon class="ml-2 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                <ArrowRight />
              </el-icon>
            </div>
          </a>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-if="links.length === 0" class="text-center py-12">
        <el-icon class="text-6xl text-gray-400 mb-4"><Link /></el-icon>
        <p class="text-gray-500 dark:text-gray-400">暂无友情链接</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api'
import { ArrowRight, Link } from '@element-plus/icons-vue'

const links = ref([])

const loadLinks = async () => {
  try {
    const response = await api.get('/links')
    links.value = response.data.filter(link => link.is_visible)
  } catch (error) {
    console.error('Failed to load links:', error)
    // 如果API不存在，使用示例数据
    links.value = [
      {
        id: 1,
        name: 'GitHub',
        url: 'https://github.com',
        logo: 'https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png',
        desc: '代码托管平台',
        is_visible: true,
      },
      {
        id: 2,
        name: 'Vue.js',
        url: 'https://vuejs.org',
        logo: 'https://vuejs.org/logo.svg',
        desc: '渐进式 JavaScript 框架',
        is_visible: true,
      },
    ]
  }
}

onMounted(() => {
  loadLinks()
})
</script>

