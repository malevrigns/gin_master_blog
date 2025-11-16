<template>
  <header class="sticky top-0 z-50 bg-slate-950/80 backdrop-blur border-b border-emerald-500/20">
    <nav
      class="max-w-7xl mx-auto px-4 py-3 mt-3 rounded-2xl bg-slate-900/90 border border-emerald-500/30 shadow-xl shadow-emerald-900/40 flex items-center justify-between"
    >
      <!-- 左侧：站点 Logo + 标题 -->
      <div class="flex items-center space-x-3">
        <router-link to="/" class="flex items-center space-x-3 group">
          <div
            class="w-10 h-10 rounded-full bg-gradient-to-br from-emerald-400 via-teal-400 to-cyan-400 flex items-center justify-center text-slate-900 font-bold shadow-lg group-hover:scale-105 transition-transform"
          >
            F
          </div>
          <div class="flex flex-col leading-tight">
            <span class="text-lg font-semibold text-slate-50 tracking-wide">
              我的博客
            </span>
            <span class="text-xs text-emerald-300/80">
              Modern · Firefly Style
            </span>
          </div>
        </router-link>
      </div>

      <!-- 中间：导航菜单 -->
      <div class="hidden lg:flex items-center space-x-6">
        <button
          v-for="item in primaryNav"
          :key="item.path"
          class="inline-flex items-center px-3 py-1.5 rounded-full text-sm font-medium transition-all duration-200"
          :class="[
            isActive(item.path)
              ? 'bg-emerald-500 text-slate-900 shadow-md'
              : 'text-slate-200/80 hover:text-white hover:bg-slate-800/80',
          ]"
          @click="handlePrimaryNavClick(item.path)"
        >
          <el-icon class="mr-1.5 text-base">
            <component :is="item.icon" />
          </el-icon>
          <span>{{ item.label }}</span>
        </button>

        <!-- 关于 下拉菜单 -->
        <el-dropdown trigger="hover">
          <span
            class="inline-flex items-center px-3 py-1.5 rounded-full text-sm font-medium cursor-pointer transition-all duration-200"
            :class="[
              isActive('/about')
                ? 'bg-emerald-500 text-slate-900 shadow-md'
                : 'text-slate-200/80 hover:text-white hover:bg-slate-800/80',
            ]"
          >
            <el-icon class="mr-1.5 text-base">
              <InfoFilled />
            </el-icon>
            <span>关于</span>
            <el-icon class="ml-1 text-xs">
              <ArrowDown />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="router.push('/about')">
                <el-icon class="mr-2"><UserFilled /></el-icon>
                关于本站
              </el-dropdown-item>
              <el-dropdown-item @click="openSponsor">
                <el-icon class="mr-2"><StarFilled /></el-icon>
                赞助支持
              </el-dropdown-item>
              <el-dropdown-item @click="openPlan">
                <el-icon class="mr-2"><Collection /></el-icon>
                番组计划
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>

      <!-- 右侧：搜索 + 工具 + 用户 -->
      <div class="flex items-center space-x-3">
        <!-- 搜索框 -->
        <div class="hidden md:block w-56">
          <el-input
            v-model="searchQuery"
            placeholder="搜索文章..."
            size="small"
            class="rounded-full overflow-hidden bg-slate-800/80"
            @keyup.enter="handleSearch"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </div>

        <el-button
          v-if="canManageContent"
          size="small"
          class="rounded-full border-none bg-emerald-500/90 text-slate-900 hover:bg-emerald-400/90 shadow-lg shadow-emerald-900/30"
          @click="goToArticleCreator"
        >
          <el-icon class="mr-1">
            <EditPen />
          </el-icon>
          写文章
        </el-button>

        <!-- 设置 -->
        <el-tooltip content="设置 (Ctrl+,)" placement="bottom">
          <el-button
            :icon="Setting"
            circle
            size="small"
            class="border-none bg-slate-800/80 text-slate-200 hover:bg-slate-700"
            @click="openSettings"
          />
        </el-tooltip>

        <!-- 主题切换 -->
        <el-tooltip content="切换主题" placement="bottom">
          <el-button
            :icon="themeStore.isDark ? Sunny : Moon"
            circle
            size="small"
            class="border-none bg-slate-800/80 text-slate-200 hover:bg-slate-700"
            @click="themeStore.toggleTheme"
          />
        </el-tooltip>

        <!-- 用户菜单 -->
        <template v-if="authStore.user">
          <el-dropdown>
            <span class="flex items-center cursor-pointer">
              <el-avatar :src="authStore.user.avatar" :size="32" />
              <span class="ml-2 text-slate-100 text-sm">
                {{ authStore.user.username }}
              </span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="$router.push('/admin')">
                  管理后台
                </el-dropdown-item>
                <el-dropdown-item divided @click="handleLogout">
                  退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <template v-else>
          <el-button
            size="small"
            class="rounded-full px-4 border-emerald-500/60 text-emerald-300 bg-transparent hover:bg-emerald-500 hover:text-slate-900"
            @click="openLoginDialog"
          >
            登录
          </el-button>
        </template>
      </div>
    </nav>

    <!-- 登录弹窗 -->
    <el-dialog
      v-model="showLoginDialog"
      :append-to-body="true"
      width="420px"
      destroy-on-close
      @closed="closeLoginDialog"
    >
      <template #header>
        <div class="text-xl font-semibold text-slate-100">
          {{ authMode === 'login' ? '登录到控制台' : '注册新账号' }}
        </div>
      </template>
      <el-tabs v-model="authMode" stretch>
        <el-tab-pane label="登录" name="login">
          <el-form :model="loginForm" label-width="80px">
            <el-form-item label="用户名">
              <el-input v-model="loginForm.username" autocomplete="username" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="loginForm.password" type="password" autocomplete="current-password" show-password />
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="注册" name="register">
          <el-form :model="registerForm" label-width="80px">
            <el-form-item label="用户名">
              <el-input v-model="registerForm.username" autocomplete="username" />
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="registerForm.email" type="email" autocomplete="email" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="registerForm.password" type="password" autocomplete="new-password" show-password />
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
      <template #footer>
        <el-button @click="showLoginDialog = false">取消</el-button>
        <el-button type="primary" :loading="authLoading" @click="handleAuthSubmit">
          {{ authMode === 'login' ? '登录' : '注册' }}
        </el-button>
      </template>
    </el-dialog>
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useThemeStore } from '../../stores/theme'
import { useAuthStore } from '../../stores/auth'
import { ElMessage } from 'element-plus'
import {
  Search,
  Setting,
  HomeFilled,
  Collection,
  Link as LinkIcon,
  InfoFilled,
  ArrowDown,
  UserFilled,
  StarFilled,
  Sunny,
  Moon,
  EditPen,
} from '@element-plus/icons-vue'

