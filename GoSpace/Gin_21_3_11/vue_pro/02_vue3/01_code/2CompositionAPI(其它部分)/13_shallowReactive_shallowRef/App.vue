<template>
  <h2>shallowReactive与shallowRef</h2>
  <h3>m1: {{ m1 }}</h3>
  <h3>m2: {{ m2 }}</h3>
  <h3>m3: {{ m3 }}</h3>
  <h3>m4: {{ m4 }}</h3>
  <button @click="update">Update</button>
</template>

<script lang='ts'>
// shallowReactive与shallowRef
//   shallowReactive: 只处理了对象内最外层属性的响应式(也就是浅响应式)
//   shallowRef: 只处理了value的响应式, 不进行对象的reactive处理
// 总结:
//   reactive与ref实现的是深度响应式, 而shallowReactive与shallowRef是浅响应式
//   什么时候用浅响应式呢?
//     一般情况下使用ref和reactive即可,
//     如果有一个对象数据, 结构比较深, 但变化时只是外层属性变化 ===> shallowReactive
//     如果有一个对象数据, 后面会产生新的对象来替换 ===> shallowRef

import {
  defineComponent,
  reactive,
  shallowReactive,
  ref,
  shallowRef,
} from "vue";

export default defineComponent({
  name: "App",

  setup() {
    // name, car name all can change
    const m1 = reactive({
      name: "mark",
      age: 30,
      car: {
        name: "aodi",
        color: "black",
      },
    });
    // name can change, car name can not change
    const m2 = shallowReactive({
      name: "mark",
      age: 30,
      car: {
        name: "aodi",
        color: "black",
      },
    });

  // value: name can change, car name can not change
    const m3 = ref({
      name: "mark",
      age: 30,
      car: {
        name: "aodi",
        color: "black",
      },
    });
    // name, car name all not can change
    const m4 = shallowRef({
      name: "mark",
      age: 30,
      car: {
        name: "aodi",
        color: "black",
      },
    });
    const update = () => {
      // m1.name += "++"
      // m1.car.name += "++"
      // m2.name += "++";
      // m2.car.name += "++";
      // m3.value.name += "++";
      // m3.value.car.name += "++";
      m4.value.name += "++"; // value is obj, not proxy_obj
      // m4.value.car.name += "++";
      console.log(m3, m4);
    };
    return {
      m1,
      m2,
      m3,
      m4,
      update,
    };
  },
});
</script>
