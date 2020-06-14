<template>   
    <nav class="navbar">
        <div class="container">
            <div class="navbar-brand">
                <a class="navbar-item" href="/"><i class="fab fa-instagram"></i>&nbsp;Instagram Bot</a>
            </div>
            <div id="navbarMenu" class="navbar-menu">
                <div class="navbar-end">
                    <a class="navbar-item is-active">Dashboard</a>
                    <a class="navbar-item" @click.prevent="logout()">Logout</a>
                </div>
            </div>
        </div>
    </nav>
</template>


<script>
import { mapGetters } from 'vuex'

export default {
    name: 'Nav',
   
    methods: {
        async logout() {
            try {
                await this.$store.dispatch('auth/logout')
            } catch(e) {
                console.log(e)
            }
        },
    },
     computed: {
        ...mapGetters('auth', {
        getUser: 'getUser',
        })
    },
}
</script>


