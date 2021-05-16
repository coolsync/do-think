(()=>{
    class User {
        firstname: string;
        lastname: string;
        fullName: string;

        constructor (firstname:string, lastname:string) {
            this.firstname = firstname;
            this.lastname = lastname
            this.fullName = this.firstname + ' ' + this.lastname
        }
    }

    interface Person {
        firstname: string;
        lastname: string;
    }

    function showFullName(p: Person) {
        return p.firstname + ' ' + p.lastname
    }

    let user = new User("jerry", "18")

    console.log(showFullName(user) === user.fullName)
})()