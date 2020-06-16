<template>   
    <section class="hero is-light is-fullheight">

        <form class="form-signin" method="post" @submit.prevent="login()">
            <h1 class="h3 mb-3 font-weight-normal">Sign in</h1>
            
            <label for="inputEmail" class="sr-only">Email address</label>
            <input v-model="form.email" type="email" placeholder="Email" class="form-control" required autofocus>
            
            <label for="inputPassword" class="sr-only">Password</label>
            <input v-model="form.password" type="password" placeholder="Password" class="form-control" required>
            
            <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>

            <p v-if="error">{{error}}</p>
        </form>


        
    </section>
</template>

<script>

export default {
    name: 'Login',
    data() {
        return {
            form: {
                email: "",
                password: ""
            },
            error: null
        }
    },
    methods: {
        async login() {
            try {
                await this.$store.dispatch('auth/login', this.form)
            } catch(e) {
                this.error = e
            }
        },
    }
}
</script>

<style scoped>
/* html,
body {
  height: 100%;
}

body {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
} */

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
