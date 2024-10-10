export default defineCachedEventHandler(async () => {
  const config = useRuntimeConfig()
  const data = await $fetch(`${config.apiBase}/random`)
  return data
},{ maxAge: 60 ,base: 'redis', swr: false})