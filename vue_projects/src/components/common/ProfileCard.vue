<template>
  <div class="profile-card glass-card card-hover">
    <div class="text-center relative z-10">
      <div class="avatar-shell mx-auto mb-4">
        <img
          :key="avatarUrl"
          :src="avatarUrl"
          :alt="name"
          class="avatar-image"
          @error="handleImageError"
        />
        <span class="status-dot"></span>
      </div>
      <h2 class="profile-name gradient-text mb-2">
        {{ name }}
      </h2>
      <p class="profile-greeting mb-5">{{ greeting }}</p>
      <div class="flex justify-center space-x-3">
        <a
          v-for="link in socialLinks"
          :key="link.name"
          :href="link.url"
          target="_blank"
          rel="noopener noreferrer"
          class="social-button"
          :title="link.name"
        >
          <component :is="link.icon" class="w-5 h-5" />
        </a>
      </div>
    </div>
    <div class="card-orbit card-orbit--one"></div>
    <div class="card-orbit card-orbit--two"></div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Camera, Refresh } from '@element-plus/icons-vue'
import { useSettingsStore } from '../../stores/settings'

const settingsStore = useSettingsStore()

const name = ref('我的博客')
const greeting = ref('用文字和代码记录灵感。')

const fallbackAvatar =
  'data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200"><defs><linearGradient id="g" x1="0%" y1="0%" x2="100%" y2="100%"><stop offset="0%" stop-color="%23ff61e6"/><stop offset="100%" stop-color="%233cf2ff"/></linearGradient></defs><rect width="200" height="200" rx="100" fill="url(%23g)"/><circle cx="150" cy="40" r="12" fill="%23ffb347"/><circle cx="60" cy="150" r="10" fill="%238ef9a5"/></svg>'

const resolveMediaUrl = (url) => {
  if (!url) return null
  const trimmed = url.trim()
  if (!trimmed) return null
  if (trimmed.startsWith('http')) return trimmed
  if (trimmed.startsWith('/')) return trimmed
  return '/' + trimmed
}

const avatarUrl = computed(() => {
  return resolveMediaUrl(settingsStore.avatar) || fallbackAvatar
})

const socialLinks = ref([
  { name: 'GitHub', url: 'https://github.com', icon: Camera },
  { name: 'Email', url: 'mailto:your@email.com', icon: Refresh },
])

const handleImageError = (e) => {
  e.target.src = fallbackAvatar
}
</script>

<style scoped>
.profile-card {
  padding: 2.5rem 1.5rem;
  overflow: hidden;
}

.avatar-shell {
  position: relative;
  width: 116px;
  height: 116px;
  border-radius: 999px;
  padding: 3px;
  background: linear-gradient(135deg, rgba(255, 97, 230, 0.7), rgba(60, 242, 255, 0.7));
}

.avatar-image {
  width: 100%;
  height: 100%;
  border-radius: 999px;
  object-fit: cover;
  display: block;
  background: linear-gradient(135deg, rgba(5, 9, 32, 0.9), rgba(5, 13, 40, 0.9));
}

.status-dot {
  position: absolute;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: var(--accent-lime);
  border: 3px solid rgba(3, 6, 24, 0.95);
  bottom: 8px;
  right: 6px;
  box-shadow: 0 0 12px rgba(142, 249, 165, 0.5);
}

.profile-name {
  font-size: 1.8rem;
  font-weight: 700;
}

.profile-greeting {
  color: rgba(235, 238, 254, 0.7);
  font-size: 0.95rem;
}

.social-button {
  width: 44px;
  height: 44px;
  border-radius: 14px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #f5f5f5;
  background: rgba(255, 255, 255, 0.08);
  transition: all 0.25s ease;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.social-button:hover {
  background: linear-gradient(135deg, rgba(255, 97, 230, 0.5), rgba(60, 242, 255, 0.5));
  transform: translateY(-2px);
  box-shadow: 0 15px 30px rgba(0, 0, 0, 0.35);
}

.card-orbit {
  position: absolute;
  border-radius: 50%;
  pointer-events: none;
  opacity: 0.35;
}

.card-orbit--one {
  width: 280px;
  height: 280px;
  top: -160px;
  left: 50%;
  transform: translateX(-50%);
  background: radial-gradient(circle, rgba(255, 97, 230, 0.35), transparent 60%);
}

.card-orbit--two {
  width: 220px;
  height: 220px;
  bottom: -130px;
  right: -40px;
  background: radial-gradient(circle, rgba(60, 242, 255, 0.25), transparent 60%);
}
</style>
