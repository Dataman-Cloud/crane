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
                controller: 'NodeContainerConfigCtrl as nodeContainerConfigCtrl',
                resolve: {
                    diffs: diffContainer
                }
            })
            .state('node.containerDetail.log', {
                url: '/log',
                templateUrl: '/src/node/container-detail/log.html',
                controller: 'NodeContainerLogCtrl as nodeContainerLogCtrl'
            });

        /* @ngInject */
        function listNodes(nodeBackend) {
            return nodeBackend.listNodes()
        }

        /* @ngInject */
        function getNode(nodeBackend, $stateParams) {
            return nodeBackend.getNode($stateParams.node_id);
            //return {"ID":"LWG7:Y66S:IJKU:I4VX:WFUO:EP5U:OQ6Y:ISNG:GYOM:5DBL:QNZG:Y3TA","Containers":1,"ContainersRunning":1,"ContainersPaused":0,"ContainersStopped":0,"Images":1,"Driver":"devicemapper","DriverStatus":[["Pool Name","docker-253:0-68225258-pool"],["Pool Blocksize","65.54 kB"],["Base Device Size","10.74 GB"],["Backing Filesystem","xfs"],["Data file","/dev/loop0"],["Metadata file","/dev/loop1"],["Data Space Used","270.5 MB"],["Data Space Total","107.4 GB"],["Data Space Available","27.27 GB"],["Metadata Space Used","1.044 MB"],["Metadata Space Total","2.147 GB"],["Metadata Space Available","2.146 GB"],["Thin Pool Minimum Free Space","10.74 GB"],["Udev Sync Supported","true"],["Deferred Removal Enabled","false"],["Deferred Deletion Enabled","false"],["Deferred Deleted Device Count","0"],["Data loop file","/var/lib/docker/devicemapper/devicemapper/data"],["Metadata loop file","/var/lib/docker/devicemapper/devicemapper/metadata"],["Library Version","1.02.107-RHEL7 (2015-10-14)"]],"SystemStatus":null,"Plugins":{"Volume":["local"],"Network":["host","bridge","null","overlay"],"Authorization":null},"MemoryLimit":true,"SwapLimit":true,"KernelMemory":true,"CpuCfsPeriod":true,"CpuCfsQuota":true,"CPUShares":true,"CPUSet":true,"IPv4Forwarding":true,"BridgeNfIptables":true,"BridgeNfIp6tables":true,"Debug":false,"NFd":47,"OomKillDisable":true,"NGoroutines":155,"SystemTime":"2016-07-11T12:05:56.446607594-04:00","ExecutionDriver":"","LoggingDriver":"json-file","CgroupDriver":"cgroupfs","NEventsListener":0,"KernelVersion":"3.10.0-327.el7.x86_64","OperatingSystem":"CentOS Linux 7 (Core)","OSType":"linux","Architecture":"x86_64","IndexServerAddress":"https://index.docker.io/v1/","NCPU":1,"MemTotal":2682290176,"DockerRootDir":"/var/lib/docker","HttpProxy":"","HttpsProxy":"","NoProxy":"","Name":"localhost","Labels":null,"ExperimentalBuild":false,"ServerVersion":"1.12.0-rc3","ClusterStore":"","ClusterAdvertise":""}
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
    }
})();
