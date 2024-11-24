import { defineStore } from 'pinia'
import { useGqlStore } from '~/stores/gql-store'
import {useApiStore} from "~/stores/api-store";
import useMeStore from "~/stores/me-store";

// ログインページにリダイレクトとか？
export const useAuthStore = defineStore('auth', {
  actions: {
    async existUserId (userId: string): Promise<boolean> {
      const { client } = useGqlStore()
      const { existUserId, } = await client.existUserId({ id: userId })
      return existUserId
    },

    async login (userId: string, password: string) {
      const apiStore = useApiStore()
      await apiStore.login({ userId, password })
      const meStore = useMeStore()
      await meStore.getMe()
    },

    async logout () {
      const apiStore = useApiStore()
      await apiStore.logout()
      const meStore = useMeStore()
      meStore.clear()
    },

    async createUser (userId: string, name: string, password: string) {
      const { client } = useGqlStore()
      await client.createUser({ userId, password, name })
      await this.login(userId, password)
    },
  },
})
