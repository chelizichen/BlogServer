<script lang="ts">
export default {
  name: "home-component",
};
</script>
<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="name" label="名称" align="center" />
      <el-table-column prop="path" label="预览" align="center">
        <template #default="scoped">
          <img :src="scoped.row.path" style="height: 200px; width: 100%" />
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" align="center">
        <template #default="scoped">
          <el-button type="text" @click="deletePic(scoped.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { getArticleList, getPics, delPic } from "@/api/article";
import { ElMessage, ElMessageBox } from "element-plus";
import { onMounted, ref } from "vue";

onMounted(async () => {
  const data = await getPics(params.value);
  list.value = data.Data;
});
const params = ref({
  offset: 0,
  size: 10,
  keyword: "",
});
const list = ref();

function deletePic(row: any) {
  ElMessageBox.prompt("Are you sure to delete this img", "Confirm", {
    confirmButtonText: "OK",
    cancelButtonText: "Cancel",
    inputPlaceholder: "input password",
  }).then(async ({ value }) => {
    if (value != "0504") {
      return false;
    }
    const name = row.name;
    delPic({ imgPath: name });
    ElMessage({
        type: 'success',
        message: `Delete Success`
      })
  
  }).catch(()=>{
    ElMessage({
      type: 'info',
      message: 'Delete canceled'
    })
  });
}
</script>

<style scoped></style>
