<template>
  <div>
    <h1>用户登录</h1>
    <form @submit.prevent="login">
      <input v-model="email" placeholder="邮箱" />
      <input v-model="password" type="password" placeholder="密码" />
      <button type="submit">登录</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';

const email = ref('');
const password = ref('');

const login = async () => {
  try {
    const response = await axios.post('/api/users/login', {
      email: email.value,
      password: password.value,
    });
    console.log(response.data);
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped>
</style>