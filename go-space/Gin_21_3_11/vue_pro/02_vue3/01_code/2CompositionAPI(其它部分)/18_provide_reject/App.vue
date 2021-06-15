<template>
  <h1>provide 与 inject</h1>
  <h2>App 组件</h2>
  <p>当前颜色: {{ color }}</p>
  <button @click="color = 'red'">红</button>
  <button @click="color = 'yellow'">黄</button>
  <button @click="color = 'blue'">蓝</button>
  <hr />
  <Son />
</template>

<script lang='ts'>
import { defineComponent, provide, ref } from "vue";
import Son from "./components/Son.vue";
export default defineComponent({
  name: "App",
  components: {
    Son,
  },
  setup() {
    const color = ref("red");
    provide("color", color);
    return {
      color,
    };
  },
});
</script>
