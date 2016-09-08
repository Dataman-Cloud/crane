(function () {
    'use strict';
    angular.module('app.misc')
        .factory('miscBackend', miscBackend);


    /* @ngInject */
    function miscBackend(gHttp) {
        return {
          craneConfig: craneConfig
        };

        function craneConfig() {
            return gHttp.Resource('misc.config').get();
        }
    }
})();
