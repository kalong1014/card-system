<template>
  <div class="page-builder">
    <div class="page-builder-header">
      <el-button type="primary" @click="savePage">保存页面</el-button>
      <el-button @click="previewPage">预览</el-button>
      <el-button @click="publishPage">发布</el-button>
      <el-button @click="undo">撤销</el-button>
      <el-button @click="redo">重做</el-button>
    </div>
    
    <div class="page-builder-content">
      <!-- 左侧组件库 -->
      <div class="components-panel">
        <h3>组件库</h3>
        <div class="component-item" v-for="component in components" :key="component.id" 
             draggable="true" @dragstart="onDragStart(component)">
          {{ component.name }}
        </div>
      </div>
      
      <!-- 中间编辑区域 -->
      <div class="editor-area" @dragover="onDragOver" @drop="onDrop">
        <h3>页面预览</h3>
        <div class="page-container" :style="{ minHeight: '500px', border: '1px dashed #ccc', padding: '10px' }">
          <div v-for="(element, index) in pageElements" :key="element.id" 
               class="page-element" :style="getElementStyle(element)"
               @click="selectElement(index)"
               @dragstart="onElementDragStart(index)"
               draggable="true">
            <div v-if="element.type === 'text'" class="text-element" :style="element.style">
              {{ element.content || '文本内容' }}
            </div>
            <div v-if="element.type === 'image'" class="image-element">
              <img :src="element.src || 'https://picsum.photos/400/200'" :style="element.style" alt="图片">
            </div>
            <div v-if="element.type === 'button'" class="button-element" :style="element.style">
              {{ element.content || '按钮' }}
            </div>
            <div v-if="element.selected" class="element-selected-indicator"></div>
          </div>
        </div>
      </div>
      
      <!-- 右侧属性面板 -->
      <div class="properties-panel" v-if="selectedElementIndex !== -1">
        <h3>属性设置</h3>
        
        <el-form :model="selectedElement" label-width="80px">
          <el-form-item label="元素类型">
            <el-input v-model="selectedElement.type" disabled></el-input>
          </el-form-item>
          
          <el-form-item label="内容" v-if="selectedElement.type === 'text' || selectedElement.type === 'button'">
            <el-input v-model="selectedElement.content" type="textarea"></el-input>
          </el-form-item>
          
          <el-form-item label="图片URL" v-if="selectedElement.type === 'image'">
            <el-input v-model="selectedElement.src"></el-input>
          </el-form-item>
          
          <el-form-item label="宽度">
            <el-input v-model="selectedElement.style.width"></el-input>
          </el-form-item>
          
          <el-form-item label="高度">
            <el-input v-model="selectedElement.style.height"></el-input>
          </el-form-item>
          
          <el-form-item label="背景颜色">
            <el-color-picker v-model="selectedElement.style.backgroundColor"></el-color-picker>
          </el-form-item>
          
          <el-form-item label="字体颜色" v-if="selectedElement.type !== 'image'">
            <el-color-picker v-model="selectedElement.style.color"></el-color-picker>
          </el-form-item>
          
          <el-form-item label="字体大小" v-if="selectedElement.type !== 'image'">
            <el-input v-model="selectedElement.style.fontSize"></el-input>
          </el-form-item>
          
          <el-form-item>
            <el-button type="primary" @click="updateElement">应用更改</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, nextTick, toRefs } from 'vue'
import { v4 as uuidv4 } from 'uuid'

// 组件库
const components = ref([
  { id: 'text', name: '文本' },
  { id: 'image', name: '图片' },
  { id: 'button', name: '按钮' },
  { id: 'product-list', name: '商品列表' },
  { id: 'carousel', name: '轮播图' }
])

// 页面元素
const pageElements = ref([])

// 选中的元素索引
const selectedElementIndex = ref(-1)

// 历史记录
const history = ref([])
const historyIndex = ref(-1)

