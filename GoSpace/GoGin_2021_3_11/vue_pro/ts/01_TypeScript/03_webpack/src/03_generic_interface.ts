(() => {
    // def generic interface
    interface IBase<T> {
        // store user info array
        data: Array<T>;
        // add user
        add: (t: T) => T;
        // by id query user
        getUserId: (number) => T;

        deleteUser: (number, T) => T;
    }
    // def user info class
    class User {
        id?: number;
        name: string;
        age: number;
        constructor(name: string, age: number) {
            this.name = name;
            this.age = age;
        }
    }

    // def user crud class, impl IBase generic inter
    class UserCRUD implements IBase<User> {
        data = [];

        add(user: User): User {
            user.id = Date.now() + Math.random();
            this.data.push(user);
            return user
        }

        getUserId(id: number): User {
            return this.data.find(user => user.id === id);
        }

        deleteUser(id: number,): User {
            // let node： int;  //要移除的对象
            // nodes： int[];
            // this.nodes = this.nodes.filter(item => item !== node);
            let user = this.getUserId(id);
            user.id = id;
            this.data = this.data.filter(item => item !== user);
            return user
        }
    }

    // instance UserCRUD
    const userCRUD: UserCRUD = new UserCRUD();

    userCRUD.add(new User("bob", 30));
    userCRUD.add(new User("paul", 20));
    let { id } = userCRUD.add(new User("jerry", 18));
    userCRUD.add(new User("mark", 31));

    console.log(userCRUD.data);
    console.log(userCRUD.getUserId(id));

    console.log(id)
    userCRUD.deleteUser(id);
    console.log(userCRUD.data);
})()