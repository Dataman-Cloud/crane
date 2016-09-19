(function () {
    'use strict';
    angular.module('app.registryAuth')
        .config(route);

    /* @ngInject */
    function route($stateProvider) {
        $stateProvider
            .state('registryAuth', {
                url: '/registryAuth',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: "{/'Registry Auth' | translate/}"
                }
            })
            .state('registryAuth.list', {
                url: '/list',
                templateUrl: '/src/registry-auth/list/list.html',
                controller: 'RegistryAuthListCtrl as regAuthListCtrl',
                resolve: {
                    reAuths: listReAuth
                },
                ncyBreadcrumb: {
                    label: "{/'Registry Auth list' | translate/}"
                }
            })
            .state('registryAuth.create', {
                url: '/create',
                templateUrl: '/src/registry-auth/create/create.html',
                controller: 'RegistryAuthCreateCtrl as regAuthCreateCtrl',
                ncyBreadcrumb: {
                    label: "{/'Add registry auth' | translate/}"
                }
            });

        /* @ngInject */
        function listReAuth(registryAuthBackend) {
            return registryAuthBackend.listRegAuth();
        }
    }
})();
