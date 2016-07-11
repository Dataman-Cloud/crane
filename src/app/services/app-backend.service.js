(function () {
    'use strict';
    angular.module('app.app')
        .factory('appBackend', appBackend);


    /* @ngInject */
    function appBackend(gHttp) {
        return {
            listApps: listApps
        };

        function listApps(params, loading) {
            return gHttp.Resource('service.services').get({params: params, "loading": loading});
        }
    }
})();