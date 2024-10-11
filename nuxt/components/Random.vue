<template>
  <div class="neonText mt-12 text-center text-9xl">
    {{ code }}
  </div>
</template>

<script setup>
let code = ''
await callOnce(async () => {
  code = random(6)
})
</script>
