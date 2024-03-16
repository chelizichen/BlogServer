<template>
  <el-dialog
    v-model="props.centerDialogVisible"
    :title="`选择专栏 [${get(currentRow, 'column.title') || ''}]`"
    width="80%"
    center
    @close="emits('closeDialogVisible')"
  >
    <el-table
      :data="props.chooseColumnList"
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
        <el-button @click="emits('closeDialogVisible')">取消</el-button>
        <el-button type="primary" @click="submitSaveIntoColumn"> 确定 </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { saveArticleInColumn } from '@/api/column';
import { ElMessage, type ElTable } from 'element-plus';
import { nextTick, ref } from 'vue';
import { get } from 'lodash'

const props = defineProps<{
    chooseId:number | string,
    chooseColumnList:any[],
    centerDialogVisible:boolean
}>()
const emits = defineEmits(["closeDialogVisible"])

const currentRow = ref({ column: {id:''} })
const handleCurrentChange = (val:any) => {
  currentRow.value = val
}
async function submitSaveIntoColumn() {
  const body = {
    id: props.chooseId,
    cid: currentRow.value.column.id
  }
  const data = await saveArticleInColumn(body)
  if(data.Code){
    ElMessage.error(data.Message)
    return
  }
  ElMessage.success("设置成功")
  emits('closeDialogVisible')
}
const singleTableRef = ref<InstanceType<typeof ElTable>>()

const setCurrent = (id) => {
    console.log('chooseColumnList',props.chooseColumnList);
    console.log('chooseId',props.chooseId);
    const cRow = props.chooseColumnList.find((v) => v.column.id == id)
    if(!cRow){
      return
    }
    singleTableRef.value!.setCurrentRow(cRow)
}

defineExpose({
  setCurrent
})

</script>

<style scoped>

</style>
