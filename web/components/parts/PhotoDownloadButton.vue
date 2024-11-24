<template>
  <v-menu>
    <template v-slot:activator="{props}">
      <v-btn color="primary" v-bind="props" @click="onMenuButtonClick">{{$t('download')}}</v-btn>
    </template>

    <v-list v-if="props.files.length > 1">
      <v-list-item v-for="item in props.files" :key="item.id" @click="onItemClick(item)">
        <v-list-item-title>{{ item.fileType }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script setup lang="ts">

interface File {
  id: string
  fileType: string
  fileHash: string
  fileName: string
}

interface Props {
  files: File[]
}

const props = withDefaults(defineProps<Props>(), {
  files: () => []
})

const emit = defineEmits<{
  download: [id: string, fileName: string]
}>()

function onMenuButtonClick() {
  if (props.files.length !== 1) {
    return
  }
  emit('download', props.files[0].id, props.files[0].fileName)
}

function onItemClick(file: File) {
  emit('download', file.id, file.fileName)
}

</script>