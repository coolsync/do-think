<template>
  <h2>App Father Component</h2>
  <button @click="isShow=!isShow">trigger</button>
  <hr>
  <Child  v-if="isShow" />
</template>

<script lang='ts'>
// intro Child Component
import Child from './components/Child.vue'

import {defineComponent, ref} from 'vue'

export default defineComponent({
  name: 'App',

  // Register Component
  components: {
    Child,
  },

  setup() {
    const isShow = ref(true)
    return {
      isShow: isShow
    }
  }
})
</script>
