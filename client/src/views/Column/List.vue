<script lang="ts">
export default {
  name:"Column-List"
}
</script>
<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="column.id" label="标识" align="center" />
      <el-table-column prop="column.title" label="标题" align="center" />
      <el-table-column prop="column.desc" label="描述" align="center" />
      <el-table-column prop="article_len" label="文章总数" align="center" />
      <el-table-column label="操作" width="180" align="center">
      <template #default="scoped">
        <el-button type="text" style="color: red;" @click="DelColumnById(scoped.row)">删除</el-button>
      </template>
       </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { delColumnById, getColumns } from '@/api/column'
import { ElMessageBox } from 'element-plus';
import { onMounted, ref } from 'vue'
onMounted(async () => {
  const data = await getColumns(params.value)
  list.value = data.Data.list
})
const params = ref({
  offset: 0,
  size: 10,
  keyword: ''
})
const list = ref()
function DelColumnById(row:any){
  ElMessageBox.prompt("Are you sure to delete this", "Confirm", {
    confirmButtonText: "OK",
    cancelButtonText: "Cancel",
    inputPlaceholder: "input password",
  }).then(async ({ value }) => {
    if (value != "0504") {
      return false;
    }
    delColumnById({id:row.column.id})
  })
}
</script>

<style scoped></style>
