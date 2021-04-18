<template>
  <h2>Custom Hook Function</h2>
  <h2>x: {{ x }}, y: {{ y }}</h2>
  <hr />
  <h3 v-if="loading">loading...</h3>
  <h3 v-else-if="errMsg">{{ errMsg }}</h3>
  <ul v-else>
    <li>id: {{ ret.id }}</li>
    <li>name: {{ ret.name }}</li>
    <li>distance: {{ ret.distance }}</li>
  </ul>
  <!-- array -->
  <hr />
  <ul v-for="p in ret" :key="p.id">
    <li>id: {{ p.id }}</li>
    <li>title: {{ p.title }}</li>
    <li>price: {{ p.price }}</li>
  </ul>
</template>

<script lang='ts'>
import { defineComponent, watch } from "vue";

// 使用Vue3的组合API封装的可复用的功能函数
// 自定义hook的作用类似于vue2中的mixin技术
// 自定义Hook的优势: 很清楚复用功能代码的来源, 更清楚易懂
import useMousePosition from "./hooks/useMousePosition";
import useRequset from "./hooks/useReq";

export default defineComponent({
  name: "App",

  // setup
  setup() {
    // 需求1: 收集用户鼠标点击的页面坐标
    const { x, y } = useMousePosition();

    // 需求2: 封装发ajax请求的hook函数
    interface AddrRequest {
      id: number;
      name: string;
      distance: string;
      length: number;
    }

    interface ProductRequest {
      id: string;
      title: string;
      price: number;
      length: number;
    }
    // const { loading, errMsg, ret } = useRequset<AddrRequest>("/data/addr.json"); // get obj
    const { loading, errMsg, ret } = useRequset<ProductRequest>("/data/product.json"); // get arr

    watch(ret, () => {
      if (ret.value) {  // if ret.value is not null
        console.log(ret.value.length);
        // console.log(ret.value);  
      }
    });

    return {
      x,
      y,
      loading,
      errMsg,
      ret,
    };
  },
});
</script>
