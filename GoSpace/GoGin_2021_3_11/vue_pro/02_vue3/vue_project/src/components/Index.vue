<template>
  <div class="div_index">
    <!-- <router-link to="http://localhost:9000/chapter11/api_axios">
    <el-button>测试 axios api</el-button>
    </router-link>
    <router-view></router-view> -->
    <el-button @click="getAxios">测试 axios api</el-button>
  </div>
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import { ElButton } from "element-plus";
import axios from "axios";
export default defineComponent({
  name: "Index",
  components: {
    ElButton,
  },

  setup() {
    const getAxios = ()=> {
       axios
      .get("http://localhost:9000/chapter11/api_axios")
      .then((response) => {
        console.log(response.data);
      })
      .catch((e) => {
        console.log(e.message || "未知错误");
      });
    }
   

    return {
      getAxios,
    };
  },
});
</script>

<style scoped>
.div_index {
  display: flex;
  justify-content: center;
}
/* 
.inline-block-center {
    text-align: center;
}
.inline-block-center div {
    display: inline-block;
}
 */
</style>