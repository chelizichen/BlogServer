<script lang="ts">
export default {
  name: "login-vue",
};
</script>
<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { LogIn } from "@/api/login";
const router = useRouter();
const name = ref("");
const password = ref("");

async function saveToken() {
  const data = {
    name: name.value,
    password: password.value,
  };
  const ret = await LogIn(data);
  if (ret.Code) {
    ElMessage.error("Please enter a valid token.");
  } else {
    router.push("/effs");
    localStorage.setItem("blog_server_token", ret.Data.token);
  }
}
</script>
<template>
  <div class="body">
    <div class="container">
      <div class="hello">Welcome Blog</div>
      <el-form :inline="true">
        <el-form-item label="username">
          <el-input v-model="name" />
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="password" />
        </el-form-item>
      </el-form>
      <el-button type="primary" @click="saveToken()">Submit</el-button>
    </div>
  </div>
</template>

<style lang="less">
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
  background-color: #ffffff;
  .container {
    text-align: center;
    width: 25vw;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    .hello {
      margin-bottom: 10px;
      font-size: 26px;
      color: var(--el-button-hover-bg-color);
    }
  }
}
</style>
