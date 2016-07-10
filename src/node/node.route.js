(function () {
    'use strict';
    angular.module('glance.node')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('node', {
                url: '/node',
                template: '<ui-view/>',
                targetState: 'list'
            })
            .state('node.list', {
                url: '/list',
                templateUrl: '/src/node/list/list.html',
                controller: 'NodeListCtrl as nodeListCtrl',
                resolve: {
                    nodes: getNodes
                }
            })
            .state('node.create', {
                url: '/create',
                templateUrl: '/src/node/create/create.html',
                controller: 'CreateCtrl as createCtrl'
            });

        /* @ngInject */
        function getNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }
    }
})();
