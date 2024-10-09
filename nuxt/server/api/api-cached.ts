export default defineCachedEventHandler(async () => {
  const config = useRuntimeConfig()
  const data = await $fetch(`${config.apiBase}/random`)
  return data
},{ maxAge: 30 ,base: 'redis'})