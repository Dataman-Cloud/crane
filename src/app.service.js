(function () {
    'use strict';
    angular.module('app')
        .factory('appCommon', appCommon);


    /* @ngInject */
    function appCommon(gHttp) {
        return {
            logout: logout
        };

        function logout(){
            return gHttp.Resource('auth.logout').post();
        }
    }
})();