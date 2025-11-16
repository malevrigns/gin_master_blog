<template>
  <el-drawer v-model="visible" title="个性化设置" size="400px">
    <div class="space-y-6">
      <!-- 主题色 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">主题色</h3>
        <el-slider
          v-model="settingsStore.themeHue"
          :min="0"
          :max="360"
          :step="1"
          @change="settingsStore.setThemeHue"
          show-input
        />
        <div class="mt-2 text-sm text-gray-500">
          色相值: {{ settingsStore.themeHue }}°
        </div>
      </div>

      <!-- 壁纸模式 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">壁纸模式</h3>
        <el-radio-group v-model="settingsStore.wallpaperMode" @change="settingsStore.setWallpaperMode">
          <el-radio label="solid">纯色背景</el-radio>
          <el-radio label="banner">横幅壁纸</el-radio>
          <el-radio label="fullscreen">全屏壁纸</el-radio>
        </el-radio-group>
      </div>

      <!-- 布局模式 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">布局模式</h3>
        <el-radio-group v-model="settingsStore.layoutMode" @change="settingsStore.setLayoutMode">
          <el-radio label="list">列表布局</el-radio>
          <el-radio label="grid">网格布局</el-radio>
        </el-radio-group>
      </div>

      <!-- 字体设置 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">字体</h3>
        <el-select v-model="settingsStore.fontFamily" @change="settingsStore.setFontFamily">
          <el-option label="系统默认" value="system" />
          <el-option label="思源黑体" value="'Noto Sans SC', sans-serif" />
          <el-option label="思源宋体" value="'Noto Serif SC', serif" />
          <el-option label="等宽字体" value="'Courier New', monospace" />
        </el-select>
      </div>

      <!-- 头像设置 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">头像</h3>
        <el-alert
          v-if="!authStore.token"
          type="warning"
          :closable="false"
          class="mb-4"
        >
          <template #title>
            <span>需要登录后才能上传图片</span>
          </template>
        </el-alert>
        <ImageUploader
          :model-value="settingsStore.avatar"
          @update:model-value="settingsStore.setAvatar"
        />
      </div>

      <!-- Hero背景图片 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">欢迎横幅背景图片</h3>
        <el-alert
          v-if="!authStore.token"
          type="warning"
          :closable="false"
          class="mb-4"
        >
          <template #title>
            <span>需要登录后才能上传图片</span>
          </template>
        </el-alert>
        <ImageUploader
          :model-value="settingsStore.heroBackground"
          @update:model-value="settingsStore.setHeroBackground"
        />
      </div>

      <!-- 樱花特效 -->
      <div>
        <h3 class="text-lg font-semibold mb-4">樱花特效</h3>
        <el-switch
          v-model="settingsStore.sakuraEnabled"
          @change="settingsStore.toggleSakura"
        />
        <div v-if="settingsStore.sakuraEnabled" class="mt-4">
          <el-slider
            v-model="settingsStore.sakuraCount"
            :min="5"
            :max="50"
            :step="5"
            @change="settingsStore.setSakuraCount"
            show-input
          />
          <div class="mt-2 text-sm text-gray-500">
            樱花数量: {{ settingsStore.sakuraCount }}
          </div>
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useSettingsStore } from '../../stores/settings'
import { useAuthStore } from '../../stores/auth'
import ImageUploader from '../common/ImageUploader.vue'

const settingsStore = useSettingsStore()
const authStore = useAuthStore()
const visible = ref(false)

defineExpose({
  open: () => { visible.value = true },
  close: () => { visible.value = false },
})
</script>

