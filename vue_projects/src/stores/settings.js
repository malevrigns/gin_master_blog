import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  // 主题色（色相值 0-360）
  const themeHue = ref(200)
  
  // 壁纸模式：banner, fullscreen, solid
  const wallpaperMode = ref('solid')
  
  // 布局模式：list, grid
  const layoutMode = ref('list')
  
  // 字体设置
  const fontFamily = ref('system')
  
  // 头像和背景图片
  const avatar = ref('')
  const heroBackground = ref('')
  
  // 樱花特效
  const sakuraEnabled = ref(true) // 默认开启
  const sakuraCount = ref(20) // 默认数量

  const initSettings = () => {
    const saved = localStorage.getItem('blogSettings')
    if (saved) {
      const settings = JSON.parse(saved)
      themeHue.value = settings.themeHue ?? 200
      wallpaperMode.value = settings.wallpaperMode ?? 'solid'
      layoutMode.value = settings.layoutMode ?? 'list'
      fontFamily.value = settings.fontFamily ?? 'system'
      avatar.value = settings.avatar ?? ''
      heroBackground.value = settings.heroBackground ?? ''
      sakuraEnabled.value = settings.sakuraEnabled ?? true
      sakuraCount.value = settings.sakuraCount ?? 20
    }
    applySettings()
  }

  const saveSettings = () => {
    const settings = {
      themeHue: themeHue.value,
      wallpaperMode: wallpaperMode.value,
      layoutMode: layoutMode.value,
      fontFamily: fontFamily.value,
      avatar: avatar.value,
      heroBackground: heroBackground.value,
      sakuraEnabled: sakuraEnabled.value,
      sakuraCount: sakuraCount.value,
    }
    localStorage.setItem('blogSettings', JSON.stringify(settings))
    applySettings()
  }

  const applySettings = () => {
    // 应用主题色
    document.documentElement.style.setProperty('--theme-hue', themeHue.value)
    
    // 应用字体
    if (fontFamily.value !== 'system') {
      document.documentElement.style.setProperty('--font-family', fontFamily.value)
    }
    
    // 应用壁纸模式
    document.documentElement.setAttribute('data-wallpaper', wallpaperMode.value)
    
    // 应用布局模式
    document.documentElement.setAttribute('data-layout', layoutMode.value)
  }

  const setThemeHue = (hue) => {
    themeHue.value = hue
    saveSettings()
  }

  const setWallpaperMode = (mode) => {
    wallpaperMode.value = mode
    saveSettings()
  }

  const setLayoutMode = (mode) => {
    layoutMode.value = mode
    saveSettings()
  }

  const setFontFamily = (font) => {
    fontFamily.value = font
    saveSettings()
  }

  const toggleSakura = () => {
    sakuraEnabled.value = !sakuraEnabled.value
    saveSettings()
  }

  const setSakuraCount = (count) => {
    sakuraCount.value = count
    saveSettings()
  }

  const setAvatar = (url) => {
    avatar.value = url
    saveSettings()
  }

  const setHeroBackground = (url) => {
    heroBackground.value = url
    saveSettings()
  }

  watch([themeHue, wallpaperMode, layoutMode, fontFamily], () => {
    applySettings()
  })

  return {
    themeHue,
    wallpaperMode,
    layoutMode,
    fontFamily,
    avatar,
    heroBackground,
    sakuraEnabled,
    sakuraCount,
    initSettings,
    setThemeHue,
    setWallpaperMode,
    setLayoutMode,
    setFontFamily,
    setAvatar,
    setHeroBackground,
    toggleSakura,
    setSakuraCount,
  }
})

