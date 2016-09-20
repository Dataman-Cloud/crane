(function () {
    'use strict';
    angular.module('app.stack')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $urlRouterProvider) {

        //warning: otherwise(url) will be redirect loop on state with errored resolve
        $urlRouterProvider.otherwise(function($injector) {
            var $state = $injector.get('$state');
            $state.go('stack');
        });

        $stateProvider
            .state('stack', {
                url: '/stack',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: "{/'Stack' | translate/}"
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
                    label: "{/'Stacks' | translate/}"
                }
            })
            .state('stack.createByJson', {
                url: '/createByJson',
                templateUrl: '/src/stack/create/create-by-Json.html',
                controller: 'StackCreateByJsonCtrl as stackCreateByJsonCtrl',
                ncyBreadcrumb: {
                    label: "{/'Create Stack' | translate/}"
                }
            })
            .state('stack.createByForm', {
                url: '/createByForm',
                templateUrl: '/src/stack/create/create-by-form.html',
                controller: 'StackCreateByFormCtrl as stackCreateByFormCtrl',
                ncyBreadcrumb: {
                    label: "{/'Create Stack' | translate/}"
                }
            })
            .state('stack.serviceUpdate', {
                url: '/:stack_name/:service_id/serviceUpdate',
                templateUrl: '/src/stack/service-update/update.html',
                controller: 'ServiceUpdateCtrl as serviceUpdateCtrl',
                ncyBreadcrumb: {
                    label:  "{/'Service Update' | translate/}"
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
                    label: "{/'Stack Detail' | translate/}"
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
                    skip: true
                }
            })
            .state('stack.serviceDetail', {
                url: '/serviceDetail/:stack_name/:service_id',
                templateUrl: '/src/stack/service-detail/detail.html',
                controller: 'ServiceDetailCtrl as serviceDetailCtrl',
                targetState: 'task',
                resolve: {
                    service: getService
                },
                ncyBreadcrumb: {
                    label: "{/'Service Detail' | translate/}",
                    parent: 'stack.detail'
                }
            })
            .state('stack.serviceDetail.config', {
                url: '/config',
                templateUrl: '/src/stack/service-detail/config.html',
                controller: 'ServiceConfigCtrl as serviceConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
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
                    skip: true
                }
            })
            .state('stack.serviceDetail.log', {
                url: '/log',
                templateUrl: '/src/stack/service-detail/log.html',
                controller: 'ServiceLogCtrl as serviceLogCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('stack.serviceDetail.stats', {
                url: '/stats',
                templateUrl: '/src/stack/service-detail/stats.html',
                controller: 'ServiceStatsCtrl as serviceStatsCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('stack.serviceDetail.discovery', {
                url: '/discovery',
                templateUrl: '/src/stack/service-detail/discovery.html',
                controller: 'ServiceDiscoveryCtrl as serviceDiscoveryCtrl',
                resolve: {
                    service: getService,
                    nodes: listNodes
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('stack.serviceDetail.cd', {
                url: '/cd',
                templateUrl: '/src/stack/service-detail/cd.html',
                controller: 'ServiceCdCtrl as serviceCdCtrl',
                resolve: {
                  serviceCdUrl: getServiceCDUrl
                },
                ncyBreadcrumb: {
                    skip: true
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

    /* @ngInject */
    function listNodes(nodeBackend) {
        return nodeBackend.listNodes()
    }

    /*@ngInject*/
    function getService(stackBackend, $stateParams) {
        return stackBackend.getService($stateParams.stack_name, $stateParams.service_id);
    }

    /*@ngInject*/
    function listServiceTasks(stackBackend, $stateParams) {
        return stackBackend.listServiceTasks($stateParams.stack_name, $stateParams.service_id);
    }

    /*@ngInject*/
    function getServiceCDUrl(stackBackend, $stateParams) {
        return stackBackend.getServiceCDUrl($stateParams.stack_name, $stateParams.service_id);
    }
})();
