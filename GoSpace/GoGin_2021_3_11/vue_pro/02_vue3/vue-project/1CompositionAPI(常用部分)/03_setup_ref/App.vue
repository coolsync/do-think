<template>
  <div id="nav">hhaa</div>
  <h1>{{count}}</h1>
  <button @click="updateCount">update data</button>
</template>

<script lang='ts'>
import {defineComponent, ref} from 'vue'

export default defineComponent ({
  name: 'App',
  setup() {
    const count = ref(0);   // return a obj,
    console.log(count); // Object { _rawValue: 0, _shallow: false, __v_isRef: true, _value: 0 }

    function updateCount() {
      count.value++ // use Ref obj value property impl Self-increasing
    }

    return {  // return a obj
      count,
      updateCount
    }
  }
})
</script>


