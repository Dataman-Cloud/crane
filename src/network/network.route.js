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
                controller: 'NetworkListCtrl as networkListCtrl'
            })
            .state('network.create', {
                url: '/create',
                templateUrl: '/src/network/create/create.html',
                controller: 'CreateCtrl as createCtrl'
            });
    }
})();
