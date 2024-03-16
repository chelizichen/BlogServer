<script lang="ts">
export default {
  name: "Column-List",
};
</script>
<template>
  <div>
    <el-table :data="list" style="width: 100%" border>
      <el-table-column type="index" width="90" label="序号" align="center" />
      <el-table-column prop="title" label="标题" align="center" />
      <el-table-column prop="content" label="内容" align="center" />
      <el-table-column prop="status" label="状态" align="center">
        <template #default="scoped">
          <div>{{ statusMap[scoped.row.status] || "其他" }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="createTime" label="创建时间" align="center" />
      <el-table-column prop="startTime" label="开始时间" align="center" />
      <el-table-column prop="endTime" label="结束时间" align="center" />
      <el-table-column prop="realEndTime" label="实际结束时间" align="center" />
      <el-table-column prop="content" label="内容" align="center" />

      <el-table-column label="操作" width="180" align="center">
        <template #default="scoped">
          <el-button type="text" @click="EditForm(scoped.row)">编辑</el-button>
          <el-button type="text" style="color: red" @click="DelColumnById(scoped.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="editDataVisible" title="状态流转" center width="80%">
      <div style="display: flex">
        <div style="flex: 1; margin: 0 10px">
          <el-form :model="editData" label-width="auto" style="max-width: 600px">
            <el-form-item label="流转状态">
              <el-select
                v-model="editData.status"
                placeholder="Select"
                size="large"
                style="width: 240px"
              >
                <el-option
                  v-for="item in statusArr"
                  :key="item.key"
                  :label="item.key"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="结束时间" v-if="['2', '3'].includes(editData.status)">
              <el-col :span="11">
                <el-date-picker
                  v-model="editData.realEndTime"
                  type="date"
                  placeholder="Pick a date"
                  style="width: 100%"
                />
              </el-col>
            </el-form-item>
            <el-form-item label="实际支付" v-if="['2', '3'].includes(editData.status)">
              <el-input-number v-model="editData.realEventPay" />
            </el-form-item>
            <el-form-item label="评价">
              <div style="border: 1px solid #ccc">
                <Toolbar
                  style="border-bottom: 1px solid #ccc"
                  :editor="editorRef"
                  :defaultConfig="toolbarConfig"
                  :mode="mode"
                />
                <Editor
                  style="height: 300px; overflow-y: hidden"
                  v-model="editData.content"
                  :defaultConfig="editorConfig"
                  :mode="mode"
                  @onCreated="handleCreated"
                />
              </div>
              <!-- <el-input v-model="editData.content" type="textarea" /> -->
            </el-form-item>
          </el-form>
        </div>
        <div style="flex: 1; margin: 0 10px">
          <el-table :data="commentList" border style="width: 100%" height="250">
            <el-table-column prop="content" label="内容" width="180" />
            <el-table-column prop="createTime" label="创建时间" width="180" />
            <el-table-column prop="status" label="状态">
              <template #default="scoped">
                <div>{{ statusMap[scoped.row.status] || "其他" }}</div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button
            style="background-color: rgb(207, 15, 124); border: none"
            type="primary"
            @click="SubmitEditForm"
            >Create</el-button
          >
          <el-button @click="editDataVisible = false">Cancel</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { changeStatus, getCommentsByEventId, getEvents } from "@/api/event";
import { ElMessage, ElMessageBox, dayjs } from "element-plus";
import { cloneDeep } from "lodash";
import moment from "moment";
import { onMounted, ref, shallowRef } from "vue";
import { toolbarConfig, editorConfig } from "@/utils/editor";
// @ts-ignore
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import "@wangeditor/editor/dist/css/style.css"; // 引入 css

const mode = "simple";
const editorRef = shallowRef();
const handleCreated = (editor) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};
const editData = ref();
const editDataVisible = ref(false);
const commentList = ref([
  {
    date: "2016-05-03",
    name: "Tom",
    address: "No. 189, Grove St, Los Angeles",
  },
  {
    date: "2016-05-02",
    name: "Tom",
    address: "No. 189, Grove St, Los Angeles",
  },
  {
    date: "2016-05-04",
    name: "Tom",
    address: "No. 189, Grove St, Los Angeles",
  },
  {
    date: "2016-05-01",
    name: "Tom",
    address: "No. 189, Grove St, Los Angeles",
  },
]);
async function SubmitEditForm() {
  const body = editData.value;
  const resp = await changeStatus(body);
  if (resp.Code) {
    ElMessage.error(resp.Message);
    editDataVisible.value = false;
  }
  ElMessage.success("评价成功");
  editDataVisible.value = false;
}
function EditForm(row) {
  getCommentsByEventId({ id: row.id }).then((res) => {
    commentList.value = res.Data;
  });
  editDataVisible.value = true;
  editData.value = cloneDeep(row);
  editData.value.status = String(editData.value.status);
  editData.value.content = "";
  console.log("editData", editData.value);
}

const statusMap = {
  "0": "已创建",
  "1": "进行中",
  "2": "已完成",
  "3": "已延期",
  "-1": "无法完成",
};
const statusArr = Object.entries(statusMap).reduce((pre, curr) => {
  const obj = {
    key: curr[1],
    value: curr[0],
  };
  pre.push(obj);
  return pre;
}, []);

onMounted(async () => {
  const data = await getEvents(params.value);
  list.value = data.Data.map((v) => {
    v.createTime = !v.createTime
      ? "--"
      : moment(v.createTime).format("YYYY-MM-DD HH:mm:ss");
    v.realEndTime = !v.realEndTime
      ? "--"
      : moment(v.realEndTime).format("YYYY-MM-DD HH:mm:ss");
    v.endTime = !v.endTime ? "--" : moment(v.endTime).format("YYYY-MM-DD HH:mm:ss");
    v.startTime = !v.startTime ? "--" : moment(v.startTime).format("YYYY-MM-DD HH:mm:ss");
    return v;
  });
});
const params = ref({
  offset: 0,
  size: 10,
  keyword: "",
});
const list = ref();
function DelColumnById(row: any) {
  ElMessageBox.prompt("Are you sure to delete this", "Confirm", {
    confirmButtonText: "OK",
    cancelButtonText: "Cancel",
    inputPlaceholder: "input password",
  }).then(async ({ value }) => {
    if (value != "0504") {
      return false;
    }
  });
}
</script>

<style scoped></style>
