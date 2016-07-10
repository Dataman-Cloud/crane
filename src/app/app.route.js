(function () {
    'use strict';
    angular.module('glance.app')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('app', {
                url: '/app',
                template: '<ui-view/>',
                targetState: 'list'
            })
            .state('app.list', {
                url: '/list',
                templateUrl: '/src/app/list/list.html',
                controller: 'AppListCtrl as appListCtrl',
                resolve: {
                    apps: getApps
                }
            })
            .state('app.create', {
                url: '/create',
                templateUrl: '/src/app/createupdate/create-update.html',
                controller: 'CreateUpdateCtrl as createUpdateCtrl',
                resolve: {
                    target: function () {
                        return 'create'
                    }
                }
            })
            .state('app.update', {
                url: '/update',
                templateUrl: '/src/app/createupdate/create-update.html',
                controller: 'CreateUpdateCtrl as createUpdateCtrl',
                resolve: {
                    target: function () {
                        return 'update'
                    }
                }
            });

        /* @ngInject */
        function getApps(appBackend) {
            return appBackend.listApps()
        }
    }
})();
