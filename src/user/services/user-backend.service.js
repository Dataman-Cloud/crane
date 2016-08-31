(function () {
    'use strict';

    angular.module('app.user').factory('userBackend', userBackend);

    /* @ngInject */
    function userBackend(gHttp) {

        return {
            logout: logout,
            aboutMe: aboutMe,
            listGroup: listGroup,
            checkLicense: checkLicense,
            registerLicense: registerLicense
        };

        function logout() {
            return gHttp.Resource('auth.logout').post();
        }

        function aboutMe() {
            return gHttp.Resource('auth.aboutme').get();
        }

        function listGroup(id) {
            return gHttp.Resource('auth.groups', {account_id: id}).get();
        }

        function checkLicense() {
            return gHttp.Resource('license.license').get({ignoreCodes: [16003]});
        }

        function registerLicense(license) {
            var data = {
                License: license
            };

            return gHttp.Resource('license.license').post(data);
        }
    }
})();