<script lang="ts">
export default {
  name: 'Column-List'
}
</script>
<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="title" label="标题" align="center" />
      <el-table-column prop="content" label="内容" align="center" />
      <el-table-column prop="status" label="状态" align="center">
        <template #default="scoped">
          <div>{{ statusMap[scoped.row.status] || '其他' }}</div>
        </template>
      </el-table-column>
      <el-table-column
        prop="createTime"
        label="创建时间"
        align="center"
      />
      <el-table-column
        prop="startTime"
        label="开始时间"
        align="center"
      />
      <el-table-column
        prop="endTime"
        label="结束时间"
        align="center"
      />
      <el-table-column
        prop="realEndTime"
        label="实际结束时间"
        align="center"
      />
      <el-table-column prop="content" label="内容" align="center" />

      <el-table-column label="操作" width="180" align="center">
        <template #default="scoped">
          <el-button type="text" style="color: red" @click="DelColumnById(scoped.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { getEvents } from '@/api/event'
import { ElMessageBox, dayjs } from 'element-plus'
import moment from 'moment'
import { onMounted, ref } from 'vue'

const statusMap = {
  0: '已创建',
  1: '进行中',
  2: '已完成',
  3: '已延期'
}

onMounted(async () => {
  const data = await getEvents(params.value)
  list.value = data.Data.map(v=>{
    v.createTime =  !v.createTime  ? "--":moment(v.createTime).format('YYYY-MM-DD HH:mm:ss')
    v.realEndTime = !v.realEndTime  ? "--":moment(v.realEndTime).format('YYYY-MM-DD HH:mm:ss')
    v.endTime = !v.endTime  ? "--":moment(v.endTime).format('YYYY-MM-DD HH:mm:ss')
    v.startTime  = !v.startTime ? "--":moment(v.startTime).format('YYYY-MM-DD HH:mm:ss')
    return v
  })
})
const params = ref({
  offset: 0,
  size: 10,
  keyword: ''
})
const list = ref()
function DelColumnById(row: any) {
  ElMessageBox.prompt('Are you sure to delete this', 'Confirm', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputPlaceholder: 'input password'
  }).then(async ({ value }) => {
    if (value != '0504') {
      return false
    }
  })
}
</script>

<style scoped></style>
