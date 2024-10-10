import { defineStore } from 'pinia'

export const useCounter = defineStore('counter', () => {
  const count = ref(60)
  function down() {
    if (count.value > 0) {
      count.value--
    }

  }
  function set(v: number) {
    count.value = v
  }
  return { count, down, set }
})