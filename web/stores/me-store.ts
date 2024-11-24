import {type MeQuery} from "~/types/api-gql";

interface State {
    me: MeQuery['me']
}

const useMeStore = defineStore('me', {
    state: (): State => ({
        me: null
    }),
    getters: {
        userId: (state) => {
            return state.me ? state.me.userId : ''
        },
        name: (state) => {
            return state.me ? state.me.name : ''
        }
    },
    actions: {
        async getMe() {
            try {
                const { client } = useGqlStore()
                const res = await client.me()
                this.me = res.me
            } catch(err){
                this.me = null
            }
        },
        async isLoggedIn() {
            if (this.me !== null) {
                return true
            }
            await this.getMe()
            return this.me !== null
        },
        async updateMe(name: string) {
            const { client } = useGqlStore()
            const res = await client.updateMe({name})
            this.me = res.updateMe
        },
        clear() {
            this.me = null
        }
    }
})

export default useMeStore
