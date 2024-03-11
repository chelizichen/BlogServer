<script lang="ts">
export default {
  name: 'detail-component'
}
</script>
<template>
  <div>
    <el-card class="card">
      <template #header>{{ ArticleStore.articleTitle }}</template>
      <div v-html="ArticleStore.articleContent"></div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { getArticle } from '@/api/article'
import { useArticleStore } from '@/stores/counter'
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
const ArticleStore = useArticleStore()
const route = useRoute()
onMounted(async () => {
  if (!ArticleStore.articleContent && !ArticleStore.articleTitle) {
    const id = route.query.id
    if (!id) {
      return
    }
    const query = { id }
    const res = await getArticle(query)
    ArticleStore.setArticle(res.Data.title, res.Data.content)
  }
})
</script>

<style scoped>
.card >>> img {
  width: 100%;
  height: 100%;
}
</style>
