(() => {
    // 1 类实现接口
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
})()