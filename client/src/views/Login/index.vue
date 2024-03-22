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
import { useUserStore } from "@/stores/counter";
import { localSet, constants } from "@/utils/local";
const router = useRouter();
const name = ref("");
const password = ref("");

const userStore = useUserStore();
async function saveToken() {
  const data = {
    name: name.value,
    password: password.value,
  };
  const ret = await LogIn(data);
  if (ret.Code) {
    ElMessage.error("Please enter a valid token.");
  } else {
    localSet(constants.BLOG_TOKEN, `${name.value}|${ret.Data.token}`);
    localSet(constants.USER_NAME, `${name.value}`);
    localSet(constants.ENCRYPT_PASSWORD, `${ret.Data.password}`);
    userStore.userInfo = ret.Data;
    router.push("/home");
  }
}
</script>
<template>
  <div class="body">
    <div class="container">
      <div class="hello">START</div>
      <el-form>
        <el-form-item label="username">
          <el-input v-model="name" />
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="password" type="password" />
        </el-form-item>
        <el-button
          type="primary"
          style="background-color: rgb(207, 15, 124); border: none"
          @click="saveToken()"
          >Submit</el-button
        >
      </el-form>
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
    width: 20vw;
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
