<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-6xl mx-auto">
      <h1 class="text-3xl font-bold mb-8 text-gray-900 dark:text-gray-100">
        标签: {{ tag?.name || route.params.slug }}
      </h1>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <article
          v-for="article in articles"
          :key="article.id"
          class="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow cursor-pointer"
          @click="$router.push(`/article/${article.id}`)"
        >
          <div v-if="article.cover_image" class="h-48 overflow-hidden">
            <img
              :src="article.cover_image"
              :alt="article.title"
              class="w-full h-full object-cover"
            />
          </div>
          <div class="p-6">
            <h2 class="text-xl font-bold mb-2 text-gray-900 dark:text-gray-100">
              {{ article.title }}
            </h2>
            <p class="text-gray-600 dark:text-gray-400 mb-4 line-clamp-2">
              {{ article.excerpt || article.content.substring(0, 100) + '...' }}
            </p>
            <div class="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
              <span>{{ formatDate(article.published_at || article.created_at) }}</span>
              <span>{{ article.views }} 次浏览</span>
            </div>
          </div>
        </article>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '../api'

const route = useRoute()

const tag = ref(null)
const articles = ref([])

const loadArticles = async () => {
  try {
    const response = await api.get('/articles', {
      params: {
        tag: route.params.slug,
        status: 'published',
      },
    })
    articles.value = response.data.articles || []
    
    // 尝试获取标签信息
    const tagsResponse = await api.get('/tags')
    tag.value = tagsResponse.data.find(t => t.slug === route.params.slug)
  } catch (error) {
    console.error('Failed to load articles:', error)
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN')
}

onMounted(() => {
  loadArticles()
})
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>

