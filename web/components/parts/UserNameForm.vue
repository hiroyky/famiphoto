<template>
  <v-card>
    <v-form v-model="validate" @submit.prevent="$emit('commit')">
      <v-card-text>
        <p v-if="loginId">
          <v-list>
            <v-list-item>{{ loginId }}</v-list-item>
          </v-list>
        </p>
        <v-text-field
            class="mt-4"
            label="表示名"
            :autofocus=true
            :error-messages="errorMessage"
            :rules="rules"
            v-model="value"
        />
      </v-card-text>
      <v-card-actions>
        <v-btn @click="$emit('back')">
          戻る
        </v-btn>
        <v-spacer />
        <v-btn :disabled="!validate" @click="$emit('commit')" color="primary" variant="elevated">
          次へ
        </v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<script setup lang="ts">
const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  errorMessage: {
    type: String || null,
    default: null,
  },
  loginId: {
    type: String || null,
    default: null,
  },
})

const emit = defineEmits<{
  commit: [  ]
  "update:modelValue": [ val: string ]
  back: [ ]
}>()


const validate =ref(false)
const rules = computed(() => [(v: string) => !! v && v.length > 0 || '入力してください'])
const value = computed({
  get() {
    return props.modelValue
  },
  set(val) {
    emit('update:modelValue', val)
  }
})
</script>