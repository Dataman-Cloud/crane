(function () {
    'use strict';
    angular.module('glance.layout')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('layout', {
                url: '/layout',
                template: '<ui-view/>',
                targetState: 'list'
            })
            .state('layout.list', {
                url: '/list',
                templateUrl: '/src/applayout/list/list.html',
                controller: 'LayoutListCtrl as layoutListCtrl'
            })
            .state('layout.create', {
                url: '/create',
                templateUrl: '/src/applayout/createupdate/create-update.html',
                controller: 'LayoutCreateCtrl as layoutCreateCtrl',
                resolve: {
                    target: function () {
                        return 'create'
                    }
                }
            })
            .state('layout.update', {
                url: '/update/:cluster_id/:stack_id',
                templateUrl: '/src/applayout/createupdate/create-update.html',
                controller: 'LayoutCreateCtrl as layoutCreateCtrl',
                resolve: {
                    target: function () {
                        return 'update'
                    }
                }
            });
    }
})();
