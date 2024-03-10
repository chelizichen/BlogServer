<script lang="ts">
export default {
  name: 'phone-component'
}
</script>
<style scoped>
.container {
  padding: 10px;
}
.card {
  margin-bottom: 20px;
}
.content {
  max-height: 200px;
}
</style>
<template>
  <div class="container">
    <el-card class="card" v-for="item in articleList" :key="item.id">
      <template #header
        ><b @click="toDetail(item)">{{ item.title }}</b></template
      >
      <div v-html="item.content" class="content"></div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { getArticleList } from '@/api/article'
import { onMounted, ref } from 'vue'
import moment from 'moment'
import router from '@/router'
import { useArticleStore } from '@/stores/counter'

onMounted(async () => {
  const data = await getArticleList(params.value)
  articleList.value = data.Data.list
  total.value = data.Data.total
})
const ArticleStore = useArticleStore()
const params = ref({
  offset: 0,
  size: 10,
  keyword: ''
})
const articleList = ref()
const total = ref()

function toDetail(item) {
  ArticleStore.setArticle(item.title, item.content)
  router.push({
    path: '/detail',
    query: {
      id: item.id
    }
  })
}
</script>

<style scoped></style>
