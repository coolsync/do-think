<template>
  <h2>Child component</h2>
  <h2>msg:{{ msg }}</h2>
  <h2>count:{{ count }}</h2>
  <!-- <h2>showMsg2:{{ showMsg2() }}</h2>
  <h2>showMsg1:{{ showMsg1() }}</h2> -->
</template>

<script lang='ts'>
import { defineComponent } from "vue";
/* 
setup细节:
1. setup执行的时机:
在beforeCreate之前执行(一次), 此时组件对象还没有创建
this是undefined, 不能通过this来访问data/computed/methods / props
其实所有的composition API相关回调函数中也都不可以
2. setup的返回值:
一般都返回一个对象: 为模板提供数据, 也就是模板中可以直接使用此对象中的所有属性/方法
返回对象的属性 会与 data函数返回对象的属性 合并成为组件对象的属性
返回对象中的方法 会与 methods中的方法 合并成功组件对象的方法
如果有重名, setup优先
注意:
一般不要混合使用: methods中可以访问setup提供的属性和方法, 但在setup方法中不能访问data和methods
setup不能是一个async函数: 因为返回值不再是return的对象, 而是promise, 模板看不到return对象中的属性数据
*/

export default defineComponent({
  name: "Child",
  props: ["msg"],

  // life cycle render before
  beforeCreate() {
    // beforeCreate after setup run
    console.log("beforeCreate run ... ");
  },
  mounted() {
    console.log("mounted run ... ");
  },
  setup() {
    console.log("setup run ... ", this); // obj not create, this is undefined
    const showMsg2 = () => {
      console.log("showMsg2 method in setup ... ");
    };
    return {
      showMsg2,
    };
  },
  data() {
    console.log("data run ... ");
    const count = 10;
    console.log(this)
    return {
      count,
    };
  },
  methods: {
    showMsg1: () => {
      console.log("showMsg1 method in methods ... ");
    },
  },
});
</script>

<style scoped lang='scss'>
</style>