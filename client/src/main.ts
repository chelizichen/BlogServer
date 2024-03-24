import 'element-plus/dist/index.css'
import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { level } from './directive/level'

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.use(ElementPlus, { zIndex: 3000 })
app.use(createPinia())
app.directive(level.name,{
  mounted:level.hooks.mounted,
})
app.use(router)
app.mount('#app')
