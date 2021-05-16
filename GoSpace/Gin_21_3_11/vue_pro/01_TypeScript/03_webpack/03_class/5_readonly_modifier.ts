(() => {
    /* 
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
    */

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
})()