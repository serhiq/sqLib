import axios from "axios";
import ImagesService from "src/services/ImagesService";

export default {

  state: {
    tags: [],
    images: [],

    filterOrVisible: true,
    filterOR: [],

    filterAnd: [],
    filterAndVisible: false,
  },

  mutations: {
    togglePanelFilterOR(state) {
      state.filterOrVisible = !state.filterOrVisible
    },

    toggleFilterAndPanel(state) {
      state.filterAndVisible = !state.filterAndVisible
    },

    setImageResponse(state, payload) {
      state.images = payload
    },

    setTagsResponse(state, payload) {
      state.tags = payload
    },

    addFilterOR(state, text) {
      if (text.size === 0) return
      if (state.filterOR.includes(text)) return

      state.filterOR.push(text)
    },

    addFilterAnd(state, text) {
      if (text.size === 0) return
      if (state.filterAnd.includes(text)) return
      state.filterAnd.push(text)
    },

    deleteFilterOR(state, text) {
      const index = state.filterOR.map(item => item).indexOf(text);
      state.filterOR.splice(index, 1);
    },

    deleteFilterAnd(state, text) {
      const index = state.filterAnd.map(item => item).indexOf(text);
      state.filterAnd.splice(index, 1);
    },

    clearFilterWord(state) {
      this.filterOR = []
    }
  },

  getters: {
    GET_ALL_IMAGES: (state) => state.images,
    GET_TAGS: (state) => state.tags,

    GET_FILTER_PANEL_OR_VISIBLE: (state) => state.filterOrVisible,

    GET_FILTER_AND_PANEL_VISIBLE: (state) => state.filterAndVisible,

    GET_FILTER_WORD: (state) => state.filterOR,
    GET_FILTER_AND_WORD: (state) => state.filterAnd,
    FILTERED_IMAGES: (state, getters) => {
      let filterAndConstraint = [];

      if (state.filterOR.length === 0 && state.filterAnd.length === 0) {
        return getters.GET_ALL_IMAGES
      }

      for (let i = 0; i < state.images.length; i++) {
        const image = state.images[i];
        if (state.filterAnd.every(r => image.tags.includes(r))){

        } else {
          continue
        }

        if (state.filterOR.length === 0){
          filterAndConstraint.push(image)
        }

        outer_loop_OR:
          for (let j = 0; j < image.tags.length; j++) {
            const imageTag = image.tags[j];

            for (let k = 0; k < state.filterOR.length; k++) {
              const tag = state.filterOR[k];

              if (imageTag.includes(tag)) {
                filterAndConstraint.push(image)
                break outer_loop_OR;
              }
            }
          }
      }

      return filterAndConstraint
    },
  },

  actions: {
    fetchImages(context) {
     return  ImagesService.getImages().then(
        response => context.commit("setImageResponse", response.data)
      )
    },

    fetchTags(context) {
      return ImagesService.getTags().then(
        response => context.commit("setTagsResponse", response.data)
      )
    },

    addTagForPhoto({commit}, payload) {
      return ImagesService.addTagsForImages(payload)
        .catch(({response}) => {
          console.log(response.data);
          console.log(response.status);
          console.log(response.headers);
        })
    },

    deleteImages({commit}, payload) {
      return ImagesService.deleteImages(payload)
          .catch(({response}) => {
          console.log(response.data);
          console.log(response.status);
          console.log(response.headers);
        })
    },
  }
}
