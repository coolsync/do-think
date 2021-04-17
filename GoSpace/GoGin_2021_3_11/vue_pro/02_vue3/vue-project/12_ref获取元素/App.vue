<template>
  <h2>ref, 让输入框自动获取焦点</h2>
  <input type="text"><br>
  <input type="text" ref="inputRef" />
</template>

<script lang='ts'>
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  name: "App",
  // ref获取元素: 利用ref函数获取组件中的标签元素
  // 功能需求: 让输入框自动获取焦点

  setup() {
    const inputRef = ref<HTMLElement | null>(null);

    onMounted(() => {
      inputRef.value && inputRef.value.focus(); // auto get focus
    });
    return {
      inputRef,
    };
  },
});
</script>
