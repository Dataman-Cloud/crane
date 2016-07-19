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
            .state('node.createNetwork', {
                url: '/createNetwork/:node_id',
                templateUrl: '/src/node/create-network/create.html',
                controller: 'NodeCreateNetworkCtrl as nodeCreateNetworkCtrl'
            })
            .state('node.detail', {
                url: '/detail/:node_id',
                templateUrl: '/src/node/detail/detail.html',
                controller: 'NodeDetailCtrl as nodeDetailCtrl',
                targetState: 'container',
                resolve: {
                    node: getNode
                }
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
                controller: 'NodeNetworkCtrl as nodeNetworkCtrl',
                resolve: {
                    networks: listNetworks
                }
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
            })
            .state('node.imageDetail', {
                url: '/imageDetail/:node_id/:image_id',
                templateUrl: '/src/node/image-detail/detail.html',
                controller: 'NodeImageDetailCtrl as nodeImageDetailCtrl',
                targetState: 'config',
                resolve: {
                    image: getImage
                }
            })
            .state('node.imageDetail.config', {
                url: '/config',
                templateUrl: '/src/node/image-detail/config.html',
                controller: 'NodeImageConfigCtrl as nodeImageConfigCtrl'
            })
            .state('node.imageDetail.layer', {
                url: '/layer',
                templateUrl: '/src/node/image-detail/layer.html',
                controller: 'NodeImageLayerCtrl as nodeImageLayerCtrl',
                resolve: {
                    layer: getLayer
                }
            })
            .state('node.containerDetail', {
                url: '/containerDetail/:node_id/:container_id',
                templateUrl: '/src/node/container-detail/detail.html',
                controller: 'NodeContainerDetailCtrl as nodeContainerDetailCtrl',
                targetState: 'config',
                resolve: {
                    container: getContainer
                }

            })
            .state('node.containerDetail.config', {
                url: '/config',
                templateUrl: '/src/node/container-detail/config.html',
                controller: 'NodeContainerConfigCtrl as nodeContainerConfigCtrl'
            })
            .state('node.containerDetail.log', {
                url: '/log',
                templateUrl: '/src/node/container-detail/log.html',
                controller: 'NodeContainerLogCtrl as nodeContainerLogCtrl'
            })
            .state('node.containerDetail.stats', {
                url: '/stats',
                templateUrl: '/src/node/container-detail/stats.html',
                controller: 'NodeContainerStatsCtrl as nodeContainerStatsCtrl'
            })
            .state('node.containerDetail.diff', {
                url: '/diff',
                templateUrl: '/src/node/container-detail/diff.html',
                controller: 'NodeContainerDiffCtrl as nodeContainerDiffCtrl',
                resolve: {
                    diffs: diffContainer
                }
            })
            .state('node.networkDetail', {
                url: '/networkDetail/:node_id/:network_id',
                templateUrl: '/src/node/network-detail/detail.html',
                controller: 'NodeNetworkDetailCtrl as nodeNetworkDetailCtrl',
                targetState: 'config',
                resolve: {
                    network: getNetwork
                }

            })
            .state('node.networkDetail.config', {
                url: '/config',
                templateUrl: '/src/node/network-detail/config.html',
                controller: 'NodeNetworkConfigCtrl as nodeNetworkConfigCtrl'
            });

        /* @ngInject */
        function listNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }

        /* @ngInject */
        function getNode(nodeBackend, $stateParams) {
            return nodeBackend.getNode($stateParams.node_id);
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

        /* @ngInject */
        function listNetworks(nodeBackend, $stateParams) {
            return nodeBackend.listNetworks($stateParams.node_id);
        }

        /* @ngInject */
        function getImage(nodeBackend, $stateParams) {
            return nodeBackend.getImage($stateParams.node_id, $stateParams.image_id)
        }

        /* @ngInject */
        function getLayer(nodeBackend, $stateParams) {
            return nodeBackend.getImageHistory($stateParams.node_id, $stateParams.image_id)
        }

        /* @ngInject */
        function getContainer(nodeBackend, $stateParams) {
            return nodeBackend.getContainer($stateParams.node_id, $stateParams.container_id)
        }

        /* @ngInject */
        function diffContainer(nodeBackend, $stateParams) {
            return nodeBackend.diffContainer($stateParams.node_id, $stateParams.container_id);
        }

        /* @ngInject */
        function getNetwork(nodeBackend, $stateParams) {
            return nodeBackend.getNetwork($stateParams.node_id, $stateParams.network_id)
        }
    }
})();
