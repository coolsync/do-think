(()=> {
    //根据指定的数量 `count` 和数据 `value` , 创建一个包含 `count` 个 `value` 的数组 

    // fn declare
    function getArr<T>(value:T, count:number): Array<T> {
        const arr: Array<T> = []
        for (let i = 0; i < count; i++) {
            arr.push(value)
        }

        return arr
    }

    // call fn
    let arr1 = getArr<number>(100.12345, 5);
    let arr2 = getArr<string>('abcdefg', 5);

    console.log(arr1)
    console.log(arr2)
    console.log(arr1[0].toFixed(2));
    console.log(arr2[0].split(''));
})()