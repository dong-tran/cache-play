import redis from "unstorage/drivers/redis";

export default defineNitroPlugin(() => {
  const storage = useStorage();
  const config = useRuntimeConfig()
  storage.mount(
    "/redis",
    redis({
      host: config.redisHost,
      port: Number(config.redisPort),
    })
  );
});