(() => {
    // define class
    class Person {
        // define property
        name: string;
        age: number;
        gender: string;

        // define constructer fn
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