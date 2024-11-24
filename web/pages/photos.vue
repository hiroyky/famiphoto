<template>
  <div>
    <navigation>
      <data-tree-list
        :years="dateTimeOriginalYears"
        :months="dateTimeOriginalMonths"
        :dates="dateTimeOriginalDates"
        @expand-year="onExpandYear"
        @expand-month="onExpandMonth"
        @item-click="onDataTreeItemClick">
      </data-tree-list>
    </navigation>
    <nuxt-page />
  </div>
</template>

<script setup lang="ts">

import type {AggregateDate, AggregateMonth} from "~/types/api-gql-alias";
import DataTreeList from "~/components/modules/DataTreeList.vue";
import type {YearMonthDateNum} from "~/types/types";
import Navigation from "~/components/parts/Navigation.vue";

definePageMeta({
  layout: 'photos'
})

useHead({
  title: useAppConfig().AppName,
})


const photoListStore = usePhotoListStore()
const photoQueryStore = usePhotoQueryStore()

const dateTimeOriginalMonths = ref<AggregateMonth>([])
const dateTimeOriginalDates = ref<AggregateDate>([])

const { data: dateTimeOriginalYears } = useAsyncData(async ()=> {
  const years = await photoListStore.aggregateDateTimeOriginalYear()
  return years ? years : []
})

useAsyncData(async () => {
  const q = photoQueryStore.parseQuery()
  await photoListStore.getPhotos(q)
})


async function onExpandYear(year: number) {
  const res = await photoListStore.aggregateDateTimeOriginalYearMonth(year)
  res.map((i) => {
    const exist = dateTimeOriginalMonths.value.find((v) => v.year == i.year && v.month === i.month)
    if (exist) {
      exist.num = i.num
    } else {
      dateTimeOriginalMonths.value.push(i)
    }
  })
}

async function onExpandMonth(val: {year:number, month:number}) {
  const res = await photoListStore.aggregateDateTimeOriginalYearMonthDate(val.year, val.month)
  res.map((i) => {
    const exist = dateTimeOriginalDates.value.find((v) => v.year == i.year && v.month === i.month && v.date === i.date)
    if (exist) {
      exist.num = i.num
    } else {
      dateTimeOriginalDates.value.push(i)
    }
  })
}

async function onDataTreeItemClick(val: YearMonthDateNum) {
  await photoQueryStore.appendRefine({
    year: val.year,
    month: val.month,
    date: val.date
  })
  await photoListStore.getPhotos(photoQueryStore.parseQuery())
}

</script>
