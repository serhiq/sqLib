<template>
  <q-layout class="bg-grey-1" view="hHh lpR fFf">
    <q-header class="bg-white text-grey-8" height-hint="64">
      <q-toolbar class="main_toolbar">
        <q-btn icon="scatter_plots" :ripple="false" class="q-ml-sm q-px-md" size="lg" dense flat label="qSLib" no-caps no-wrap   @click="navigateToMainScreen()" />
        <q-space/>
        <q-space/>

        <q-btn v-if="haveUser"  class="q-ml-sm q-px-md" dense flat label="Edit" no-caps no-wrap      @click="navigateToEdit()"/>
        <q-btn v-if="haveUser"  class="q-ml-sm q-px-md" dense flat label="Upload" no-caps no-wrap   @click="navigateToUpload()" />

        <q-btn v-if="!haveUser" class="q-ml-sm q-px-md" dense flat label="Login" no-caps no-wrap   @click="navigateToLogin()" />
        <q-btn v-if="haveUser" class="q-ml-sm q-px-md" dense flat label="Logout" no-caps no-wrap   @click="logout()" />
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view v-slot="{ Component }">
<!--        <keep-alive>-->
          <component :is="Component" :key="$route.fullPath"></component>
<!--        </keep-alive>-->
      </router-view>
    </q-page-container>
  </q-layout>
</template>

<script>

export default {
  computed: {
    haveUser() {
      return this.$store.getters.haveUser
    }
  },

  methods: {
    navigateToUpload() {
      this.$router.push('/upload')
    },

    navigateToEdit() {
      this.$router.push('/edit')
    },

    navigateToLogin() {
      this.$router.push('/login')
    },

    navigateToMainScreen() {
      this.$router.push('/')
    },

    logout(){
      this.$store.dispatch('logout')
      this.$router.push('/')
    }
  }}
</script>

<style lang="sass" scoped>
.main_toolbar
  height: 64px
</style>
