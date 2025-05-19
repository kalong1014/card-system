<template>
  <div class="card-secret-generate">
    <el-card class="box-card">
      <template #header>
        <div class="clearfix">
          <span>生成卡密</span>
          <el-button style="float: right; padding: 3px 0" type="text">
            <router-link to="/merchant/card-secrets">返回列表</router-link>
          </el-button>
        </div>
      </template>
      
      <el-form :model="form" :rules="rules" ref="formRef" label-width="120px">
        <el-form-item label="选择商品" prop="productId">
          <el-select v-model="form.productId" placeholder="请选择商品">
            <el-option
              v-for="item in products"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="生成数量" prop="count">
          <el-input v-model.number="form.count" type="number" placeholder="请输入生成数量"></el-input>
        </el-form-item>
        
        <el-form-item label="卡密长度" prop="length">
          <el-input v-model.number="form.length" type="number" placeholder="请输入卡密长度"></el-input>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="submitForm">立即生成</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    
    <el-card class="box-card" v-if="generatedCardSecrets.length > 0">
      <template #header>
        <div>生成结果</div>
      </template>
      
      <el-table :data="generatedCardSecrets" stripe style="width: 100%">
        <el-table-column prop="secret" label="卡密"></el-table-column>
        <el-table-column prop="status" label="状态">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">{{getStatusText(scope.row.status)}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="mini" @click="copySecret(scope.row.secret)">复制</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <el-button type="primary" @click="downloadCardSecrets" style="margin-top: 20px">
        下载卡密文件
      </el-button>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { generateCardSecrets, getMerchantProducts } from '@/api/merchant'

const router = useRouter()
const formRef = ref(null)
const form = reactive({
  productId: '',
  count: 10,
  length: 16
})

const rules = reactive({
  productId: [
    { required: true, message: '请选择商品', trigger: 'change' }
  ],
  count: [
    { required: true, message: '请输入生成数量', trigger: 'blur' },
    { type: 'number', min: 1, max: 1000, message: '数量范围为1-1000', trigger: 'blur' }
  ],
  length: [
    { required: true, message: '请输入卡密长度', trigger: 'blur' },
    { type: 'number', min: 8, max: 32, message: '长度范围为8-32', trigger: 'blur' }
  ]
})

const products = ref([])
const generatedCardSecrets = ref([])

// 获取商户商品列表
const fetchProducts = async () => {
  try {
    const res = await getMerchantProducts()
    products.value = res.data
  } catch (error) {
    ElMessage.error(error.message || '获取商品列表失败')
  }
}

// 提交表单
const submitForm = async () => {
  await formRef.value.validate()
  
  try {
    const res = await generateCardSecrets(form.productId, form.count, form.length)
    generatedCardSecrets.value = res.data.card_secrets
    
    ElMessage.success(`成功生成 ${res.data.count} 个卡密`)
  } catch (error) {
    ElMessage.error(error.message || '生成卡密失败')
  }
}

// 重置表单
const resetForm = () => {
  formRef.value.resetFields()
  generatedCardSecrets.value = []
}

// 复制卡密
const copySecret = (secret) => {
  navigator.clipboard.writeText(secret).then(() => {
    ElMessage.success('复制成功')
  }).catch(() => {
    ElMessage.error('复制失败')
  })
}

// 下载卡密文件
const downloadCardSecrets = () => {
  if (generatedCardSecrets.value.length === 0) {
    ElMessage.warning('没有可下载的卡密')
    return
  }
  
  const content = generatedCardSecrets.value.map(item => item.secret).join('\n')
  const blob = new Blob([content], { type: 'text/plain;charset=utf-8' })
  const url = URL.createObjectURL(blob)
  
  const a = document.createElement('a')
  a.href = url
  a.download = `card_secrets_${new Date().getTime()}.txt`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = {
    0: '未使用',
    1: '已使用',
    2: '已锁定'
  }
  return statusMap[status] || '未知状态'
}

// 获取状态标签类型
const getStatusType = (status) => {
  const typeMap = {
    0: 'success',
    1: 'info',
    2: 'warning'
  }
  return typeMap[status] || 'default'
}

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped>
.box-card {
  margin-bottom: 20px;
}
</style>    