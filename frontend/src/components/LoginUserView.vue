<template>
  <div class="q-pa-md" style="max-width: 300px">

    <q-form
      class="q-gutter-md"
      @reset="onReset"
      @submit="login">

      <q-input ref="emailRef" v-model="email" :rules="[ val => val && val.length > 3 || 'Please type email']" filled
               hint="Email"></q-input>
      <q-input ref="passwordRef" v-model="password" :type="isPwd ? 'password' : 'text'" filled hint="password">
        <template v-slot:append>
          <q-icon
            :name="isPwd ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="isPwd = !isPwd"
          ></q-icon>
        </template>
      </q-input>
      <div class="q-pa-md">
        <router-link to="/signup">
          Don't have an account? Register.
        </router-link>
      </div>
      <div>
        <q-btn color="primary" label="Login" type="submit"></q-btn>
        <q-btn class="q-ml-sm" color="primary" flat label="Reset" type="reset"></q-btn>
      </div>
      <p> {{ error }}</p>
    </q-form>

  </div>
</template>

<script>

import {useQuasar} from "quasar";
import {ref} from "vue";
import axios from "axios";
import ImagesService from "src/services/ImagesService";

const $q = useQuasar()

export default {
  name: 'LoginUser',
  data: () => ({
    password: "",
    isPwd: true,
    email: "",
    emailRef: ref(null),
    passwordRef: ref(null),

    status: null,
    error: null

  }),

  methods: {
    login() {
      const credentials = {
        mail: this.email,
        password: this.password
      }

      ImagesService.login(credentials)
        .then(({data}) => {
          this.$store.commit('SET_USER_DATA', data)
          this.$router.push({name: 'root'})
        }).catch(err => {
        this.status = err.response.status

        this.$q.notify({
          color: 'red-5',
          textColor: 'white',
          icon: 'warning',
          message: err.response.data
        })
      })
    },

    onReset() {
      this.password = null
      this.email = null
      this.passwordRef.value.resetValidation()
      this.emailRef.value.resetValidation()
    }
  }
}
</script>

<style scoped>

</style>
