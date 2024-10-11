<template>
  <div class="neonText mt-12 text-center text-9xl">
    {{ data }}
  </div>
</template>

<script setup>
const {data} = useFetch('/api/api-cached')
</script>
