<template>
  <div class="container mx-auto px-4 py-8">
    <div class="max-w-7xl mx-auto">
      <h1 class="text-3xl font-bold mb-8 text-gray-900 dark:text-gray-100">管理后台</h1>

      <el-tabs v-model="activeTab">
        <!-- 文章管理 -->
        <el-tab-pane label="文章管理" name="articles">
          <div class="mb-4 flex justify-between">
            <el-button type="primary" @click="goToArticleEditor">新建文章</el-button>
          </div>
          <el-table :data="articles" style="width: 100%">
            <el-table-column prop="title" label="标题" />
            <el-table-column prop="category.name" label="分类" />
            <el-table-column prop="status" label="状态" />
            <el-table-column prop="views" label="浏览数" />
            <el-table-column label="操作">
              <template #default="{ row }">
                <el-button size="small" @click="editArticle(row)">编辑</el-button>
                <el-button size="small" type="danger" @click="deleteArticle(row.id)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>

        <!-- 音乐管理 -->
        <el-tab-pane label="音乐管理" name="music">
          <div class="mb-4 flex justify-between">
            <el-button type="primary" @click="showMusicDialog = true">添加音乐</el-button>
          </div>
          <el-table :data="musics" style="width: 100%">
            <el-table-column prop="title" label="标题" />
            <el-table-column prop="artist" label="艺术家" />
            <el-table-column prop="play_count" label="播放次数" />
            <el-table-column label="操作">
              <template #default="{ row }">
                <el-button size="small" @click="editMusic(row)">编辑</el-button>
                <el-button size="small" type="danger" @click="deleteMusic(row.id)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>

        <!-- 评论管理 -->
        <el-tab-pane label="评论管理" name="comments">
          <el-table :data="pendingComments" style="width: 100%">
            <el-table-column prop="author" label="作者" />
            <el-table-column prop="content" label="内容" />
            <el-table-column prop="article.title" label="文章" />
            <el-table-column label="操作">
              <template #default="{ row }">
                <el-button size="small" type="success" @click="approveComment(row.id)">通过</el-button>
                <el-button size="small" type="danger" @click="rejectComment(row.id)">拒绝</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>

      <!-- 文章对话框 -->
      <el-dialog v-model="showArticleDialog" title="文章编辑" width="960px">
        <el-form :model="articleForm" label-width="90px" class="space-y-2">
          <el-form-item label="标题">
            <el-input v-model="articleForm.title" />
          </el-form-item>
          <el-form-item label="摘要">
            <el-input v-model="articleForm.excerpt" type="textarea" :rows="3" />
          </el-form-item>
          <el-form-item label="封面">
            <ImageUploader v-model="articleForm.cover_image" />
          </el-form-item>
          <el-form-item label="内容">
            <MarkdownEditor v-model="articleForm.content" />
          </el-form-item>
          <el-form-item label="分类">
            <el-select v-model="articleForm.category_id">
              <el-option
                v-for="cat in categories"
                :key="cat.id"
                :label="cat.name"
                :value="cat.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="标签">
            <el-select v-model="articleForm.tag_ids" multiple placeholder="选择标签">
              <el-option
                v-for="tag in tags"
                :key="tag.id"
                :label="tag.name"
                :value="tag.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="articleForm.status">
              <el-option label="草稿" value="draft" />
              <el-option label="已发布" value="published" />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="showArticleDialog = false">取消</el-button>
          <el-button type="primary" :loading="isSavingArticle" @click="saveArticle">保存</el-button>
        </template>
      </el-dialog>

      <!-- 音乐对话框 -->
      <el-dialog v-model="showMusicDialog" title="音乐编辑" width="600px">
        <el-form :model="musicForm" label-width="80px">
          <el-form-item label="标题">
            <el-input v-model="musicForm.title" />
          </el-form-item>
          <el-form-item label="艺术家">
            <el-input v-model="musicForm.artist" />
          </el-form-item>
          <el-form-item label="URL">
            <el-input v-model="musicForm.url" />
          </el-form-item>
          <el-form-item label="封面">
            <el-input v-model="musicForm.cover" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="showMusicDialog = false">取消</el-button>
          <el-button type="primary" @click="saveMusic">保存</el-button>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '../api'
import { ElMessage, ElMessageBox } from 'element-plus'
import MarkdownEditor from '../components/articles/MarkdownEditor.vue'
import ImageUploader from '../components/common/ImageUploader.vue'

const route = useRoute()
const router = useRouter()

const activeTab = ref('articles')

const articles = ref([])
const musics = ref([])
const pendingComments = ref([])
const categories = ref([])
const tags = ref([])

