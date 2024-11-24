<template></template>
<v-progress-linear indeterminate />
<script setup lang="ts">
import {definePageMeta, useAuthStore} from "#imports";
import {navigateTo, useAsyncData} from "#app";

definePageMeta({
  layout:"simple"
})

const authStore = useAuthStore()

await authStore.logout()
navigateTo('/')

</script>