<template>
  <h2>this books page</h2>
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import $axios from "../axios";
export default defineComponent({
  name: "Books",
  created() {
    this.getBooks();
  },
  methods: {
    getBooks() {
      $axios
        .get("/chapter11/get_books")
        .then((res) => {
          console.log(res);
        })
        .catch((e) => {
          console.log(e.message || "未知错误");
        });
    },
  },
});
</script>

<style scoped>
h2 {
  text-align: center;
}
</style>