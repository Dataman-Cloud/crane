(function () {
    'use strict';
    angular.module('app.misc')
        .factory('miscBackend', miscBackend);


    /* @ngInject */
    function miscBackend(gHttp) {
        return {
          rolexConfig: rolexConfig
        };

        function rolexConfig() {
            return gHttp.Resource('misc.config').get();
        }
    }
})();
