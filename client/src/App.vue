<script setup lang="ts">
import layoutComponent from '@/components/layout.vue'
import { isApp } from './utils/tools'
import { onMounted, ref } from 'vue'
import router from './router'
const IsPhone = ref(false)
onMounted(() => {
  IsPhone.value = isApp()
  if (!IsPhone.value) {
    return
  }
  // router.push('/phone')
})
</script>

<template>
  <div class="wrapper" v-if="!IsPhone">
    <layout-component></layout-component>
  </div>
  <router-view v-else></router-view>
</template>

<style scoped></style>
