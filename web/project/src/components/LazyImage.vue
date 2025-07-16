<template>
    <div class="lazy-image-container">
      <!-- 加载中状态 -->
      <div v-if="!isLoaded" class="image-placeholder">
        <div class="loading-spinner"></div>
      </div>
      
      <!-- 图片加载完成后显示 -->
      <img 
        v-else 
        :src="imageSrc" 
        :alt="alt"
        :class="imgClass"
      />
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  
  // 定义props
  const props = defineProps({
    src: {
      type: String,
      required: true
    },
    alt: {
      type: String,
      default: ''
    },
    imgClass: {
      type: String,
      default: ''
    }
  })
  
  // 定义emits
  const emit = defineEmits(['load', 'error'])
  
  // 响应式数据
  const isLoaded = ref(false)
  const imageSrc = ref('')
  
  // 方法
  const loadImage = () => {
    // 创建图片对象
    const img = new Image()
    
    // 图片加载成功
    img.onload = () => {
      isLoaded.value = true
      imageSrc.value = props.src
      emit('load')
    }
    
    // 图片加载失败
    img.onerror = () => {
      emit('error')
      console.error('图片加载失败:', props.src)
    }
    
    // 开始加载图片
    img.src = props.src
  }
  
  // 组件挂载时开始加载图片
  onMounted(() => {
    loadImage()
  })
  </script>
  
  <style scoped>
  </style>