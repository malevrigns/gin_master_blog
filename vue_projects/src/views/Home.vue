<template>
    <div class="container mx-auto px-4 py-10">
    <!-- Hero Banner -->
    <div class="mb-10 relative h-64 md:h-96 rounded-2xl overflow-hidden group glass-card card-hover border border-slate-700/60">
      <div
        class="absolute inset-0 bg-gradient-to-r from-sky-500 via-fuchsia-500 to-emerald-400 flex items-center justify-center transition-all duration-500 group-hover:scale-105"
        :style="{
          backgroundImage: heroImage ? `url(${getImageUrl(heroImage)})` : 'none',
          backgroundSize: 'cover',
          backgroundPosition: 'center',
        }"
      >
        <!-- 闪烁效果 -->
        <SparkleEffect :count="15" />
        
        <div class="text-center text-white z-10 relative animate-fade-in">
          <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-lg animate-slide-up gradient-text">
            <TypewriterText :text="heroTitle" :speed="100" />
          </h1>
          <p class="text-xl md:text-2xl drop-shadow-md animate-slide-up-delay">
            {{ heroSubtitle }}
          </p>
        </div>
        <div class="absolute inset-0 bg-black/30 group-hover:bg-black/20 transition-colors"></div>
        
        <!-- 装饰性元素 -->
        <div class="absolute top-4 right-4 w-20 h-20 border-2 border-white/30 rounded-full animate-pulse"></div>
        <div class="absolute bottom-4 left-4 w-16 h-16 border-2 border-white/20 rounded-full animate-pulse delay-300"></div>
      </div>
    </div>

    <!-- 实验室模块 -->
    <section class="mb-10 grid gap-4 md:grid-cols-3">
      <div
        v-for="lab in labModules"
        :key="lab.slug || lab.id"
        class="glass-card card-hover rounded-2xl p-5 border border-slate-800/60 cursor-pointer group"
        @click="$router.push(`/labs/${lab.slug}`)"
      >
        <div class="flex items-center justify-between mb-3">
          <h2 class="text-sm font-semibold text-slate-100">
            {{ lab.title }}
          </h2>
          <span
            class="text-[11px] px-2 py-0.5 rounded-full uppercase tracking-wide border"
            :style="{ color: lab.badge_color, borderColor: lab.badge_color || '#38bdf8' }"
          >
            {{ lab.badge }}
          </span>
        </div>
        <p class="text-xs text-slate-300/80 leading-relaxed mb-3">
          {{ lab.subtitle }}
        </p>
        <p class="text-[11px] text-slate-400 flex items-center">
          <span
            class="inline-block w-1.5 h-1.5 rounded-full mr-2"
            :style="{ backgroundColor: lab.badge_color || '#22d3ee' }"
          ></span>
          {{ lab.focus }}
        </p>
      </div>
    </section>

    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
      <!-- 侧边栏 -->
      <aside class="lg:col-span-1 space-y-6 order-2 lg:order-1">
        <!-- 个人资料卡片 -->
        <ProfileCard />

        <!-- 公告栏 -->
        <AnnouncementCard />

        <!-- 统计信息 -->
        <div class="glass-card rounded-2xl p-6 hover:shadow-xl transition-all duration-300 border border-slate-700/60">
          <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-gray-100 flex items-center">
            <el-icon class="mr-2"><DataAnalysis /></el-icon>
            统计
          </h3>
          <div class="space-y-4">
            <div class="flex items-center justify-between p-3 rounded-lg bg-gray-50 dark:bg-gray-700/50 hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-colors group">
              <span class="text-gray-600 dark:text-gray-400 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                文章总数
              </span>
              <span class="text-primary-600 dark:text-primary-400 font-bold text-lg">
                <AnimatedNumber :value="stats.totalArticles" />
              </span>
            </div>
            <div class="flex items-center justify-between p-3 rounded-lg bg-gray-50 dark:bg-gray-700/50 hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-colors group">
              <span class="text-gray-600 dark:text-gray-400 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                分类数量
              </span>
              <span class="text-primary-600 dark:text-primary-400 font-bold text-lg">
                <AnimatedNumber :value="stats.totalCategories" />
              </span>
            </div>
            <div class="flex items-center justify-between p-3 rounded-lg bg-gray-50 dark:bg-gray-700/50 hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-colors group">
              <span class="text-gray-600 dark:text-gray-400 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                标签数量
              </span>
              <span class="text-primary-600 dark:text-primary-400 font-bold text-lg">
                <AnimatedNumber :value="stats.totalTags" />
              </span>
            </div>
          </div>
        </div>

        <!-- 分类 -->
        <div class="glass-card rounded-2xl p-6 border border-slate-700/60">
          <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-gray-100 flex items-center">
            <el-icon class="mr-2"><Folder /></el-icon>
            分类
          </h3>
          <div class="space-y-2">
            <div
              v-for="category in categories"
              :key="category.id"
              @click="$router.push(`/category/${category.id}`)"
              class="flex items-center justify-between p-2 rounded hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer transition-colors group"
            >
              <span class="text-gray-700 dark:text-gray-300 group-hover:text-primary-600 dark:group-hover:text-primary-400">
                {{ category.name }}
              </span>
              <el-badge :value="category.articles?.length || 0" class="item" />
            </div>
          </div>
        </div>

        <!-- 标签云 -->
        <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
          <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-gray-100 flex items-center">
            <el-icon class="mr-2"><PriceTag /></el-icon>
            标签
          </h3>
          <div class="flex flex-wrap gap-2">
            <el-tag
              v-for="tag in tags"
              :key="tag.id"
              @click="$router.push(`/tag/${tag.slug}`)"
              class="cursor-pointer hover:scale-105 transition-transform"
              effect="plain"
            >
              #{{ tag.name }}
            </el-tag>
          </div>
        </div>

        <!-- 热门文章 -->
        <div class="glass-card rounded-2xl p-6 border border-slate-700/60">
          <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-gray-100 flex items-center">
            <el-icon class="mr-2"><TrendCharts /></el-icon>
            热门文章
          </h3>
          <div class="space-y-3">
            <div
              v-for="(article, index) in hotArticles"
              :key="article.id"
              @click="$router.push(`/article/${article.id}`)"
              class="cursor-pointer group"
            >
              <div class="flex items-start space-x-2">
                <span class="text-primary-600 dark:text-primary-400 font-bold text-sm w-6">
                  {{ index + 1 }}
                </span>
                <div class="flex-1 min-w-0">
                  <p class="text-sm truncate group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                    {{ article.title }}
                  </p>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                    <el-icon class="text-xs"><View /></el-icon>
                    {{ article.views }} 次浏览
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </aside>

      <!-- 主内容区 -->
      <div class="lg:col-span-3 order-1 lg:order-2 space-y-8">
        <!-- 加载状态 -->
        <div v-if="loading" class="space-y-6">
          <SkeletonCard v-for="i in 3" :key="i" />
        </div>

        <!-- 文章列表 -->
        <div v-else class="space-y-6">
          <article
            v-for="(article, index) in articles"
            :key="article.id"
            class="glass-card rounded-2xl overflow-hidden transition-all duration-300 cursor-pointer group border border-slate-700/60 hover:border-sky-400/60 animate-fade-in"
            :style="{ animationDelay: `${index * 0.1}s` }"
            @click="$router.push(`/article/${article.id}`)"
          >
            <div class="flex flex-col md:flex-row">
              <!-- 封面图 -->
              <div
                class="md:w-64 h-48 md:h-auto overflow-hidden flex-shrink-0 article-cover"
                :class="{ 'article-cover--fallback': article.coverFallback || !article.coverUrl }"
                :style="article.coverFallback || !article.coverUrl ? { backgroundImage: placeholderBackground(article.coverSeed) } : {}"
              >
                <img
                  v-if="article.coverUrl && !article.coverFallback"
                  :src="article.coverUrl"
                  :alt="article.title"
                  class="w-full h-full object-cover transition-transform duration-300"
                  @error.stop.prevent="handleCoverError(article)"
                />
                <div
                  v-if="article.coverFallback || !article.coverUrl"
                  class="article-cover__placeholder"
                >
                  <span>{{ coverPlaceholderText(article.title) }}</span>
                </div>
              </div>

              <!-- 内容 -->
              <div class="flex-1 p-6">
                <div class="flex items-center space-x-2 mb-3 flex-wrap gap-2">
                  <el-tag
                    v-if="article.is_top"
                    type="danger"
                    size="small"
                    effect="dark"
                  >
                    <el-icon class="mr-1"><Top /></el-icon>
                    置顶
                  </el-tag>
                  <el-tag type="primary" size="small" effect="plain">
                    <el-icon class="mr-1"><Folder /></el-icon>
                    {{ article.category?.name }}
                  </el-tag>
                  <span class="text-sm text-gray-500 dark:text-gray-400 flex items-center">
                    <el-icon class="mr-1 text-xs"><Calendar /></el-icon>
                    {{ formatDate(article.published_at || article.created_at) }}
                  </span>
                </div>

                <h2 class="text-2xl font-bold mb-3 text-gray-900 dark:text-gray-100 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                  {{ article.title }}
                </h2>

                <p class="text-gray-600 dark:text-gray-400 mb-4 line-clamp-2 leading-relaxed">
                  {{ article.excerpt || article.content.substring(0, 150) + '...' }}
                </p>

                <div class="flex items-center justify-between flex-wrap gap-3">
                  <div class="flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400">
                    <span class="flex items-center hover:text-primary-600 dark:hover:text-primary-400 transition-colors">
                      <el-icon class="mr-1"><User /></el-icon>
                      {{ article.author?.username }}
                    </span>
                    <span class="flex items-center">
                      <el-icon class="mr-1"><View /></el-icon>
                      {{ article.views }}
                    </span>
                    <span class="flex items-center">
                      <el-icon class="mr-1"><Star /></el-icon>
                      {{ article.likes }}
                    </span>
                  </div>
                  <div class="flex flex-wrap gap-2">
                    <el-tag
                      v-for="tag in article.tags"
                      :key="tag.id"
                      size="small"
                      effect="plain"
                      class="hover:bg-primary-100 dark:hover:bg-primary-900 transition-colors"
                    >
                      #{{ tag.name }}
                    </el-tag>
                  </div>
                </div>
              </div>
            </div>
          </article>
        </div>

        <!-- 分页 -->
        <div class="mt-8 flex justify-center">
          <el-pagination
            v-model:current-page="pagination.page"
            :page-size="pagination.pageSize"
            :total="pagination.total"
            layout="prev, pager, next"
            @current-change="loadArticles"
          />
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import api from '../api'
import {
  User,
  View,
  Star,
  Folder,
  PriceTag,
  TrendCharts,
  Calendar,
  Top,
  DataAnalysis,
} from '@element-plus/icons-vue'
import ProfileCard from '../components/common/ProfileCard.vue'
import AnnouncementCard from '../components/common/AnnouncementCard.vue'
import SkeletonCard from '../components/common/SkeletonCard.vue'
import LoadingSpinner from '../components/common/LoadingSpinner.vue'
import TypewriterText from '../components/common/TypewriterText.vue'
import AnimatedNumber from '../components/common/AnimatedNumber.vue'
import SparkleEffect from '../components/common/SparkleEffect.vue'

