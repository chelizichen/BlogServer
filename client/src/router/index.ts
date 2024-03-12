import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/Article/List.vue')
    },
    {
      path: '/edit',
      name: 'edit',
      component: () => import('../views/Article/Edit.vue')
    },
    {
      path: '/phone',
      name: 'phone',
      component: () => import('../views/Phone.vue')
    },
    {
      path: '/Detail',
      name: 'detail',
      component: () => import('../views/Detail.vue')
    },
    {
      path: '/Pics',
      name: 'pics',
      component: () => import('../views/Pics.vue')
    },
    {
      path: '/columns',
      name: 'columns',
      component: () => import('../views/Column/List.vue')
    },
    {
      path: '/edit_column',
      name: 'edit_column',
      component: () => import('../views/Column/Edit.vue')
    },
    {
      path: '/app_column',
      name: 'app_column',
      component: () => import('../views/Column/App.vue')
    }
  ]
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

export default router
