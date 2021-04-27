<template>
  <div class="div_index">
    <!-- <router-link to="http://localhost:9000/chapter11/api_axios">
    <el-button>测试 axios api</el-button>
    </router-link>
    <router-view></router-view> -->
    <el-button @click="getAxios">测试 axios api</el-button>
    <br />
    <span>id: {{ id }}</span
    ><br />
    <span>name: {{ name }}</span
    ><br />
    <span>age: {{ age }}</span
    ><br />
    <span>code: {{ code }}</span
    ><br />
    <span>msg: {{ msg }}</span
    ><br />

    <span v-for="(i, v) in arrs" :key="i">
      <span>{{ i }}: {{ v }}, </span>
    </span>
    <br />

    <span v-for="arr in arrs_s" :key="arr">
      <span>{{ arr.id }}</span>
      <span>{{ arr.name }}</span>
      <span>{{ arr.age }}</span>
      <br />
    </span>
    <br />
    <span>
      <span>{{ map_s.user }}</span>

      <!-- <span>{{map_s.user.id }}</span>,  -->
      <!-- <span>{{map_s.user.name}}</span>, -->
      <!-- <span>{{map_s.user.age}}</span>, -->

      <br />
    </span>
  </div>
</template>

<script lang='ts'>
import { defineComponent} from "vue";
import { ElButton } from "element-plus";
import $axios from "../axios";
// import router from "@/router";

export default defineComponent({
  name: "Index",
  components: {
    ElButton,
  },
  data: () => {
    return {
      code: "",
      msg: "",
      id: "",
      name: "",
      age: "",
      arrs: [],
      arrs_s: [], // arrs struct
      map_s: {},
    };
  },
  mounted() {
    this.getAxios();
  },

  methods: {
    getAxios() {
      // axios
      //  .get("http://localhost:9000/chapter11/api_axios")
      $axios
        .get("/chapter11/api_axios")
        .then((res) => {
          console.log(res.data.arrs_s);
          console.log(res.data.map_s);
          this.code = res.data.code;
          this.msg = res.data.msg;
          this.id = res.data.user.id;
          this.name = res.data.user.name;
          this.age = res.data.user.age;
          this.arrs = res.data.arrs;
          this.arrs_s = res.data.arrs_s;
          this.map_s = res.data.map_s;
          // router.push('/home')
          // window.location.href = '/home';
        })
        .catch((e) => {
          alert(e.message || "未知错误");
        });
    },
  },

  setup() {
    return {};
  },
});
</script>

<style scoped>
.div_index {
  text-align: center;
}
/* .div_index {
  display: flex;
  justify-content: center;
} */
/* .inline-block-center div {
  display: inline-block;
} */
</style>