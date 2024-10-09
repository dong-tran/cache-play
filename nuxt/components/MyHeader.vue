<template>
  <div class="h-12 bg-gradient-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90%">
    <div class="h-full flex flex-row items-center justify-center space-x-3 mx-5">
      <a href="/" class="shadow-md rounded-md bg-slate-300 px-3">Index (No Cache)</a>
      <a href="/proxied" class="shadow-md rounded-md  bg-slate-300 px-3">Proxy Cache</a>
      <a href="/cached" class="shadow-md rounded-md  bg-slate-300 px-3">Page Cache</a>
      <a href="/nitro" class="shadow-md rounded-md  bg-slate-300 px-3">Nitro(Node Server) Cache</a>
      <a href="/api-cached" class="shadow-md rounded-md  bg-slate-300 px-3">API Response Cache</a>
      <a href="/redis-cached" class="shadow-md rounded-md  bg-slate-300 px-3">Go API(Redis) Cache</a>
    </div>
  </div>
</template>
