(function () {
    'use strict';
    angular.module('app.registryAuth')
        .factory('registryAuthBackend', registryAuthBackend);


    /* @ngInject */
    function registryAuthBackend(gHttp) {
        return {
            createRegAuth: createRegAuth,
            deleteRegAuth: deleteRegAuth,
            listRegAuth: listRegAuth
        };

        function createRegAuth(data, form) {
            return gHttp.Resource('registryauth.registryauths').post(data, {form: form});
        }

        function deleteRegAuth(name) {
            return gHttp.Resource('registryauth.registryauth', {'regauth_name': name}).delete();
        }

        function listRegAuth() {
            return gHttp.Resource('registryauth.registryauths').get();
        }
    }
})();
