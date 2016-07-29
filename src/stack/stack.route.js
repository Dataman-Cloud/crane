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
                targetState: 'list',
                ncyBreadcrumb: {
                    label: '编排'
                }
            })
            .state('stack.list', {
                url: '/list',
                templateUrl: '/src/stack/list/list.html',
                controller: 'StackListCtrl as stackListCtrl',
                resolve: {
                    stacks: listStacks
                },
                ncyBreadcrumb: {
                    label: '编排列表'
                }
            })
            .state('stack.create', {
                url: '/create',
                templateUrl: '/src/stack/create/create-nav.html',
                ncyBreadcrumb: {
                    label: '选择创建方式'
                }
            })
            .state('stack.createByJson', {
                url: '/createByJson',
                templateUrl: '/src/stack/create/create-by-Json.html',
                controller: 'StackCreateByJsonCtrl as stackCreateByJsonCtrl',
                ncyBreadcrumb: {
                    label: '创建编排'
                }
            })
            .state('stack.createByForm', {
                url: '/createByForm',
                templateUrl: '/src/stack/create/create-by-form.html',
                controller: 'StackCreateByFormCtrl as stackCreateByFormCtrl',
                ncyBreadcrumb: {
                    label: '创建编排'
                }
            })
            .state('stack.serviceUpdate', {
                url: '/:stack_name/:service_id/serviceUpdate',
                templateUrl: '/src/stack/service-update/update.html',
                controller: 'ServiceUpdateCtrl as serviceUpdateCtrl',
                ncyBreadcrumb: {
                    label: '服务更新'
                },
                resolve: {
                    service: getService
                }
            })
            .state('stack.detail', {
                url: '/detail/:stack_name',
                templateUrl: '/src/stack/detail/detail.html',
                controller: 'StackDetailCtrl as stackDetailCtrl',
                targetState: 'service',
                resolve: {
                    stack: getStack
                },
                ncyBreadcrumb: {
                    label: '编排详情'
                }
            })
            .state('stack.detail.service', {
                url: '/service',
                templateUrl: '/src/stack/detail/service.html',
                controller: 'StackServiceCtrl as stackServiceCtrl',
                resolve: {
                    services: listStackServices 
                },
                ncyBreadcrumb: {
                    label: '服务列表'
                }
            })
            .state('stack.serviceDetail', {
                url: '/serviceDetail/:stack_name/:service_id',
                templateUrl: '/src/stack/service-detail/detail.html',
                controller: 'ServiceDetailCtrl as serviceDetailCtrl',
                targetState: 'config',
                resolve: {
                    service: getService
                },
                ncyBreadcrumb: {
                    label: '服务详情',
                    parent: 'stack.detail'
                }
            })
            .state('stack.serviceDetail.config', {
                url: '/config',
                templateUrl: '/src/stack/service-detail/config.html',
                controller: 'ServiceConfigCtrl as serviceConfigCtrl',
                ncyBreadcrumb: {
                    label: '配置'
                }
            })
            .state('stack.serviceDetail.task', {
                url: '/task',
                templateUrl: '/src/stack/service-detail/task.html',
                controller: 'ServiceTaskCtrl as serviceTaskCtrl',
                resolve: {
                    tasks: listServiceTasks
                },
                ncyBreadcrumb: {
                    label: '任务列表'
                }
            })
            .state('stack.serviceDetail.log', {
                url: '/log',
                templateUrl: '/src/stack/service-detail/log.html',
                controller: 'ServiceLogCtrl as serviceLogCtrl',
                ncyBreadcrumb: {
                    label: '日志'
                }
            })
            .state('stack.serviceDetail.stats', {
                url: '/stats',
                templateUrl: '/src/stack/service-detail/stats.html',
                controller: 'ServiceStatsCtrl as serviceStatsCtrl',
                ncyBreadcrumb: {
                    label: '统计'
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
