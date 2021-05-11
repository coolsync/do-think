
函数是 JavaScript 应用程序的基础，它帮助你实现抽象层，模拟类，信息隐藏和模块。在 TypeScript 里，虽然已经支持类，命名空间和模块，但函数仍然是主要的定义行为的地方。TypeScript 为 JavaScript 函数添加了额外的功能，让我们可以更容易地使用。

## [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#基本示例)基本示例

和 JavaScript 一样，TypeScript 函数可以创建有名字的函数和匿名函数。你可以随意选择适合应用程序的方式，不论是定义一系列 API 函数还是只使用一次的函数。

通过下面的例子可以迅速回想起这两种 JavaScript 中的函数：

```javascript
// 命名函数
function add(x, y) {
  return x + y
}

// 匿名函数
let myAdd = function(x, y) { 
  return x + y;
}
```

## [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#函数类型)函数类型

### [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#为函数定义类型)为函数定义类型

让我们为上面那个函数添加类型：

```typescript
function add(x: number, y: number): number {
  return x + y
}

let myAdd = function(x: number, y: number): number { 
  return x + y
}
```

我们可以给每个参数添加类型之后再为函数本身添加返回值类型。TypeScript 能够根据返回语句自动推断出返回值类型。

### [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#书写完整函数类型)书写完整函数类型

现在我们已经为函数指定了类型，下面让我们写出函数的完整类型。

```typescript
let myAdd2: (x: number, y: number) => number = 
function(x: number, y: number): number {
  return x + y
}
```

## [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#可选参数和默认参数)可选参数和默认参数

TypeScript 里的每个函数参数都是必须的。 这不是指不能传递 `null` 或 `undefined` 作为参数，而是说编译器检查用户是否为每个参数都传入了值。编译器还会假设只有这些参数会被传递进函数。 简短地说，传递给一个函数的参数个数必须与函数期望的参数个数一致。

JavaScript 里，每个参数都是可选的，可传可不传。 没传参的时候，它的值就是 `undefined`。 在TypeScript 里我们可以在参数名旁使用 `?` 实现可选参数的功能。 比如，我们想让 `lastName` 是可选的：

在 TypeScript 里，我们也可以为参数提供一个默认值当用户没有传递这个参数或传递的值是 `undefined` 时。 它们叫做有默认初始化值的参数。 让我们修改上例，把`firstName` 的默认值设置为 `"bob"`。

```typescript
// use ? -> option params
(()=>{
    const getFullName = function(firstName: string = 'bob', lastName?: string) {
        if (lastName) {
            return firstName + '_' + lastName;
        }
        return firstName
    }

    // no any params
    console.log(getFullName());

    // only firstName param
    console.log(getFullName('jerry'));

    // all params
    console.log(getFullName('dongfa', 'bubai'));
})()
```

### [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#剩余参数)剩余参数

必要参数，默认参数和可选参数有个共同点：它们表示某一个参数。 有时，你想同时操作多个参数，或者你并不知道会有多少参数传递进来。 在 JavaScript 里，你可以使用 `arguments` 来访问所有传入的参数。

在 TypeScript 里，你可以把所有参数收集到一个变量里：
剩余参数会被当做个数不限的可选参数。 可以一个都没有，同样也可以有任意个。 编译器创建参数数组，名字是你在省略号（ `...`）后面给定的名字，你可以在函数体内使用这个数组。

```typescript
// fn declaration
function showMsg(s: string, s1: string, ...args: string[]) {
    console.log(s)
    console.log(s1)
    console.log(args)
}

// call fn
showMsg('a', 'b', 'c', 'd');
```

## [#](https://24kcs.github.io/vue3_study/chapter2/4_function.html#函数重载)函数重载

函数重载: 函数名相同, 而形参不同的多个函数
在JS中, 由于弱类型的特点和形参与实参可以不匹配, 是没有函数重载这一说的 但在TS中, 与其它面向对象的语言(如Java)就存在此语法

```typescript
/* 
函数重载: 函数名相同, 而形参不同的多个函数
需求: 我们有一个add函数，它可以接收2个string类型的参数进行拼接，也可以接收2个number类型的参数进行相加 
*/

// fn overloading
(() => {
    // 需求: 我们有一个add函数，它可以接收2个string类型的参数进行拼接，也可以接收2个number类型的参数进行相加 

    // def fn
    function add(x: string | number, y: string | number): string | number {
        if (typeof x === 'string' && typeof y === 'string') {
            return x + y
        } else if (typeof x === 'number' && typeof y === 'number') {
            return x + y
        }
    }

    // fn call
    console.log(add('bob', 'say'));
    console.log(add(10, 20));

    // input illegal value
    console.log(add('zhengxiang', 10)); // undefined
    console.log(add(10, 'zhengxiang')); // undefined
})()
```