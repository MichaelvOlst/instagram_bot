import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import UsersView from '../views/Users.vue'

export function createRouter () {
  return new Router({
    mode: 'history',
    fallback: false,
    scrollBehavior: () => ({ y: 0 }),
    routes: [
      { path: '/dashboard/users', name: 'users', component: UsersView },
      { path: '/', redirect: '/dashboard/users' }
    ]
  })
}