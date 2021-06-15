<template>
  <div>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
    >
      <!-- @select="handleSelect" -->
      <el-menu-item index="1">
        <router-link to="/">
          <!-- <el-button type="primary">Index</el-button> -->
          Index
        </router-link>
      </el-menu-item>
      <el-menu-item index="2">
        <router-link to="/books">
          <!-- <el-button>Index</el-button> -->
          Books
        </router-link>
      </el-menu-item>
      <el-submenu index="3">
        <template #title>我的工作台</template>
        <el-menu-item index="3-1">选项1</el-menu-item>
        <el-menu-item index="3-2">选项2</el-menu-item>
        <el-menu-item index="3-3">选项3</el-menu-item>
        <el-submenu index="3-4">
          <template #title>选项4</template>
          <el-menu-item index="3-4-1">选项1</el-menu-item>
          <el-menu-item index="3-4-2">选项2</el-menu-item>
          <el-menu-item index="3-4-3">选项3</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="4" disabled>未开启连接</el-menu-item>
      <el-menu-item index="5">
        <!-- <el-button>Home</el-button> -->
        <router-link to="/home"> Home </router-link>
      </el-menu-item>
    </el-menu>
    <router-view />
  </div>
</template>

<script lang='ts'>
import { defineComponent } from "vue";
export default defineComponent({
  name: "Nav",
  components: {},
  data() {
    return {
      activeIndex: "1",
      activeIndex2: "1",
    };
  },
  methods: {
    // handleSelect(key: any, keyPath: any) {
    //   console.log(key, keyPath);
    // },
  },
});
</script>

<style scoped>
.el-menu-demo {
  display: flex;
  justify-content: center;
}
</style>

