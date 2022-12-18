import axios from "axios";
import ImagesService from "src/services/ImagesService";

export default {

  state: {
    user: null
  },

  mutations: {
    SET_USER_DATA(state, userData) {
      state.user = userData
      localStorage.setItem('user', JSON.stringify(userData))
      ImagesService.setUserToken(userData.token)
      // axios.defaults.headers.common['Authorization'] = `Bearer $userData.token`
    },

    CLEAR_USER_DATA(state) {
      state.user = null
      ImagesService.clearUserData()

      // axios.defaults.headers.common['Authorization'] = null
      localStorage.removeItem("user")
      localStorage.reload()
    }
  },

  getters: {
    haveUser(state) {
       return state.user
    }
  },

  actions: {
    logout({ commit }) {
      commit('CLEAR_USER_DATA')
    }
  },
}
