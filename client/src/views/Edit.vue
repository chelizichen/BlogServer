<script lang="ts">
export default {
  name: 'edit-component'
}
</script>
<template>
  <div class="about" v-loading="loading">
    <el-form :model="form" label-width="auto">
      <el-form-item label="标题" style="width: 400px">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="内容">
        <div style="border: 1px solid #ccc">
          <Toolbar
            style="border-bottom: 1px solid #ccc"
            :editor="editorRef"
            :defaultConfig="toolbarConfig"
            :mode="mode"
          />
          <Editor
            style="height: 500px; overflow-y: hidden"
            v-model="valueHtml"
            :defaultConfig="editorConfig"
            :mode="mode"
            @onCreated="handleCreated"
          />
        </div>
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
import { onBeforeUnmount, ref, shallowRef, onMounted, reactive } from 'vue'
// @ts-ignore
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { getArticle, saveArticle } from '@/api/article'
import { ElMessage } from 'element-plus'
import { useRoute, useRouter } from 'vue-router'

const [route, router] = [useRoute(), useRouter()]

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()
const mode = 'simple'
// 内容 HTML
const valueHtml = ref('')
const form = reactive({
  title: ''
})
async function GetById() {
  const id = route.query.id
  if (!id) {
    return
  }
  const query = { id }
  const res = await getArticle(query)
  valueHtml.value = res.Data.content
  form.title = res.Data.title
  console.log('res')
}
// 模拟 ajax 异步获取内容
onMounted(async () => {
  await GetById()
  // setTimeout(() => {
  //   valueHtml.value = "<p>模拟 Ajax 异步设置内容</p>";
  // }, 1500);
})

const toolbarConfig = {}
const editorConfig = {
  placeholder: '请输入内容...',
  MENU_CONF: {
    uploadImage: {
      server: '/blogserver/blogImg',
      // file 文件的字段名
      fieldName: 'file',
      maxFileSize: 1 * 1024 * 1024
    }
  }
}

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}
const loading = ref(false)
const onSubmit = async () => {
  loading.value = true
  if (!form.title) {
    ElMessage.error(`error! miss title`)
    return (loading.value = false)
  }
  const body = {
    id: route.query.id ? Number(route.query.id) : 0,
    title: form.title,
    content: valueHtml.value
  }
  const resp = await saveArticle(body)
  if (resp.Code) {
    ElMessage.error(`error! ${resp.Message}`)
    return (loading.value = false)
  }
  return (loading.value = false)
}
</script>
