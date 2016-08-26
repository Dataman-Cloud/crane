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
                    label: 'reAuth'
                }
            })
            .state('registryAuth.list', {
                url: '/list',
                templateUrl: '/src/registry-auth/list/list.html',
                controller: 'RegistryAuthListCtrl as regAuthListCtrl',
                ncyBreadcrumb: {
                    label: 'reAuthListCtrl'
                }
            })
            .state('registryAuth.create', {
                url: '/create',
                templateUrl: '/src/registry-auth/create/create.html',
                controller: 'RegistryAuthCreateCtrl as regAuthCreateCtrl',
                ncyBreadcrumb: {
                    label: 'reAuthCreateCtrl'
                }
            });

        /* @ngInject */
        function listReAuth() {
            //TODO
        }
    }
    
})();
