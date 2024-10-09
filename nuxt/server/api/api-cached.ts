export default defineCachedEventHandler(async () => {
  const config = useRuntimeConfig()
  const data = await $fetch(`${config.public.apiBase}/random`)
  return data
},{ maxAge: 30 ,base: 'redis'})