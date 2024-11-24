import {defineStore} from "pinia";
import useMeStore from "~/stores/me-store";
import {useGqlStore} from "~/stores/gql-store";
import type {PaginationInfo, Photo, PhotosQuery} from "~/types/api-gql";
import type {PhotoList} from "~/types/api-gql-alias";
import {limits} from "argon2";

interface State {
    photos: PhotoList
    photoQuery: PhotoGetQuery
    paginationInfo: PaginationInfo
}

export interface PhotoGetQuery {
    year?: number
    month?: number
    date?: number
}
interface Pagination {
    limit: number
    offset: number
}

async function searchPhotos(q: PhotoGetQuery, page: Pagination) {
    const {client} = useGqlStore()

    const res = await client.photos({
        year: q.year,
        month: q.month,
        date: q.date,
        limit: page.limit ,
        offset: page.offset,
    })
    const photos = res.photos.nodes
        .map(item => item.thumbnailUrl === '' ? {...item, thumbnailUrl: '/no_thumbnail.png'} : item)
        .map(item => item.previewUrl === '' ? {...item, previewUrl: '/no_thumbnail.ong'} : item)

    return {photos, res}
}

const defaultLimit = 20

export const usePhotoListStore = defineStore('photoList', {
    state: (): State => ({
        photos: [],
        photoQuery: {
        },
        paginationInfo: {
            limit: 0,
            offset: 0,
            page: 0,
            paginationLength: 0,
            count: 0,
            totalCount: 0,
            hasNextPage: false,
            hasPreviousPage: false,
        },
    }),
    actions: {
        async getPhotos(q: PhotoGetQuery) {
            try {
                this.photoQuery = q
                const {photos, res} = await searchPhotos(this.photoQuery, { limit: defaultLimit, offset: 0 })
                this.photos = photos
                this.paginationInfo = res.photos.pageInfo
            } catch(err) {
                console.error(err)
            }
        },
        async appendPhotos() {
            if (!this.paginationInfo.hasNextPage) {
                return
            }

            try {
                const {photos, res} = await searchPhotos(this.photoQuery, { limit: defaultLimit, offset: this.paginationInfo.offset + this.paginationInfo.limit})
                this.photos.push(...photos)
                this.paginationInfo = res.photos.pageInfo
            } catch(err) {
                console.error(err)
            }
        },

        async aggregateDateTimeOriginalYear() {
            const {client} = useGqlStore()
            const res = await client.aggregateDateTimeOriginalYear()
            return res.aggregateDateTimeOriginal
        },
        async aggregateDateTimeOriginalYearMonth(year: number) {
            const {client} = useGqlStore()
            const res = await client.aggregateDateTimeOriginalMonth({year})
            return res.aggregateDateTimeOriginal
        },
        async aggregateDateTimeOriginalYearMonthDate(year: number, month: number) {
            const {client} = useGqlStore()
            const res = await client.aggregateDateTimeOriginalDate({year, month})
            return res.aggregateDateTimeOriginal
        }
    }
})
