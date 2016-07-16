(function () {
    'use strict';
    angular.module('app')
        .run(run);

    /*@ngInject*/
    function run($state, $rootScope) {
        $rootScope.MESSAGE_CODE = {
            success: 0,
            dataInvalid: 10001,
            noExist: 10009,
            unknow: 10000
        };

        $rootScope.CODE_MESSAGE = {
            10001: '参数错误',
            10002: '操作失败'
        };

        $rootScope.STACK_SAMPLES = {
            singleService: {
                "Services": {
                    "redis": {
                        "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34",
                        "network": ["bridge"]
                    }
                },
                "Version": "0.1"
            },
            doubleServices: {
                "Services": {
                    "redis": {
                        "Image": "redis@sha256:b50f15d427aea5b579f9bf972ab82ff8c1c47bffc0481b225c6a714095a9ec34",
                        "network": ["ingress", "bridge"]
                    },
                    "nginx": {
                        "Image": "nginx:stable-alpine",
                        "network": ["ingress"]
                    }
                },
                "Version": "0.1"
            }
        };

        $rootScope.CONTAINER_STATUS_LABELS = {
            running: '运行中'
        };

        $rootScope.NODE_ROLE = {
            worker: '工作节点',
            manager: '管理节点'
        };

        $rootScope.BACKEND_URL = {
            node: {
                nodes: 'api/v1/nodes',
                node: 'api/v1/nodes/$node_id/info',
                leader: 'api/v1/nodes/leader_manager',
                volumes: 'api/v1/nodes/$node_id/volumes',
                volume: 'api/v1/nodes/$node_id/volumes/$volume_id',
                images: 'api/v1/nodes/$node_id/images',
                image: 'api/v1/nodes/$node_id/images/$image_id',
                imageHistory: 'api/v1/nodes/$node_id/images/$image_id/history',
                containers: 'api/v1/nodes/$node_id/containers',
                container: 'api/v1/nodes/$node_id/containers/$container_id',
                containerDiff: 'api/v1/nodes/$node_id/containers/$container_id/diff'
            },
            stack: {
                stacks: 'api/v1/stacks',
                stack: 'api/v1/stacks/$stack_name',
                services: 'api/v1/stacks/$stack_name/services',
                service: 'api/v1/stacks/$stack_name/services/$service_id'
            },
            network: {
                network: 'api/v1/networks/$network_id',
                container: 'api/v1/networks/$network_id/container',
                networks: 'api/v1/networks'
            }
        };
    }
})();