const router = useRouter()
const themeStore = useThemeStore()
const authStore = useAuthStore()

const searchQuery = ref('')
const showLoginDialog = ref(false)
const authMode = ref('login')
const authLoading = ref(false)
const loginForm = ref({
  username: '',
  password: '',
})
const registerForm = ref({
  username: '',
  email: '',
  password: '',
})

const canManageContent = computed(() => authStore.user?.role === 'admin')

const primaryNav = [
  { label: '首页', path: '/', icon: HomeFilled },
  { label: '归档', path: '/archive', icon: Collection },
  { label: '友链', path: '/links', icon: LinkIcon },
]

const isActive = (path) => {
  if (path === '/') {
    return router.currentRoute.value.path === '/'
  }
  return router.currentRoute.value.path.startsWith(path)
}

const handlePrimaryNavClick = (path) => {
  if (router.currentRoute.value.path !== path) {
    router.push(path)
  }
}

const handleSearch = () => {
  if (searchQuery.value) {
    router.push({ name: 'Home', query: { search: searchQuery.value } })
  }
}

const goToArticleCreator = () => {
  router.push({ name: 'ArticleEditor' })
}

const handleAuthSubmit = async () => {
  try {
    authLoading.value = true
    if (authMode.value === 'login') {
      await authStore.login(loginForm.value.username, loginForm.value.password)
      ElMessage.success('登录成功')
    } else {
      await authStore.register(
        registerForm.value.username,
        registerForm.value.email,
        registerForm.value.password
      )
      ElMessage.success('注册成功')
    }
    closeLoginDialog()
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '操作失败')
  } finally {
    authLoading.value = false
  }
}

const handleLogout = () => {
  authStore.logout()
  ElMessage.success('已退出登录')
  router.push('/')
}

const openLoginDialog = () => {
  authMode.value = 'login'
  showLoginDialog.value = true
}

const closeLoginDialog = () => {
  showLoginDialog.value = false
  loginForm.value = { username: '', password: '' }
  registerForm.value = { username: '', email: '', password: '' }
}

const openSettings = () => {
  window.dispatchEvent(new CustomEvent('open-settings'))
}

const openSponsor = () => {
  // 可以在这里替换为你的赞助链接
  window.open('https://github.com', '_blank')
}

const openPlan = () => {
  // 番组计划等外部或内部链接
  router.push('/archive')
}

// 监听打开登录对话框事件
const handleOpenLogin = () => {
  openLoginDialog()
}

onMounted(() => {
  window.addEventListener('open-login', handleOpenLogin)
})

onUnmounted(() => {
  window.removeEventListener('open-login', handleOpenLogin)
})
</script>
