(() => {
    // 需求: 创建人的对象, 需要对人的属性进行一定的约束
    // id是number类型, 必须有, 只读的
    // name是string类型, 必须有
    // age是number类型, 必须有
    // sex是string类型, 可以没有

    interface IPerson {
        readonly id: number,  // 只读属性: readonly
        name: string,
        age: number,
        sex?: string    // 可选属性: ?
    }

    const person: IPerson = {
        id: 1,
        name: "bob",
        age: 20,
        // sex: '男'
    };

    // person.id = 100;    // err: Cannot assign to 'id' because it is a read-only property.
    console.log(person);

    

})()