const showArticleDialog = ref(false)
const showMusicDialog = ref(false)
const isSavingArticle = ref(false)

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

const musicForm = ref({
  title: '',
  artist: '',
  url: '',
  cover: '',
})

const loadArticles = async () => {
  try {
    const response = await api.get('/articles', { params: { page_size: 100 } })
    articles.value = response.data.articles
  } catch (error) {
    console.error('Failed to load articles:', error)
  }
}

const loadMusics = async () => {
  try {
    const response = await api.get('/music', { params: { page_size: 100 } })
    musics.value = response.data.musics
  } catch (error) {
    console.error('Failed to load musics:', error)
  }
}

const loadPendingComments = async () => {
  try {
    const response = await api.get('/comments/pending')
    pendingComments.value = response.data.comments
  } catch (error) {
    console.error('Failed to load comments:', error)
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

const loadTags = async () => {
  try {
    const response = await api.get('/tags')
    tags.value = response.data
  } catch (error) {
    console.error('Failed to load tags:', error)
  }
}

const resetArticleForm = () => {
  articleForm.value = defaultArticleForm()
}

const prepareArticlePayload = () => {
  const payload = {
    title: articleForm.value.title,
    excerpt: articleForm.value.excerpt,
    cover_image: articleForm.value.cover_image,
    content: articleForm.value.content,
    category_id: articleForm.value.category_id,
    status: articleForm.value.status,
    tag_ids: articleForm.value.tag_ids,
  }
  return payload
}

const saveArticle = async () => {
  try {
    isSavingArticle.value = true
    const payload = prepareArticlePayload()

    if (articleForm.value.id) {
      await api.put(`/articles/${articleForm.value.id}`, payload)
      ElMessage.success('文章更新成功')
    } else {
      await api.post('/articles', payload)
      ElMessage.success('文章创建成功')
    }
    showArticleDialog.value = false
    resetArticleForm()
    await loadArticles()
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '保存失败')
  } finally {
    isSavingArticle.value = false
  }
}

const editArticle = (article) => {
  articleForm.value = {
    id: article.id,
    title: article.title,
    excerpt: article.excerpt,
    cover_image: article.cover_image,
    content: article.content,
    category_id: article.category_id || article.category?.id,
    status: article.status,
    tag_ids: article.tags?.map((tag) => tag.id) || [],
  }
  showArticleDialog.value = true
}

const openArticleDialog = () => {
  resetArticleForm()
  showArticleDialog.value = true
}

const goToArticleEditor = () => {
  router.push({ name: 'ArticleEditor' })
}

const clearRouteActionParam = () => {
  if (!route.query.action) return
  const nextQuery = { ...route.query }
  delete nextQuery.action
  router.replace({ path: route.path, query: nextQuery })
}

const syncRouteState = () => {
  const tab = route.query.tab
  if (typeof tab === 'string' && tab) {
    activeTab.value = tab
  }
  if (route.query.action === 'create-article') {
    openArticleDialog()
    clearRouteActionParam()
  }
}

watch(
  () => route.query,
  () => {
    syncRouteState()
  },
  { immediate: true }
)

const deleteArticle = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这篇文章吗？', '提示', {
      type: 'warning',
    })
    await api.delete(`/articles/${id}`)
    ElMessage.success('删除成功')
    await loadArticles()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const saveMusic = async () => {
  try {
    if (musicForm.value.id) {
      await api.put(`/admin/music/${musicForm.value.id}`, musicForm.value)
      ElMessage.success('音乐更新成功')
    } else {
      await api.post('/admin/music', musicForm.value)
      ElMessage.success('音乐添加成功')
    }
    showMusicDialog.value = false
    musicForm.value = { title: '', artist: '', url: '', cover: '' }
    await loadMusics()
  } catch (error) {
    ElMessage.error('保存失败')
  }
}

const editMusic = (music) => {
  musicForm.value = { ...music }
  showMusicDialog.value = true
}

const deleteMusic = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这首音乐吗？', '提示', {
      type: 'warning',
    })
    await api.delete(`/admin/music/${id}`)
    ElMessage.success('删除成功')
    await loadMusics()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const approveComment = async (id) => {
  try {
    await api.put(`/comments/${id}/status`, { status: 'approved' })
    ElMessage.success('评论已通过')
    await loadPendingComments()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const rejectComment = async (id) => {
  try {
    await api.put(`/comments/${id}/status`, { status: 'rejected' })
    ElMessage.success('评论已拒绝')
    await loadPendingComments()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

onMounted(() => {
  loadArticles()
  loadMusics()
  loadPendingComments()
  loadCategories()
  loadTags()
})
</script>

