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
                    nodes: listNodes
                }
            })
            .state('node.create', {
                url: '/create',
                templateUrl: '/src/node/create/create.html',
                controller: 'NodeCreateCtrl as nodeCreateCtrl'
            })
            .state('node.createVolume', {
                url: '/createVolume/:node_id',
                templateUrl: '/src/node/create-volume/create.html',
                controller: 'NodeCreateVolumeCtrl as nodeCreateVolumeCtrl'
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
                controller: 'NodeContainerCtrl as nodeContainerCtrl',
                resolve: {
                    containers: listContainers
                }
            })
            .state('node.detail.network', {
                url: '/network',
                templateUrl: '/src/node/detail/network.html',
                controller: 'NodeNetworkCtrl as nodeNetworkCtrl'
            })
            .state('node.detail.volume', {
                url: '/volume',
                templateUrl: '/src/node/detail/volume.html',
                controller: 'NodeVolumeCtrl as nodeVolumeCtrl',
                resolve: {
                    volumes: listVolumes
                }
            })
            .state('node.detail.image', {
                url: '/image',
                templateUrl: '/src/node/detail/image.html',
                controller: 'NodeImageCtrl as nodeImageCtrl',
                resolve: {
                    images: listImages
                }
            });

        /* @ngInject */
        function listNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }

        /* @ngInject */
        function listVolumes(nodeBackend, $stateParams) {
            return nodeBackend.listVolumes($stateParams.node_id)
        }

        /* @ngInject */
        function listImages(nodeBackend, $stateParams) {
            return nodeBackend.listImages($stateParams.node_id)
        }
        
        /* @ngInject */
        function listContainers(nodeBackend, $stateParams) {
            return nodeBackend.listContainers($stateParams.node_id);
        }
    }
})();
