<template>
  <v-card>
    <v-form v-model="validate" @submit.prevent="$emit('commit')">
      <v-card-text>
        <v-text-field
            label="ユーザID"
            autofocus
            :model-value="modelValue"
            :error-messages="errorMessage"
            :rules="rules"
            :loading="loading"
            @update:modelValue="(v) => $emit('update:modelValue', v)"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn :disabled="!validate" type="submit" color="primary" variant="elevated">
          次へ
        </v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<script setup lang="ts">

defineProps({
  modelValue: {
    type: String,
    default: '',
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

defineEmits<{
  commit: [  ]
  "update:modelValue": [ val: string ]
}>()


const validate = ref(false)
const rules = computed(() => [(v: string) => !! v && v.length > 0 || '入力してください'])

</script>