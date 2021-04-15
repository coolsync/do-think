(()=>{
    // fn declare
    function showMsg<K, V>(val1: K, val2: V): [K, V] {
        return [val1, val2]
    }

    // call fn
    let msg = showMsg<string, number>('abcd', 100.12345)
    console.log(msg);
    console.log(msg[0].split(''));
    console.log(msg[1].toFixed(2))
})()