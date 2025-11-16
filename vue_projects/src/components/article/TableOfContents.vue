<template>
  <div
    v-if="headings.length > 0"
    class="toc-container bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 sticky top-4"
  >
    <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-gray-100 flex items-center">
      <el-icon class="mr-2"><List /></el-icon>
      目录
    </h3>
    <nav class="toc-nav">
      <ul class="space-y-2">
        <li
          v-for="heading in headings"
          :key="heading.id"
          :class="[
            'toc-item',
            `toc-level-${heading.level}`,
            { 'toc-active': activeId === heading.id }
          ]"
        >
          <a
            :href="`#${heading.id}`"
            @click.prevent="scrollToHeading(heading.id)"
            class="block py-1 px-2 rounded hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
            :class="{
              'text-primary-600 dark:text-primary-400 font-semibold': activeId === heading.id,
              'text-gray-600 dark:text-gray-400': activeId !== heading.id
            }"
          >
            {{ heading.text }}
          </a>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { List } from '@element-plus/icons-vue'

const props = defineProps({
  content: {
    type: String,
    default: '',
  },
})

const headings = ref([])
const activeId = ref('')

const extractHeadings = () => {
  if (!props.content) return []
  
  const headingRegex = /^#{1,6}\s+(.+)$/gm
  const matches = []
  let match
  
  while ((match = headingRegex.exec(props.content)) !== null) {
    const level = match[0].match(/^#+/)[0].length
    const text = match[1].trim()
    const id = text.toLowerCase()
      .replace(/[^\w\s-]/g, '')
      .replace(/\s+/g, '-')
    
    matches.push({ level, text, id })
  }
  
  return matches
}

const scrollToHeading = (id) => {
  const element = document.getElementById(id)
  if (element) {
    element.scrollIntoView({ behavior: 'smooth', block: 'start' })
    activeId.value = id
  }
}

const updateActiveHeading = () => {
  const headingElements = headings.value.map(h => ({
    id: h.id,
    element: document.getElementById(h.id)
  })).filter(h => h.element)
  
  const scrollPosition = window.scrollY + 100
  
  for (let i = headingElements.length - 1; i >= 0; i--) {
    const { id, element } = headingElements[i]
    if (element.offsetTop <= scrollPosition) {
      activeId.value = id
      break
    }
  }
}

onMounted(() => {
  headings.value = extractHeadings()
  window.addEventListener('scroll', updateActiveHeading)
  updateActiveHeading()
})

onUnmounted(() => {
  window.removeEventListener('scroll', updateActiveHeading)
})
</script>

<style scoped>
.toc-level-1 {
  padding-left: 0;
  font-size: 1rem;
}

.toc-level-2 {
  padding-left: 1rem;
  font-size: 0.95rem;
}

.toc-level-3 {
  padding-left: 2rem;
  font-size: 0.9rem;
}

.toc-level-4 {
  padding-left: 3rem;
  font-size: 0.85rem;
}

.toc-level-5 {
  padding-left: 4rem;
  font-size: 0.8rem;
}

.toc-level-6 {
  padding-left: 5rem;
  font-size: 0.75rem;
}

.toc-active {
  border-left: 3px solid var(--el-color-primary);
}
</style>

