<template>
  <v-list>
    <template v-if="modelValue && modelValue.length > 0">
      <v-list-item v-for="item in modelValue" :key="item.file.name">
        <template #prepend v-if="item.status === 'standby'"><v-icon>mdi-image</v-icon></template>
        <template #prepend v-if="item.status === 'failed'"><v-icon color="error">mdi-alert-circle-outline</v-icon></template>
        <template #prepend v-if="item.status === 'completed'"><v-icon color="success">mdi-check</v-icon></template>
        <template #prepend v-if="item.status === 'progressing'"><v-icon color="primary">mdi-image</v-icon></template>
        <template #title>{{ item.file.name }}</template>
        <template #subtitle v-if="item.status==='progressing'">
          アップロード中..
        </template>
      </v-list-item>
    </template>
    <template v-if="!modelValue || modelValue.length === 0">
      <v-list-item>
        {{ $t('noUploadingProcesses') }}
      </v-list-item>
    </template>
  </v-list>
</template>

<script setup lang="ts">

import type {FileUploadQueueItem} from "~/stores/photo-upload-store";

interface Prop {
  modelValue?: FileUploadQueueItem[]
}

defineProps<Prop>()

</script>