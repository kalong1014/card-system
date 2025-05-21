<template>
  <div>
    <h1>卡密列表</h1>
    <ul>
      <li v-for="card in cardSecrets" :key="card.id">{{ card.secret }}</li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const cardSecrets = ref([]);

const getCardSecrets = async () => {
  try {
    const response = await axios.get('/api/cards/product/1');
    cardSecrets.value = response.data.data;
  } catch (error) {
    console.error(error);
  }
};

onMounted(() => {
  getCardSecrets();
});
</script>

<style scoped>
</style>