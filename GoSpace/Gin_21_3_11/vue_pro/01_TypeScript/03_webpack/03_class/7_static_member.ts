// 静态 member
(() => {
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
})()