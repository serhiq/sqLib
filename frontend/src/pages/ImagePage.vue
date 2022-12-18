<template>
  <div class="container_preview_image ">
    <div class="sidebar ">

      <div v-show="!$store.getters.GET_PANEL_EDIT_TAG_VISISBLE" class="panel_display_image_info ">
        <div class="display_tags_panel">
          <h5>Теги </h5>
          <div class="tags">
            <q-btn v-for="tag in $store.getters.GET_CURRENT_IMAGE.tags" v-bind:key="tag" :label="tag" class="tags_button" color="primary"
                   flat
                   rounded></q-btn>
          </div>
        </div>
      </div>

      <q-btn v-if="haveUser" v-show="!$store.getters.GET_PANEL_EDIT_TAG_VISISBLE" class="edit_button" flat @click="onEditTagClicked()">
        edit
      </q-btn>

      <div v-show="$store.getters.GET_PANEL_EDIT_TAG_VISISBLE" class="filter_container">
        <div>
          <h5>{{ this.$route.params.id }} </h5>
          <p> {{ $store.getters.GET_EDITED_TAGS_STRING }}</p>
          <div class="edit_tag_panel">
            <q-btn unelevated rounded color="primary" @click="onSaveTagClicked()"> сохранить</q-btn>
            <q-btn  flat @click="onCancelEditClicked()"> отмена</q-btn>
          </div>
        </div>

        <div class="input_tag_container" style="height: 50vh">
          <q-input v-model="newTag" :dense="false" bottom-slots>
            <template v-slot:prepend>
              <q-icon name="tag"></q-icon>
            </template>
            <template v-slot:append>
              <q-btn v-if="newTag !== ''" dense flat icon="add" round @click="onAddTagClicked"></q-btn>
            </template>
          </q-input>

          <q-item v-for="tag in $store.getters.GET_EDITED_TAGS_CURRENT_PHOTO" v-bind:key="tag" dense>
            <q-item-section>{{ tag }}</q-item-section>
            <q-btn dense flat icon="clear" round @click="onDeleteTagClicked(tag)"></q-btn>
          </q-item>
        </div>

        <div class="tags_container" style="height: 20vh">
          <div class="tags_group">
            <q-btn v-for="tag in $store.getters.GET_TAGS" v-bind:key="tag.ID" :label="tag.Title" class="tags_button" color="primary"
                   flat
                   rounded @click="onButtonTagClicked(tag.Title)"></q-btn>
          </div>
        </div>
      </div>
    </div>

    <div class="content">
      <q-img :src="$store.getters.GET_CURRENT_IMAGE.url" class="img_preview" fit="fill"/>
    </div>
  </div>
</template>

<script>

import {useQuasar} from "quasar";

const $q = useQuasar()

export default {
  name: "ImageDisplay",
  data: () => ({
    newTag: ""
  }),

  computed: {
    haveUser() {
      return this.$store.getters.haveUser
    }
  },

  mounted() {
    this.$store.dispatch("fetchImage", this.$route.params.id)
  },

  methods: {
    onEditTagClicked() {
      this.$store.commit("showPanelEditTag")
    },

    onCancelEditClicked() {
      this.$store.commit("hidePanelEditTag")
    },

    onSaveTagClicked() {
      const send_tags = []
      for (let i = 0; i < this.$store.getters.GET_EDITED_TAGS_CURRENT_PHOTO.length; i++) {
        send_tags.push(this.$store.getters.GET_EDITED_TAGS_CURRENT_PHOTO[i])
      }

      const payload =  {
        tags: send_tags,
        fileName: this.$route.params.id
      }

        this.$store.dispatch('replaceTagsForImage', payload)
          .finally(()=>
          this.$q.notify({
            icon: 'done',
            color: 'positive',
            message: 'Теги назначены'
          }))
          },

    onDeleteTagClicked(tag) {
      this.$store.commit("deleteImageTag", tag)
    },

    onAddTagClicked() {
      if (this.newTag.size === 0) return
      this.$store.commit("addImageTag", this.newTag)
      this.newTag = ""
    },

    onButtonTagClicked(tag) {
      this.$store.commit("addImageTag", tag)
    },
  }
}
</script>

<style lang="sass" scoped>

.img_preview
  max-width: 600px

.container_preview_image
  padding: 10px
  display: grid
  grid-template-areas: 'sidebar content'
  grid-gap: 10px
  grid-template-columns: 30vh 1fr

.sidebar
  grid-area: sidebar
  display: flex
  flex-direction: column

  align-self: start
  padding-left: 16px


.content
  grid-area: content
  display: flex
  justify-content: center
  align-items: center

.item_filter_bar
  height: 90vh
  min-width: 30vh

.panel_display_image_info
  height: 87vh

.display_tags_panel
  display: flex
  flex-direction: column
  min-height: 75%

.tags
  height: fit-content
  display: flex
  flex-wrap: wrap

.edit_button
  align-self: flex-start

.title_tags
  padding: 8px

.edit_tag_panel
  display: flex
  justify-content: space-between

</style>
