import Vue from 'vue';
import Vuex from 'vuex'
import App from './App'
import store from './store'
import { createRouter } from './router'

Vue.config.productionTip = false

const router = createRouter()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')