import axios from "axios";


const apiClient = axios.create({
  baseURL: `http://localhost:7777`,
  // withCredentials: false, // This is the default
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getImages() {
    return apiClient.get('/images')
  },

  getImageInfo(id) {
    return apiClient.get('/images/' + id + '/info')
  },

  getTags() {
    return apiClient.get('/tags')
  },

  addTagsForImages(payload) {
    return apiClient.post('/images/actions/add_tag', payload)
  },

  replaceTagsForImage(payload) {
    const request = {
      tags: payload.tags
    }
    const url = "/images/" + payload.fileName + "/update_tags"

    return apiClient.post(url, request)
  },

  deleteImages(payload) {
    return apiClient.post('/images/actions/delete', payload)
  },

  uploadImage(payload) {
    return apiClient.post('/images', payload, {
      headers: {
        "Content-Type": "multipart/form-data",
        "Accept": "*/*"
      },
    })
  },

  login(credentials) {
    return apiClient.post('/login', credentials)
  },

  register(credentials) {
    return apiClient.post('/register', credentials)
  },


  setUserToken(token) {
    apiClient.defaults.headers.common['Authorization'] = `Bearer `+ token
  },

  clearUserData() {
    apiClient.defaults.headers.common['Authorization'] = null
  }
}
