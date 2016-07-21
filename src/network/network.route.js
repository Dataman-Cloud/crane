(function () {
    'use strict';
    angular.module('app.network')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('network', {
                url: '/network',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: '网络'
                }
            })
            .state('network.list', {
                url: '/list',
                templateUrl: '/src/network/list/list.html',
                controller: 'NetworkListCtrl as networkListCtrl',
                resolve: {
                    networks: listNetwork
                },
                ncyBreadcrumb: {
                    label: '网络列表'
                }
            })
            .state('network.create', {
                url: '/create',
                templateUrl: '/src/network/create/create.html',
                controller: 'NetworkCreateCtrl as networkCreateCtrl',
                ncyBreadcrumb: {
                    label: '创建网络'
                }
            })
            .state('network.detail', {
                url: '/detail/:network_id',
                templateUrl: '/src/network/detail/detail.html',
                controller: 'NetworkDetailCtrl as networkDetailCtrl',
                targetState: 'config',
                resolve: {
                    network: getNetwork
                },
                ncyBreadcrumb: {
                    label: '网络详情'
                }
            })
            .state('network.detail.container', {
                url: '/container',
                templateUrl: '/src/network/detail/container.html',
                controller: 'NetworkContainerCtrl as networkContainerCtrl',
                ncyBreadcrumb: {
                    label: '容器列表'
                }
            })
            .state('network.detail.config', {
                url: '/config',
                templateUrl: '/src/network/detail/config.html',
                controller: 'NetworkConfigCtrl as networkConfigCtrl',
                ncyBreadcrumb: {
                    label: '配置'
                }
            });

        /* @ngInject */
        function listNetwork(networkBackend) {
            return networkBackend.listNetwork()
        }

        /* @ngInject */
        function getNetwork(networkBackend, $stateParams) {
            return networkBackend.getNetwork($stateParams.network_id)
        }
    }
})();
