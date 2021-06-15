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