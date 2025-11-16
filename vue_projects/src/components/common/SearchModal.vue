<template>
  <el-dialog
    v-model="visible"
    title="搜索文章"
    width="600px"
    :before-close="handleClose"
  >
    <div class="search-container">
      <el-input
        v-model="searchQuery"
        placeholder="输入关键词搜索..."
        size="large"
        clearable
        @input="handleSearch"
        @keyup.enter="performSearch"
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
      </el-input>

      <div v-if="loading" class="mt-4 text-center py-8">
        <el-icon class="is-loading text-2xl"><Loading /></el-icon>
        <p class="mt-2 text-gray-500">搜索中...</p>
      </div>

      <div v-else-if="searchResults.length > 0" class="mt-4 space-y-3 max-h-96 overflow-y-auto">
        <div
          v-for="article in searchResults"
          :key="article.id"
          @click="goToArticle(article.id)"
          class="p-4 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-primary-500 cursor-pointer transition-all hover:shadow-md"
        >
          <h4 class="font-semibold text-gray-900 dark:text-gray-100 mb-2">
            {{ highlightText(article.title, searchQuery) }}
          </h4>
          <p class="text-sm text-gray-600 dark:text-gray-400 line-clamp-2">
            {{ highlightText(article.excerpt || article.content.substring(0, 150), searchQuery) }}
          </p>
          <div class="flex items-center mt-2 text-xs text-gray-500 dark:text-gray-400">
            <span>{{ formatDate(article.created_at) }}</span>
            <span class="mx-2">•</span>
            <span>{{ article.views }} 次浏览</span>
          </div>
        </div>
      </div>

      <div v-else-if="searchQuery && !loading" class="mt-4 text-center py-8">
        <el-icon class="text-4xl text-gray-400"><DocumentDelete /></el-icon>
        <p class="mt-2 text-gray-500">未找到相关文章</p>
      </div>

      <div v-else class="mt-4 text-center py-8 text-gray-500">
        <p>输入关键词开始搜索</p>
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Search, Loading, DocumentDelete } from '@element-plus/icons-vue'
import api from '../../api'

const router = useRouter()

const visible = ref(false)
const searchQuery = ref('')
const searchResults = ref([])
const loading = ref(false)

const handleSearch = () => {
  if (!searchQuery.value.trim()) {
    searchResults.value = []
    return
  }
  
  // 防抖搜索
  clearTimeout(window.searchTimeout)
  window.searchTimeout = setTimeout(() => {
    performSearch()
  }, 300)
}

const performSearch = async () => {
  if (!searchQuery.value.trim()) {
    searchResults.value = []
    return
  }

  loading.value = true
  try {
    const response = await api.get('/articles', {
      params: {
        search: searchQuery.value,
        status: 'published',
        page_size: 10,
      },
    })
    searchResults.value = response.data.articles || []
  } catch (error) {
    console.error('Search failed:', error)
    searchResults.value = []
  } finally {
    loading.value = false
  }
}

const highlightText = (text, query) => {
  if (!query || !text) return text
  const regex = new RegExp(`(${query})`, 'gi')
  return text.replace(regex, '<mark class="bg-yellow-200 dark:bg-yellow-800">$1</mark>')
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN')
}

const goToArticle = (id) => {
  router.push(`/article/${id}`)
  handleClose()
}

const handleClose = () => {
  visible.value = false
  searchQuery.value = ''
  searchResults.value = []
}

const open = () => {
  visible.value = true
}

defineExpose({ open })
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>

