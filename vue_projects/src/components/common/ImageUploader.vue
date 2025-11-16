<template>
  <div class="image-uploader">
    <el-upload
      :action="uploadUrl"
      :headers="headers"
      :name="uploadType === 'image' ? 'image' : 'file'"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleError"
      :before-upload="beforeUpload"
      accept="image/*"
      class="uploader"
    >
      <div class="upload-area">
        <img 
          v-if="imageUrl" 
          :key="imageUrl"
          :src="getImageUrl(imageUrl)" 
          class="preview-image"
          @error="handleImageLoadError"
        />
        <div v-else class="upload-placeholder">
          <el-icon class="text-4xl text-gray-400"><Plus /></el-icon>
          <p class="mt-2 text-sm text-gray-500">点击上传图片</p>
        </div>
      </div>
    </el-upload>
    <div v-if="imageUrl" class="mt-2 flex items-center space-x-2">
      <el-button size="small" @click="clearImage">清除</el-button>
      <span class="text-xs text-gray-500">{{ imageUrl }}</span>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import api from '../../api'
import { useAuthStore } from '../../stores/auth'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  uploadType: {
    type: String,
    default: 'image', // 'image' or 'file'
  },
})

const emit = defineEmits(['update:modelValue'])

const authStore = useAuthStore()
const imageUrl = ref(props.modelValue)
const uploadUrl = computed(() => {
  return props.uploadType === 'image' ? '/api/upload/image' : '/api/upload/file'
})

// 响应式的 headers，当 token 变化时自动更新
const headers = computed(() => {
  const h = {}
  if (authStore.token) {
    h['Authorization'] = `Bearer ${authStore.token}`
  }
  return h
})

watch(() => props.modelValue, (newVal) => {
  imageUrl.value = newVal
})

const beforeUpload = (file) => {
  // 检查是否已登录
  if (!authStore.token) {
    ElMessage.warning('请先登录后再上传图片')
    // 触发登录对话框
    window.dispatchEvent(new CustomEvent('open-login'))
    return false
  }

  const isImage = file.type.startsWith('image/')
  const isLt5M = file.size / 1024 / 1024 < 5

  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt5M) {
    ElMessage.error('图片大小不能超过 5MB!')
    return false
  }
  return true
}

const handleSuccess = (response) => {
  console.log('Upload response:', response)
  
  // Element Plus 的响应可能是 response.data 或直接是 response
  const data = response.data || response
  
  if (data && data.url) {
    // 确保URL是完整的路径
    let finalUrl = data.url
    if (!finalUrl.startsWith('http') && !finalUrl.startsWith('/')) {
      finalUrl = '/' + finalUrl
    }
    
    imageUrl.value = finalUrl
    emit('update:modelValue', finalUrl)
    ElMessage.success('上传成功!')
  } else {
    console.error('Invalid response format:', response)
    ElMessage.error('上传失败: 响应格式错误')
  }
}

const handleError = (error) => {
  console.error('Upload error:', error)
  
  // 检查是否是认证错误
  if (error.response?.status === 401 || error.message?.includes('401') || error.message?.includes('Unauthorized')) {
    ElMessage.warning('上传需要登录，请先登录')
    // 触发登录对话框
    window.dispatchEvent(new CustomEvent('open-login'))
  } else {
    ElMessage.error('上传失败，请重试')
  }
}

const clearImage = () => {
  imageUrl.value = ''
  emit('update:modelValue', '')
}

const getImageUrl = (url) => {
  if (!url) return ''
  // 确保URL格式正确
  if (url.startsWith('http')) return url
  if (url.startsWith('/')) return url
  return '/' + url
}

const handleImageLoadError = (e) => {
  console.error('Image load error in uploader:', e.target.src)
  // 不显示错误，只是记录
}
</script>

<style scoped>
.image-uploader {
  width: 100%;
}

.uploader {
  width: 100%;
}

.upload-area {
  width: 100%;
  min-height: 200px;
  border: 2px dashed #d9d9d9;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
  overflow: hidden;
}

.upload-area:hover {
  border-color: var(--el-color-primary);
}

.preview-image {
  width: 100%;
  height: auto;
  max-height: 300px;
  object-fit: contain;
}

.upload-placeholder {
  text-align: center;
  padding: 40px;
}
</style>

