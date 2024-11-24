import type {
    AggregateDateTimeOriginalDateQuery,
    AggregateDateTimeOriginalMonthQuery,
    AggregateDateTimeOriginalYearQuery,
    PhotosQuery
} from "~/types/api-gql";

export type PhotoList = PhotosQuery['photos']['nodes']
export type AggregateYear = AggregateDateTimeOriginalYearQuery['aggregateDateTimeOriginal']
export type AggregateMonth = AggregateDateTimeOriginalMonthQuery['aggregateDateTimeOriginal']
export type AggregateDate = AggregateDateTimeOriginalDateQuery['aggregateDateTimeOriginal']

