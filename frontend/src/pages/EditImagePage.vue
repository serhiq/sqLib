<template>
  <div class="container_img">
    <div  class="item_filter_bar item">
      <q-input v-model="text" :dense="false" bottom-slots>
        <template v-slot:prepend>
          <q-icon name="tag"></q-icon>
        </template>
        <template v-slot:append>
          <q-btn v-if="text !== ''" dense flat icon="add" round @click="onAddTagClicked"></q-btn>
        </template>
      </q-input>
      <p v-if=" $store.getters.GET_FILTER_WORD.length !== 0"> {{ $store.getters.GET_FILTER_WORD }} </p>

      <q-item v-for="tag in tags" v-bind:key="tag" dense>
        <q-item-section>{{ tag }}</q-item-section>
        <q-btn dense flat icon="clear" round @click="onDeleteTagClicked(tag)"></q-btn>
      </q-item>

    </div>
    <div class="doc-page ">

      <div class="q-px-sm">
        Файлы      <strong>{{ selection }}</strong>
      </div>
      <div class="flex">
        <q-btn color="gray" label="Назначить теги" dense flat icon="tag"  @click="onAddTagButtonClicked"></q-btn>

        <q-btn color="gray" label="удалить" dense flat icon="delete"  @click="onDeleteButtonClicked"></q-btn>

      </div>
      <p> {{ tags }} </p>

      <div class="text-right">
      </div>
      <div class="row imgg">
        <div v-for="(photo) in images " :key="photo.title" class="col-1  ">
          <q-card class="q-pa-md  bg-transparent" flat>

            <q-item  class="no-padding" flat tag="label">
            <q-img :src="photo.url" fit="fill"  v-ripple>
                <div class="absolute-top text-right bg-transparent">
                  <q-checkbox dense v-model="selection" v-bind:val="photo.title"  style="top: 8px; left: 8px" ></q-checkbox>
                </div>
              </q-img>
            </q-item>

          </q-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {useQuasar} from "quasar";

const $q = useQuasar()

export default {
  name: "EditImagePage",
  data: () => ({
    text: "",
    tags: [],
    selection: ([ ]),
  }),

  computed: {
    images() {
      return this.$store.getters.GET_ALL_IMAGES
    }
  },

  mounted() {
    this.$store.dispatch("fetchImages")
  },

  methods: {
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
    },

    onAddTagButtonClicked() {
      if (this.tags.length === 0) {
        this.$q.notify({
          icon: 'warning',
          color: 'warning',
          message: 'Список тегов пуст'
        })

        return
      }

      const request = {
        ids: this.selection,
        tags: this.tags
      }
      this.$store.dispatch('addTagForPhoto',request).then(
        () => {
          this.selection =[]
          this.tags = []
        }
      ).finally(()=> {
        this.$q.notify({
          icon: 'done',
          color: 'positive',
          message: 'Теги назначены'
        })
        this.$store.dispatch("fetchImages")
      })
    },
    onDeleteButtonClicked() {
      const request = {
        ids: this.selection
      }
      this.$store.dispatch('deleteImages',request).then(
        () => {
          this.selection =[]
          this.tags = []
        }
      ).finally(()=> {
        this.$q.notify({
          icon: 'done',
          color: 'positive',
          message: 'Запрос на удаление, зарегестрирован'
        })
        this.$store.dispatch("fetchImages")
      }
    )
    }
  }
}
</script>

<style lang="sass" scoped>

.imgg > div
  min-width: 300px

.doc-page > div
  padding: 16px 16px
  justify-content: left


.container_img
  display: flex

.item_filter_bar
  height: 90vh
  min-width: 30vh

</style>
