(() => {
    // define interface
    interface ISearchFn {
        (src_str: string, sub_str: string): boolean
    }
    // create interface obj
    let search_fn: ISearchFn = function (src_str: string, sub_str: string): boolean {
        return src_str.search(sub_str) > -1
    };
    // call obj func
    console.log(search_fn('abcd', 'c'));
})()