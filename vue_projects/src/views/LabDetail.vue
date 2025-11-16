<template>
  <div class="container mx-auto px-4 py-10">
    <div v-if="loading" class="text-center py-20 text-slate-400">
      <el-icon class="is-loading text-4xl mb-4"><Loading /></el-icon>
      正在载入模块...
    </div>

    <div v-else-if="lab" class="space-y-8">
      <div
        class="rounded-3xl overflow-hidden relative min-h-[260px] flex flex-col justify-end p-8 text-white shadow-2xl"
        :style="{
          backgroundImage: lab.hero_image ? `linear-gradient(135deg, rgba(0,0,0,0.65), rgba(2,6,23,0.85)), url(${lab.hero_image})` : undefined,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
        }"
      >
        <div class="mb-4 inline-flex items-center text-xs uppercase tracking-[0.4em] px-3 py-1 rounded-full bg-white/20">
          <span :style="{ color: lab.badge_color || '#34d399' }">
            {{ lab.badge }}
          </span>
        </div>
        <h1 class="text-4xl font-bold mb-2 drop-shadow-lg">{{ lab.title }}</h1>
        <p class="text-lg text-slate-100/90 max-w-2xl">{{ lab.subtitle }}</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div class="lg:col-span-2 space-y-6">
          <div class="glass-card rounded-2xl p-8 border border-emerald-500/30">
            <p class="text-sm uppercase tracking-[0.3em] text-emerald-300 mb-4">
              模块简介
            </p>
            <p class="text-lg text-slate-200 leading-relaxed">{{ lab.description }}</p>

            <div class="mt-8">
              <p class="text-sm text-emerald-200 mb-3">{{ lab.focus }}</p>
              <div v-if="lab.highlights?.length" class="flex flex-wrap gap-3">
                <button
                  v-for="highlight in lab.highlights"
                  :key="highlight.title"
                  class="topic-chip"
                  :class="{ 'topic-chip--active': activeHighlight?.title === highlight.title }"
                  @click="handleHighlightClick(highlight)"
                >
                  <span class="block font-semibold">{{ highlight.title }}</span>
                  <small class="block text-[11px] text-slate-300/70">{{ highlight.description }}</small>
                </button>
              </div>
              <p v-else class="text-slate-400 text-sm">尚未配置研究方向。</p>
            </div>
          </div>

          <div class="glass-card rounded-2xl p-8 border border-slate-700/60" v-if="lab.content">
            <div class="prose prose-invert max-w-none" v-html="renderMarkdown(lab.content)"></div>
          </div>

          <div
            v-if="activeHighlight"
            class="glass-card rounded-2xl p-6 border border-slate-700/60"
          >
            <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-4 gap-3">
              <div>
                <p class="text-xs tracking-[0.4em] uppercase text-emerald-200">相关内容</p>
                <h3 class="text-xl font-semibold text-white">
                  {{ activeHighlight.title }}
                </h3>
              </div>
              <div class="flex items-center gap-2">
                <el-button
                  v-if="activeHighlight.tag"
                  size="small"
                  text
                  @click="goToTagPage(activeHighlight)"
                >
                  更多相关文章
                </el-button>
              </div>
            </div>
            <div v-if="articlesLoading" class="py-8 text-center text-slate-400">
              <el-icon class="is-loading text-2xl"></el-icon>
              <p class="mt-2">正在加载相关文章...</p>
            </div>
            <div v-else-if="relatedArticles.length" class="grid gap-4 md:grid-cols-2">
              <article
                v-for="article in relatedArticles"
                :key="article.id"
                class="related-card"
                @click="$router.push(`/article/${article.id}`)"
              >
                <div
                  class="related-card__cover"
                  :class="{ 'related-card__cover--placeholder': article.coverFallback || !article.coverUrl }"
                  :style="article.coverFallback || !article.coverUrl ? { backgroundImage: placeholderBackground(article.coverSeed) } : {}"
                >
                  <img
                    v-if="article.coverUrl && !article.coverFallback"
                    :src="article.coverUrl"
                    :alt="article.title"
                    @error.stop.prevent="article.coverFallback = true"
                  />
                  <div v-else class="related-card__cover-text">
                    {{ coverPlaceholderText(article.title) }}
                  </div>
                </div>
                <div class="p-4 space-y-2">
                  <p class="text-xs uppercase tracking-[0.3em] text-slate-400">
                    {{ formatDate(article.published_at || article.created_at) }}
                  </p>
                  <h4 class="text-lg font-semibold text-slate-50 line-clamp-2">
                    {{ article.title }}
                  </h4>
                  <p class="text-sm text-slate-400 line-clamp-2">
                    {{ article.excerpt || article.content?.slice(0, 80) + '...' }}
                  </p>
                </div>
              </article>
            </div>
            <div v-else class="py-8 text-center text-slate-400">
              暂无相关的文章，可以稍后再试。
            </div>
          </div>
        </div>

        <aside class="space-y-6">
          <div class="glass-card rounded-2xl p-6 border border-emerald-500/20">
            <h3 class="text-lg font-semibold text-slate-100 mb-4 flex items-center">
              <el-icon class="mr-2 text-emerald-300"><TrendCharts /></el-icon>
              推荐资源
            </h3>
            <div v-if="lab.resources?.length" class="space-y-4">
              <a
                v-for="resource in lab.resources"
                :key="resource.title"
                :href="resource.url"
                target="_blank"
                rel="noopener"
                class="block rounded-xl p-4 bg-slate-900/60 hover:bg-slate-800/80 transition-colors border border-slate-700/60"
              >
                <div class="flex items-center mb-2">
                  <el-icon class="text-emerald-300 mr-2">
                    <component :is="resolveIcon(resource.icon)" />
                  </el-icon>
                  <p class="text-slate-100 font-semibold">{{ resource.title }}</p>
                </div>
                <p class="text-sm text-slate-400">{{ resource.desc }}</p>
              </a>
            </div>
            <p v-else class="text-slate-400 text-sm">暂无整理好的资源。</p>
          </div>
        </aside>
      </div>
    </div>

    <div v-else class="text-center py-20 text-slate-400">
      暂未找到该模块。
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github-dark.css'
import api from '../api'
import { ElMessage } from 'element-plus'
import {
  TrendCharts,
  Document,
  Monitor,
  Operation,
  StarFilled,
  Headset,
  Loading,
} from '@element-plus/icons-vue'

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
const router = useRouter()
const lab = ref(null)
const loading = ref(true)
const activeHighlight = ref(null)
const relatedArticles = ref([])
const articlesLoading = ref(false)

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
  (list || []).map((item, idx) => ({
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
  return title.slice(0, 4).toUpperCase()
}

const fetchLab = async () => {
  try {
    loading.value = true
    const response = await api.get(`/labs/${route.params.slug}`)
    lab.value = response.data
    relatedArticles.value = []
    if (lab.value.highlights?.length) {
      fetchRelatedArticles(lab.value.highlights[0])
    } else {
      activeHighlight.value = null
    }
  } catch (error) {
    lab.value = null
    activeHighlight.value = null
    relatedArticles.value = []
    ElMessage.error(error.response?.data?.error || '加载模块失败')
  } finally {
    loading.value = false
  }
}

const fetchRelatedArticles = async (highlight) => {
  if (!highlight?.tag) {
    if (highlight?.link) {
      window.open(highlight.link, '_blank')
    }
    return
  }
  activeHighlight.value = highlight
  articlesLoading.value = true
  relatedArticles.value = []
  try {
    const response = await api.get(`/labs/${route.params.slug}/articles`, {
      params: { tag: highlight.tag },
    })
    relatedArticles.value = decorateArticles(response.data.articles)
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '加载相关文章失败')
  } finally {
    articlesLoading.value = false
  }
}

const handleHighlightClick = (highlight) => {
  fetchRelatedArticles(highlight)
}

const goToTagPage = (highlight) => {
  if (!highlight?.tag) return
  router.push(`/tag/${highlight.tag}`)
}

const renderMarkdown = (content) => {
  if (!content) return ''
  return marked.parse(content)
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

const iconMap = {
  TrendCharts,
  Document,
  Monitor,
  Operation,
  StarFilled,
  Headset,
}

const resolveIcon = (name) => iconMap[name] || Document

onMounted(fetchLab)

watch(
  () => route.params.slug,
  () => {
    fetchLab()
  }
)
</script>

<style scoped>
.prose :global(h2) {
  @apply text-2xl font-bold text-white mt-6 mb-4;
}
.prose :global(p) {
  @apply mb-4 leading-relaxed text-slate-200;
}

.topic-chip {
  @apply px-3 py-2 rounded-2xl bg-slate-900/60 border border-white/10 text-left text-slate-100 text-sm transition-all duration-200;
}

.topic-chip:hover {
  border-color: rgba(60, 242, 255, 0.6);
  transform: translateY(-1px);
}

.topic-chip--active {
  border-color: rgba(60, 242, 255, 0.9);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.35);
}

.related-card {
  @apply rounded-2xl overflow-hidden border border-white/10 bg-slate-900/70 cursor-pointer transition-transform duration-300;
}

.related-card:hover {
  transform: translateY(-4px);
  border-color: rgba(60, 242, 255, 0.5);
}

.related-card__cover {
  position: relative;
  height: 150px;
  overflow: hidden;
}

.related-card__cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.related-card__cover--placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
}

.related-card__cover-text {
  font-weight: 700;
  letter-spacing: 0.3em;
  color: rgba(0, 0, 0, 0.45);
}
</style>
