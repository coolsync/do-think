(() => {
    // named fn
    function add(x: string, y: string): string {
        return x + y
    }
    let r1 = add('111', '222');
    console.log(r1);

    // fn expr, fn value
    const add2 = function (x: number, y: number): number {
        return x + y
    }
    console.log(add2(1, 2))

    // 完整 fn type
    // fn name --> type --> fn val
    const add3: (x: number, y: number) => number = function (x: number, y: number): number {
        return x + y
    }
    console.log(add3(10, 20));
})()