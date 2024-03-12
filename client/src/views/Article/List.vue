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
            <el-button type="text" @click="chooseColumn(scoped.row)">选择专栏</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <el-dialog
    v-model="centerDialogVisible"
    :title="`选择专栏 [${get(currentRow, 'column.title') || ''}]`"
    width="80%"
    center
  >
    <el-table
      :data="chooseColumnList"
      style="width: 100%"
      border
      highlight-current-row
      @current-change="handleCurrentChange"
      ref="singleTableRef"
    >
      <el-table-column prop="column.id" label="标识" align="center" />
      <el-table-column prop="column.title" label="标题" align="center" />
      <el-table-column prop="column.desc" label="描述" align="center" />
      <el-table-column prop="article_len" label="文章总数" align="center" />
    </el-table>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="centerDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitSaveIntoColumn"> 确定 </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { getArticleList } from '@/api/article'
import { onMounted, ref } from 'vue'
import moment from 'moment'
import router from '@/router'
import { getColumns, saveArticleInColumn } from '@/api/column'
import { ElTable } from 'element-plus'
import { get } from 'lodash'
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

const singleTableRef = ref<InstanceType<typeof ElTable>>()

const setCurrent = (row) => {
  console.log('row', row)
  singleTableRef.value!.setCurrentRow(row)
}

const chooseId = ref(0)
const chooseColumnList = ref([])
const centerDialogVisible = ref(false)
async function chooseColumn(row) {
  chooseId.value = row.id
  centerDialogVisible.value = true
  const data = await getColumns(params.value)
  const cRow = data.Data.list.find((v) => v.column.id == row.column_id)
  setCurrent(cRow)
  chooseColumnList.value = data.Data.list
}
const currentRow = ref({ column: {} })
const handleCurrentChange = (val) => {
  currentRow.value = val
}

async function submitSaveIntoColumn() {
  const body = {
    id: chooseId.value,
    cid: currentRow.value.column.id
  }
  const data = await saveArticleInColumn(body)
  centerDialogVisible.value = false
}
</script>

<style scoped></style>
