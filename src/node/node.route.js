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
                targetState: 'list',
                ncyBreadcrumb: {
                    label: '主机'
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
                    label: '主机列表'
                }
            })
            .state('node.create', {
                url: '/create',
                templateUrl: '/src/node/create/create.html',
                controller: 'NodeCreateCtrl as nodeCreateCtrl',
                ncyBreadcrumb: {
                    label: '添加主机'
                }
            })
            .state('node.createVolume', {
                url: '/createVolume/:node_id',
                templateUrl: '/src/node/create-volume/create.html',
                controller: 'NodeCreateVolumeCtrl as nodeCreateVolumeCtrl',
                ncyBreadcrumb: {
                    label: '创建储存卷'
                }
            })
            .state('node.createNetwork', {
                url: '/createNetwork/:node_id',
                templateUrl: '/src/node/create-network/create.html',
                controller: 'NodeCreateNetworkCtrl as nodeCreateNetworkCtrl',
                ncyBreadcrumb: {
                    label: '创建单机网络'
                }
            })
            .state('node.detail', {
                url: '/detail/:node_id',
                templateUrl: '/src/node/detail/detail.html',
                controller: 'NodeDetailCtrl as nodeDetailCtrl',
                targetState: 'container',
                resolve: {
                    node: getNode
                },
                ncyBreadcrumb: {
                    label: '主机详情'
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
                    label: '容器列表'
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
                    label: '网络列表'
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
                    label: '存储列表'
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
                    label: '镜像列表'
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
                    label: '镜像详情',
                    parent: 'node.detail'
                }
            })
            .state('node.imageDetail.config', {
                url: '/config',
                templateUrl: '/src/node/image-detail/config.html',
                controller: 'NodeImageConfigCtrl as nodeImageConfigCtrl',
                ncyBreadcrumb: {
                    label: '详情'
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
                    label: '层级'
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
                    label: '容器详情',
                    parent: 'node.detail'
                }

            })
            .state('node.containerDetail.config', {
                url: '/config',
                templateUrl: '/src/node/container-detail/config.html',
                controller: 'NodeContainerConfigCtrl as nodeContainerConfigCtrl',
                ncyBreadcrumb: {
                    label: '详情'
                }
            })
            .state('node.containerDetail.log', {
                url: '/log',
                templateUrl: '/src/node/container-detail/log.html',
                controller: 'NodeContainerLogCtrl as nodeContainerLogCtrl',
                ncyBreadcrumb: {
                    label: '日志'
                }
            })
            .state('node.containerDetail.stats', {
                url: '/stats',
                templateUrl: '/src/node/container-detail/stats.html',
                controller: 'NodeContainerStatsCtrl as nodeContainerStatsCtrl',
                ncyBreadcrumb: {
                    label: '监控'
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
                    label: '变更'
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
                    label: '网络详情',
                    parent: 'node.detail'
                }

            })
            .state('node.networkDetail.config', {
                url: '/config',
                templateUrl: '/src/node/network-detail/config.html',
                controller: 'NodeNetworkConfigCtrl as nodeNetworkConfigCtrl',
                ncyBreadcrumb: {
                    label: '详情'
                }
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
