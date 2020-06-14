import Vue from 'vue';
import Vuex from 'vuex'
import App from './App'
import store from './store'


// import VueRouter from 'vue-router'

// Vue.use(VueRouter)

Vue.use(Vuex)

Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App),
}).$mount('#app')