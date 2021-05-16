<template>
  <div class="div_home">
    <h2>This Home 页面</h2>
  </div>
</template>

<script lang='ts'>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Home",
});
</script>

<style scoped>
.div_home {
  display: flex;
  justify-content: center;
}
</style>