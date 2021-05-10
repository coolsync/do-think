<template>
  <h2>this books page</h2>

  <div id="book">
    <table>
      <thead>
        <th>ID</th>
        <th>BookName</th>
        <th>Author</th>
      </thead>
      <tbody v-for="book in books" :key="book">
        <td>{{book.id}}</td>
        <td>
          <router-link :to="{name: 'book_detail', params: { id: book.id }}">
          {{book.name}}
          </router-link>

        </td>
        <td>{{book.author}}</td>
      </tbody>
    </table>
  </div>
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import $axios from "@/axios";
export default defineComponent({
  name: "Books",
  data() {
    return {
      books: []
    }
  },
  created() {
    this.getBooks();
  },
  methods: {
    getBooks() {
      $axios
        .get("/chapter11/get_books")
        .then((res) => {
          // console.log(res.data.books);
          this.books = res.data.books;
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

#book {
  margin: 0 45%;
}
</style>