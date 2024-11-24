<template>
  <v-card>
    <v-form v-model="validate" @submit.prevent="$emit('commit')">
      <v-card-text>
        <v-list>
          <v-list-item>{{ loginId }}</v-list-item>
          <v-list-item>{{ userName }}</v-list-item>
        </v-list>
        <v-text-field
            v-model="value"
            label="パスワード"
            :type="isShowText ? 'text' : 'password'"
            :append-icon="isShowText ? 'mdi-eye' : 'mdi-eye-off'"
            :autofocus="true"
            :error-messages="errorMessage"
            :rules="rules"
            :loading="loading"
            @click:append="isShowText = !isShowText"
        />
      </v-card-text>
      <v-card-actions>
        <v-btn @click="$emit('back')">
          戻る
        </v-btn>
        <v-spacer />
        <v-btn :disabled="!validate" type="submit" variant="elevated" color="primary">
          <slot name="nextButton">
            ログイン
          </slot>
        </v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  loginId: {
    type: String,
    default: '',
  },
  userName: {
    type: String || null,
    default: null,
  },
  errorMessage: {
    type: String || null,
    default: null,
  },
  loading: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits<{
  "update:modelValue": [val: string]
  commit: []
  back: []
}>()

const validate = ref(false)
const isShowText = ref(false)
const rules = computed(() => [(v: string) => v.length > 0 || '入力してください'])
const value = computed({
  get() {
    return props.modelValue
  },
  set(val: string) {
    emit("update:modelValue", val)
  }
})

</script>