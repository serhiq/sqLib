<template>
  <div class="main_container">
    <div v-show="$store.getters.GET_FILTER_PANEL_OR_VISIBLE" class="filter_container item">
      <div style="display:flex">
        <q-btn color="gray" flat round >   OR   </q-btn>
        <q-btn color="gray" flat round disable>   /   </q-btn>
        <q-btn color="gray" flat round @click="onToggleFilterAnd">   AND   </q-btn>
      </div>
      <div class="input_tag_container">
      <q-input v-model="inputOR" :dense="false" bottom-slots>
        <template v-slot:prepend>
          <q-icon name="tag"></q-icon>
        </template>
        <template v-slot:append>
          <q-btn v-if="inputOR !== ''" dense flat icon="add" round @click="onAddFilterOr"></q-btn>
        </template>
      </q-input>

      <q-item v-for="tag in  $store.getters.GET_FILTER_WORD " v-bind:key="tag" dense>
        <q-item-section>{{ tag }}</q-item-section>
        <q-btn dense flat icon="clear" round @click="onDeleteFilterOr(tag)"></q-btn>
      </q-item>

    </div>

      <div class="tags_container">
        <div class="tags_group">
          <q-btn class="tags_button" v-for="tag in $store.getters.GET_TAGS" v-bind:key="tag.ID" flat rounded color="primary"
                 :label="tag.Title" @click="onRemoteTagClickedOr(tag.Title)"></q-btn>
        </div>
      </div>
    </div>

    <div v-show="$store.getters.GET_FILTER_AND_PANEL_VISIBLE" class="filter_container item">
      <p>AND</p>

      <div class="input_tag_container">
        <q-input v-model="inputAnd" :dense="false" bottom-slots>
          <template v-slot:prepend>
            <q-icon name="tag"></q-icon>
          </template>
          <template v-slot:append>
            <q-btn v-if="inputAnd !== ''" dense flat icon="add" round @click="onAddTagAnd"></q-btn>
          </template>
        </q-input>

        <q-item v-for="tag in  $store.getters.GET_FILTER_AND_WORD " v-bind:key="tag" dense>
          <q-item-section>{{ tag }}</q-item-section>
          <q-btn dense flat icon="clear" round @click="onDeleteTagAnd(tag)"></q-btn>
        </q-item>

      </div>

      <div class="tags_container">
        <div class="tags_group">
          <q-btn class="tags_button" v-for="tag in $store.getters.GET_TAGS" v-bind:key="tag.ID" flat rounded color="primary"
                 :label="tag.Title" @click="onRemoteTagAndClicked(tag.Title)"></q-btn>
        </div>
      </div>
    </div>

    <div class="images_container ">
      <q-btn color="gray" dense flat icon="tag" round @click="onToggleFilterOR"></q-btn>
      <p v-if=" $store.getters.GET_FILTER_WORD.length !== 0">  or: {{ $store.getters.GET_FILTER_WORD }} ] [and : {{  $store.getters.GET_FILTER_AND_WORD }}  </p>
      <div class="row image_display">
        <div v-for="(img, index) in $store.getters.FILTERED_IMAGES " :key="index" class="col-1  ">
          <q-card  class="q-pa-md" style="background: transparent" flat  @click="openImage(img)">
            <q-img  :src="img.url" fit="fill"/>
          </q-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "ImagesPage",
  data: () => ({
    inputOR: "",
    inputAnd: ""
  }),

  mounted() {
    this.$store.dispatch("fetchImages")
    this.$store.dispatch("fetchTags")
  },

  methods: {
    onToggleFilterOR() {
      this.$store.commit("togglePanelFilterOR")
    },
    onAddFilterOr() {
      if (this.inputOR.size === 0) return
      this.$store.commit("addFilterOR", this.inputOR)
      this.inputOR = ""
    },

    onDeleteFilterOr(tag) {
      this.$store.commit("deleteFilterOR", tag)
    },

    onRemoteTagClickedOr(tag) {
      this.$store.commit("addFilterOR", tag)
    },
//////////////////////////////////////////////////////////////////
    onToggleFilterAnd() {
      this.$store.commit("toggleFilterAndPanel")
    },

    onAddTagAnd() {
      if (this.inputAnd.size === 0) return

      this.$store.commit("addFilterAnd", this.inputAnd)
      this.inputAnd = ""
    },

    onDeleteTagAnd(tag) {
      this.$store.commit("deleteFilterAnd", tag)
    },

    onRemoteTagAndClicked(tag) {
      this.$store.commit("addFilterAnd", tag)
    },

    openImage(photo){
      this.$router.push({ name: 'ImageDisplay', params: { id: photo.title } })
    }
  }
}
</script>

<style lang="sass" scoped>

.image_display > div
  min-width: 300px

.images_container
  padding: 16px 16px
  justify-content: left

.filter_container
  display: flex
  flex-direction: column
  height: 90vh
  min-width: 30vh

.input_tag_container
  flex-grow: 1

.image_display > div
  min-width: 300px

</style>
