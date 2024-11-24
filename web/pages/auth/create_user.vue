<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>{{ appConfig.AppName }}</v-card-title>
          <v-card-subtitle>アカウント作成</v-card-subtitle>

          <v-window v-model="step">
            <v-window-item>
              <login-id-form
                  v-model="loginId"
                  :loading="loading"
                  :error-message="loginIdErrorMessage"
                  @commit="onLoginIdCommit"
                  @input="loginIdErrorMessage=''"
              />
            </v-window-item>

            <v-window-item>
              <user-name-form
                  v-model="userName"
                  :login-id="loginId"
                  :error-message="userNameErrorMessage"
                  @commit="onUserNameCommit"
                  @back="onBack"
                  @input="userNameErrorMessage=''"
              />
            </v-window-item>

            <v-window-item>
              <login-password-form
                  v-model="password"
                  :loading="loading"
                  :login-id="loginId"
                  :error-message="passwordErrorMessage"
                  :user-name="userName"
                  @back="onBack"
                  @commit="onCommit"
                  @input="passwordErrorMessage=''"
              >
                <template #nextButton>
                  アカウント作成
                </template>
              </login-password-form>
            </v-window-item>

          </v-window>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">

import {definePageMeta, useAuthStore} from "#imports";
import LoginIdForm from "~/components/parts/LoginIdForm.vue";
import UserNameForm from "~/components/parts/UserNameForm.vue";
import LoginPasswordForm from "~/components/parts/LoginPasswordForm.vue";
import useAlertStore from "~/stores/alert-store";
const appConfig = useAppConfig()
definePageMeta({
  layout:"center"
})
const step = ref(0)
const loading = ref(false)
const loginId = ref('')
const loginIdErrorMessage = ref('')
const userName = ref('')
const userNameErrorMessage = ref('')
const password = ref('')
const passwordErrorMessage = ref('')

const authStore = useAuthStore()
const alertStore = useAlertStore()

async function onLoginIdCommit() {
  try {
    loading.value = true
    if (await authStore.existUserId(loginId.value)) {
      loginIdErrorMessage.value = "そのユーザIDは使われています。"
      return
    }
    step.value = 1
  } catch(err){
    alertStore.displayError('予期せぬ通信エラーが発生しました。')
  } finally {
    loading.value = false
  }

}
function onBack() {
  step.value--
}
function onUserNameCommit () {
  step.value = 2
}
async function onCommit() {
  try {
    loading.value = true
    await authStore.createUser(loginId.value, userName.value, password.value)
    navigateTo("/")
  } finally {
    loading.value = false
  }
}
</script>