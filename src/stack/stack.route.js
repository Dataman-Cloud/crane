(function () {
    'use strict';
    angular.module('app.stack')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('stack', {
                url: '/stack',
                template: '<ui-view/>',
                targetState: 'list'
            })
            .state('stack.list', {
                url: '/list',
                templateUrl: '/src/stack/list/list.html',
                controller: 'StackListCtrl as stackListCtrl'
            })
            .state('stack.create', {
                url: '/create',
                templateUrl: '/src/stack/createupdate/create-update.html',
                controller: 'StackCreateCtrl as stackCreateCtrl',
                resolve: {
                    target: function () {
                        return 'create'
                    }
                }
            })
            .state('stack.update', {
                url: '/update/:cluster_id/:stack_id',
                templateUrl: '/src/stack/createupdate/create-update.html',
                controller: 'StackCreateCtrl as stackCreateCtrl',
                resolve: {
                    target: function () {
                        return 'update'
                    }
                }
            });
    }
})();
