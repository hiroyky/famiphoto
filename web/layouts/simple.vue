<template>
  <v-app>
    <alert-snackbar />
    <v-main>
      <slot></slot>
    </v-main>
  </v-app>
</template>
<script setup lang="ts">
import AlertSnackbar from "~/components/modules/AlertSnackbar.vue";
</script>