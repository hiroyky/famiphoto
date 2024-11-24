<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" v-for="photoGroup in photosPerDateTimeOriginal" :key="photoGroup.id">
        <v-card>
          <v-card-title>{{ photoGroup.title }}</v-card-title>
          <ul class="photo_list">
            <li v-for="item in photoGroup.photos" :key="item.id" class="photo_list__item">
              <img :src="item.thumbnailUrl" :alt="item.name" class="photo_list__item__img" @click="onPhotoItemClick(item)">
            </li>
          </ul>
        </v-card>
        <div v-intersect="onIntersectNext">
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">

import type {PhotoList} from "~/types/api-gql-alias";
const { t } = useI18n({useScope: 'global'})

interface Props {
  value: PhotoList | null
}

interface PhotoGroup {
  id: string
  title: string
  photos: PhotoList
}

const props = withDefaults(defineProps<Props>(), {
  value: () => [],
})

const emit =defineEmits<{
  photoClick: [ id: string ]
  intersectNext: []
}>()

const photosPerDateTimeOriginal = computed(() => {
  if (!props.value){
    return []
  }
  const photos = props.value
  const groups: PhotoGroup[] = []
  for (let i = 0; i < photos.length; ++i) {
    if (isDateTimeOriginalBorder(photos, i)) {
      const dateTimeOriginal = new Date(photos[i].dateTimeOriginal)
      const group: PhotoGroup = {
        id: dateTimeOriginal.toDateString(),
        title: t(
            'xYearMonthDate', {
              year: dateTimeOriginal.getFullYear(),
              month: dateTimeOriginal.getMonth() + 1,
              date: dateTimeOriginal.getDate()
            }
        ),
        photos: [ ]
      }
      groups.push(group)
    }

    const group = groups[groups.length - 1]
    group.photos.push(photos[i])
  }
  return groups
})

function isDateTimeOriginalBorder(photos: PhotoList, i: number): boolean {
  if (i===0 || !photos[i] || !photos[i - 1]) {
    return true
  }
  const date1 = new Date(photos[i].dateTimeOriginal)
  const date2 = new Date(photos[i-1].dateTimeOriginal)
  return date1.getDate() !== date2.getDate()
}

function onPhotoItemClick(el :{ id: string }) {
  emit('photoClick', el.id)
}

function onIntersectNext(isIntersecting: boolean, _1: IntersectionObserverEntry[], _2: IntersectionObserver,) {
  if (isIntersecting) {
    emit('intersectNext')
  }
}

</script>

<style lang="scss" scoped>
.photo_list {
  list-style: none;
  display: flex;
  flex-wrap: wrap;
  gap:10px;
  margin: 13px auto;

  &__item {
    flex: 1 1 auto;
    height: 300px;
    position: relative;

    &__img {
      width: 100%;
      height: 100%;
      vertical-align: middle;
      object-fit: contain;
      cursor: pointer;
    }
  }
}

</style>