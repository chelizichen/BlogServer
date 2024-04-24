<script lang="ts">
export default {
  name: "detail-component",
};
</script>
<template>
  <div>
    <el-card class="card">
      <template #header>{{ ArticleStore.articleTitle }}</template>
      <div v-html="ArticleStore.articleContent" ref="articleContent"></div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { getArticle } from "@/api/article";
import { useArticleStore } from "@/stores/counter";
import { onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import "highlight.js/styles/stackoverflow-light.css";
import "highlight.js/lib/common";
import hljs from "highlight.js";
import "@wangeditor/editor/dist/css/style.css";

const ArticleStore = useArticleStore();
const route = useRoute();

const articleContent = ref<HTMLDivElement>();
onMounted(async () => {
  if (!ArticleStore.articleContent && !ArticleStore.articleTitle) {
    const id = route.query.id;
    if (!id) {
      return;
    }
    const query = { id };
    const res = await getArticle(query);
    ArticleStore.setArticle(res.data.title, res.data.content);
  }
});

watch(articleContent, function (newVal) {
  setTimeout(() => {
    const blocks = newVal!.querySelectorAll("pre code");
    blocks.forEach((block: any) => {
      block.setAttribute("style", "margin-top: 8px;");
      hljs.highlightBlock(block);
    });
  }, 50);
});
</script>

<style scoped>
.card >>> img {
  width: 100%;
  height: 100%;
}
</style>
