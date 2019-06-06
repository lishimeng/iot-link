import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/* Router Modules */

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
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
