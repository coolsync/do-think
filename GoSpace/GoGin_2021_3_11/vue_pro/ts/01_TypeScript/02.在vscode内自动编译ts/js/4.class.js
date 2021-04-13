(function () {
    var User = /** @class */ (function () {
        function User(firstname, lastname) {
            this.firstname = firstname;
            this.lastname = lastname;
            this.fullName = this.firstname + ' ' + this.lastname;
        }
        return User;
    }());
    function showFullName(p) {
        return p.firstname + ' ' + p.lastname;
    }
    var user = new User("jerry", "18");
    console.log(showFullName(user) === user.fullName);
})();
