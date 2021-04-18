<template>
    <h2>Son Component</h2>
    <hr>
    <GrandSon />
</template>

<script lang='ts'>
// import GrandSon from './components/GrandSon.vue'    // err
import GrandSon from "./GrandSon.vue";  // current dir,
import {defineComponent} from 'vue'
export default defineComponent({
  name: 'Son',

  components: {
      GrandSon,
  },

//   setup() {
//       const color = ref('red')
//       return {
//           color
//       }
//   }
})
</script>
