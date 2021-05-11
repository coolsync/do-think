
对于传统的 JavaScript 程序我们会使用`函数`和`基于原型的继承`来创建可重用的组件，但对于熟悉使用面向对象方式的程序员使用这些语法就有些棘手，因为他们用的是`基于类的继承`并且对象是由类构建出来的。 从 ECMAScript 2015，也就是 ES6 开始， JavaScript 程序员将能够使用基于类的面向对象的方式。 使用 TypeScript，我们允许开发者现在就使用这些特性，并且编译后的 JavaScript 可以在所有主流浏览器和平台上运行，而不需要等到下个 JavaScript 版本。

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#基本示例)基本示例

下面看一个使用类的例子：

```typescript
(() => {
    // define class
    class Person {
        // define property
        name: string;
        age: number;
        gender: string;

        // define constructor fn, 用来初始化时直接使用params 初始化 属性值
        constructor(name: string = '甜甜', age: number = 18, gender: string = '女') {
            this.name = name;
            this.age = age;
            this.gender = gender;
        }

        // define 实例方法
        say_hi(s: string) {
            console.log(s, ` Hi, I am ${this.name}, age: ${this.age}, gender: ${this.gender}`)
        }
    }

    // 实例化 obj
    const person = new Person(
        'bob', 30, 'man',
    );
    person.say_hi('你叫什么');
})()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#继承)Inheritance

在 TypeScript 里，我们可以使用常用的面向对象模式。 

基于类的程序设计中一种最基本的模式是允许使用继承来扩展现有的类。

例子

```typescript
(() => {
    // def class
    class Person {
        // def property
        name: string;
        age: number;
        gender: string;

        // def constructor fn
        constructor(name:string='bob', age:number=30, gender:string='man') {
            this.name = name;
            this.age = age;
            this.gender = gender;
        }

        // def instance method
        say_hi(s: string) {
            console.log(`i am ${this.name}, `, s)
        }
    }

    class Emp extends Person {
        constructor(name: string, age: number, gender: string) {
            super(name, age, gender)
        }

        say_hi() {
            console.log('in class Emp.')
            super.say_hi('are you name?')   // rewirte super class method
        }
    }

    // get obj
    const p = new Person();
    p.say_hi('hello')

    const e = new Emp('jerry', 18, 'girl');
    // call obj method
    e.say_hi();
})()
```



## Polymorshim

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#公共-私有与受保护的修饰符)公共，私有与受保护的修饰符

### [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#默认为-public)默认为 public



### [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#理解-private)理解 private

当成员被标记成 `private` 时，它就不能在声明它的类的外部访问。

### [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#理解-protected)理解 protected

`protected` 修饰符与 `private` 修饰符的行为很相似，但有一点不同，`protected`成员在派生类中仍然可以访问。例如：

```typescript
/* 
访问修饰符: 用来描述类内部的属性/方法的可访问性
  public: 默认值, 公开的外部也可以访问
  private: 只能类内部可以访问
  protected: 类内部和子类可以访问
*/
(()=> {
    // def cls
    class Person {
        // property
        // name: string;   // defualt public
        // private name: string;
        protected name: string; // only in son cls visit
        // constructor
        constructor(name: string) {
            // update property
            this.name = name;
        }

        // method
        eat() {
            console.log(`${this.name} eat something`);
        }
    }

    // son cls
    class Emp extends Person{
        constructor(name: string) {
            super(name) // private name 无法访问
        }

        play() {
            console.log(`${this.name} play something`);  // private name 无法访问
        }
    }

    const per = new Person('小黑');
    // console.log(per.name);  // private name 无法访问
    per.eat();

    const e = new Emp('小白');
    // console.log(e.name) // protected name 无法访问
    e.play();
})()
```



## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#readonly-修饰符)readonly 修饰符

你可以使用 `readonly` 关键字将属性设置为只读的。 只读属性必须在声明时或构造函数里被初始化。

```typescript
// create cls
class Person {
    // property
    readonly name: string;
    // constructor fn
    constructor(name: string = 'bob') { // constructor params: readonly ok
        this.name = name;
    }
    // method
    say_hi() {
        // this.name = 'kk';   // err: readonly
        console.log(`name: ${this.name}`);
    }
}
// get cls obj
const p = new Person('jerry'); 
console.log(p)
console.log(p.name)

