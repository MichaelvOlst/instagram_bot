<template>
  <div>
    <Header/>
    <main>
      <div class="container">
        <div class="columns">
          <div class="column is-2">
            <Sidebar/>
          </div>
          <div class="column"><router-view></router-view></div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import Header from './Header'
import Sidebar from './Sidebar'

export default {
  name: 'Dashboard',
  components: {
    Header, Sidebar
  },

  data() {
    return {
      loggedIn: true,
    }
  }
}
</script>

<style>
  main {
    /* background: #eee; */
    padding-top: 3em;
    min-height: 100vh;
  }
</style>