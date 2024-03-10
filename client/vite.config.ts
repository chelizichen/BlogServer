import { fileURLToPath, URL } from 'node:url'

import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({mode})=>{
  const env = loadEnv(mode, process.cwd(), '')

  return {
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  base:env.BASE,
  server:{
    proxy: {
      '/blogserver': {
        target: 'http://localhost:8514/',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/blogserver/, '/blogserver/') // 不可以省略rewrite
      }
    }
  }
}})
