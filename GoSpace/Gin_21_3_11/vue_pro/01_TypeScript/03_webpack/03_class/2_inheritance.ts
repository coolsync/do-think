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