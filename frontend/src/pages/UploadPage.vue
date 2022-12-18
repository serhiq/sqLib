<template>
  <div class="container">
    <div class="header">

      <div class="q-gutter-md" style="max-width: 300px">
        <q-file v-model="model" :label="labelFileInput" filled multiple>
          <template v-slot:selected="{files}">
            <p>{{ files.length }}</p>
          </template>
          <template v-if="model.length !== 0" v-slot:append>
            <q-icon class="cursor-pointer" name="cancel" @click.stop.prevent="model = []"></q-icon>
          </template>

          <template v-if="canUpload" v-slot:after>
            <q-btn
              :disable="!canUpload"
              :loading="isUploading"
              color="primary"
              dense
              icon="cloud_upload"
              round
              @click="uploadFiles"
            ></q-btn>
          </template>

        </q-file>

        <q-input v-model="text" :dense="false" bottom-slots>
          <template v-slot:prepend>
            <q-icon name="tag"></q-icon>
          </template>
          <template v-slot:append>
            <q-btn v-if="text !== ''" dense flat icon="add" round @click="onAddTagClicked"></q-btn>
          </template>
        </q-input>
      </div>
    </div>

    <div class="sidebar">
      <p v-if="tags.length !==0"> {{ tags }} </p>
      <q-item v-for="tag in tags" v-bind:key="tag" dense>
        <q-item-section>{{ tag }}</q-item-section>
        <q-btn dense flat icon="clear" round @click="onDeleteTagClicked(tag)"></q-btn>
      </q-item>
    </div>
    <div class="content">
      <div v-for="(photo, index) in selectedImage"
           v-bind:key="index"
           class="photo_uploaded">
        <img :src=photo.src class="photo_img">
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios";
import {useQuasar} from "quasar";
import ImagesService from "src/services/ImagesService";

const $q = useQuasar()

export default {
  name: "UploadPage",
  data: () => ({
    model: [],
    isUploading: false,
    tags: [],
    text: ""
  }),

  computed: {
    selectedImage() {
      let result = []
      for (let i = 0; i < this.model.length; i++) {
        let photoInPreview = {}
        photoInPreview.title = this.model[i].name
        photoInPreview.src = URL.createObjectURL(this.model[i])
        result.push(photoInPreview)
      }
      return result
    },

    labelFileInput() {
      return this.model.length === 0 ? "Select files" : "Files"
    },
    canUpload() {
      return this.model.length !== 0
    },
  },

  mounted() {
    this.$store.dispatch("fetchImages")
  },

  methods: {
    uploadFiles() {
      for (let i = 0; i < this.model.length; i++) {
        this.uploadFile(i, this.model[i]);
      }
      this.$q.notify({
        icon: 'done',
        color: 'positive',
        message: 'Отправлено'
      })
    },

    uploadFile(idx, file) {
      let formData = new FormData();
      formData.append("file", file)
      formData.append("tags", this.tags);

      ImagesService.uploadImage(formData).catch(({response}) => {
        console.log(response.data);
        console.log(response.status);
        console.log(response.headers);
        this.messageError = response.data
      }).finally(() => {
        this.model = []
        this.isUploading = false
        this.tags = []
        this.text = ""
      });
    },

    onAddTagClicked() {
      if (this.text.size === 0) return

      let index = this.tags.indexOf(this.text)
      if (index !== -1) {
        this.text = ""
        return;
      }
      this.tags.push(this.text)
      this.text = ""
    },

    onDeleteTagClicked(tag) {
      let index = this.tags.indexOf(tag)
      if (index !== -1) {
        this.tags.splice(this.tags.indexOf(tag), 1);
      }
    }
  }
}

</script>

<style lang="sass" scoped>
.container
  display: grid
  grid-template-areas: 'header header' 'sidebar content'
  grid-gap: 10px
  grid-template-columns: 150px 1fr
  padding: 16px

.header
  grid-area: header

.sidebar
  grid-area: sidebar
  align-self: start

.content
  padding: 16px
  grid-area: content
  display: flex
  justify-content: flex-start
  flex-wrap: wrap
  gap: 16px

.photo_uploaded
  height: fit-content
  width: fit-content
  min-width: 300px
  max-width: 400px
  padding: 16px
  flex-grow: 1

.photo_img
    max-width: 100%
    height: auto

label.label input[type="file"]
  position: absolute
  top: -1000px

.label
  cursor: pointer
  border-radius: 5px
  padding: 8px 15px
  margin: 5px
  display: inline-block


.label:hover
  background: #e6f2fa

.label:invalid + span
  color: #000000

.label:valid + span
  color: #ffffff

</style>
