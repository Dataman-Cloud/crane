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
                    label: "{/'Information' | translate/}"
                }
            })
            .state('misc.config', {
                url: '/config',
                templateUrl: '/src/misc/craneconfig/craneconfig.html',
                controller: 'MiscConfigCtrl as miscConfigCtrl',
                resolve: {
                    craneconfig: getConfig
                },
                ncyBreadcrumb: {
                    skip: true
                }
            });

        /* @ngInject */
        function getConfig(miscBackend) {
            return miscBackend.craneConfig()
        }
    }
})();
