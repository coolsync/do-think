<template>
  <h2>toRaw 与 markRaw</h2>
  <h3>state: {{ state }}</h3>
  <button @click="testToRaw">test toRaw</button>
  <button @click="testMarkRaw">test markRaw</button>
</template>

<script lang='ts'>
import { defineComponent, markRaw, reactive, toRaw } from "vue";

interface UserInfo {
  name: string;
  age: number;
  likes?: string[];
}

export default defineComponent({
  name: "App",

  setup() {
    const state = reactive<UserInfo>({
      name: "mark",
      age: 25,
    });

    const testToRaw = () => {
      const user = toRaw(state); // 这是一个还原方法，可用于临时读取，访问不会被代理/跟踪，写入时也不会触发界面更新
      user.name += "++";
      console.log(user.name);
    };

    const testMarkRaw = () => {
      // state.likes = ["chi", "he"];
      const likes = ["chi", "he"];
      state.likes = markRaw(likes); // 标记一个对象，使其永远不会转换为代理。返回对象本身 // likes数组就不再是响应式的了
      setInterval(() => {
        // console.log(state.likes) // obj
        state.likes[0] += "++";
        console.log("----");
      }, 1000);
    };
    return {
      state,
      testToRaw,
      testMarkRaw,
    };
  },
});
</script>
