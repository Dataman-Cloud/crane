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
                url: '/detail/:stack_name',
                templateUrl: '/src/stack/detail/detail.html',
                controller: 'StackDetailCtrl as stackDetailCtrl',
                targetState: 'service',
                resolve: {
                    stack: getStack
                }
            })
            .state('stack.detail.service', {
                url: '/service',
                templateUrl: '/src/stack/detail/service.html',
                controller: 'StackServiceCtrl as stackServiceCtrl',
                resolve: {
                    services: listStackServices 
                }
            })
            .state('stack.serviceDetail', {
                url: '/serviceDetail/:stack_name/:service_id',
                templateUrl: '/src/stack/service-detail/detail.html',
                controller: 'ServiceDetailCtrl as serviceDetailCtrl',
                targetState: 'config',
                resolve: {
                    service: getService
                }
            })
            .state('stack.serviceDetail.config', {
                url: '/config',
                templateUrl: '/src/stack/service-detail/config.html',
                controller: 'ServiceConfigCtrl as serviceConfigCtrl'
            })
            .state('stack.serviceDetail.task', {
                url: '/task',
                templateUrl: '/src/stack/service-detail/task.html',
                controller: 'ServiceTaskCtrl as serviceTaskCtrl',
                resolve: {
                    tasks: listServiceTasks
                }
            });
    }
    
    /*@ngInject*/
    function listStacks(stackBackend) {
        return stackBackend.listStacks();
    }
    
    /*@ngInject*/
    function getStack(stackBackend, $stateParams) {
        return stackBackend.getStack($stateParams.stack_name);
    }
    
    /*@ngInject*/
    function listStackServices(stackBackend, $stateParams) {
        return stackBackend.listStackServices($stateParams.stack_name);
    }

    /*@ngInject*/
    function getService(stackBackend, $stateParams) {
        return stackBackend.getService($stateParams.stack_name, $stateParams.service_id);
    }

    /*@ngInject*/
    function listServiceTasks(stackBackend, $stateParams) {
        return stackBackend.listServiceTasks($stateParams.stack_name, $stateParams.service_id);
    }
})();
