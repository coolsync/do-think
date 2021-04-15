<template>
  <h2>Father component</h2>
  <h3>msg:{{msg}}</h3>
  <button @click="fn('======')">update data</button>
  <hr>
  <child :msg="msg" />
</template>

<script lang='ts'>
import { defineComponent, ref } from "vue";

// import Child component
import child from "./components/Child.vue";

export default defineComponent({
  name: "App",

  // register component
  components: {
    child,
  },

  // setup render data
  setup() {
    // def ref obj
    const msg = ref('what are you silly');

    function fn(content: string) {
      msg.value += content 
    } 
    return {
      msg,
      fn,
    }
  }
});
</script>


