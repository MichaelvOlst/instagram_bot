<template>
  <div>
    <Nav/>
    <div class="container">
        <!-- <router-view></router-view> -->
    </div>
  </div>
</template>

<script>
import Nav from './Nav'
export default {
  name: 'Dashboard',
  components: {
    Nav
  },

  data() {
    return {
      loggedIn: true,
    }
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