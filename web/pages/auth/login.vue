<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>
            FAMIPHOTO
          </v-card-title>
          <v-card-subtitle>ログイン</v-card-subtitle>

          <v-window v-model="step">
            <v-window-item>
              <login-id-form v-model="loginId" :loading="loading" :error-message="loginIdErrorMessage" @commit="onLoginIdCommit" @input="loginIdErrorMessage=''" />
            </v-window-item>

            <v-window-item>
              <login-password-form
                  v-model="password"
                  :loading="loading"
                  :login-id="loginId"
                  :error-message="passwordErrorMessage"
                  @back="onBack"
                  @commit="onCommit"
                  @input="passwordErrorMessage=''"
              />
            </v-window-item>
          </v-window>

          <v-card-text class="text-center">
            <nuxt-link to="./create_user">
              アカウント新規作成
            </nuxt-link>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script setup lang="ts">
import {definePageMeta, navigateTo, useAuthStore} from "#imports";
import LoginIdForm from "~/components/parts/LoginIdForm.vue";
import LoginPasswordForm from "~/components/parts/LoginPasswordForm.vue";
import useAlertStore from "~/stores/alert-store";

definePageMeta({
  layout:"center"
})

const step = ref(0)
const loading = ref(false)
const loginId = ref('')
const loginIdErrorMessage = ref('')
const password = ref('')
const passwordErrorMessage = ref('')

const route = useRoute()
const to = route.query.to ? route.query.to.toString() : '/'

const authStore = useAuthStore()
const alertStore = useAlertStore()

function onBack() {
  step.value--
}

async function onLoginIdCommit() {
  try {
    if (!await authStore.existUserId(loginId.value)) {
      loginIdErrorMessage.value = 'そのユーザIDは存在しません。'
      return
    }
    step.value = 1
  } catch(err) {
    alertStore.displayError('予期せぬ通信エラーが発生しました。')
  } finally {
    loading.value = false
  }
}
async function onCommit() {
  try {
    loading.value = true
    await authStore.login(loginId.value, password.value)
    navigateTo(to.toString())
  } catch (err) {

  } finally {
    loading.value = false
  }
}
</script>