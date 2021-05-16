<template>
  <h2>Child Father Component</h2>
  <h3>msg: {{ msg }}</h3>
  <button @click="update">Update data</button>
</template>

<script lang='ts'>
import {
  defineComponent,
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  onBeforeUnmount,
  onUnmounted,
  ref,
} from "vue";

export default defineComponent({
  name: "Child",

  beforeCreate() {
    console.log("in v2.x beforeCreate()");
  },

  created() {
    console.log("in v2.x created");
  },

  beforeMount() {
    console.log("in v2.x beforeMount");
  },

  mounted() {
    console.log("in v2.x mounted");
  },

  beforeUpdate() {
    console.log("in v2.x beforeUpdate");
  },

  updated() {
    console.log("in v2.x updated");
  },

  beforeUnmount() {
    console.log("in v2.x beforeUnmount");
  },

  unmounted() {
    console.log("in v2.x unmounted");
  },

  setup() {
    const msg = ref("abc");

    const update = () => {
      msg.value += "+++";
    };

    onBeforeMount(() => {
      console.log("in v3 --onBeforeMount");
    });

    onMounted(() => {
      console.log("in v3 --onMounted");
    });

    onBeforeUpdate(() => {
      console.log("in v3 --onBeforeUpdate");
    });

    onUpdated(() => {
      console.log("in v3 --onUpdated");
    });

    onBeforeUnmount(() => {
      console.log("in v3 --onBeforeUnmount");
    });

    onUnmounted(() => {
      console.log("in v3 --onUnmounted");
    });

    return {
      msg,
      update
    };
  },
});
</script>
