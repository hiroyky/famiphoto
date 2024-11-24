<template>
  <v-app>
    <alert-snackbar />
    <v-main>
      <v-container class="fill-height align-center">
        <v-row justify="center" align="center">
          <v-col cols="12" sm="8" md="6">
            <slot></slot>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import AlertSnackbar from "~/components/modules/AlertSnackbar.vue";
</script>
