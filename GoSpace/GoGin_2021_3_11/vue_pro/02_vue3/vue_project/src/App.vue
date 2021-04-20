<template>
  <h2>App components</h2>
  <Nav />
</template>

<script lang='ts'>
import { defineComponent } from "vue";
import Nav from "./components/Nav.vue";


export default defineComponent({
  name: "App",
  components: {
    Nav,
    // Home,
    // ElButton.name: Element,
    // ElSelect,
    // ElementPlus,
    // ElMain,
    // ElMenu,
    // ElMenuItem,
    // ElMenuItemGroup,
  },
});
</script>