const route = useRoute()

const articles = ref([])
const categories = ref([])
const tags = ref([])
const hotArticles = ref([])
const loading = ref(true)

import { useSettingsStore } from '../stores/settings'

const settingsStore = useSettingsStore()

const heroTitle = ref('欢迎来到我的博客')
const heroSubtitle = ref('分享技术、生活与思考')
const heroImage = computed(() => settingsStore.heroBackground)

const stats = ref({
  totalArticles: 0,
  totalCategories: 0,
  totalTags: 0,
})

const labModules = ref([])

const placeholderPalette = [
  ['#ff61e6', '#ffc371'],
  ['#3cf2ff', '#7c5dff'],
  ['#8ef9a5', '#1cb5e0'],
  ['#ffb347', '#ffcc33'],
  ['#ff4d79', '#7f53ac'],
]

const normalizeMediaUrl = (url) => {
  if (!url) return ''
  const trimmed = url.trim()
  if (!trimmed) return ''
  if (trimmed.startsWith('http')) return trimmed
  if (trimmed.startsWith('/')) return trimmed
  return '/' + trimmed
}

const decorateArticles = (list) =>
  list.map((item, idx) => ({
    ...item,
    coverUrl: normalizeMediaUrl(item.cover_image),
    coverFallback: !item.cover_image,
    coverSeed: idx,
  }))

