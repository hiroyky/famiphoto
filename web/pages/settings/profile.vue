<template>
  <v-container>
    <v-row>
      <v-col>
          <h1>{{ $t('personalSettings') }}</h1>

        <v-container>
          <v-row>
            <v-col md="6">
              <v-card>
                <v-form>
                  <v-list>
                    <v-list-subheader>{{ $t('profile') }}</v-list-subheader>
                    <v-list-item>
                      <v-text-field :label="$t('userId')" :disabled="true" :model-value="meStore.userId"></v-text-field>
                    </v-list-item>
                    <v-list-item>
                      <v-text-field :label="$t('displayName')" v-model="displayName"></v-text-field>
                    </v-list-item>
                  </v-list>
                  <v-card-actions>
                    <v-spacer/>
                    <v-btn @click="onApplyPersonalSetting">{{ $t('apply') }}</v-btn>
                    <v-spacer/>
                  </v-card-actions>
                </v-form>
              </v-card>
            </v-col>
          </v-row>

          <v-row>
            <v-col md="6">
              <v-list>
                <v-list-item to="/auth/change_password">
                  <v-list-item-title>{{ $t('changePassword') }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-col>
          </v-row>

          <v-row>
            <v-col md="6">
              <v-card>
              <v-list>
                <v-list-subheader>{{ $t('deviceSettings') }}</v-list-subheader>
                <v-list-item>
                  <v-list-item-title>
                    <v-select label="Select Language" :items="localeSelection" item-value="value" item-title="text" :model-value="locale" @update:model-value="onApplyLocale"></v-select>
                  </v-list-item-title>
                </v-list-item>
              </v-list>
              </v-card>
            </v-col>
          </v-row>

        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import {definePageMeta} from "#imports";
import useMeStore from "~/stores/me-store";
import {assertTSMethodSignature} from "@babel/types";

const {t, locales, locale, setLocale } = useI18n({useScope: 'global'})
const meStore = useMeStore()

definePageMeta({
  layout: "default"
})

const displayName = ref(meStore.name)

const localeSelection = computed(() => {
  return locales.value.map(l => ({
    value: l.code,
    text: l.name,
  }))
})


async function onApplyLocale(val: any) {
  await setLocale(val)
}

function onApplyPersonalSetting() {
  meStore.updateMe(displayName.value)
}

</script>