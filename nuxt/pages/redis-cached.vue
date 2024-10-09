<template>
  <div class="flex flex-col items-center mt-5">
    <UPagination v-model="page" :page-count="pageCount" :total="total" show-last show-first :max="10" />
    <UTable :rows="data" class="w-5/6 text-left max-h-screen" :loading="status === 'pending'" :columns="columns" />
  </div>

</template>

<script setup>
const page = ref(1)
const pageCount = 100
const total = 1000000
let apiBase = useRuntimeConfig().public.apiBase
if (import.meta.server) {
  apiBase = useRuntimeConfig().apiBase
}
const path = computed(() => apiBase + '/users/' + page.value)
const { data, status } = useFetch(path,{default:()=>[]})
const columns = [{
  label: 'ID',
  key: 'ID'
}, {
  label: 'Email',
  key: 'Email'
}, {
  label: 'Password',
  key: 'Password'
}]
</script>
