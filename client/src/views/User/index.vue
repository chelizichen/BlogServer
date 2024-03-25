<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="name" label="名称" align="center" />
      <el-table-column prop="level" label="等级" align="center" />
      <el-table-column prop="create_time" label="创建时间" align="center" />
      <el-table-column label="操作" align="center">
        <template #default="scoped">
          <el-button type="text" @click="changeLevelVisible(scoped.row)"
            >修改状态</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="visibleState.changeLevelVisible" title="修改等级" width="500">
      <el-form :model="TableRowItem">
        <el-form-item label="Level">
          <el-input v-model.number="TableRowItem.level" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="visibleState.changeLevelVisible = false">Cancel</el-button>
          <el-button type="primary" @click="submitChangeLevel"> Confirm </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { getUserList, changeUserLevel } from "@/api/user";
import { dayjs } from "element-plus";
const list = ref();

const visibleState = reactive({
  changeLevelVisible: false,
});

async function init() {
  const data = await getUserList(params.value);
  list.value = data.Data.list
    .filter((v) => v)
    .map((v) => {
      v.create_time = dayjs(v.create_time).format("YYYY-MM-DD HH:mm:ss");
      return v;
    });
}

onMounted(async () => {
  await init();
});
const params = ref({
  offset: 0,
  size: 10,
  keyword: "",
});
const TableRowItem = ref({
  id: "",
  level: 0,
});
function changeLevelVisible(item) {
  visibleState.changeLevelVisible = true;
  console.log(item);
  TableRowItem.value.id = item.id;
  TableRowItem.value.level = item.level;
}
async function submitChangeLevel() {
  await changeUserLevel(TableRowItem.value);
  visibleState.changeLevelVisible = false;
  await init();
}
</script>

<style scoped></style>
