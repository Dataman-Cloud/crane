(function () {
    'use strict';

    angular.module('app.auth').factory('authCurd', authCurd);

    /* ngInject */
    function authCurd(authBackend, $cookies, $window) {

        return {
            login: login
        };

        function login(data, form, returnTo) {
            if (!returnTo) {
                returnTo = "/index.html?timestamp="+new Date().getTime();
            }
            return authBackend.login(data, form).then(function (data) {
                $cookies.put('token', data);
                $window.location.href = returnTo;
            });
        }
    }
})();