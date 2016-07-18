(function () {
    'use strict';
    angular.module('app.misc')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('misc', {
                url: '/misc',
                template: '<ui-view/>',
                targetState: 'config'
            })
            .state('misc.config', {
                url: '/config',
                templateUrl: '/src/misc/rolexconfig/rolexconfig.html',
                controller: 'MiscConfigCtrl as miscConfigCtrl',
                resolve: {
                    rolexconfig: getConfig
                }
            });

        /* @ngInject */
        function getConfig(misc) {
            return misc.rolexConfig()
        }
    }
})();
