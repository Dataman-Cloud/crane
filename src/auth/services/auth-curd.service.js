(function () {
    'use strict';

    angular.module('app.auth').factory('authCurd', authCurd);

    /* ngInject */
    function authCurd(authBackend, $window) {

        return {
            login: login
        };

        function login(data, form, returnTo) {
            if (!returnTo) {
                returnTo = "/index.html?timestamp="+new Date().getTime();
            }
            return authBackend.login(data, form).then(function (data) {
                $window.localStorage .setItem('token', data);
                $window.location.href = returnTo;
            });
        }
    }
})();