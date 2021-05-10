<template>
  <div>Detail Page</div>
  <!-- <span>{{this.$route.param.id}}</span> -->
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import $axios from "../axios";
export default defineComponent({
  name: "book_detail",

  mounted() {
    this.get_book_detail();
  },

  methods: {
    get_book_detail() {
      $axios
        .get("/chapter11/book_detail", {
            params: {
                id: 1,
            }
        })
        .then((res) => {
          console.log(res);
        })
        .catch((e) => {
          console.log(e.message || "Unknown");
        });
    },
  },
});
</script>