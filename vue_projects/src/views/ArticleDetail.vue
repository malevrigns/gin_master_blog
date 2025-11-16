<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-4xl mx-auto">
      <!-- 加载状态 -->
      <div v-if="loading" class="text-center py-12">
        <el-icon class="is-loading" size="48"><Loading /></el-icon>
      </div>

      <!-- 文章内容 -->
      <article v-else-if="article" class="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden">
        <!-- 封面图 -->
        <div v-if="article.cover_image" class="h-64 md:h-96 overflow-hidden">
          <img
            :src="article.cover_image"
            :alt="article.title"
            class="w-full h-full object-cover"
          />
        </div>

        <!-- 文章头部 -->
        <div class="p-8">
          <div class="flex items-center space-x-2 mb-4">
            <el-tag v-if="article.is_top" type="danger" size="small">置顶</el-tag>
            <el-tag type="primary" size="small">{{ article.category?.name }}</el-tag>
            <span class="text-sm text-gray-500 dark:text-gray-400">
              {{ formatDate(article.published_at || article.created_at) }}
            </span>
          </div>

          <h1 class="text-4xl font-bold mb-4 text-gray-900 dark:text-gray-100">
            {{ article.title }}
          </h1>

          <div class="flex items-center space-x-4 mb-6 text-sm text-gray-500 dark:text-gray-400">
            <span class="flex items-center">
              <el-icon class="mr-1"><User /></el-icon>
              {{ article.author?.username }}
            </span>
            <span class="flex items-center">
              <el-icon class="mr-1"><View /></el-icon>
              {{ article.views }} 次浏览
            </span>
            <span class="flex items-center cursor-pointer" @click="handleLike">
              <el-icon class="mr-1"><Star /></el-icon>
              {{ article.likes }} 个赞
            </span>
          </div>

          <!-- 标签 -->
          <div class="flex flex-wrap gap-2 mb-6">
            <el-tag
              v-for="tag in article.tags"
              :key="tag.id"
              @click="$router.push(`/tag/${tag.slug}`)"
              class="cursor-pointer"
            >
              {{ tag.name }}
            </el-tag>
          </div>

          <!-- 文章内容 -->
          <div
            class="prose dark:prose-invert max-w-none mb-8"
            v-html="renderMarkdown(article.content)"
          ></div>
        </div>
      </article>

      <!-- 评论区域 -->
      <div class="mt-8 bg-white dark:bg-gray-800 rounded-lg shadow-md p-8">
        <h3 class="text-2xl font-bold mb-6 text-gray-900 dark:text-gray-100">评论</h3>

        <!-- 评论表单 -->
        <el-form :model="commentForm" label-width="80px" class="mb-8">
          <el-form-item label="姓名">
            <el-input v-model="commentForm.author" placeholder="请输入您的姓名" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="commentForm.email" placeholder="请输入您的邮箱" />
          </el-form-item>
          <el-form-item label="评论内容">
            <el-input
              v-model="commentForm.content"
              type="textarea"
              :rows="4"
              placeholder="请输入评论内容"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitComment">提交评论</el-button>
          </el-form-item>
        </el-form>

        <!-- 评论列表 -->
        <div class="space-y-6">
          <div
            v-for="comment in comments"
            :key="comment.id"
            class="border-b border-gray-200 dark:border-gray-700 pb-6 last:border-0"
          >
            <div class="flex items-start space-x-4">
              <el-avatar :size="40">{{ comment.author[0] }}</el-avatar>
              <div class="flex-1">
                <div class="flex items-center space-x-2 mb-2">
                  <span class="font-semibold text-gray-900 dark:text-gray-100">
                    {{ comment.author }}
                  </span>
                  <span class="text-sm text-gray-500 dark:text-gray-400">
                    {{ formatDate(comment.created_at) }}
                  </span>
                </div>
                <p class="text-gray-700 dark:text-gray-300 whitespace-pre-wrap">
                  {{ comment.content }}
                </p>
                <!-- 回复 -->
                <div
                  v-if="comment.replies && comment.replies.length > 0"
                  class="mt-4 ml-8 space-y-4"
                >
                  <div
                    v-for="reply in comment.replies"
                    :key="reply.id"
                    class="flex items-start space-x-3"
                  >
                    <el-avatar :size="32">{{ reply.author[0] }}</el-avatar>
                    <div class="flex-1">
                      <div class="flex items-center space-x-2 mb-1">
                        <span class="font-semibold text-sm text-gray-900 dark:text-gray-100">
                          {{ reply.author }}
                        </span>
                        <span class="text-xs text-gray-500 dark:text-gray-400">
                          {{ formatDate(reply.created_at) }}
                        </span>
                      </div>
                      <p class="text-sm text-gray-700 dark:text-gray-300 whitespace-pre-wrap">
                        {{ reply.content }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github-dark.css'
import api from '../api'
import { ElMessage } from 'element-plus'
import { User, View, Star, Loading } from '@element-plus/icons-vue'

marked.setOptions({
  breaks: true,
  highlight(code, lang) {
    if (lang && hljs.getLanguage(lang)) {
      return hljs.highlight(code, { language: lang }).value
    }
    return hljs.highlightAuto(code).value
  },
})

const route = useRoute()

const article = ref(null)
const comments = ref([])
const loading = ref(true)

const commentForm = ref({
  author: '',
  email: '',
  content: '',
})

const loadArticle = async () => {
  try {
    loading.value = true
    const response = await api.get(`/articles/${route.params.id}`)
    article.value = response.data
    await loadComments()
  } catch (error) {
    console.error('Failed to load article:', error)
    ElMessage.error('加载文章失败')
  } finally {
    loading.value = false
  }
}

const loadComments = async () => {
  try {
    const response = await api.get('/comments', {
      params: { article_id: route.params.id },
    })
    comments.value = response.data
  } catch (error) {
    console.error('Failed to load comments:', error)
  }
}

const submitComment = async () => {
  if (!commentForm.value.author || !commentForm.value.content) {
    ElMessage.warning('请填写姓名和评论内容')
    return
  }

  try {
    await api.post('/comments', {
      article_id: parseInt(route.params.id),
      ...commentForm.value,
    })
    ElMessage.success('评论提交成功，等待审核')
    commentForm.value = { author: '', email: '', content: '' }
    await loadComments()
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '提交评论失败')
  }
}

