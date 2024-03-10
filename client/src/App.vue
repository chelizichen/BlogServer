<script setup lang="ts">
import layoutComponent from "@/components/layout.vue";
import { isApp } from "./utils/tools";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
const IsPhone = ref(false);
const [route, router] = [useRoute(), useRouter()];
onMounted(() => {
  IsPhone.value = isApp();
  if (!IsPhone.value) {
    return;
  }
  const whiteList = ["/phone", "/detail"];
  const matchedNext = route.matched.map((v) => v.path);
  console.log('matchs',matchedNext);
  
  if (matchedNext.every((v) => whiteList.findIndex((e) => e === v) == -1)){
    router.push('/phone')
  }
});
</script>

<template>
  <div class="wrapper" v-if="!IsPhone">
    <layout-component></layout-component>
  </div>
  <router-view v-else></router-view>
</template>

<style scoped></style>
