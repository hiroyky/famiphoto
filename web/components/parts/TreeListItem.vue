<template>
  <v-list-group
      :value="modelValue"
      collapse-icon=""
      prepend-icon=""
      expand-icon=""
      :subgroup="!!subGroup"
  >
    <template #activator="{props, isOpen}">
      <v-list-item :active="false">
        <template #title>
          <div @click="onClick">
            <span><slot name="title"></slot></span></div>
        </template>
        <template #subtitle><div @click="onClick"><slot name="subtitle"></slot></div></template>
        <template #append>
          <span class="count">(<slot name="count">100</slot>)</span>
          <v-btn
              :disabled="!!terminate"
              :icon="buttonIcon(isOpen)"
              size="small"
              variant="text"
              density="comfortable"
              v-bind="props"
          ></v-btn>
        </template>
      </v-list-item>
    </template>
    <template #default>
      <slot></slot>
    </template>
  </v-list-group>
</template>

<script setup lang="ts">

interface Props {
  modelValue: any
  subGroup?: boolean
  terminate?: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{
  click: [any]
}>()

function buttonIcon(isOpen: boolean): string {
  return isOpen ? 'mdi-menu-up' : 'mdi-menu-down'
}

function onClick() {
  console.log(props.modelValue)
  emit('click', props.modelValue)
}
</script>

<style lang="scss" scoped>
.count {
  padding-left: 1em;
  font-size: 80%;
}
</style>