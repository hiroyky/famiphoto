<template>
  <v-navigation-drawer v-model:model-value="navigation.display.value" :temporary="navigation.temporary()">
    <slot></slot>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>

import {useNavigation} from "~/composables/useNavigation";

const navigation = useNavigation()

onMounted(() => {
  navigation.init()
})


</script>
