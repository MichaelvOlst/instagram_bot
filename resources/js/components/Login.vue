<template>   
     <section class="hero is-light is-fullheight">
         {{form }}
        <div class="hero-body">
          <div class="container">
            <div class="columns is-centered">
              <div class="column is-5-tablet is-4-desktop is-3-widescreen">
               
                <form action="" class="box">
                    <h3 class="title has-text-centered">Login</h3>
                    <hr/>
                    <div class="field">
                        <div class="control">
                            <input v-model="form.email" type="email" placeholder="Email" class="input" required>
                        </div>
                    </div>
                    <div class="field">
                        <div class="control">
                            <input v-model="form.password" type="password" placeholder="Password" class="input" required>
                        </div>
                    </div>
                    <div class="field">
                        <button class="button is-success">Login</button>
                    </div>
                </form>
              </div>
            </div>
          </div>
        </div>
    </section>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
            form: {
                email: "",
                password: ""
            }
        }
    },
}
</script>
