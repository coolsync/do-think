(function () {
    function hello(p) {
        return 'hi ' + p.firstname + p.lastname;
    }
    var user = {
        firstname: 'bob',
        lastname: ' say'
    };
    console.log(hello(user));
    console.log;
})();
