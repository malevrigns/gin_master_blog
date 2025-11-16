<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-6xl mx-auto">
      <h1 class="text-4xl font-bold mb-8 text-gray-900 dark:text-gray-100">归档</h1>

      <div class="space-y-8">
        <div
          v-for="(posts, year) in archiveData"
          :key="year"
          class="relative"
        >
          <!-- 年份标题 -->
          <div class="flex items-center mb-4">
            <div class="flex items-center space-x-3">
              <div class="w-3 h-3 rounded-full bg-primary-500"></div>
              <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-100">
                {{ year }}
              </h2>
              <span class="text-gray-500 dark:text-gray-400 text-sm">
                {{ posts.length }} 篇文章
              </span>
            </div>
          </div>

          <!-- 时间线 -->
          <div class="ml-6 border-l-2 border-primary-200 dark:border-primary-800 pl-8 space-y-6">
            <div
              v-for="article in posts"
              :key="article.id"
              class="relative group cursor-pointer"
              @click="$router.push(`/article/${article.id}`)"
            >
              <!-- 时间点 -->
              <div class="absolute -left-11 top-2">
                <div class="w-4 h-4 rounded-full bg-primary-500 border-4 border-white dark:border-gray-800 group-hover:scale-125 transition-transform"></div>
              </div>

              <!-- 文章信息 -->
              <div class="bg-white dark:bg-gray-800 rounded-lg p-4 shadow-md hover:shadow-lg transition-all group-hover:border-primary-500 border border-transparent">
                <div class="flex items-center space-x-3 mb-2">
                  <span class="text-sm text-gray-500 dark:text-gray-400">
                    {{ formatMonthDay(article.published_at || article.created_at) }}
                  </span>
                  <el-tag
                    v-if="article.category"
                    type="primary"
                    size="small"
                    effect="plain"
                  >
                    {{ article.category.name }}
                  </el-tag>
                </div>
                <h3 class="text-xl font-semibold text-gray-900 dark:text-gray-100 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors mb-2">
                  {{ article.title }}
                </h3>
                <div class="flex flex-wrap gap-2">
                  <el-tag
                    v-for="tag in article.tags"
                    :key="tag.id"
                    size="small"
                    effect="plain"
                    class="text-xs"
                  >
                    #{{ tag.name }}
                  </el-tag>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-if="Object.keys(archiveData).length === 0" class="text-center py-12">
        <el-icon class="text-6xl text-gray-400 mb-4"><Document /></el-icon>
        <p class="text-gray-500 dark:text-gray-400">暂无文章</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '../api'
import { Document } from '@element-plus/icons-vue'

const articles = ref([])

const archiveData = computed(() => {
  const grouped = {}
  articles.value.forEach((article) => {
    const date = new Date(article.published_at || article.created_at)
    const year = date.getFullYear()
    if (!grouped[year]) {
      grouped[year] = []
    }
    grouped[year].push(article)
  })

  // 按年份降序排序，每年内的文章按日期降序排序
  const sorted = {}
  Object.keys(grouped)
    .sort((a, b) => b - a)
    .forEach((year) => {
      sorted[year] = grouped[year].sort(
        (a, b) =>
          new Date(b.published_at || b.created_at) -
          new Date(a.published_at || a.created_at)
      )
    })

  return sorted
})

const loadArticles = async () => {
  try {
    const response = await api.get('/articles', {
      params: {
        status: 'published',
        page_size: 1000, // 获取所有文章
      },
    })
    articles.value = response.data.articles || []
  } catch (error) {
    console.error('Failed to load articles:', error)
  }
}

const formatMonthDay = (date) => {
  if (!date) return ''
  const d = new Date(date)
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${month}-${day}`
}

onMounted(() => {
  loadArticles()
})
</script>

<style scoped>
</style>

