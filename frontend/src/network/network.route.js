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
                    label: "{/'Network' | translate/}"
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
                    label: "{/'Networks' | translate/}"
                }
            })
            .state('network.create', {
                url: '/create',
                templateUrl: '/src/network/create/create.html',
                controller: 'NetworkCreateCtrl as networkCreateCtrl',
                ncyBreadcrumb: {
                    label: "{/'Create Network' | translate/}"
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
                    label: "{/'Network Detail' | translate/}"
                }
            })
            .state('network.detail.config', {
                url: '/config',
                templateUrl: '/src/network/detail/config.html',
                controller: 'NetworkConfigCtrl as networkConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
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
