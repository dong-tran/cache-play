<template>
  <NuxtLayout>
    <MyHeader />
    <div class="min-h-96">
      <div class="absolute inset-0 opacity-30 bg-noise"></div>
      <NuxtPage />
    </div>
    <MyFooter />
  </NuxtLayout>
</template>

<style>
.bg-noise {
  background-image: url('/bg.png');
}
</style>
