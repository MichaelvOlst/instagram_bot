<template>  
    <div>
        <div class="panel">
            <li class="panel-block" v-for="n in 10" :key="n">
                <div> User {{n}}</div>
               
                <div style="margin-left: auto;">
                    <a href="" style="margin-right: .5em;"><i class="fas fa-edit"></i></a>
                    <a href=""><i class="fas fa-trash-alt"></i></a>
                </div>
            </li>
        </div>
    </div> 
</template>


<script>
import { mapGetters, mapActions } from 'vuex'

export default {
    name: 'User-index',
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
    },
     watch: {
        // call again the method if the route changes
        '$route': 'getUsers'
    },

}
</script>


