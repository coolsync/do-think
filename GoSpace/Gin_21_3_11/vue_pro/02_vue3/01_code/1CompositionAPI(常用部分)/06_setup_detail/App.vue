<template>
  <h2>App Component</h2>
  <h3>msg: {{ msg }}</h3>
  <button @click="msg += '++'">Update Data</button>
  <hr />
  <!-- msg 同步到 Child -->
  <Child :msg="msg" msg2="cba" @fn="fn"/> 
</template>

<script lang='ts'>
import { defineComponent, ref } from "vue";

// 引入 Child component
import Child from "./components/Child.vue";
// use component

export default defineComponent({


  name: "App",

  // register component
  components: {
    Child,
  },

  setup() {
    const msg = ref("what are you no sai lei");

    const fn = (content: string) => {
      msg.value += content
    }

    return {
      msg,
      fn,
    };
  },
});
</script>