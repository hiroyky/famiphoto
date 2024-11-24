import type {PhotoGetQuery} from "~/stores/photo-list-store";
import type {LocationQueryValue} from "vue-router";

interface RefineQuery {
    year?: number | LocationQueryValue
    month?: number | LocationQueryValue
    date?: number | LocationQueryValue
}
export const usePhotoQueryStore = defineStore('photo-query', {
    actions: {
        async appendRefine(q: RefineQuery) {
            const router = useRouter()
            const route = useRoute()
            await router.push({
                path: '/photos',
                query: {
                    ...route.query,
                    ...q
                }
            })
        },

        parseQuery(): PhotoGetQuery {
            const route = useRoute()
            const query: RefineQuery = route.query

            const q: PhotoGetQuery = {
                year: toNumber(query.year),
                month: toNumber(query.month),
                date: toNumber(query.date),
            }
            return q
        }
    }
})
