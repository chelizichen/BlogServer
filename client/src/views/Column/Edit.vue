<script lang="ts">
export default {
  name: 'edit-component'
}
</script>
<template>
  <div class="about">
    <el-form :model="form" label-width="auto">
      <el-form-item label="标题" style="width: 400px">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="描述" style="width: 400px">
        <el-input v-model="form.desc" />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          style="background-color: rgb(207, 15, 124); border: none"
          @click="onSubmit"
          >Create</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped></style>

<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import {  onMounted, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { saveColumn } from '@/api/column'

const [route, router] = [useRoute(), useRouter()]

const form = reactive({
  title: '',
  desc:'',
})
async function GetById() {
  const id = route.query.id
  if (!id) {
    return
  }
  const query = { id }
}
// 模拟 ajax 异步获取内容
onMounted(async () => {
  await GetById()
  // setTimeout(() => {

})

async function onSubmit(){
    const body = form
    const data = await saveColumn(body)
}

</script>
