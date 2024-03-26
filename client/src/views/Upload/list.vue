<template>
  <el-table :data="list" style="width: 100%" border>
    <el-table-column type="index" width="90" label="Index" align="center" />
    <el-table-column prop="ID" label="Primary" align="center" />
    <el-table-column prop="Hash" label="Hash" align="center" />
    <el-table-column prop="ServerName" label="ServerName" align="center" />
    <el-table-column prop="SpendTime" label="SpendTime" align="center" />
    <el-table-column prop="create_time" label="createTime" align="center" />
  </el-table>
</template>

<script setup lang="ts">
import API from "@/api/upload";
import { onMounted, ref } from "vue";

const list = ref([]);
const keyword = ref("");
onMounted(async () => {
  const data = await API.getUploadInfoList(keyword.value);
  list.value = data.Data.list;
});
</script>