// 初始化页面
const initPage = () => {
  // 加载已保存的页面配置
  const savedPage = localStorage.getItem('pageConfig')
  if (savedPage) {
    pageElements.value = JSON.parse(savedPage)
  }
}

// 拖拽相关事件
const onDragStart = (component) => {
  event.dataTransfer.setData('component-type', component.id)
}

const onDragOver = (event) => {
  event.preventDefault()
}

const onDrop = (event) => {
  event.preventDefault()
  const componentType = event.dataTransfer.getData('component-type')
  
  if (!componentType) return
  
  // 创建新元素
  const newElement = createNewElement(componentType)
  
  // 添加到历史记录
  saveHistory()
  
  // 添加到页面
  pageElements.value.push(newElement)
  
  // 选中新元素
  selectedElementIndex.value = pageElements.value.length - 1
}

// 创建新元素
const createNewElement = (type) => {
  const elementId = uuidv4()
  
  const baseElement = {
    id: elementId,
    type: type,
    selected: false,
    style: {
      width: '200px',
      height: 'auto',
      margin: '10px',
      padding: '10px',
      borderRadius: '4px',
      border: '1px solid #ddd'
    }
  }
  
  // 根据组件类型设置不同的默认属性
  switch (type) {
    case 'text':
      return {
        ...baseElement,
        content: '这是一段文本',
        style: {
          ...baseElement.style,
          fontSize: '14px',
          color: '#333',
          backgroundColor: '#fff'
        }
      }
    case 'image':
      return {
        ...baseElement,
        src: 'https://picsum.photos/400/200',
        style: {
          ...baseElement.style,
          width: '400px',
          height: '200px',
          backgroundColor: '#f5f5f5'
        }
      }
    case 'button':
      return {
        ...baseElement,
        content: '点击按钮',
        style: {
          ...baseElement.style,
          fontSize: '14px',
          color: '#fff',
          backgroundColor: '#409eff',
          textAlign: 'center',
          cursor: 'pointer'
        }
      }
    case 'product-list':
      return {
        ...baseElement,
        style: {
          ...baseElement.style,
          width: '100%',
          backgroundColor: '#fff'
        }
      }
    case 'carousel':
      return {
        ...baseElement,
        style: {
          ...baseElement.style,
          width: '100%',
          height: '300px',
          backgroundColor: '#f5f5f5'
        }
      }
    default:
      return baseElement
  }
}

// 选中元素
const selectElement = (index) => {
  // 取消之前选中的元素
  if (selectedElementIndex.value !== -1) {
    pageElements.value[selectedElementIndex.value].selected = false
  }
  
  // 选中当前元素
  pageElements.value[index].selected = true
  selectedElementIndex.value = index
}

// 获取选中的元素
const selectedElement = computed(() => {
  if (selectedElementIndex.value === -1) {
    return null
  }
  return pageElements.value[selectedElementIndex.value]
})

// 获取元素样式
const getElementStyle = (element) => {
  return {
    ...element.style,
    position: 'relative'
  }
}

// 更新元素
const updateElement = () => {
  if (selectedElementIndex.value === -1) return
  
  // 添加到历史记录
  saveHistory()
  
  // 更新元素
  pageElements.value[selectedElementIndex.value] = {
    ...selectedElement.value
  }
  
  // 保存页面
  savePage()
}

// 保存页面
const savePage = () => {
  localStorage.setItem('pageConfig', JSON.stringify(pageElements.value))
  ElMessage.success('页面已保存')
}

// 预览页面
const previewPage = () => {
  // 创建预览窗口
  const previewWindow = window.open('', '_blank')
  if (!previewWindow) {
    ElMessage.warning('请允许弹出窗口以查看预览')
    return
  }
  
  // 生成预览HTML
  const previewHtml = generatePreviewHtml()
  previewWindow.document.write(previewHtml)
  previewWindow.document.close()
}

