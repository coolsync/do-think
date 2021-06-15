# 5. Generic

指在定义函数、接口或类的时候，不预先指定具体的类型，而在使用的时候再指定具体类型的一种特性。

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#引入) Intro

下面创建一个函数, 实现功能: 根据指定的数量 `count` 和数据 `value` , 创建一个包含 `count` 个 `value` 的数组 不用泛型的话，这个函数可能是下面这样：

```typescript
//根据指定的数量 `count` 和数据 `value` , 创建一个包含 `count` 个 `value` 的数组 

// fn declare
function getArr<T>(value:T, count:number): Array<T> {
    const arr: Array<T> = []
    for (let i = 0; i < count; i++) {
        arr.push(value)
    }

    return arr
}

// call fn
let arr1 = getArr<number>(100.12345, 5);
let arr2 = getArr<string>('abcdefg', 5);

console.log(arr1)
console.log(arr2)
console.log(arr1[0].toFixed(2));
console.log(arr2[0].split(''));
```

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#使用函数泛型) Function generic

```typescript
(()=> {
    //根据指定的数量 `count` 和数据 `value` , 创建一个包含 `count` 个 `value` 的数组 

    // fn declare
    function getArr<T>(value:T, count:number): Array<T> {
        const arr: Array<T> = []
        for (let i = 0; i < count; i++) {
            arr.push(value)
        }

        return arr
    }

    // call fn
    let arr1 = getArr<number>(100.12345, 5);
    let arr2 = getArr<string>('abcdefg', 5);

    console.log(arr1)
    console.log(arr2)
    console.log(arr1[0].toFixed(2));
    console.log(arr2[0].split(''));
})()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#多个泛型参数的函数)多个泛型参数的函数

一个函数可以定义多个泛型参数

```typescript
// fn declare
function showMsg<K, V>(val1: K, val2: V): [K, V] {
    return [val1, val2]
}

// call fn
let msg = showMsg<string, number>('abcd', 100.12345)
console.log(msg);
console.log(msg[0].split(''));
console.log(msg[1].toFixed(2))

function swap <K, V> (a: K, b: V): [K, V] {
  return [a, b]
}
const result = swap<string, number>('abc', 123)
console.log(result[0].length, result[1].toFixed())

```

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#泛型接口)generic interface

在定义接口时, 为接口中的属性或方法定义泛型类型
在使用接口时, 再指定具体的泛型类型

```typescript
(() => {
    // def generic interface
    interface IBase<T> {
        // store user info array
        data: Array<T>;
        // add user
        add: (t: T) => T;
        // by id query user
        getUserId: (number) => T;

        deleteUser: (number) => T;
    }
    // def user info class
    class User {
        id?: number;
        name: string;
        age: number;
        constructor(name: string, age: number) {
            this.name = name;
            this.age = age;
        }
    }

    // def user crud class, impl IBase generic inter
    class UserCRUD implements IBase<User> {
        data = [];

        add(user: User): User {
            user.id = Date.now() + Math.random();
            this.data.push(user);
            return user
        }

        getUserId(id: number): User {
            return this.data.find(user => user.id === id);
        }

        deleteUser(id: number): User {
            // let node： int;  //要移除的对象
            // nodes： int[];
            // this.nodes = this.nodes.filter(item => item !== node);
            let user = this.getUserId(id);
            user.id = id;
            this.data = this.data.filter(item => item !== user);
            return user
        }
    }

    // instance UserCRUD
    const userCRUD: UserCRUD = new UserCRUD();

    userCRUD.add(new User("bob", 30));
    userCRUD.add(new User("paul", 20));
    let { id } = userCRUD.add(new User("jerry", 18));
    userCRUD.add(new User("mark", 31));

    console.log(userCRUD.data);
    console.log(userCRUD.getUserId(id));

    console.log(id)
    userCRUD.deleteUser(id);
    console.log(userCRUD.data);
})()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#泛型类)generic class

在定义类时, 为类中的属性或方法定义泛型类型 在创建类的实例时, 再指定特定的泛型类型

```typescript
(() => {
    // when def class, class property, method params, return value is 不确定， 可以使用 generic class

    // def generic class
    class GenericNumber<T> {
        // def generic class property, method
        defaultValue: T;
        add: (t1: T, t2: T) => T;
    }

    // instance generic class obj 后， 在确定 type
    // number type
    const g1: GenericNumber<number> = new GenericNumber();
    g1.defaultValue = 100;
    g1.add = function (x, y) {
        return x + y
    }
    console.log(g1.add(g1.defaultValue, 10))

    // string type
    const g2: GenericNumber<string> = new GenericNumber();
    g2.defaultValue = 'haha';
    g2.add = function (x, y) {
        return x + y
    }
    console.log(g2.add(g2.defaultValue, ' xiaoyan'))
})()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/5_generic.html#泛型约束)generic constraints 

如果我们直接对一个泛型参数取 `length` 属性, 会报错, 因为这个泛型根本就不知道它有这个属性

我们可以使用generic constraints 来实现

```typescript
(() => {
    // 如果我们直接对一个泛型参数取 length 属性, 会报错, 因为这个泛型根本就不知道它有这个属性
    // function getLen<T>(x:T):number {
    //     return x.length	// err
    // }
    
    
    // def generic constraints, 对 future some type process constraints
    interface ILen {
        length: number;
    }
    
    function getLen<T extends ILen>(x: T): number {
        return x.length
    }

    console.log(getLen<string>('abcd')) // 4
    // console.log(getLen<number>(12345)) // err number没有length属性
})()
```

