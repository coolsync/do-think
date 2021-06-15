<template>
  <h2>Child Component</h2>
  <h3>msg: {{ msg }}</h3>
  <h3>count: {{ count }}</h3>
  <!-- Get passing msg2 -->
  <h3>msg2: {{ $attrs.msg2 }}</h3>
  <h3>==================</h3>
  <h3>m: {{ m }}</h3>
  <h3>n: {{ n }}</h3>
  <!-- Get App 分发的 fn -->
  <button @click="update_num">update Child num</button>
</template>

<script lang='ts'>
import { defineComponent, ref } from "vue";

// setup 细节：
// 1. setup 在beforeCreate之前执行(一次), 此时组件对象还没有创建, this = undefined, 不能使用
// 2. this是undefined, 不能通过this来访问data/computed/methods / props
// 3. 其实所有的composition API相关回调函数中都不能

// setup的返回值:
// 1. setup 一般返回一个对象: 为模板提供数据, 模板能 直接使用 setup对象所有属性/方法
// 2. setup 返回对象的属性 与 data_Function 返回对象的属性 合并成 组件对象属性
// 3. setup 返回对象的方法 与 methods 方法 合并成 组件对象方法
// 4. 如果有重名, setup优先
// 注意:
// 5. 不要混合使用: methods 能访问 setup 属性和方法, setup 方法不能访问 data 和 methods
// 6. setup不能是一个async Fn: 返回对象 不再是 return 对象, 而是promise, 模板看不到return对象中的属性数据

// setup parameters:
// 1. setup(props, context) / setup(props, {attrs, slots, emit})
// 2. props: 包含props配置声明且传入了的所有属性的对象
// 3. attrs: 包含没有在props配置中声明的属性的对象, 相当于 this.$attrs
// 4. slots: 包含所有传入的插槽内容的对象, 相当于 this.$slots
// 5. emit: 用来分发自定义事件的函数, 相当于 this.$emit

export default defineComponent({
  name: "Child",
  props: ["msg"], // Get App_component passing_msg
  emits: ["fn"], // 可选的, 声明了利于阅读, 而且对分发的事件数据进行校验

  // Lifecycle hook, 在实例初始化之后，数据观测 (data observer) 和 event/watcher 事件配置之前被调用。
  beforeCreate() {
    console.log("beforeCreate run ... ");
    // console.log(this) // 此时 this 能用
  },

  // return 6. async setup() { // setup不能是一个async
  setup(prop, context) {
    // setup parameters:
    // console.log('prop: ', prop);    // Proxy { <target>: Proxy, <handler>: {…} }
    // console.log('context: ',context);   // Object { attrs: Getter, slots: Getter, emit: Getter, expose: expose(exposed)

    // console.log('prop.msg: ', prop.msg);    // prop.msg:  what are you no sai lei
    console.log("context.attrs: ", context.attrs);
    console.log(context.emit); // function emit(event, args)

    const m = ref(2);
    const n = ref(3);

    function update_num() {
        // m, n 不受 App Father component 的影响
        m.value += 1;
      n.value += 2;
      
      context.emit('fn', '++'); // 获取 App 分发自定义事件, 两个 comp 能同时控制 msg 的显示
    }

    // setup的返回值:
    console.log("setup run ... ", this); // 此时 this 不能用
    const showMsg1 = () => {
      console.log("setup use showMsg1");
    };

    return {
      showMsg1,
      update_num,
      m,
      n,
    };
  },

  // 实例被挂载后调用
  mounted() {
    console.log(this);
  },

  data() {
    console.log("data run ... ");
    const count = 10;
    return {
      count,
    };
  },

  methods: {
    showMsg2: function () {
      console.log("methods use showMsg2");
    },
  },
});
// data, Type： Function 返回组件实例的 data 对象的 Function。

// methods, Type: { [key: string]: Function }  methods 将被混入到组件实例中。
// 可以直接通过 VM 实例访问这些方法，或者在指令表达式中使用。方法中的 this 自动绑定为组件实例。
</script>