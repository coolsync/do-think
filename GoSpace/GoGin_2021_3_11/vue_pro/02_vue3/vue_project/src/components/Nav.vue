<template>
  <!-- <h2>App components</h2> -->
  <!-- <div>
    <router-link to="/">
      <el-button>Index</el-button>
    </router-link>
    <router-link to="/home">
      <el-button type="primary">Home</el-button>
    </router-link>
    <router-view />
  </div> -->
  <div>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">
        <router-link to="/">
          <!-- <el-button>Index</el-button> -->
          Index
        </router-link>
      </el-menu-item>
      <el-submenu index="2">
        <template #title>我的工作台</template>
        <el-menu-item index="2-1">选项1</el-menu-item>
        <el-menu-item index="2-2">选项2</el-menu-item>
        <el-menu-item index="2-3">选项3</el-menu-item>
        <el-submenu index="2-4">
          <template #title>选项4</template>
          <el-menu-item index="2-4-1">选项1</el-menu-item>
          <el-menu-item index="2-4-2">选项2</el-menu-item>
          <el-menu-item index="2-4-3">选项3</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="3" disabled>未开启连接</el-menu-item>
      <el-menu-item index="4">
        <!-- <a href="https://www.ele.me" target="_blank">订单管理</a> -->
        <el-menu-item index="4">
          <!-- <el-button type="primary">Home</el-button> -->
          <router-link to="/home"> Home </router-link>
        </el-menu-item>
      </el-menu-item>
    </el-menu>
    <router-view />
  </div>
</template>
<script lang='ts'>
import { defineComponent } from "vue";
// import Home from "./components/Home.vue";
// import {
//   ElButton,
//   ElSelect,
//   ElMain,
//   ElMenu,
//   ElMenuItem,
//   ElMenuItemGroup,
// } from "element-plus";
export default defineComponent({
  name: "Nav",
  components: {
    // ElButton.name: Element,
    // ElSelect,
    // ElementPlus,
    // ElMain,
    // ElMenu,
    // ElMenuItem,
    // ElMenuItemGroup,
  },
  data() {
    return {
      activeIndex: "1",
      activeIndex2: "1",
    };
  },
  methods: {
    handleSelect(key:any, keyPath:any) {
      console.log(key, keyPath);
    },
  },
});
</script>

<style scoped>
/* body {
    margin: 100px auto;
  } */
.el-menu-demo {
  /* margin: 100px auto; */
  display: flex;
  justify-content: center;
}
</style>

