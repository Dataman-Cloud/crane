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
                targetState: 'list'
            })
            .state('network.list', {
                url: '/list',
                templateUrl: '/src/network/list/list.html',
                controller: 'NetworkListCtrl as networkListCtrl',
                resolve: {
                    networks: getNetworks
                }
            })
            .state('network.create', {
                url: '/create',
                templateUrl: '/src/network/create/create.html',
                controller: 'NetworkCreateCtrl as networkCreateCtrl'
            })
            .state('network.detail', {
                url: '/detail/:network_id',
                templateUrl: '/src/network/detail/detail.html',
                controller: 'NetworkDetailCtrl as NetworkDetailCtrl',
                targetState: 'service'
            })
            .state('network.detail.service', {
                url: '/service',
                templateUrl: '/src/network/detail/service.html',
                controller: 'NetworkServiceCtrl as networkServiceCtrl'
            });

        /* @ngInject */
        function getNetworks(networkBackend) {
            return networkBackend.listNetwork()
        }
    }
})();
