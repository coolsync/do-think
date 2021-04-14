// 抽象类做为其它派生类的基类使用。 
// 它们不能被实例化。不同于接口，抽象类可以包含成员的实现细节。 
// `abstract` 关键字是用于定义抽象类和在抽象类内部定义抽象方法。

(() => {
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
})()