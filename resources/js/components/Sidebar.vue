<template>
    <aside class="menu">
        <ul class="menu-list">
            <li>
                <router-link :to="{name: 'users'}" v-slot="{ isActive }">
                    <a :class="{'is-active':isActive}">Users</a>
                </router-link>
            </li>
            <li><a>API tokens</a></li>
            <li><a @click.prevent="logout()">Logout</a></li>
        </ul>
    </aside>
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


<style>
    aside {
        border-right: 1px solid #d8d6d6;
        min-height: 100vh;
        padding-right: 2em;
    }
</style>


