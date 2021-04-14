# 2. Interface

TypeScript 的核心原则之一是对值所具有的结构进行类型检查。我们使用接口（Interfaces）来定义对象的类型。`接口是对象的状态(属性)和行为(方法)的抽象(描述)`

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#接口初探)接口初探

需求: 创建人的对象, 需要对人的属性进行一定的约束

```text
id是number类型, 必须有, 只读的
name是string类型, 必须有
age是number类型, 必须有
sex是string类型, 可以没有
```

下面通过一个简单示例来观察接口是如何工作的：

```typescript
/* 
在 TypeScript 中，我们使用接口（Interfaces）来定义对象的类型
接口: 是对象的状态(属性)和行为(方法)的抽象(描述)
接口类型的对象
    多了或者少了属性是不允许的
    可选属性: ?
    只读属性: readonly
*/
// 需求: 创建人的对象, 需要对人的属性进行一定的约束
// id是number类型, 必须有, 只读的
// name是string类型, 必须有
// age是number类型, 必须有
// sex是string类型, 可以没有

interface IPerson {
    readonly id: number,  // 只读属性: readonly
        name: string,
            age: number,
                sex?: string    // 可选属性: ?
}

            const person: IPerson = {
                id: 1,
                name: "bob",
                age: 20,
                // sex: '男'
            };

// person.id = 100;    // err: Cannot assign to 'id' because it is a read-only property.
console.log(person);
```

类型检查器会查看对象内部的属性是否与IPerson接口描述一致, 如果不一致就会提示类型错误。

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#可选属性)可选 Property

接口里的属性不全都是必需的。 有些是只在某些条件下存在，或者根本不存在。

```typescript
interface IPerson {
  id: number
  name: string
  age: number
  sex?: string
}
```

带有可选属性的接口与普通的接口定义差不多，只是在可选属性名字定义的后面加一个 `?` 符号。

好处之一是可以对可能存在的属性进行预定义，

好处之二是可以捕获引用了不存在的属性时的错误。

```typescript
const person2: IPerson = {
  id: 1,
  name: 'tom',
  age: 20,
  // sex: '男' // 可以没有
}
```

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#只读属性)Readonly Property

一些对象属性只能在对象创建的时候修改其值。可以在属性名前用 `readonly` 来指定只读属性:

```typescript
interface IPerson {
  readonly id: number
  name: string
  age: number
  sex?: string
}
```

一旦赋值后再也不能被改变了。

```typescript
const person2: IPerson = {
  id: 2,
  name: 'tom',
  age: 20,
  // sex: '男' // 可以没有
  // xxx: 12 // error 没有在接口中定义, 不能有
}
person2.id = 2 // error
```

### [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#readonly-vs-const)readonly vs const

最简单判断该用 `readonly` 还是 `const` 的方法是看要把它做为变量使用还是做为一个属性。 做为变量使用的话用 `const`，若做为属性则使用 `readonly`。

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#函数类型)Method Type

```typescript
/* 
接口可以描述函数类型(参数的类型与返回的类型)
*/

interface SearchFunc {
  (source: string, subString: string): boolean
}
```

这样定义后，我们可以像使用其它接口一样使用这个函数类型的接口。

下例展示了如何创建一个函数类型的变量，并将一个同类型的函数赋值给这个变量。

```typescript
// define interface
interface ISearchFn {
    (src_str: string, sub_str: string): boolean
}
// create interface obj
let search_fn: ISearchFn = function (src_str: string, sub_str: string): boolean {
    return src_str.search(sub_str) > -1
};
// call obj func
console.log(search_fn('abcd', 'c'));
```

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#类类型)Class Type

​    // 1 类实现接口



​    // 2. class 实现 多个 interface

​    // define inter

​    *interface* ISwim {

​        swim()

​    }



​    // a class imple mul interface

​    *class* Person2 implements IFly, ISwim {

​        fly() {

​            console.log('I can fly, 2');

​        }

​        swim() {

​            console.log('哈哈， I can swim 2')

​        }

​    }



​    // 实例化 obj

​    *let* p2 = new Person2;

​    p2.fly();

​    p2.swim();



​    // 3. interface 继承 多个 inter

​    *interface* IMyFlySwim extends IFly, ISwim {}



​    *class* Person3 implements IMyFlySwim {

​        fly() {

​            console.log('I can fly, 3');

​        }

​        swim() {

​            console.log('哈哈， I can swim 3')

​        }

​    }



​    *let* p3 = new Person3;

​    p3.fly();

​    p3.swim();



### [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#类实现接口)Class impl Interface

能够用它来明确的强制一个类去符合某种契约。

```typescript
// define interface
interface IFly {
    fly()
}
// create class impl interface
class Person implements IFly {
    fly() {
        console.log('I can fly');
    }
}
// 实例化 obj, call inter method
let p = new Person();
p.fly();
```

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#一个类可以实现多个接口)A class impl multiple interfaces

```typescript
// 2. class 实现 多个 interface
// define inter
interface ISwim {
    swim()
}

// a class imple mul interface
class Person2 implements IFly, ISwim {
    fly() {
        console.log('I can fly, 2');
    }
    swim() {
        console.log('哈哈， I can swim 2')
    }
}

// 实例化 obj
let p2 = new Person2;
p2.fly();
p2.swim();
```

## [#](https://24kcs.github.io/vue3_study/chapter2/2_interface.html#接口继承接口)Interface inherits interface

```typescript
// 3. interface 继承 多个 inter
interface IMyFlySwim extends IFly, ISwim {}

class Person3 implements IMyFlySwim {
    fly() {
        console.log('I can fly, 3');
    }
    swim() {
        console.log('哈哈， I can swim 3')
    }
}

let p3 = new Person3;
p3.fly();
p3.swim();
```