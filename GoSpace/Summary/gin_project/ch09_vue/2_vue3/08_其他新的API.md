#  2. 其他新的API

## [#](https://24kcs.github.io/vue3_study/chapter5/02_其他新API.html#全新的全局api) 全新的全局API

- createApp()
- defineProperty()
- defineAsyncComponent()
- nextTick()

## [#](https://24kcs.github.io/vue3_study/chapter5/02_其他新API.html#将原来的全局api转移到应用对象) 将原来的全局API转移到应用对象

- app.component()
- app.config()
- app.directive()
- app.mount()
- app.unmount()
- app.use()

## [#](https://24kcs.github.io/vue3_study/chapter5/02_其他新API.html#模板语法变化) 模板语法变化

- v-model的本质变化
  - prop：value -> modelValue；
  - event：input -> update:modelValue；
- .sync修改符已移除, 由v-model代替
  - 
- v-if优先v-for解析