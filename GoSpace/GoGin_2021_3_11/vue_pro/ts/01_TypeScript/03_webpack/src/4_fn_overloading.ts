// fn overloading
(() => {
    // 需求: 我们有一个add函数，它可以接收2个string类型的参数进行拼接，也可以接收2个number类型的参数进行相加 

    // def fn
    function add(x: string | number, y: string | number): string | number {
        if (typeof x === 'string' && typeof y === 'string') {
            return x + y
        } else if (typeof x === 'number' && typeof y === 'number') {
            return x + y
        }
    }

    // fn call
    console.log(add('bob', 'say'));
    console.log(add(10, 20));

    // input illegal value
    console.log(add('zhengxiang', 10)); // undefined
    console.log(add(10, 'zhengxiang')); // undefined
})()