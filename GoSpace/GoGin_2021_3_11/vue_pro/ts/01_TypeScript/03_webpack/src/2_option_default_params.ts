// use ? -> option params
(()=>{
    const getFullName = function(firstName: string = 'bob', lastName?: string) {
        if (lastName) {
            return firstName + '_' + lastName;
        }
        return firstName
    }

    // no any params
    console.log(getFullName());

    // only firstName param
    console.log(getFullName('jerry'));

    // all params
    console.log(getFullName('dongfa', 'bubai'));
})()