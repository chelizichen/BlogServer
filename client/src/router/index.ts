import {  LoginByCache } from '@/api/login'
import { useUserStore } from '@/stores/counter'
import { constants, localGet, localSet } from '@/utils/local'
import { constant } from 'lodash'
import { createRouter, createWebHistory } from 'vue-router'

export const routes = [
  {
    path: '/home',
    name: '文章列表',
    component: () => import('../views/Article/List.vue'),
    level:4
  },
  {
    path: '/edit',
    name: '编辑文章',
    component: () => import('../views/Article/Edit.vue'),
    level:5
  },
  {
    path: '/Pics',
    name: '图片审核',
    component: () => import('../views/Pics.vue'),
    level:5
  },
  {
    path: '/columns',
    name: '专栏列表',
    component: () => import('../views/Column/List.vue'),
    level:4
  },
  {
    path: '/edit_column',
    name: '编辑专栏',
    component: () => import('../views/Column/Edit.vue'),
    level:5
  },
  {
    path: '/effs',
    name: '事件列表',
    component: () => import('../views/Events/List.vue'),
    level:4
  },
  {
    path: '/create_event',
    name: '创建事件',
    component: () => import('../views/Events/Edit.vue'),
    level:5
  },
  {
    path: '/phone',
    name: '移动端',
    component: () => import('../views/Phone.vue')
  },
  {
    path: '/Detail',
    name: 'detail',
    component: () => import('../views/Detail.vue'),
    show: false
  },
  {
    path: '/app_column',
    name: 'app_column',
    component: () => import('../views/Column/App.vue'),
    show: false
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/Login/index.vue'),
    show: false
  },
  {
    path: '/upload',
    name: '服务上传',
    component: () => import('../views/Upload/index.vue'),
    level:0
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// router.beforeEach((to,from,next)=>{
//   const IsApp = isApp()
//   console.log('isApp',IsApp);
//   if(!IsApp){
//     next(to)
//   }else{
//     next('/app')
//   }
// })

const whileList = ["/login","/upload"]

router.beforeEach(async (to,from,next)=>{
  const tkn = localGet(constants.BLOG_TOKEN)
  const userStore = useUserStore()

  if(whileList.includes(to.path)){
    console.log('to',to.path);
    return next()
  }else if (tkn && (!userStore.userInfo.name && !userStore.userInfo.password && !userStore.userInfo.token)){
    const data = await LoginByCache({
      name:'',
      password:''
    })
    if(!data.Data){
      return next("/login")
    }
    userStore.userInfo = data.Data
    next()
  }else if(tkn){
    next()
  }else{
    next({path:"/login"})
  }
})

export default router
