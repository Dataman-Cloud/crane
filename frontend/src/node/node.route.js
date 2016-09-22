(function () {
    'use strict';
    angular.module('app.node')
        .config(route);

    /* @ngInject */
    function route($stateProvider) {
        $stateProvider
            .state('node', {
                url: '/node',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: "{/'Node' | translate/}"
                }
            })
            .state('node.list', {
                url: '/list',
                templateUrl: '/src/node/list/list.html',
                controller: 'NodeListCtrl as nodeListCtrl',
                resolve: {
                    nodes: listNodes
                },
                ncyBreadcrumb: {
                    label: "{/'Nodes' | translate/}"
                }
            })
            .state('node.createVolume', {
                url: '/createVolume/:node_id',
                templateUrl: '/src/node/create-volume/create.html',
                controller: 'NodeCreateVolumeCtrl as nodeCreateVolumeCtrl',
                ncyBreadcrumb: {
                    label: "{/'Create Volume' | translate/}"
                }
            })
            .state('node.createNetwork', {
                url: '/createNetwork/:node_id',
                templateUrl: '/src/node/create-network/create.html',
                controller: 'NodeCreateNetworkCtrl as nodeCreateNetworkCtrl',
                ncyBreadcrumb: {
                    label: "{/'Create Host-Only network' | translate/}"
                }
            })
            .state('node.detail', {
                url: '/detail/:node_id',
                templateUrl: '/src/node/detail/detail.html',
                controller: 'NodeDetailCtrl as nodeDetailCtrl',
                targetState: 'config',
                resolve: {
                    node: getNode
                },
                ncyBreadcrumb: {
                    label: "{/'Node Detail' | translate/}"
                }
            })
            .state('node.detail.config', {
                url: '/config',
                templateUrl: '/src/node/detail/config.html',
                controller: 'NodeConfigCtrl as nodeConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.detail.container', {
                url: '/container',
                templateUrl: '/src/node/detail/container.html',
                controller: 'NodeContainerCtrl as nodeContainerCtrl',
                resolve: {
                    containers: listContainers
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.detail.network', {
                url: '/network',
                templateUrl: '/src/node/detail/network.html',
                controller: 'NodeNetworkCtrl as nodeNetworkCtrl',
                resolve: {
                    networks: listNetworks
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.detail.volume', {
                url: '/volume',
                templateUrl: '/src/node/detail/volume.html',
                controller: 'NodeVolumeCtrl as nodeVolumeCtrl',
                resolve: {
                    volumes: listVolumes
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.detail.image', {
                url: '/image',
                templateUrl: '/src/node/detail/image.html',
                controller: 'NodeImageCtrl as nodeImageCtrl',
                resolve: {
                    images: listImages
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.imageDetail', {
                url: '/imageDetail/:node_id/:image_id',
                templateUrl: '/src/node/image-detail/detail.html',
                controller: 'NodeImageDetailCtrl as nodeImageDetailCtrl',
                targetState: 'config',
                resolve: {
                    image: getImage
                },
                ncyBreadcrumb: {
                    label: "{/'Image Detail' | translate/}",
                    parent: 'node.detail'
                }
            })
            .state('node.imageDetail.config', {
                url: '/config',
                templateUrl: '/src/node/image-detail/config.html',
                controller: 'NodeImageConfigCtrl as nodeImageConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.imageDetail.layer', {
                url: '/layer',
                templateUrl: '/src/node/image-detail/layer.html',
                controller: 'NodeImageLayerCtrl as nodeImageLayerCtrl',
                resolve: {
                    layer: getLayer
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.containerDetail', {
                url: '/containerDetail/:node_id/:container_id',
                templateUrl: '/src/node/container-detail/detail.html',
                controller: 'NodeContainerDetailCtrl as nodeContainerDetailCtrl',
                targetState: 'config',
                resolve: {
                    container: getContainer
                },
                ncyBreadcrumb: {
                    label: "{/'Container Detail' | translate/}",
                    parent: 'node.detail'
                }

            })
            .state('node.containerDetail.config', {
                url: '/config',
                templateUrl: '/src/node/container-detail/config.html',
                controller: 'NodeContainerConfigCtrl as nodeContainerConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.containerDetail.log', {
                url: '/log',
                templateUrl: '/src/node/container-detail/log.html',
                controller: 'NodeContainerLogCtrl as nodeContainerLogCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.containerDetail.stats', {
                url: '/stats',
                templateUrl: '/src/node/container-detail/stats.html',
                controller: 'NodeContainerStatsCtrl as nodeContainerStatsCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.containerDetail.diff', {
                url: '/diff',
                templateUrl: '/src/node/container-detail/diff.html',
                controller: 'NodeContainerDiffCtrl as nodeContainerDiffCtrl',
                resolve: {
                    diffs: diffContainer
                },
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.containerDetail.terminal', {
                url: '/terminal',
                templateUrl: '/src/node/container-detail/terminal.html',
                controller: 'NodeContainerTerminalCtrl as nodeContainerTerminalCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.networkDetail', {
                url: '/networkDetail/:node_id/:network_id',
                templateUrl: '/src/node/network-detail/detail.html',
                controller: 'NodeNetworkDetailCtrl as nodeNetworkDetailCtrl',
                targetState: 'config',
                resolve: {
                    network: getNetwork
                },
                ncyBreadcrumb: {
                    label: "{/'Network Detail' | translate/}",
                    parent: 'node.detail'
                }

            })
            .state('node.networkDetail.config', {
                url: '/config',
                templateUrl: '/src/node/network-detail/config.html',
                controller: 'NodeNetworkConfigCtrl as nodeNetworkConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            })
            .state('node.volumeDetail', {
                url: '/volumeDetail/:node_id/:volume_name',
                templateUrl: '/src/node/volume-detail/detail.html',
                controller: 'NodeVolumeDetailCtrl as nodeVolumeDetailCtrl',
                targetState: 'config',
                resolve: {
                    volume: getVolume
                },
                ncyBreadcrumb: {
                    label: "{/'Volume Detail' | translate/}",
                    parent: 'node.detail'
                }

            })
            .state('node.volumeDetail.config', {
                url: '/config',
                templateUrl: '/src/node/volume-detail/config.html',
                controller: 'NodeVolumeConfigCtrl as nodeVolumeConfigCtrl',
                ncyBreadcrumb: {
                    skip: true
                }
            });

        /* @ngInject */
        function listNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }

        /* @ngInject */
        function getNode(nodeCurd, $stateParams, nodeBackend, $q) {
            //TODO: which is a workaround. The desired solution is abstract all of the exception as a layer,
            //per error_code_XXX can trigger the special error_XXX_handle_func
            return nodeBackend.getNode($stateParams.node_id).catch(function (res) {
                var deferred = $q.defer();
                if (res.data && angular.isObject(res.data) && res.code && NODE_CONN_ERROR_CODE.indexOf(res.code) != -1) {
                    nodeCurd.updateEndpoint(res.data.ID, res.data.Endpoint)
                }
                deferred.reject();
                return deferred.promise;
            });
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
        function getContainer(nodeBackend, $stateParams, $q, nodeCurd) {
            return  nodeBackend.getContainer($stateParams.node_id, $stateParams.container_id).catch(function (res) {
                var deferred = $q.defer();
                if (res.data && angular.isObject(res.data) && res.code && NODE_CONN_ERROR_CODE.indexOf(res.code) != -1) {
                    nodeCurd.updateEndpoint(res.data.ID, res.data.Endpoint)
                }
                deferred.reject();
                return deferred.promise;
            });
        }

        /* @ngInject */
        function getVolume(nodeBackend, $stateParams) {
            return nodeBackend.getVolume($stateParams.node_id, $stateParams.volume_name)
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
