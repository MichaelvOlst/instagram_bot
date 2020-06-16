<template>
    <div class="menu">
        <ul class="nav nav-pills flex-column">
            <li class="nav-item">
                <router-link :to="{name: 'users'}" v-slot="{ isActive }">
                    <a class="nav-link" :class="{'active': isActive}">Users</a>
                </router-link>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">API tokens</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" @click.prevent="logout()" href="#">Logout</a>
            </li>
        </ul>
    </div>
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


<style scoped>
    .menu {
        border-right: 1px solid #d8d6d6;
        min-height: 100vh;
        padding-right: 2em;
    }
</style>


