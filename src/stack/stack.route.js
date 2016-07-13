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
                controller: 'StackListCtrl as stackListCtrl',
                resolve: {
                    stacks: listStacks
                }
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
                url: '/update/:stack_name',
                templateUrl: '/src/stack/createupdate/create-update.html',
                controller: 'StackCreateCtrl as stackCreateCtrl',
                resolve: {
                    target: function () {
                        return 'update'
                    }
                }
            })
            .state('stack.detail', {
                url: '/detail/:stack_name'
            });
    }
    
    /*@ngInject*/
    function listStacks(stackBackend) {
        return stackBackend.listStacks();
    }
})();
