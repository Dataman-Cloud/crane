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
                targetState: 'config',
                ncyBreadcrumb: {
                    label: '设置'
                }
            })
            .state('misc.config', {
                url: '/config',
                templateUrl: '/src/misc/rolexconfig/rolexconfig.html',
                controller: 'MiscConfigCtrl as miscConfigCtrl',
                resolve: {
                    rolexconfig: getConfig
                },
                ncyBreadcrumb: {
                    label: '配置'
                }
            });

        /* @ngInject */
        function getConfig(misc) {
            return misc.rolexConfig()
        }
    }
})();
