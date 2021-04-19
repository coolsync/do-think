<template>
  <h2>AsyncAddress component</h2>
  <h3>data: {{ data }}</h3>
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import axios from "axios";
export default defineComponent({
  name: "AsyncAddress",
  async setup() {
    const result = await axios.get("/data/address.json");
    return {
      data: result.data,
    };
  },
});
</script>
