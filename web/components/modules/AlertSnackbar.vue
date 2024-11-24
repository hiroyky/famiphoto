<template>
  <v-snackbar
      outlined
      right
      top
      app
      :value="isDisplay"
      :color="color"
      @input="onDisplayChange"
  >
    <span>{{ body }}</span>
    <v-progress-linear absolute bottom :value="progressPercentage" :color="color" />
  </v-snackbar>
</template>

<script setup lang="ts">

import {computed} from "vue";
import useAlertStore from "~/stores/alert-store";

const props = defineProps({
  timeout: {
    type: Number,
    default: 5000,
  },
})

const alertStore = useAlertStore()

const currentTime = ref(0)
const isDisplay = ref(false)
const body = ref('')
const color = ref('success')

const progressPercentage = computed(() => {
  if (props.timeout <= 0) {
    return 0
  }
  return Math.floor(100 * (currentTime.value / props.timeout))
})

const head = computed(() => {
  if (alertStore.alertQueue.length === 0) {
    return undefined
  }
  return alertStore.alertQueue[0]
})

watch(head, (v) => {
  if (!isDisplay.value && v) {
    setTimeout(() => {
      isDisplay.value = true
      color.value = calcColor()
      body.value = head.value ? head.value.body : ''
      syncProgressbar()
    }, 300)
  }
})

function calcColor () {
  if (head.value === undefined) {
    return 'message'
  }
  switch (head.value.type) {
    case 'message':
      return 'success'
    case 'warn':
      return 'warning'
    case 'error':
    case 'fatal':
      return 'error'
  }
  return 'error'
}

function syncProgressbar () {
  setTimeout(() => {
    currentTime.value += 110

    if (isDisplay.value) {
      syncProgressbar()
    }
  }, 100)
}

function onDisplayChange (v : boolean) {
  if (!v) {
    isDisplay.value = false
    alertStore.deque()
    currentTime.value = 0
  }
}

</script>