import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

export const constantRoutes = [

  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: 'Dashboard', icon: 'dashboard', affix: true }
      }
    ]
  },
  {
    path: '/profile',
    component: Layout,
    redirect: '/profile/index',
    hidden: true,
    children: [
      {
        path: 'index',
        component: () => import('@/views/profile/index'),
        name: 'Profile',
        meta: { title: 'Profile', icon: 'user', noCache: true }
      }
    ]
  },
  {
    path: '/application',
    component: Layout,
    redirect: '/application/index',
    hidden: false,
    children: [
      {
        path: '',
        component: () => import('@/views/application/index'),
        name: 'Applications',
        meta: { title: 'Applications', icon: 'tree', noCache: true }
      },
      {
        path: ':id(\\w+)',
        component: () => import('@/views/application/info'),
        name: 'Application',
        meta: { title: 'Application', noCache: true, activeMenu: '/application' },
        hidden: true
      }
    ]
  },
  {
    path: '/logic',
    component: Layout,
    hidden: false,
    children: [
      {
        path: ':id(\\w+)',
        component: () => import('@/views/logic/index'),
        name: 'Logic',
        meta: { title: 'Logic', noCache: true },
        hidden: true
      }
    ]
  },
  {
    path: '/codec',
    component: Layout,
    hidden: false,
    children: [
      {
        path: ':id(\\w+)',
        component: () => import('@/views/codec/index'),
        name: 'Codec',
        meta: { title: 'Codec', noCache: true },
        hidden: true
      }
    ]
  },
  {
    path: '/dp',
    component: Layout,
    hidden: false,
    children: [
      {
        path: ':id(\\w+)',
        component: () => import('@/views/datapoint/index'),
        name: 'DataPoint',
        meta: { title: 'DataPoint', noCache: true },
        hidden: true
      }
    ]
  },
  {
    path: '/connector',
    component: Layout,
    redirect: '/application/index',
    hidden: false,
    children: [
      {
        path: '',
        component: () => import('@/views/connector/index'),
        name: 'Connectors',
        meta: { title: 'Connectors', icon: 'list', noCache: true }
      },
      {
        path: ':id(\\w+)',
        component: () => import('@/views/connector/info'),
        name: 'Connector',
        meta: { title: 'Connector', noCache: true, activeMenu: '/connector' },
        hidden: true
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [

  /** when your routing map is too long, you can split it into small modules **/

  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'https://github.com/PanJiaChen/vue-element-admin',
        meta: { title: 'External Link', icon: 'link' }
      }
    ]
  }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
