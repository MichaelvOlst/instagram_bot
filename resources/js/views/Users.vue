<template>  
    <div>
        <div class="header">
            <h3>Users</h3>
            <b-button variant="primary" v-b-modal.create-user>Create</b-button>
        </div>

        <p v-if="error">{{error}}</p>
       
        <b-modal id="create-user" title="Add User" @cancel="clearInput()" @ok="addUser()">
            <form method="post" action="" @submit.prevent="addUser()">
                <div class="form-group">
                    <label for="email">Email address</label>
                    <input type="email" v-model="form.email" class="form-control" id="email" required>
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" v-model="form.password" class="form-control" id="password" required>
                </div>
                    
            </form>
        </b-modal>


        <ul class="list-group" v-if="users.length">
            <li class="list-group-item" v-for="user in users" :key="user.id" >
                <div class="title">{{ user.email }}</div>
                <div class="buttons right">
                    <a href="#" @click.prevent="editUser(user)"><i class="far fa-edit"></i></a>
                    <a href="#" @click.prevent="deleteUser(user)"><i class="far fa-trash-alt"></i></a>
                </div>
            </li>
        </ul>

        <p v-if="!users.length">No users found</p>

    </div> 
</template>


<script>
import { mapGetters, mapActions } from 'vuex'

export default {
    name: 'User-index',

    data() {
        return {
            form: {
                email: "",
                password: "",
            }, 
            error: null,
        }
    },

    mounted() {
        this.getUsers()
    },
    computed: {
        ...mapGetters('users', {
            users: 'getUsers',
        })
    },
    methods: {
        ...mapActions('users', [
            'getUsers'
        ]),

        async addUser() {           
            try {
                await this.$store.dispatch('users/addUser', this.form)
            } catch(e) {
                this.error = e 
            }
        },

        clearInput() {
            this.form.email = "";
            this.form.password = "";
        }
    },
     watch: {
        // call again the method if the route changes
        '$route': 'getUsers'
    },

}
</script>


<style scoped>

    .header {
        display: flex;
        justify-content: space-between;
        margin-bottom: 1em;
    }

    .list-group-item {
        display: flex;
    }

    .list-group-item .buttons {
        margin-left: auto;
    }
</style>