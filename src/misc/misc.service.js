(function () {
    'use strict';
    angular.module('app.misc')
        .factory('misc', misc);


    /* @ngInject */
    function misc(gHttp) {
        return {
          rolexConfig: rolexConfig
        };

        function rolexConfig() {
            return gHttp.Resource('misc.config').get();
        }
    }
})();