// 生成预览HTML
const generatePreviewHtml = () => {
  // 这里简化处理，实际项目中应该根据页面配置生成完整的HTML
  let elementsHtml = ''
  
  pageElements.value.forEach(element => {
    let elementHtml = ''
    
    switch (element.type) {
      case 'text':
        elementHtml = `<div style="${formatStyle(element.style)}">${element.content}</div>`
        break
      case 'image':
        elementHtml = `<img src="${element.src}" style="${formatStyle(element.style)}" alt="图片">`
        break
      case 'button':
        elementHtml = `<button style="${formatStyle(element.style)}">${element.content}</button>`
        break
      case 'product-list':
        elementHtml = `<div style="${formatStyle(element.style)}">商品列表区域</div>`
        break
      case 'carousel':
        elementHtml = `<div style="${formatStyle(element.style)}">轮播图区域</div>`
        break
    }
    
    elementsHtml += `<div style="margin: 10px;">${elementHtml}</div>`
  })
  
  return `
    <!DOCTYPE html>
    <html>
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>页面预览</title>
      <style>
        body {
          margin: 0;
          padding: 20px;
          font-family: Arial, sans-serif;
        }
      </style>
    </head>
    <body>
      ${elementsHtml}
    </body>
    </html>
  `
}

// 格式化样式对象为CSS字符串
const formatStyle = (styleObj) => {
  let styleStr = ''
  for (const [key, value] of Object.entries(styleObj)) {
    styleStr += `${key}: ${value}; `
  }
  return styleStr
}

// 保存历史记录
const saveHistory = () => {
  // 如果当前不是历史记录的最后一项，删除后面的所有历史记录
  if (historyIndex.value < history.value.length - 1) {
    history.value = history.value.slice(0, historyIndex.value + 1)
  }
  
  // 保存当前状态
  history.value.push(JSON.parse(JSON.stringify(pageElements.value)))
  historyIndex.value = history.value.length - 1
}

// 撤销
const undo = () => {
  if (historyIndex.value <= 0) return
  
  historyIndex.value--
  pageElements.value = JSON.parse(JSON.stringify(history.value[historyIndex.value]))
}

// 重做
const redo = () => {
  if (historyIndex.value >= history.value.length - 1) return
  
  historyIndex.value++
  pageElements.value = JSON.parse(JSON.stringify(history.value[historyIndex.value]))
}

// 发布页面
const publishPage = () => {
  // 实际项目中应该将页面配置保存到数据库
  // 并生成对应的前端页面文件
  ElMessage.success('页面已发布')
}

// 元素拖拽排序
const onElementDragStart = (index) => {
  event.dataTransfer.setData('element-index', index)
}

// 元素拖拽排序的drop事件
const onElementDrop = (event) => {
  event.preventDefault()
  const fromIndex = parseInt(event.dataTransfer.getData('element-index'))
  if (isNaN(fromIndex) || fromIndex === selectedElementIndex.value) return
  
  // 添加到历史记录
  saveHistory()
  
  // 移动元素
  const element = pageElements.value.splice(fromIndex, 1)[0]
  pageElements.value.splice(selectedElementIndex.value, 0, element)
  
  // 更新选中元素的索引
  if (fromIndex < selectedElementIndex.value) {
    selectedElementIndex.value--
  }
}

// 初始化页面
initPage()
</script>

<style scoped>
.page-builder {
  display: flex;
  height: calc(100vh - 60px);
}

.page-builder-header {
  padding: 10px;
  border-bottom: 1px solid #eee;
}

.page-builder-content {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.components-panel {
  width: 200px;
  padding: 10px;
  border-right: 1px solid #eee;
  overflow-y: auto;
}

.component-item {
  padding: 8px;
  margin-bottom: 10px;
  background-color: #f5f5f5;
  border-radius: 4px;
  cursor: move;
  text-align: center;
}

.editor-area {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
}

.properties-panel {
  width: 300px;
  padding: 10px;
  border-left: 1px solid #eee;
  overflow-y: auto;
}

.page-element {
  position: relative;
  margin-bottom: 15px;
}

.element-selected-indicator {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border: 2px dashed #409eff;
  pointer-events: none;
}
</style>    