import { fileURLToPath, URL } from 'node:url'
import { chunkSplitPlugin } from 'vite-plugin-chunk-split'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [
      vue(),
      chunkSplitPlugin({
        strategy: 'default',
        customSplitting: {
          'wangeditor-chunk': [/wangeditor/, /@wangeditor\/editor/, /@wangeditor\/editor-for-vue/],
          elements: [/element-plus/, /@element-plus\/icons-vue/]
        }
      })
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    base: env.BASE,
    server: {
      proxy: {
        '/blogserver': {
          target: 'http://localhost:8514/',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/blogserver/, '/blogserver/') // 不可以省略rewrite
        },
        '/ltsnodeserver': {
          target: 'http://localhost:18517/',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/ltsnodeserver/, '/ltsnodeserver/') // 不可以省略rewrite
        }
      }
    },
    build: {
      // rollupOptions:{
      //   output:{
      //     manualChunks:{
      //       "lodash-chunk":["lodash"],
      //       "wangeditor-chunk":["wangeditor","@wangeditor/editor","@wangeditor/editor-for-vue"]
      //     }
      //   }
      // }
    }
  }
})
