<script lang="ts">
export default {
  name: "home-component",
};
</script>
<template>
  <div>
    <div style="margin: 10px 10px 10px 0; display: flex">
      <el-input v-model="params.keyword" style="width: 200px; margin: 0 10px 0 0" />
      <el-button @click="getList">搜索</el-button>
    </div>
    <el-table :data="articleList" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="id" label="编号" width="90" align="center" />
      <el-table-column prop="title" label="标题" width="180" align="center" />
      <el-table-column label="内容" align="center">
        <template #default="scoped">
          <div
            class="imgs"
            style="height: 32px; overflow: hidden"
            v-html="scoped.row.content"
          ></div>
        </template>
      </el-table-column>
      <el-table-column prop="type" label="类型" align="center" />
      <el-table-column prop="create_time" label="创建时间" width="180" align="center">
        <template #default="scoped">
          {{ moment(scoped.row.create_time).format("YYYY-MM-DD HH:mm:ss") }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" align="center">
        <template #default="scoped">
          <div>
            <el-button type="text" style="color: red" v-level="5">删除</el-button>
            <el-button type="text" v-level="5" @click="Edit(scoped.row)">编辑</el-button>
            <el-button type="text" v-level="5" @click="chooseColumn(scoped.row)"
              >选择专栏</el-button
            >
          </div>
        </template>
      </el-table-column>
    </el-table>
    <div style="float: right; margin: 10px">
      <el-pagination
        :current-page="params.page + 1"
        :page-size="params.size"
        small
        background
        layout="prev, pager, next"
        :total="total"
        @current-change="handleCurrentChange"
      />
    </div>
    <chooseColumnComponent
      :chooseId="chooseId"
      :chooseColumnList="chooseColumnList"
      :centerDialogVisible="centerDialogVisible"
      @closeDialogVisible="() => (centerDialogVisible = false)"
      ref="childComp"
    ></chooseColumnComponent>
  </div>
</template>

<script setup lang="ts">
import { getArticleList } from "@/api/article";
import { computed, nextTick, onMounted, ref, watch } from "vue";
import moment from "moment";
import router from "@/router";
import { getColumns } from "@/api/column";
import { ElTable } from "element-plus";
import chooseColumnComponent from "./ChooseColumn.vue";
import { useSearchStore } from "@/stores/counter";

const Types: Record<string, string> = {
  "0": "正常",
  "-1": "已删除",
};
async function getList() {
  const data = await getArticleList(pagination.value);
  articleList.value = data.Data.list.map((v: any) => {
    v.type = Types[v.type];
    return v;
  });
  total.value = data.Data.total;
}

onMounted(() => {
  getList();
});

const params = ref({
  offset: 0,
  size: 10,
  keyword: "",
  page: 0,
});

const pagination = computed(() => {
  return {
    offset: params.value.page * params.value.size,
    size: params.value.size,
    keyword: params.value.keyword,
  };
});
const articleList = ref();
const total = ref();

function handleCurrentChange(page) {
  params.value.page = page - 1;
  getList();
}

function Edit(item) {
  router.push({
    path: "/edit",
    query: {
      id: item.id,
    },
  });
}

const chooseId = ref(0);
const chooseColumnList = ref([]);
const centerDialogVisible = ref(false);
const childComp = ref<any>(null);
async function chooseColumn(row: any) {
  centerDialogVisible.value = true;
  chooseId.value = row.id;
  const data = await getColumns(params.value);
  chooseColumnList.value = data.Data.list;
  nextTick(() => {
    childComp.value.setCurrent(row.column_id);
  });
}
</script>

<style scoped>
.imgs >>> img {
  display: none;
}
</style>
