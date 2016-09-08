(function () {
    'use strict';
    angular.module('app.misc')
        .config(route);

    /* @ngInject */
    function route($stateProvider) {
        $stateProvider
            .state('misc', {
                url: '/misc',
                template: '<ui-view/>',
                targetState: 'config',
                ncyBreadcrumb: {
                    label: '信息'
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
                    skip: true
                }
            });

        /* @ngInject */
        function getConfig(miscBackend) {
            return miscBackend.rolexConfig()
        }
    }
})();
