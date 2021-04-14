(() => {
    // def base class
    class Animal {
        // def property
        name: string;
        // def constructor fn
        constructor(name: string) {
            this.name = name;
        }
        // def instance method
        run(distance: number = 0) {
            console.log(`run ${distance} away, ${this.name}`)
        }
    }

    // def son class
    class Dog extends Animal {
        // def constructor fn
        constructor(name: string) {
            super(name);
        }
        // def instance method
        run(distance: number = 5) {
            console.log(`run ${distance} away, ${this.name}`)
        }
    }
     // def son class
     class Pig extends Animal {
        // def constructor fn
        constructor(name: string) {
            super(name);
        }
        // def instance method
        run(distance: number = 10) {
            console.log(`run ${distance} away, ${this.name}`)
        }
    }

    // get obj
    const ani: Animal = new Animal('animal');
    // call obj method
    ani.run();

    const dog: Dog = new Dog('da huang');
    dog.run();

    const pig: Pig = new Pig('pig')   
    pig.run();


    const dog1: Animal = new Dog('xiao huang');
    dog1.run();

    const pig1: Animal = new Pig('hei pig');
    pig1.run();

    function get_run(ani:Animal) {
        ani.run();
    }

    get_run(dog1);
    get_run(pig1);
})()
