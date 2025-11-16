import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import ArticleDetail from '../views/ArticleDetail.vue'
import Category from '../views/Category.vue'
import Tag from '../views/Tag.vue'
import About from '../views/About.vue'
import Admin from '../views/Admin.vue'
import ArticleEditor from '../views/ArticleEditor.vue'
import LabDetail from '../views/LabDetail.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/article/:id',
    name: 'ArticleDetail',
    component: ArticleDetail,
  },
  {
    path: '/category/:id',
    name: 'Category',
    component: Category,
  },
  {
    path: '/tag/:slug',
    name: 'Tag',
    component: Tag,
  },
  {
    path: '/about',
    name: 'About',
    component: About,
  },
  {
    path: '/archive',
    name: 'Archive',
    component: () => import('../views/Archive.vue'),
  },
  {
    path: '/links',
    name: 'Links',
    component: () => import('../views/Links.vue'),
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
    meta: { requiresAuth: true },
  },
  {
    path: '/editor',
    name: 'ArticleEditor',
    component: ArticleEditor,
    meta: { requiresAuth: true },
  },
  {
    path: '/labs/:slug',
    name: 'LabDetail',
    component: LabDetail,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  },
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next({ name: 'Home' })
  } else {
    next()
  }
})

export default router

