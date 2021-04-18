<template>
  <h2>GrandSon Component</h2>
  <h3 :style="{ color }">GrandSon: {{ color }}</h3>
</template>

<script lang='ts'>
import { defineComponent, inject } from "vue";

export default defineComponent({
  name: "GrandSon",
  setup() {
    const color = inject("color");
    return {
      color,
    };
  },
});
</script>
