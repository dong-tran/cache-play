export default defineEventHandler(async () => {
  const config = useRuntimeConfig()
  const data = await $fetch(`${config.apiBase}/time`)
  return data
})