import axios from "axios";
import ImagesService from "src/services/ImagesService";

export default {

  state: {
    currentImage: {},
    editImageTagVisible: false,
    editedTags: [],
  },

  mutations: {
    setCurrentImage(state, payload) {
      state.currentImage = payload
      state.editedTags = payload.tags
      state.editImageTagVisible = false
    },

    showPanelEditTag(state) {
      state.editImageTagVisible = true
    },

    hidePanelEditTag(state) {
      state.editImageTagVisible = false
    },

    addImageTag(state, text) {
      if (text.size === 0) return
      if (state.editedTags.includes(text)) return

      state.editedTags.push(text)
    },

    deleteImageTag(state, text) {
      const index = state.editedTags.map(item => item).indexOf(text);
      state.editedTags.splice(index, 1);
    },
  },

  getters: {
    GET_PANEL_EDIT_TAG_VISISBLE: (state) => state.editImageTagVisible,
    GET_CURRENT_IMAGE: (state) => state.currentImage,
    GET_EDITED_TAGS_CURRENT_PHOTO: (state) => state.editedTags,
    GET_EDITED_TAGS_STRING: (state) => {
      if (state.editedTags.length === 0) {
        return ""
      }
      return state.editedTags.join(", ")
    }

  },
  actions: {
    fetchImage(context, id) {
     ImagesService.getImageInfo(id).then(
        response => context.commit("setCurrentImage", response.data)
      )
    },

    replaceTagsForImage({commit}, payload) {
      return ImagesService.replaceTagsForImage(payload)
        .then(() => {
          commit('hidePanelEditTag')
        }).catch(({response}) => {
          console.log(response.data);
          console.log(response.status);
          console.log(response.headers);
        })
    },




  }
}
