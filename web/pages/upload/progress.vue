<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>写真アップロードの状況</h1>

        <v-container>
          <v-row>
            <v-col>
              <v-btn color="primary" to="/upload">アップロード</v-btn>
              <v-btn class="mx-1" :disabled="!isProgressing">中断</v-btn>
              <v-btn>クリア</v-btn>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-card :loading="isProgressing">
                <upload-progress-list :model-value="queue" />
              </v-card>
            </v-col>
          </v-row>
        </v-container>

      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import UploadProgressList from "~/components/modules/UploadProgressList.vue";

const photoUploadStore = usePhotoUploadStore()

const queue = computed(() => photoUploadStore.queue)
const isProgressing = computed(() => photoUploadStore.progressing)

</script>