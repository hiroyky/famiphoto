<template>
  <v-list density="compact" :nav="false" @click:open="onExpandCollapse">
    <v-list-item><slot name="title">{{ $t('dateTimeOriginal') }}</slot></v-list-item>

    <tree-list-item :model-value="year" v-for="year in years" :key="year.year" @click="onItemClick" :terminate="year.year===0">
      <template #title>{{ displayYear(year) }}</template>
      <template #count>{{year.num}}</template>

      <template v-for="month in months" :key="`${month.year}-${month.month}`">
        <tree-list-item :model-value="month" :sub-group="true" v-if="month.year === year.year" @click="onItemClick">
          <template #title>{{ $t('xMonth', { month: month.month }) }}</template>
          <template #count>{{month.num}}</template>

          <template v-for="date in dates" :key="`${date.year}-${date.month}-${date.date}`">
            <tree-list-item :model-value="date" :terminate="true" v-if="isTheDate(month, date)" @click="onItemClick(date)">
              <template #title>{{ t('xMonthDate', { month: date.month, date: date.date }) }}</template>
              <template #count>{{date.num}}</template>
            </tree-list-item>
          </template>
        </tree-list-item>
      </template>
    </tree-list-item>

  </v-list>
</template>

<script setup lang="ts">

import type {AggregateDate, AggregateMonth, AggregateYear} from "~/types/api-gql-alias";
import TreeListItem from "~/components/parts/TreeListItem.vue";
import type {YearMonthDateNum} from "~/types/types";

const {t} = useI18n({useScope: 'global'})

interface Props {
  years: AggregateYear| null
  months: AggregateMonth | null
  dates: AggregateDate | null
}

interface OpenCloseEvent {
  id: YearMonthDateNum
  path: YearMonthDateNum[]
  value: boolean
}

defineProps<Props>()
const emit = defineEmits<{
  expandYear: [number]
  expandMonth: [{ year: number, month: number }]
  itemClick: [YearMonthDateNum]
}>()

function onItemClick(val: YearMonthDateNum) {
  console.log(val)
  emit('itemClick', val)
}

function onExpandCollapse({id, path, value}:OpenCloseEvent) {
  if (!value) {
    return
  }
  if (path.length === 1) {
    emit('expandYear', id.year)
    return
  }
  if (path.length === 2 && id.month) {
    emit('expandMonth', { year: id.year, month:id.month})
    return
  }
}

function displayYear(year: {year: number}) {
  return year.year > 0 ? t('xYear', {year: year.year}) : t('unknown')
}

function isTheDate(a: {year: number, month: number}, b: {year: number, month: number}) {
  return a.year === b.year && a.month === b.month
}

</script>