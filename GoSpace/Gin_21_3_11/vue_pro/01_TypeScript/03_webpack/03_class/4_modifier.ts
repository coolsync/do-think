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