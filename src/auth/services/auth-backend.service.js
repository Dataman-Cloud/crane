(function () {
    'use strict';

    angular.module('app.auth').factory('authBackend', authBackend);

    /* @ngInject */
    function authBackend(gHttp) {

        return {
            login: login
        };

        /////////
        function login(data, form) {
            return gHttp.Resource('auth.login').post(data, {form: form});
        }
    }
})();