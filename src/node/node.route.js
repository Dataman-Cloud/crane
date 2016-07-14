(function () {
    'use strict';
    angular.module('app.node')
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
                controller: 'NodeCreateCtrl as nodeCreateCtrl'
            })
            .state('node.detail', {
                url: '/detail/:node_id',
                templateUrl: '/src/node/detail/detail.html',
                controller: 'NodeDetailCtrl as nodeDetailCtrl',
                targetState: 'container'
            })
            .state('node.detail.container', {
                url: '/container',
                templateUrl: '/src/node/detail/container.html',
                controller: 'NodeContainerCtrl as nodeContainerCtrl'
            });

        /* @ngInject */
        function getNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }
    }
})();
