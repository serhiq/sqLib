import { boot } from 'quasar/wrappers'

export default ({ app, router, store }) => {
  const userSting = localStorage.getItem('user')

  if (userSting) {
    const userData = JSON.parse(userSting)
    store.commit('SET_USER_DATA', userData)
  }
}


