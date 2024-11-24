<template>
  <v-form>
    <v-list>
      <v-list-item>
        <v-btn
            @click="onUploadClick"
            color="primary"
            :disabled="uploadButtonDisabled"
            :loading="loading"
        >{{$t('upload')}}</v-btn>
      </v-list-item>
      <v-list-item>
        <v-file-input
            :multiple="true"
            accept="image/*,.arw"
            :label="$t('selectPhotos')"
            :counter="true"
            :disabled="loading"
            v-model="files"
            :loading="loading"
        ></v-file-input>

      </v-list-item>
    </v-list>
  </v-form>
</template>

<script setup lang="ts">

interface Prop {
  modelValue?: string[]
  loading?: boolean
}

const props = defineProps<Prop>()

const emit = defineEmits<{
  'update:modelValue': [File[]]
}>()

const files = ref<File[]>([])

function onUploadClick() {
  emit('update:modelValue', files.value)
}

const uploadButtonDisabled = computed(() => {
  if (files.value.length == 0) {
    return true
  }
  return false
})
</script>