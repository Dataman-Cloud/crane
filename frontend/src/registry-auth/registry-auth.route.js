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
                    label: '仓库认证'
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
                    label: '管理列表'
                }
            })
            .state('registryAuth.create', {
                url: '/create',
                templateUrl: '/src/registry-auth/create/create.html',
                controller: 'RegistryAuthCreateCtrl as regAuthCreateCtrl',
                ncyBreadcrumb: {
                    label: '添加认证'
                }
            });

        /* @ngInject */
        function listReAuth(registryAuthBackend) {
            return registryAuthBackend.listRegAuth();
        }
    }
})();
