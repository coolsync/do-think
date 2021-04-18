<template>
  <h2>ref and reactive detail</h2>
  <h3>m1: {{ m1 }}</h3>
  <h3>m2: {{ m2 }}</h3>
  <h3>m3: {{ m3 }}</h3>
  <button @click="update">update data</button>
</template>

<script lang='ts'>
import { defineComponent, ref, reactive } from "vue";

export default defineComponent({
  name: "App",

// Vue3的 composition API中2个最重要的响应式API (ref, reactive)
// ref用来处理基本类型数据, reactive用来处理对象(递归深度响应式)
// 如果用ref对象/数组, 内部会自动将对象/数组转换为reactive的代理对象
// ref内部: 通过给value属性添加getter/setter来实现对数据的劫持
// reactive内部: 通过使用Proxy来实现对对象内部所有数据的劫持, 并通过Reflect操作对象内部数据
// ref的数据操作: 在js中要.value, 在模板中不需要(内部解析模板时会自动添加.value)

  setup() {
    const m1 = ref("abc");
    const m2 = reactive({
      name: "mark",
      wife: {
        name: "jerry",
      },
    });
    const m3 = ref({
      name: "mark",
      wife: {
        name: "jerry",
      },
    });

    const update = () => {
      m1.value += "==";
      m2.wife.name += "==";
      m3.value.wife.name += "==";
      // ref store obj, 通过 reactive 处理后 成 Proxy_obj, 在通过 Reflect 处理 obj 内部 data 
      console.log(m3.value);
    };
    return {
      m1,
      m2,
      m3,
      update,
    };
  },
});
</script>