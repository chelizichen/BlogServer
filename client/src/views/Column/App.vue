<template>
  <div class="app">
    <el-card>
      <el-collapse v-model="activeId" accordion @change="handleChange">
        <el-collapse-item
          v-for="item in articles"
          :title="item.title"
          :name="item.id"
          :key="item.id"
        >
          <div v-if="activeId == item.id">
            <div v-html="content" ref="articleContent"></div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { getArticle } from "@/api/article";
import { getColumnDetail } from "@/api/column";
import { onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import "highlight.js/styles/stackoverflow-light.css";
import "highlight.js/lib/common";
import hljs from "highlight.js";
import "@wangeditor/editor/dist/css/style.css";
const [route, router] = [useRoute(), useRouter()];

const articles = ref();
const activeId = ref();
const content = ref();

const articleContent = ref<HTMLDivElement[]>();

async function GetById() {
  const id = route.query.id as string;
  if (!id) {
    return;
  }
  const data = await getColumnDetail({ id });
  articles.value = data.Data.articles;
}

const handleChange = async (val: string) => {
  if (!val) {
    return;
  }
  const res = await getArticle({ id: val });
  content.value = res.Data.content;
  activeId.value = val;
};

onMounted(() => {
  GetById();
});

watch(activeId, function (newVal) {
  if (!newVal) {
    return;
  }
  setTimeout(() => {
    console.log("newVal", newVal);
    console.log("articleContent.value", articleContent.value);
    const blocks = articleContent.value![0].querySelectorAll("pre code");
    blocks.forEach((block: any) => {
      block.setAttribute("style", "margin-top: 8px;");
      hljs.highlightBlock(block);
    });
  }, 50);
});
</script>

<style scoped>
.app >>> img {
  width: 100%;
  height: 100%;
}
</style>
