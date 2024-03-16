import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useArticleStore = defineStore('article', () => {
  const articleTitle = ref('')
  const articleContent = ref('')
  function setArticle(t: string, c: string) {
    articleTitle.value = t
    articleContent.value = c
  }

  return { articleTitle, articleContent, setArticle }
})

export const useSearchStore = defineStore('search', () => {
  const searchKeyword = ref('')
  return { searchKeyword }
})

