<template>
  <h1>use reactive</h1>
  <h2>name: {{ user.name }}</h2>
  <h2>age: {{ user.age }}</h2>
  <h2>wife: {{ user.wife }}</h2>
  <button @click="update_user">update user data</button>
</template>

<script lang='ts'>
import { defineComponent, reactive } from "vue";

/* 
作用: 定义多个数据的响应式
const proxy = reactive(obj): 接收一个普通对象然后返回该普通对象的响应式代理器对象
响应式转换是“深层的”：会影响对象内部所有嵌套的属性
内部基于 ES6 的 Proxy 实现，通过代理对象操作源对象内部数据都是响应式的
*/

export default defineComponent({
  name: "App",

  setup() {
    // class User {
    //   name: string;
    //   age: number;
    //   wife: User;
    // }

    const obj = {
      name: "paul",
      age: 30,
      wife: {
        name: "甜甜",
        age: 18,
        cars: ["bengci", "baoma", "aodi"],
      },
    };

    const user = reactive<any>(obj); // obj 是 target 被代理对象， user 是 handler 代理对象

    const update_user = () => {
      // user.name+='======';
      // user.age += 2
      // user.wife.name += '+++'
      // user.wife.age += 2
      // user.wife.cars[0] = 'sanmaladi'

      // only operate proxy obj user, 才能 render page
      user.name = 'mark'; // by user proxy_obj modify target_obj property value
      delete user.age;

      // 通过 proxy_obj modify target_obj operate child_obj data
      user.wife.cars[1] = 'samaladi';
      user.wife.cars[3] = 'aotuo';
    };
    return {
      user,
      update_user
    };
  },
});
</script>


