import { store } from 'quasar/wrappers'
import { createStore } from 'vuex'
import photoModule from "src/store/image";
import authModule from "src/store/auth";
import imageModule  from "src/store/images";

export default store(function (/* { ssrContext } */) {
  const Store = createStore({
    modules: {
      photoModule,
      authModule,
      imageModule
    },

    // enable strict mode (adds overhead!)
    // for dev mode and --debug builds only
    strict: process.env.DEBUGGING
  })

  return Store
})
