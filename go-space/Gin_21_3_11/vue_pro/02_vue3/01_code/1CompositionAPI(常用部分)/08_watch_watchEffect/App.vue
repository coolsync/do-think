<template>
  <h2>Computing attributes and monitoring</h2>
  <fieldset>
    <legend>姓名操作</legend>
    first_name:
    <input type="text" placeholder="firstName" v-model="user.firstName" /><br />
    last_name:
    <input type="text" placeholder="lastName" v-model="user.lastName" /><br />
  </fieldset>
  <fieldset>
    <legend>Computing attributes and monitoring 演示</legend>
    full_name1:
    <input type="text" placeholder="fullName" v-model="fullName1" /><br />
    full_name2:
    <input type="text" placeholder="fullName" v-model="fullName2" /><br />
    full_name3:
    <input type="text" placeholder="fullName" v-model="fullName3" /><br />
  </fieldset>
</template>

<script lang='ts'>
import {
  computed,
  defineComponent,
  reactive,
  ref,
  watch,
  watchEffect,
} from "vue";

export default defineComponent({
  name: "App",

  setup() {
    // def responsive obj
    const user = reactive({
      firstName: "dongfang",
      lastName: "bubai",
    });

    // fullName 1

    // computed param only callback fn, Means get
    const fullName1 = computed(() => {
      // console.log(fullName1)
      return user.firstName + "_" + user.lastName;
    });
    console.log(fullName1);

    // fullName 2
    // computed函数:
    // 与computed配置功能一致
    // 只有getter
    // 有getter和setter
    const fullName2 = computed({
      get() {
        return user.firstName + "_" + user.lastName;
      },
      set(val: string) {
        // console.log("#####",val)
        const names = val.split("_");
        user.firstName = names[0];
        user.lastName = names[1];
      },
    });

    // watch Function
    // 与watch配置功能一致
    // 监视指定的一个或多个响应式数据, 一旦数据变化, 就自动执行监视回调
    // 默认初始时不执行回调, 但可以通过配置immediate为true, 来指定初始时立即执行第一次
    // 通过配置deep为true, 来指定深度监视
    // fullName 3
    const fullName3 = ref("");
    watch(
      // user,() => {
      //   // console.log(val)
      //   fullName3.value = user.firstName + '_' + user.lastName
      // },

      // fullName3, (val) => {
      //   const names = val.split('_');
      //   user.firstName = names[0];
      //   user.lastName = names[1];
      // },

      user,
      ({ firstName, lastName }) => {
        fullName3.value = firstName + "_" + lastName;
      },
      { immediate: true, deep: true } //
    );

    // watchEffect函数
    // 不用直接指定要监视的数据, 回调函数中使用的哪些响应式数据就监视哪些响应式数据
    // 默认初始时就会执行第一次, 从而可以收集需要监视的数据
    // 监视数据发生变化时回调

// watchEffect(() => {
    //     fullName3.value = user.firstName + '_' + user.lastName
    //   },
    // );

    watchEffect(() => {
      const names = fullName3.value.split("_");
      user.firstName = names[0];
      user.lastName = names[1];
    });

    // watch not responsive data
    watch([() => user.firstName, () => user.lastName], () => {
      console.log("++++++w");
    });
    return {
      user,
      fullName1,
      fullName2,
      fullName3,
    };
  },
});
</script>
