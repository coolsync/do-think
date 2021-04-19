<template>
  <div></div>
  <ModalButton />
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import ModalButton from "./ModalButton.vue";

export default defineComponent({
  name: "App",

  components: {
    ModalButton,
  },

  setup() {
    return {};
  },
});
</script>
