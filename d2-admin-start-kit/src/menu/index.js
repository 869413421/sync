import { uniqueId } from 'lodash'

/**
 * @description 给菜单数据补充上 path 字段
 * @description https://github.com/d2-projects/d2-admin/issues/209
 * @param {Array} menu 原始的菜单数据
 */
function supplementPath (menu) {
  return menu.map(e => ({
    ...e,
    path: e.path || uniqueId('d2-menu-empty-'),
    ...e.children ? {
      children: supplementPath(e.children)
    } : {}
  }))
}

export const menuHeader = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '权限管理',
    icon: 'folder-o',
    children: [
      { path: '/user', title: '用户' },
      { path: '/role', title: '角色' },
      { path: '/permission', title: '权限' }
    ]
  }
])

export const menuAside = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '权限管理',
    icon: 'folder-o',
    children: [
      { path: '/user', title: '用户' },
      { path: '/role', title: '角色' },
      { path: '/permission', title: '权限' }
    ]
  }
])
