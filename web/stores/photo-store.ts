import type {PhotoQuery} from "~/types/api-gql";

interface State {
    photo: PhotoQuery['photo']
}

export const usePhotoStore = defineStore('photo', {
    state: (): State => ({
        photo: null
    }),

    actions: {
        async getPhoto(id: string) {
            const { client } = useGqlStore()
            const res = await client.photo({ id })
            this.photo = res.photo
            return this.photo
        },
        download(fileId: string) {
            console.log(fileId)
        }
    }
})
