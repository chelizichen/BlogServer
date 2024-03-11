<script lang="ts">
export default {
  name: 'home-component'
}
</script>
<template>
  <div>
    <el-table :data="articleList" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="id" label="编号" width="90" align="center" />
      <el-table-column prop="title" label="标题" width="180" align="center" />
      <el-table-column prop="content" label="内容" align="center" />
      <el-table-column prop="type" label="类型" align="center" />
      <el-table-column prop="create_time" label="创建时间" width="180" align="center">
        <template #default="scoped">
          {{ moment(scoped.row.create_time).format('YYYY-MM-DD HH:mm:ss') }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" align="center">
        <template #default="scoped">
          <div>
            <el-button type="text" style="color: red">删除</el-button>
            <el-button type="text" @click="Edit(scoped.row)">编辑</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { getArticleList } from '@/api/article'
import { onMounted, ref } from 'vue'
import moment from 'moment'
import router from '@/router'

const Types: Record<string, string> = {
  '0': '正常',
  '-1': '已删除'
}

onMounted(async () => {
  const data = await getArticleList(params.value)
  articleList.value = data.Data.list.map((v: any) => {
    v.type = Types[v.type]
    return v
  })
  total.value = data.Data.total
})
const params = ref({
  offset: 0,
  size: 10,
  keyword: ''
})
const articleList = ref()
const total = ref()

function Edit(item) {
  router.push({
    path: '/edit',
    query: {
      id: item.id
    }
  })
}
</script>

<style scoped></style>
