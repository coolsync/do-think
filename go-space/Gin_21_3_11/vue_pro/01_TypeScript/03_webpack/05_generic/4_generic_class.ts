(() => {
    // when def class, class property, method params, return value is 不确定， 可以使用 generic class

    // def generic class
    class GenericNumber<T> {
        // def generic class property, method
        defaultValue: T;
        add: (t1: T, t2: T) => T;
    }

    // instance generic class obj 后， 在确定 type
    // number type
    const g1: GenericNumber<number> = new GenericNumber();
    g1.defaultValue = 100;
    g1.add = function (x, y) {
        return x + y
    }
    console.log(g1.add(g1.defaultValue, 10))

    // string type
    const g2: GenericNumber<string> = new GenericNumber();
    g2.defaultValue = 'haha';
    g2.add = function (x, y) {
        return x + y
    }
    console.log(g2.add(g2.defaultValue, ' xiaoyan'))
})()