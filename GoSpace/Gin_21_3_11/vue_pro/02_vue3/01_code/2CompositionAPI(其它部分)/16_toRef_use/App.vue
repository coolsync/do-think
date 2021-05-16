<template>
  <h2>toRef And Ref</h2>
  <h3>state: {{ state }}</h3>
  <h3>age: {{ age }}</h3>
  <h3>money: {{ money }}</h3>
  <button @click="update">update data</button>
  <hr />
  <Child :age="age" />
</template>

<script lang='ts'>
// toRef:
// 为源响应式对象上的某个属性创建一个 ref对象, 二者内部操作的是同一个数据值, 更新时二者是同步的
// ref 区别: 拷贝了一份新的数据值单独操作, 更新时相互不影响
// 应用: 当要将某个 prop 的 ref 传递给复合函数时，toRef 很有用
import Child from "./components/Child.vue";
import { defineComponent, reactive, ref, toRef } from "vue";

export default defineComponent({
  name: "App",

  components: {
    Child,
  },

  setup() {
    const state = reactive({
      name: "mark",
      age: 28,
      money: 1000,
    });

    // use ref
    const age = ref(state.age);
    // use toRef
    const money = toRef(state, "money");

    const update = () => {
      age.value += 2;
      money.value += 100;
    }
    return {
      state,
      age,
      money,
      update
    };
  },
});
</script>

