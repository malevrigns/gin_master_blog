<template>
  <div class="min-h-screen bg-slate-950 text-slate-50">
    <div
      class="sticky top-0 z-30 bg-slate-950/90 border-b border-emerald-400/30 backdrop-blur px-6 py-4 flex flex-wrap items-center justify-between gap-4"
    >
      <div>
        <div class="flex items-center gap-3">
          <el-button text class="text-emerald-300" @click="handleBack">
            <el-icon class="mr-2">
              <ArrowLeftBold />
            </el-icon>
            返回后台
          </el-button>
          <span class="text-sm text-slate-400">沉浸式写作 · 专注内容创作</span>
        </div>
        <h1 class="text-2xl font-semibold mt-2">写文章</h1>
      </div>
      <div class="flex items-center gap-4">
        <span class="text-sm text-slate-400">字数：{{ wordCount }}</span>
        <el-button
          :loading="isSaving && pendingStatus === 'draft'"
          @click="handleSaveDraft"
        >
          保存草稿
        </el-button>
        <el-button
          type="primary"
          class="bg-emerald-500 border-none"
          :loading="isSaving && pendingStatus === 'published'"
          @click="handlePublish"
        >
          发布文章
        </el-button>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-6 py-10 grid gap-6 lg:grid-cols-[minmax(0,1fr)_320px]">
      <section
        class="rounded-3xl bg-slate-900/70 border border-slate-800 shadow-2xl shadow-emerald-900/20 px-6 py-6 space-y-6"
      >
        <el-form label-position="top" :model="articleForm">
          <el-form-item label="标题">
            <el-input v-model="articleForm.title" placeholder="请输入文章标题" size="large" />
          </el-form-item>
          <el-form-item label="摘要">
            <el-input
              v-model="articleForm.excerpt"
              type="textarea"
              :rows="3"
              placeholder="一句话概括文章，方便在列表和分享中展示"
            />
          </el-form-item>
          <el-form-item label="封面图">
            <ImageUploader v-model="articleForm.cover_image" class="max-w-md" />
          </el-form-item>
          <div>
            <div class="flex items-center justify-between mb-2">
              <label class="text-sm text-slate-300">正文内容</label>
              <span class="text-xs text-slate-500">支持 Markdown / Emoji / 代码块</span>
            </div>
            <MarkdownEditor v-model="articleForm.content" :height="640" />
          </div>
        </el-form>
      </section>
      <aside class="space-y-6">
        <div class="rounded-3xl bg-slate-900/70 border border-slate-800 p-5 space-y-4">
          <h2 class="text-lg font-semibold">发布设置</h2>
          <el-form label-position="top" :model="articleForm" class="space-y-3">
            <el-form-item label="文章状态">
              <el-select v-model="articleForm.status">
                <el-option label="草稿" value="draft" />
                <el-option label="公开" value="published" />
              </el-select>
            </el-form-item>
            <el-form-item label="所属分类">
              <el-select v-model="articleForm.category_id" placeholder="选择分类">
                <el-option
                  v-for="cat in categories"
                  :key="cat.id"
                  :label="cat.name"
                  :value="cat.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="标签">
              <el-select
                v-model="articleForm.tag_ids"
                multiple
                filterable
                collapse-tags
                placeholder="选择多个标签"
              >
                <el-option
                  v-for="tag in tags"
                  :key="tag.id"
                  :label="tag.name"
                  :value="tag.id"
                />
              </el-select>
            </el-form-item>
          </el-form>
        </div>

        <div class="rounded-3xl bg-slate-900/70 border border-slate-800 p-5 space-y-3 text-sm">
          <h2 class="text-lg font-semibold">写作建议</h2>
          <ul class="space-y-2 text-slate-400">
            <li>• 标题不超过 50 字，突出主题</li>
            <li>• 摘要用 1-2 句话说明价值</li>
            <li>• 合理使用小标题、列表和代码块</li>
            <li>• 适当插入图片，增强阅读体验</li>
          </ul>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeftBold } from '@element-plus/icons-vue'
import api from '../api'
import MarkdownEditor from '../components/articles/MarkdownEditor.vue'
import ImageUploader from '../components/common/ImageUploader.vue'

const router = useRouter()

const categories = ref([])
const tags = ref([])
const isSaving = ref(false)
const pendingStatus = ref(null)

const defaultArticleForm = () => ({
  id: null,
  title: '',
  excerpt: '',
  cover_image: '',
  content: '',
  category_id: null,
  status: 'draft',
  tag_ids: [],
})

const articleForm = ref(defaultArticleForm())

const wordCount = computed(() => {
  const raw = articleForm.value.content || ''
  return raw.replace(/\s+/g, '').length
})

const loadCategories = async () => {
  try {
    const response = await api.get('/categories')
    categories.value = response.data
  } catch (error) {
    console.error('Failed to load categories', error)
  }
}

const loadTags = async () => {
  try {
    const response = await api.get('/tags')
    tags.value = response.data
  } catch (error) {
    console.error('Failed to load tags', error)
  }
}

const preparePayload = (status) => ({
  title: articleForm.value.title,
  excerpt: articleForm.value.excerpt,
  cover_image: articleForm.value.cover_image,
  content: articleForm.value.content,
  category_id: articleForm.value.category_id,
  status: status || articleForm.value.status,
  tag_ids: articleForm.value.tag_ids,
})

const saveArticle = async (targetStatus) => {
  try {
    isSaving.value = true
    pendingStatus.value = targetStatus
    const payload = preparePayload(targetStatus)

    if (articleForm.value.id) {
      await api.put(`/articles/${articleForm.value.id}`, payload)
      ElMessage.success(targetStatus === 'published' ? '文章已更新' : '草稿已更新')
    } else {
      await api.post('/articles', payload)
      ElMessage.success(targetStatus === 'published' ? '发布成功' : '草稿保存成功')
      if (targetStatus === 'draft') {
        articleForm.value = { ...articleForm.value, status: 'draft' }
      } else {
        router.push({ name: 'Admin', query: { tab: 'articles' } })
      }
    }
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '保存失败，请检查必填信息')
  } finally {
    isSaving.value = false
    pendingStatus.value = null
  }
}

const handleSaveDraft = () => {
  saveArticle('draft')
}

const handlePublish = () => {
  saveArticle('published')
}

const handleBack = () => {
  router.push({ name: 'Admin', query: { tab: 'articles' } })
}

onMounted(() => {
  loadCategories()
  loadTags()
})
</script>