const handleLike = async () => {
  try {
    await api.post(`/articles/${route.params.id}/like`)
    article.value.likes++
    ElMessage.success('点赞成功')
  } catch (error) {
    ElMessage.error('点赞失败')
  }
}

const renderMarkdown = (content) => {
  if (!content) return ''
  return marked.parse(content)
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

onMounted(() => {
  loadArticle()
})
</script>

<style scoped>
.prose {
  @apply text-gray-700 dark:text-gray-300;
}

.prose h1,
.prose h2,
.prose h3,
.prose h4,
.prose h5,
.prose h6 {
  @apply font-bold text-gray-900 dark:text-gray-100 mt-6 mb-4;
}

.prose p {
  @apply mb-4 leading-7;
}

.prose code {
  @apply bg-gray-100 dark:bg-gray-700 px-1 py-0.5 rounded text-sm;
}

.prose pre {
  @apply bg-gray-100 dark:bg-gray-700 p-4 rounded overflow-x-auto mb-4;
}

.prose pre code {
  @apply bg-transparent p-0;
}

.prose img {
  @apply rounded-lg my-4;
}

.prose a {
  @apply text-primary-600 dark:text-primary-400 hover:underline;
}

.prose blockquote {
  @apply border-l-4 border-primary-500 pl-4 italic my-4;
}

.prose ul,
.prose ol {
  @apply ml-6 mb-4;
}

.prose li {
  @apply mb-2;
}
</style>

