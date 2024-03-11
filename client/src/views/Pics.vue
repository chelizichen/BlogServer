<script lang="ts">
export default {
  name: 'home-component'
}
</script>
<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="name" label="名称" align="center" />
      <el-table-column prop="path" label="预览" align="center" >
        <template #default="scoped">
            <img :src="scoped.row.path" style="height: 200px;width:100%"/>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" align="center">
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { getArticleList, getPics } from '@/api/article'
import { onMounted, ref } from 'vue'
import moment from 'moment'
import router from '@/router'


onMounted(async () => {
  const data = await getPics(params.value)
  list.value = data.Data
})
const params = ref({
  offset: 0,
  size: 10,
  keyword: ''
})
const list = ref()
</script>

<style scoped></style>
