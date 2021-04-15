(() => {
    // 如果我们直接对一个泛型参数取 length 属性, 会报错, 因为这个泛型根本就不知道它有这个属性
    // function getLen<T>(x:T):number {
    //     return x.length
    // }

    // def generic constraints, 对 future some type process constraints
    interface ILen {
        length: number;
    }
    
    function getLen<T extends ILen>(x: T): number {
        return x.length
    }

    console.log(getLen<string>('abcd')) // 4
    // console.log(getLen<number>(12345)) // err number没有length属性
})()