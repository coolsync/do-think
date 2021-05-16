<template>
  <h2>Child Component</h2>
  <h3>age: {{ age }}</h3>
  <h3>length: {{ length }}</h3>
</template>

<script lang='ts'>
import { computed, defineComponent, Ref, toRef } from "vue";

function useGetLength(age: Ref) { // if out requerst is Ref obj, not value type
  return computed(() => age.value.toString().length);
}

export default defineComponent({
  name: "Child",

  props: {
    age: {
      type: Number,
      required: true,
    },
  },
  setup(props) {
    // console.log(props.age.toString().length)// value type
    const length = useGetLength(toRef(props, "age"));
    return {
      length,
    };
  },
});
</script>

