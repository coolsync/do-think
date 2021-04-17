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
    full_name3: <input type="text" placeholder="fullName" v-model="fullName3" /><br />
  </fieldset>
</template>

<script lang='ts'>
import { computed, defineComponent, reactive, ref, watch, watchEffect } from "vue";

export default defineComponent({
  name: "App",

  setup() {
    // def responsive obj
    const user = reactive({
      firstName: "dongfang",
      lastName: "bubai",
    });

    // First fullName
    // computed param only callback fn, Means get
    const fullName1 = computed(() => {
      // console.log(fullName1)
      return user.firstName + "_" + user.lastName;
    });
    console.log(fullName1);

    // Second fullName
    // computed get set
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
    // Third fullName
    // watch
    const fullName3 = ref('');
    // watch(
    //   // user,() => {
    //   //   // console.log(val)
    //   //   fullName3.value = user.firstName + '_' + user.lastName
    //   // },
      
    //   // fullName3, (val) => {
    //   //   const names = val.split('_');
    //   //   user.firstName = names[0];
    //   //   user.lastName = names[1];
    //   // },

    //   user, ({firstName, lastName}) => {
    //     fullName3.value = firstName + '_' + lastName
    //   },{ immediate: true, deep: true },
    // );

    watchEffect(() => {
        fullName3.value = user.firstName + '_' + user.lastName
      },
    );
    return {
      user,
      fullName1,
      fullName2,
      fullName3,
    };
  },
});
</script>
