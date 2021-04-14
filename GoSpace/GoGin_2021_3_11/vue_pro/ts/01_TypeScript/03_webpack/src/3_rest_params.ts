// rest params 剩余参数

(() => {
    // fn declaration
    function showMsg(s: string, s1: string, ...args: string[]) {
        console.log(s)
        console.log(s1)
        console.log(args)
    }

    // call fn
    showMsg('a', 'b', 'c', 'd');
})()