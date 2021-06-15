<template>
  <h2>use toRefs</h2>
  <!-- <h3>name: {{ state.name }}</h3>
  <h3>age: {{ state.age }}</h3> -->

  <h3>name: {{ name }}</h3>
  <h3>age: {{ age }}</h3>
</template>

<script lang='ts'>
import { defineComponent, reactive, toRefs } from "vue";

// 把一个响应式对象转换成普通对象，该普通对象的每个 property 都是一个 ref
// 应用: 当从合成函数返回响应式对象时，toRefs 非常有用，这样消费组件就可以在不丢失响应式的情况下对返回的对象进行分解使用
// 问题: reactive 对象取出的所有属性值都是非响应式的
// 解决: 利用 toRefs 可以将一个响应式 reactive 对象的所有原始属性转换为响应式的 ref 属性

function useRefs() {
  const state3 = reactive({
    name: "mark",
    age: 30,
  });
  return toRefs(state3);
}

export default defineComponent({
  name: "App",
  setup() {
    const state = reactive({
      name: "mark",
      age: 30,
    });
    const state2 = toRefs(state);

    // out intro
    const {name, age} = useRefs();

    setInterval(() => {
      // state.name += "=="
      // state2.name.value += "=="
      name.value += '++'
      console.log("----------");
    }, 1000);

    return {
      // state,
      // ...state // {name, age}
      // ...state2,
      name,
      age,
    };
  },
});
</script>
