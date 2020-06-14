<template>
  <div>

    {{ loggedIn }}

    {{ getUser }}

    <Dashboard v-if="loggedIn"/>
    <Login v-else />

    <Nav/>
    <!-- <router-view></router-view> -->

  </div>
</template>

<script>

import { mapGetters, mapActions } from 'vuex'


import Dashboard from './components/Dashboard'
import Login from './components/Login'
export default {
  name: 'App',
  components: {
    Dashboard, Login
  },
  mounted() {
    this.checkLogin()
  },
  computed: {
    ...mapGetters('user', {
      loggedIn: 'loggedIn',
      getUser: 'getUser',
    })
  },
  methods: {
    ...mapActions('user', [
      'checkLogin'
    ]),
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>