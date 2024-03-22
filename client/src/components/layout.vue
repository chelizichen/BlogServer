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
              <el-icon style="color: rgb(207, 90, 124); font-size: 36px"
                ><Help
              /></el-icon>
              <template #title>
                <div style="color: rgb(207, 90, 124); font-size: 18px">Blog</div>
              </template>
            </el-menu-item>
            <template v-for="(item, index) in routes" :key="index">
              <el-menu-item
                class="app-text-center"
                :index="item"
                @click="handleOpen(item)"
                v-if="[item.show !== false].every((v) => v === true)"
              >
                <el-icon class="app-not-show"><Menu /></el-icon>
                <template #title>{{ item.name }}</template>
              </el-menu-item>
            </template>
          </el-menu>
        </el-card>
      </el-aside>
      <el-container>
        <el-header>
          <el-card>
            <el-input v-model="searchKeyword" style="width: 200px"></el-input>
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
              BlogServer Started From SimpCloud
            </div>
          </el-divider>
        </el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { routes } from "@/router";
import { ref } from "vue";
import { RouterView, useRoute, useRouter } from "vue-router";
import { useUserStore } from "@/stores/counter";
const userStore = useUserStore();
const [_, router] = [useRoute(), useRouter()];
const searchKeyword = ref();
function handleOpen(item) {
  router.push(item.path);
  console.log("item", item);
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
