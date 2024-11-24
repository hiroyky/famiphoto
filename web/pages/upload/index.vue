<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>{{ $t('uploadPhotos') }}</h1>

        <v-container>
          <v-row>
            <v-col>
              <v-card>
                <photo-upload-form @update:model-value="onCommit" />
              </v-card>
            </v-col>
          </v-row>
        </v-container>

      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import {usePhotoUploadStore} from "~/stores/photo-upload-store";

const { t } = useI18n({useScope: 'global'})
import PhotoUploadForm from "~/components/parts/photoUploadForm.vue";

useHead({
  title: `${t("uploadPhotos")} - ${useAppConfig().AppName}`,
})

const photoUploadStore = usePhotoUploadStore()

function onCommit(files: File[]) {
  photoUploadStore.enqueueToUploadFiles(files)
  navigateTo('/upload/progress')
}

</script>