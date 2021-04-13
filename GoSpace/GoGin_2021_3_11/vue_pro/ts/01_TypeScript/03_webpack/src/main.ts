import './01.basic_type'

let htmlStr = `<div><span>'bob'</span></div>`;
const element: HTMLElement = document.querySelector('body') as HTMLElement
element.innerHTML = htmlStr;