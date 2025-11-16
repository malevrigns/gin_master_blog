<template>
  <div class="md-editor-wrapper">
    <MdEditor
      v-model="localValue"
      :theme="theme"
      :editorId="editorId"
      :preview="preview"
      :style="{ height: resolvedHeight }"
      :toolbars="toolbars"
      :footers="['markdownTotal', '=', 0]"
      :outline="false"
      :tabWidth="2"
      :autoDetectCode="true"
      placeholder="在这里书写内容，支持 Markdown / Emoji / 代码块"
      @change="handleChange"
    />
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  height: {
    type: [Number, String],
    default: 520,
  },
  preview: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits(['update:modelValue'])

const editorId = `md-editor-${Math.random().toString(36).slice(2, 8)}`
const localValue = ref(props.modelValue)

const resolvedHeight = computed(() =>
  typeof props.height === 'number' ? `${props.height}px` : props.height
)

const theme = computed(() => (document.documentElement.classList.contains('dark') ? 'dark' : 'light'))

const toolbars = [
  'bold',
  'underline',
  'italic',
  'strikeThrough',
  '-',
  'title',
  'quote',
  'unorderedList',
  'orderedList',
  'taskList',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  '-',
  'revoke',
  'next',
  'preview',
  'fullscreen',
]

watch(
  () => props.modelValue,
  (val) => {
    if (val !== localValue.value) {
      localValue.value = val || ''
    }
  }
)

const handleChange = (val) => {
  emit('update:modelValue', val)
}
</script>

<style scoped>
.md-editor-wrapper :deep(.md-editor) {
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  background: rgba(10, 10, 18, 0.9);
}

.md-editor-wrapper :deep(.md-toolbar-wrapper) {
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  background: rgba(20, 20, 36, 0.85);
}

.md-editor-wrapper :deep(.md-editor-content) {
  background: transparent;
}
</style>
