import Vue from 'vue';
import Vuex from 'vuex'
import App from './App'
import store from './store'
import { createRouter } from './router'

import { BootstrapVue } from 'bootstrap-vue'
Vue.use(BootstrapVue)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false



const router = createRouter()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')