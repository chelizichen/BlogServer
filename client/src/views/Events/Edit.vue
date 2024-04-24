<template>
  <el-form :model="form" label-width="auto" style="max-width: 600px">
    <el-form-item label="StartTime">
      <el-col :span="11">
        <el-date-picker
          v-model="form.startTime"
          type="date"
          placeholder="Pick a date"
          style="width: 100%"
        />
      </el-col>
    </el-form-item>
    <el-form-item label="EndTime">
      <el-col :span="11">
        <el-date-picker
          v-model="form.endTime"
          type="date"
          placeholder="Pick a date"
          style="width: 100%"
        />
      </el-col>
    </el-form-item>
    <el-form-item label="Title">
      <el-input v-model="form.title" />
    </el-form-item>
    <el-form-item label="Content">
      <el-input v-model="form.content" type="textarea" />
    </el-form-item>
    <el-form-item label="Pay">
      <el-input-number v-model="form.eventPay" />
    </el-form-item>
    <el-form-item>
      <el-button
        style="background-color: rgb(207, 15, 124); border: none"
        type="primary"
        @click="onSubmit"
        >Create</el-button
      >
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { saveEvent } from "@/api/event";
import router from "@/router";
import { ElMessage } from "element-plus";
import { reactive } from "vue";

// do not use same name with ref
const form = reactive({
  startTime: "",
  endTime: "",
  title: "",
  content: "",
  eventPay: 0,
});

const onSubmit = () => {
  saveEvent(form)
    .then((data) => {
      console.log("data", data);
      if (data.code) {
        ElMessage.error(data.message);
        return;
      }
      router.push("/effs");
    })
    .catch((e) => {
      console.log("e", e);
      ElMessage.error(e.message);
    });
};
</script>
