<template>
  <div
    class="flex flex-row items-center h-12 bg-gradient-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90%">
    <UChip :text="counter.count" size="2xl" :ui="{ base: 'mt-2'}">
      <UIcon name="subway:alam" class="text-blue-200 w-8 h-8 mx-3" @click="reset" />
    </UChip>
    <div class="h-full flex flex-row grow items-center justify-center space-x-3 mx-5">
      <a href="/" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">Index (No Cache)</a>
      <a href="/proxied" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">Proxy Cache</a>
      <a href="/cached" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">Page Cache</a>
      <a href="/nitro" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">Nitro(Node Server) Cache</a>
      <a href="/api-cached" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">API Response Cache</a>
      <a href="/redis-cached" class="hover:animate-bounce shadow-md rounded-md bg-gray-50 hover:bg-emerald-200 px-3">Go API(Redis) Cache</a>
    </div>
  </div>
</template>

<script setup>
const counter = useCounter()

const reset = () => {
  const t = useCookie('init', { maxAge: 60 })
  t.value = undefined
  counter.set(0)
}
onMounted(() => {
  const t = useCookie('init', { maxAge: 60 })
  let count = 60
  const cur = Math.round(Date.now() / 1000)
  if (!t.value) {
    t.value = cur
  } else {
    count = count - (cur - t.value)
  }
  counter.set(count)
  setInterval(counter.down, 1000)
})
</script>
