<template>
  <v-btn :icon="icon" :to="to" :active="false">
    <slot v-if="!loading"></slot>
    <slot name="loading" v-if="loading">
      <v-progress-circular indeterminate width="2">
        <slot></slot>
      </v-progress-circular>
    </slot>
  </v-btn>
</template>

<script lang="ts" setup>
import type {RouteLocationRaw} from "vue-router";

interface Prop {
  defaultTo?: RouteLocationRaw
  loadingTo?: RouteLocationRaw
  loading?: boolean
  icon?: boolean
}

const prop = defineProps<Prop>()

const to = computed(() => {
  return prop.loading ? prop.loadingTo : prop.defaultTo
})
</script>