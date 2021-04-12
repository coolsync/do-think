(()=> {
    interface Person {
        firstname: string
        lastname: string
    }

    function hello(p: Person) {
        return 'hi ' + p.firstname + p.lastname
    }

    let user = {
        firstname: 'bob',
        lastname: ' say'
    }

    console.log(hello(user))
    console.log
})()