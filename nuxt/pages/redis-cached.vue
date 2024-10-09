<template>
  <div class="flex flex-col items-center mt-5">
    <UPagination v-model="page" :page-count="pageCount" :total="total" show-last show-first :max="10" />
    <UTable :rows="data" class="w-5/6 text-left max-h-screen" :loading="status === 'pending'"/>
  </div>

</template>

<script setup>
const page = ref(1)
const pageCount = ref(100)
const total = ref(1000000)
const config = useRuntimeConfig()
const path = computed(() => config.public.apiBase + '/users/' + page.value)
const { data, status } = useFetch(path)
</script>
