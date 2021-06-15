<template>
  <div>
    <h1>响应式数据的判断</h1>
  </div>
</template>

<script lang='ts'>
// - isRef: 检查一个值是否为一个 ref 对象
// - isReactive: 检查一个对象是否是由 `reactive` 创建的响应式代理
// - isReadonly: 检查一个对象是否是由 `readonly` 创建的只读代理
// - isProxy: 检查一个对象是否是由 `reactive` 或者 `readonly` 方法创建的代理
import {
  defineComponent,
  isRef,
  ref,
  reactive,
  isReactive,
  isReadonly,
  readonly,
  isProxy,
} from "vue";

export default defineComponent({
  name: "App",
  setup() {
    // - isRef: 检查一个值是否为一个 ref 对象
    console.log(isRef(ref({}))); // true

    // - isReactive: 检查一个对象是否是由 `reactive` 创建的响应式代理
    console.log(isReactive(reactive({}))); // true

    // - isReadonly: 检查一个对象是否是由 `readonly` 创建的只读代理
    console.log(isReadonly(readonly({}))); // true

    // - isProxy: 检查一个对象是否是由 `reactive` 或者 `readonly` 方法创建的代理
    console.log(isProxy(reactive({})), isProxy(readonly({}))); // true, true



    // Creates a readonly copy of the original object. Note the returned copy is not made reactive, but readonly can be called on an already reactive object.
    //创建原始对象的只读副本。请注意，返回的副本不是被动式的，但是可以在已经被动式的对象上调用只读。
  },
});
</script>
