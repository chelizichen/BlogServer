<script lang="ts">
export default {
  name: "layout-component",
};
</script>
<template>
  <el-container>
    <el-container>
      <el-aside width="200px" style="min-height: 100vh">
        <!-- <div class="app-bigger-size title">
            <el-icon style="color: rgb(207, 90, 124); font-size: 36px"><Help /></el-icon>
            Blog
          </div> -->
        <el-card style="height: 100vh">
          <el-menu
            class="el-menu-vertical-demo"
            active-text-color="rgb(207, 15, 124)"
            style="border: none"
          >
            <el-menu-item>
              <el-icon style="color: rgb(207, 90, 124); font-size: 24px">
                <Expand />
              </el-icon>
              <template #title>
                <div style="color: rgb(207, 90, 124); font-size: 18px; font-weight: 900">
                  SefficiencY
                </div>
              </template>
            </el-menu-item>
            <template v-for="(item, index) in routes" :key="index">
              <el-menu-item
                class="app-text-center"
                :index="item"
                @click="handleOpen(item)"
                v-if="
                  [item.show !== false, userLevel >= item.level].every((v) => v === true)
                "
              >
                <el-icon class="app-not-show"><Menu /></el-icon>
                <template #title>{{ item.name }}</template>
              </el-menu-item>
            </template>
          </el-menu>
        </el-card>
      </el-aside>
      <el-container>
        <el-header style="margin-bottom: 20px">
          <el-card>
            <div
              style="display: flex; justify-content: space-between; align-items: center"
            >
              <div>
                <el-input v-model="searchKeyword" style="width: 200px"></el-input>
              </div>
              <div>
                <el-menu
                  class="el-menu-demo"
                  mode="horizontal"
                  :ellipsis="false"
                  @select="handleSelect"
                >
                  <el-sub-menu index="2">
                    <template #title>
                      <el-avatar
                        src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                    /></template>
                    <el-menu-item index="2-0" v-if="!userLevel" v-level="0"
                      >登陆</el-menu-item
                    >
                    <el-menu-item index="2-1" v-if="userLevel" v-level="1"
                      >注销</el-menu-item
                    >
                    <el-menu-item index="2-2" v-if="userLevel" v-level="1"
                      >修改信息</el-menu-item
                    >
                  </el-sub-menu>
                </el-menu>
              </div>
            </div>
          </el-card>
        </el-header>
        <el-main>
          <el-card>
            <RouterView />
          </el-card>
        </el-main>
        <el-footer>
          <el-divider content-position="center">
            <div style="color: rgb(207, 15, 124); font-size: 16px">
              Copyright © 2023-2024
            </div>
          </el-divider>
          <el-divider content-position="center">
            <div style="color: rgb(207, 15, 124); font-size: 14px">
              SefficiencY Deployed On SimpCloud
            </div>
          </el-divider>
        </el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { routes } from "@/router";
import { ref, computed } from "vue";
import { RouterView, useRoute, useRouter } from "vue-router";
import { useUserStore } from "@/stores/counter";
import { Logout } from "@/api/login";
import { constants, localDel, localGet } from "@/utils/local";
const userStore = useUserStore();
const userLevel = computed(() => {
  console.log("userLevel", userStore.userInfo);
  return userStore.userInfo.level;
});

const [_, router] = [useRoute(), useRouter()];
const searchKeyword = ref();
function handleOpen(item) {
  router.push(item.path);
  console.log("item", item);
}

function handleSelect(item) {
  if (item == "2-1") {
    const token = localGet(constants.BLOG_TOKEN);
    Logout({ token }).then((res) => {
      router.push("/login");
      localDel(constants.BLOG_TOKEN);
    });
  }
  if (item == "2-0") {
    router.push("/login");
    localDel(constants.BLOG_TOKEN);
  }
  console.log(item);
}
</script>

<style scoped>
.title {
  color: rgb(207, 15, 124);
  display: flex;
  align-items: center;
  font-size: 30px;
  width: 200px;
  justify-content: center;
  cursor: pointer;
}
</style>
