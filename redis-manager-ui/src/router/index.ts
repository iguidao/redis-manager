import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'
const routes = [
    {
        path: '/',
        redirect: '/login'
      },
      {
        path: '/login',
        name: 'login',
        meta: {
          title: '登录'
        },
        component: () => import('../pages/login/Index.vue')
      },
  
      {
        path: '/home',
        meta: {
          title: 'Redis Manager'
        },
        component: () => import('../pages/home/Index.vue'),
        redirect: '/dashboard',
        children: [{
            path: '/dashboard',
            meta: {
              title: '概览'
            },
            component: () => import('../pages/dashboard/Index.vue')
          },
          {
            path: '/user/index',
            meta: {
              title: '用户管理'
            },
            component: () => import('../pages/user/Index.vue'),
          },
          {
            path: '/aliredis/index',
            meta: {
              title: '阿里redis'
            },
            component: () => import('../pages/redis/aliredis/Index.vue'),
          },
          {
            path: '/redis/index',
            meta: {
              title: '自建cluster'
            },
            component: () => import('../pages/redis/cluster/Index.vue'),
          },
          {
            path: '/codis/index',
            meta: {
              title: '自建codis'
            },
            component: () => import('../pages/redis/codis/Index.vue'),
          },
          {
            path: '/txredis/index',
            meta: {
              title: '腾讯redis'
            },
            component: () => import('../pages/redis/txredis/Index.vue'),
          },
          {
            path: '/command/query',
            meta: {
              title: '查询数据'
            },
            component: () => import('../pages/command/Index.vue'),
          },
          {
            path: '/history/index',
            meta: {
              title: '历史记录'
            },
            component: () => import('../pages/history/Index.vue'),
          },
          {
            path: '/setting/index',
            meta: {
              title: '系统设置'
            },
            component: () => import('../pages/setting/Index.vue'),
          },
          {
            path: '/setting/rule',
            meta: {
              title: '权限配置'
            },
            component: () => import('../pages/setting/Rule.vue'),
          },
        ]
    },
]
const router = createRouter({
    history: createWebHistory(),
    // history: createWebHashHistory(),
    routes
})
// 挂载路由导航守卫：to表示将要访问的路径，from表示从哪里来，next是下一个要做的操作
router.beforeEach((to, from, next) => {
    // 修改页面 title
    // if (to.meta.title) {
    //   document.title = 'Redis-Manager平台 - ' + to.meta.title
    // }
    // 放行登录页面
    if (to.path === '/login') {
      return next()
    }
    // 获取token
    const token = sessionStorage.getItem('Authorization')
    if (!token) {
      return next('/login')
    } else {
      next()
    }
    return next()
  })
  
// 导出路由
export default router