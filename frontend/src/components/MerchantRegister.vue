<template>
  <div>
    <h1>商户注册</h1>
    <form @submit.prevent="register">
      <input v-model="name" placeholder="商户名称" />
      <input v-model="email" placeholder="邮箱" />
      <input v-model="phone" placeholder="电话" />
      <input v-model="address" placeholder="地址" />
      <button type="submit">注册</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';

const name = ref('');
const email = ref('');
const phone = ref('');
const address = ref('');

const register = async () => {
  try {
    const response = await axios.post('/api/merchants/register', {
      name: name.value,
      email: email.value,
      phone: phone.value,
      address: address.value,
    });
    console.log(response.data);
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped>
</style>