// p.name = 'paul'  // err: readonly

// call obj method
p.say_hi()
```

### [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#参数属性)参数属性

constructor params readonly, public, private, protected  modifier

```typescript
// create cls
class Person {
    // constructor fn
    // constructor params has modifier, auto product name property.
    constructor(readonly name: string = 'bob') {
        this.name = name;
    }

    // constructor(public name: string = 'bob') {
    //     this.name = name;
    // }

    // constructor(private name: string = 'bob') {
    //     this.name = name;
    // }

    // constructor(protected name: string = 'bob') {  // in base cls and son cls use
    //     this.name = name;
    // }
}
// get cls obj
const p = new Person('jerry');

// p.name = 'pual' // readonly err, public ok, private err, protected err
console.log(p.name)
```

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#存取器)存取器

`TypeScript` 支持通过 `getters/setters` 来截取对对象成员的访问， 帮助你有效的控制对对象成员的访问。

下面来看如何把一个简单的类改写成使用 `get` 和 `set`。 首先，我们从一个没有使用存取器的例子开始。

```typescript
// 存取器 支持通过 getters/setters 来截取对对象成员的访问, 帮助你有效的控制对对象成员的访问。
(() => {
    // create cls
    class Person {
        first_name: string;
        last_name: string;
        // full_name: string;

        constructor(first_name: string, last_name: string) {
            // update property values
            this.first_name = first_name;
            this.last_name = last_name;
        }
        // get 
        get full_name() {
            console.log('in get ...');
            return this.first_name + '_' + this.last_name;
        }
        // set 
        set full_name(val) {
            console.log('in set ...');
            let names = val.split('_');
            this.first_name = names[0];
            this.last_name = names[1];
        }
    }

    // instance obj
    const p = new Person('dongfa', 'bubai');
    
    console.log(p.full_name);   // dongfa_bubai
    p.full_name = 'bob_sim';
    console.log(p.full_name);   // bob_sim
    console.log(p.first_name);  // bob
})()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#静态属性)静态属性

到目前为止，我们只讨论了类的实例成员，那些仅当类被实例化的时候才会被初始化的属性。 我们也可以创建类的静态成员，这些属性存在于类本身上面而不是类的实例上。 在这个例子里，我们使用 `static` 定义 `origin`，因为它是所有网格都会用到的属性。 每个实例想要访问这个属性的时候，都要在 `origin` 前面加上类名。 如同在实例属性上使用 `this.xxx` 来访问属性一样，这里我们使用 `Grid.xxx` 来访问静态属性。

```typescript
/* 
静态属性, 是类对象的属性
非静态属性, 是类的实例对象的属性
*/
// def cls
class Person {
    static name1: string;

    constructor(public name: string) {
        // this.name1 = name;  // can not assign val to static property
    }

    static say_hi() {
        console.log('hello');
    }
}

// instance obj
const p = new Person('bob');
// console.log(p.name);    // can not by instance obj call static property
// p.say_hi();  // can not by instance obj call static property

console.log(p.name)

// call static member
console.log(Person.name1)
Person.say_hi()
```

## [#](https://24kcs.github.io/vue3_study/chapter2/3_class.html#抽象类)抽象类

抽象类做为其它派生类的基类使用。 它们不能被实例化。不同于接口，抽象类可以包含成员的实现细节。 `abstract` 关键字是用于定义抽象类和在抽象类内部定义抽象方法。

```typescript
/* 
抽象类
  不能创建实例对象, 只有实现类才能创建实例
  可以包含未实现的抽象方法
*/

// def abstract class
abstract class Animal {
    // def abstract property, pass
    // def abstract method
    abstract eat();
    // def instance method
    say_hi() {
        console.log('hello, brother');
    }
}

// def son cls
class Dog extends Animal {
    name: string;
    eat() {
        console.log('dog eat, haha');
    }
}

// instance dog obj
const dog: Dog = new Dog();
dog.eat();
dog.say_hi();
```