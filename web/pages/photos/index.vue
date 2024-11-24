<template>
  <v-container :fluid="true">
    <v-row justify="center" align="center">
      <v-col cols="12">
        <photo-list :value="photos" @photo-click="onPhotoClick" @intersect-next="onIntersectNext"></photo-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">

import PhotoList from "~/components/modules/PhotoList.vue";
import {usePhotoListStore} from "~/stores/photo-list-store";

const photoListStore = usePhotoListStore()

const photos = computed(() => photoListStore.photos)

function onPhotoClick(id: string) {
  navigateTo(`photos/${id}`)
}

function onIntersectNext() {
  photoListStore.appendPhotos()
}

</script>