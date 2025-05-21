<template>
  <div>
    <h1>多商户发卡系统</h1>
    <UserLogin /> <!-- 确保组件名称与引用一致 -->
    <MerchantRegister />
    <CardList />
  </div>
</template>

<script setup>
import UserLogin from '../components/UserLogin.vue'; // 确保导入路径正确
import MerchantRegister from '../components/MerchantRegister.vue';
import CardList from '../components/CardList.vue';
</script>