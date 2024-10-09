<template>
  <div style="font-size: xx-large;" class="neonText mt-12 text-center">
    {{ code }}
  </div>
</template>

<script setup>
let code = ''
await callOnce(async () => {
  code = random(10)
})
</script>

<style>
.neonText {
  color: hsl(244, 83%, 53%);
  text-shadow:
    0 0 7px #fff,
    0 0 10px #fff,
    0 0 21px #210fee,
    0 0 42px #0fa,
    0 0 82px #0fa,
    0 0 92px #0fa,
    0 0 102px #0fa,
    0 0 151px #0fa;
}
</style>