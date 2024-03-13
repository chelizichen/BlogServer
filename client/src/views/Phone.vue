<script lang="ts">
export default {
  name: 'phone-component'
}
</script>
<style scoped>
.container {
  padding: 10px;
}
.container >>> img {
  width: 100%;
  height: 100%;
}
.card {
  margin-bottom: 20px;
}
.content {
  max-height: 10vh;
}
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 10vh;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
.carousel {
  margin-bottom: 20px;
}
</style>
<template>
  <div class="container">
    <div class="carousel">
      <el-carousel :interval="2000" type="card" height="10vh" indicator-position="none" arrow="always">
        <el-carousel-item v-for="item in columnsList" :key="item">
          <h3 text="2xl" justify="center" @click="toAppColumn(item)">{{ item.column.title }}</h3>
        </el-carousel-item>
      </el-carousel>
    </div>
    <el-card class="card" v-for="item in articleList" :key="item.id">
      <template #header
        ><b @click="toDetail(item)">{{ item.title }}</b>
        <div style="float: right; font-size: 12px">
          {{ moment(item.create_time).format('YYYYMMDD HH:mm:ss') }}
        </div></template
      >
      <div v-code v-html="item.content" class="content"></div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { getArticleList } from '@/api/article'
import { onMounted, ref } from 'vue'
import moment from 'moment'
import router from '@/router'
import { useArticleStore } from '@/stores/counter'
import { getColumns } from '@/api/column'

const columnsList = ref([])

onMounted(async () => {
  const data = await getArticleList(params.value)
  articleList.value = data.Data.list
  total.value = data.Data.total

  const columns = await getColumns({})
  columnsList.value = columns.Data.list
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

function toAppColumn(item: any) {
  router.push({
    path: '/app_column',
    query: {
      id: item.column.id
    }
  })
}
</script>

<style scoped></style>