const placeholderBackground = (seed = 0) => {
  const palette = placeholderPalette[seed % placeholderPalette.length]
  return `linear-gradient(135deg, ${palette[0]}, ${palette[1]})`
}

const coverPlaceholderText = (title) => {
  if (!title) return 'ARTICLE'
  return title.slice(0, 6).toUpperCase()
}

const handleCoverError = (article) => {
  article.coverFallback = true
}

const pagination = ref({
  page: 1,
  pageSize: 10,
  total: 0,
})

const loadArticles = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.value.page,
      page_size: pagination.value.pageSize,
      status: 'published',
    }

    if (route.query.search) {
      params.search = route.query.search
    }

    const response = await api.get('/articles', { params })
    articles.value = decorateArticles(response.data.articles || [])
    pagination.value.total = response.data.total
  } catch (error) {
    console.error('Failed to load articles:', error)
  } finally {
    loading.value = false
  }
}

const loadCategories = async () => {
  try {
    const response = await api.get('/categories')
    categories.value = response.data
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
}

const loadLabs = async () => {
  try {
    const response = await api.get('/labs')
    labModules.value = response.data
  } catch (error) {
    console.error('Failed to load labs:', error)
  }
}

const loadTags = async () => {
  try {
    const response = await api.get('/tags')
    tags.value = response.data
  } catch (error) {
    console.error('Failed to load tags:', error)
  }
}

const loadHotArticles = async () => {
  try {
    const response = await api.get('/articles', {
      params: {
        status: 'published',
        page: 1,
        page_size: 5,
      },
    })
    hotArticles.value = response.data.articles.sort((a, b) => b.views - a.views).slice(0, 5)
  } catch (error) {
    console.error('Failed to load hot articles:', error)
  }
}

const loadStats = async () => {
  try {
    const [articlesRes, categoriesRes, tagsRes] = await Promise.all([
      api.get('/articles', { params: { status: 'published', page_size: 1 } }),
      api.get('/categories'),
      api.get('/tags'),
    ])
    stats.value = {
      totalArticles: articlesRes.data.total || 0,
      totalCategories: categoriesRes.data.length || 0,
      totalTags: tagsRes.data.length || 0,
    }
  } catch (error) {
    console.error('Failed to load stats:', error)
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

const getImageUrl = (url) => normalizeMediaUrl(url)

onMounted(() => {
  loadArticles()
  loadCategories()
  loadTags()
  loadHotArticles()
  loadStats()
  loadLabs()
})

watch(() => route.query, () => {
  pagination.value.page = 1
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

.article-cover {
  position: relative;
  border-radius: 1.25rem 1.25rem 0 0;
  background: linear-gradient(135deg, rgba(255, 97, 230, 0.35), rgba(60, 242, 255, 0.35));
}

.article-cover img {
  width: 100%;
  height: 100%;
  display: block;
  transition: transform 0.45s ease;
}

.article-cover__placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: rgba(5, 4, 20, 0.85);
  font-weight: 700;
  letter-spacing: 0.2em;
  font-size: 0.9rem;
  text-transform: uppercase;
}

.article-cover--fallback {
  border: none;
}
</style>

