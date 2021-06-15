// boolean
let isDone: boolean = false;
isDone = true;
console.log(isDone)
// isDone = 2 // error

// number
let a1: number = 10 // 十进制
let a2: number = 0b1010  // 二进制
let a3: number = 0o12 // 八进制
let a4: number = 0xa // 十六进制
console.log(a1, a2, a3, a4)

// string
let st1: string = '窗前明月光';
let st2: string = '小明来开窗'
let st3: string = '遇到一耳光'
let st4: string = '牙齿掉光光'
console.log(`${st1}, ${st2}, st3, ${st4}`);

// let age: number = 15;
// let nameone: string = 'hehe';
// console.log(`$(nameone), $(age)`)

// number and string
let user_name: string = 'jim'
user_name = 'mark'
// user_name = 12 // error
let age: number = 12
const info = `My name is ${user_name}, I am ${age} years old!`
console.log(info)

// undefined, null
let u: undefined = undefined
let n: null = null
console.log(`${u}`, n)

// array
let num1: string = undefined
let num2: number = null
console.log(num1, num2)

let arr: number[] = [10, 20, 30, 40, 50]
console.log(arr)

let arr2: Array<number> = [100, 200, 300]
console.log(arr2)

// tuple
let tp: [string, number, boolean] = ['小甜甜', 100.12345, true]
console.log(tp[0].split(''))
console.log(tp[1].toFixed(2))


// enum
enum Color {
    red = 1,
    green,
    blue,
}
let color: Color = Color.red
console.log(color)
console.log(Color.red, Color.green, Color.blue);

enum Gender {
    男,
    女,
}
console.log(Gender.男)

// any
let any_type: any = 100;
any_type = 'hahahhah';
console.log(any_type)
// [string, number, boolean]
let any_tuple: any[] = [100.1234, 'a str', true];
// console.log(any_tuple[0].split(''))


// void
function showMsg(): void {
    console.log('show msg')
    return
    // return undefined
    // return null
}

console.log(showMsg())


// 声明一个 void 类型的变量没有什么大用，因为你只能为它赋予 undefined 和 null：
let unusable: void = undefined

// object
function getObj(obj: object): object {
    console.log(obj)
    return {
        user_name: 'bob',
        age: '30',
    }
}

// console.log(getObj({name: 'paul', gender: 'man'}))
// console.log(getObj('1234')) // err
// console.log(getObj(String))
console.log(getObj(new String))

// 联合类型
// 联合类型（Union Types）表示取值可以为多种类型中的一种
// 需求1: 定义一个函数得到一个数字或字符串值的字符串形式值
function get_str(str: number | string): string {
    return str.toString()
}
console.log(get_str('123'))

// 类型断言
// 需求2: 定义一个函数得到一个数字或字符串值的长度
function get_len(v: number | string): number {
    // return v.toString().length

    if ((<string>v).length) {
        // return (<string>v).length
        return (v as string).length
    }
    return v.toString().length

}
console.log("len: ", get_len('13456'), get_len(13456))

// 类型推断
// let txt = 100;
// txt = 'hhhaha'  // err
// console.log(txt)

let txt2;   // any type
txt2 = 100;
txt2 = 'hehhehh';
console.log(txt2)