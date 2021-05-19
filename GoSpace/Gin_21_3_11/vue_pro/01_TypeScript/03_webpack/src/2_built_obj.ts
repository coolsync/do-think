(() => {
    /* 1. ECMAScript 的内置对象 */
    // let b: Boolean = new Boolean(1)
    // let n: Number = new Number(true)
    // let s: String = new String('abc')
    // let d: Date = new Date()
    // let r: RegExp = /^1/
    // let e: Error = new Error('error message')
    // b = true
    // // let bb: boolean = new Boolean(2)  // error
    // let bb: Boolean = new Boolean(2) 
    // console.log(bb)
    // let nn: Number = new Number(false);
    // console.log(nn)

    /* 2 BOM 和 DOM 的内置对象 */
    // Window
    // Document
    // HTMLElement
    // DocumentFragment
    // Event
    // NodeList

    const div: HTMLElement = document.getElementById('test')
    const divs: NodeList = document.querySelectorAll('div')
    document.addEventListener('click', (event: MouseEvent) => {
        console.dir(event.target)
    })
    const fragment: DocumentFragment = document.createDocumentFragment()
})()