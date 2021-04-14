#  1. 基础类型

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#布尔值)Boolean

```typescript
let isDone: boolean = false;
isDone = true;
// isDone = 2 // error
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#数字)Number

```typescript
let a1: number = 10 // 十进制
let a2: number = 0b1010  // 二进制
let a3: number = 0o12 // 八进制
let a4: number = 0xa // 十六进制
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#字符串)String

JavaScript 程序的另一项基本操作是处理网页或服务器端的文本数据。

```typescript
// string
let st1: string = '窗前明月光';
let st2: string = '小明来开窗'
let st3: string = '遇到一耳光'
let st4: string = '牙齿掉光光'
console.log(`${st1}, ${st2}, st3, ${st4}`);
```

```typescript
// let age: number = 15;
// let nameone: string = 'hehe';
// console.log(`$(nameone), $(age)`)

// number and string
let user_name: string = 'jim'
user_name = 'mark'
// user_name = 12 // error
let age: number = 12
const info = `My name is ${user_name}, I am ${age} years old!`
console.log(info)
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#undefined-和-null)undefined And null

类型用处不是很大：

```typescript
// undefined, null
let u: undefined = undefined
let n: null = null
console.log(`${u}`, n)

// array
let num1: string = undefined
let num2: number = null
console.log(num1, num2)
```

默认情况下 `null` 和 `undefined` 是所有类型的子类型。 就是说你可以把 `null` 和 `undefined` 赋值给 `number` 类型的变量。

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#数组)Array

```typescript
let num1: string = undefined
let num2: number = null
console.log(num1, num2)

// 1
let arr: number[] = [10, 20, 30, 40, 50]
console.log(arr)

// 2 generic
let arr2: Array<number> = [100, 200, 300]
console.log(arr2)
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#元组-tuple)Tuple

```typescript
// tuple
let tp: [string, number, boolean] = ['小甜甜', 100.12345, true]
console.log(tp[0].split(''))
console.log(tp[1].toFixed(2))
```

```typescript
console.log(t1[0].substring(1)) // OK
console.log(t1[1].substring(1)) // Error, 'number' 不存在 'substring' 方法
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#枚举)Enum

默认情况下，从 `0` 开始为元素编号。

```typescript
// enum
enum Color {
    red = 1,	// 可以手动的指定成员的数值
    green,
    blue,
}
let color: Color = Color.red
console.log(color)
console.log(Color.red, Color.green, Color.blue);

enum Gender {
    男,
    女,
}
console.log(Gender.男)
```

全部手动赋值：

```typescript
enum Color {Red = 1, Green = 2, Blue = 4}
let c: Color = Color.Green
```

枚举类型提供的一个便利是你可以由枚举的值得到它的名字。 例如，我们知道数值为 2，但是不确定它映射到 Color 里的哪个名字，我们可以查找相应的名字：

```typescript
enum Color {Red = 1, Green, Blue}
let colorName: string = Color[2]

console.log(colorName)  // 'Green'
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#any)any



```typescript
// any
let any_type: any = 100;
any_type = 'hahahhah';
console.log(any_type)
// [string, number, boolean]
let any_tuple: any[] = [100.1234, 'a str', true];
// console.log(any_tuple[0].split(''))
```



## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#void)void

```typescript
// void
function showMsg(): void {
    console.log('show msg')
    return
    // return undefined
    // return null
}

console.log(showMsg())


// 声明一个 void 类型的变量没有什么大用，因为你只能为它赋予 undefined 和 null：
let unusable: void = undefined
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#object)object

```typescript
// object
function getObj(obj: object): object {
    console.log(obj)
    return {
        user_name: 'bob',
        age: '30',
    }
}

// console.log(getObj({name: 'paul', gender: 'man'}))
// console.log(getObj('1234')) // err
// console.log(getObj(String))
console.log(getObj(new String))
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#联合类型)Union Types

联合类型

```typescript
// 需求1: 定义一个函数得到一个数字或字符串值的字符串形式值
function get_str(str: number | string): string {
    return str.toString()
}
console.log(get_str('123'))
```

需求2: 定义一个一个函数得到一个数字或字符串值的长度

```typescript
function getLength(x: number | string) {

  // return x.length // error

  if (x.length) { // error
    return x.length
  } else {
    return x.toString().length
  }
}
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#类型断言)Type Assertion

类型断言: 手动指定一个值的类型

告诉编译器: “相信我，我知道自己在干什么”

语法:
   一: <类型>值
   二: 值 as 类型  tsx中只能用这种方式



```typescript

// 需求2: 定义一个函数得到一个数字或字符串值的长度
function get_len(v: number | string): number {
    // return v.toString().length

    if ((<string>v).length) {
        // return (<string>v).length
        return (v as string).length
    }
    return v.toString().length

}
console.log("len: ", get_len('13456'), get_len(13456))
```

## [#](https://24kcs.github.io/vue3_study/chapter2/1_type.html#类型推断)类型推断

```typescript
/* 定义变量时赋值了, 推断为对应的类型 */
// let txt = 100;
// txt = 'hhhaha'  // err
// console.log(txt)

/* 定义变量时没有赋值, 推断为any类型 */
let txt2;   // any type
txt2 = 100;
txt2 = 'hehhehh';
console.log(txt2)
```
