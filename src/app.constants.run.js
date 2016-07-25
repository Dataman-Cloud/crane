(function () {
    'use strict';
    angular.module('app')
        .run(run);

    /*@ngInject*/
    function run($rootScope) {
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
            running: '运行中',
            paused: '暂停中',
            dead: '已崩溃',
            restarting: '正在重启中',
            created: '初始化中',
            exited: '已退出'
        };
        
        $rootScope.CONTAINER_STATS_POINT_NUM = 180;
        
        $rootScope.DIFF_KIND = {
            0: '修改',
            1: '添加',
            2: '删除'
        };

        $rootScope.NODE_ROLE = {
            worker: '工作节点',
            manager: '管理节点'
        };

        $rootScope.NODE_AVAILABILITY = {
            drain: '停止调度',
            active: '正常调度',
            pause: '暂停调度'
        };

        $rootScope.NODE_STATE = {
            unknown: '未知',
            down: '下线',
            ready: '就绪',
            disconnected: '失联'
        };

        $rootScope.TASK_STATE = {
            new: '初始化',
            allocated: '资源已确认',
            pending: '排队中',
            assigned: '任务已派发',
            accepted: '已接受',
            preparing: '准备中',
            ready: '准备就绪',
            starting: '启动中',
            running: '运行中',
            complete: '已完成',
            shutdown: '已关闭',
            failed: '失败',
            rejected: '拒绝'
        };

        $rootScope.BACKEND_URL = {
            node: {
                nodes: 'api/v1/nodes',
                nodeInfo: 'api/v1/nodes/$node_id/info',
                node: 'api/v1/nodes/$node_id',
                leader: 'api/v1/nodes/leader_manager',
                volumes: 'api/v1/nodes/$node_id/volumes',
                volume: 'api/v1/nodes/$node_id/volumes/$volume_id',
                images: 'api/v1/nodes/$node_id/images',
                image: 'api/v1/nodes/$node_id/images/$image_id',
                imageHistory: 'api/v1/nodes/$node_id/images/$image_id/history',
                containers: 'api/v1/nodes/$node_id/containers',
                container: 'api/v1/nodes/$node_id/containers/$container_id',
                containerDiff: 'api/v1/nodes/$node_id/containers/$container_id/diff',
                containerLog: 'api/v1/nodes/$node_id/containers/$container_id/logs',
                containerStats: 'api/v1/nodes/$node_id/containers/$container_id/stats',
                containerTerminal: 'api/v1/nodes/$node_id/containers/$container_id/terminal',
                networks: 'api/v1/nodes/$node_id/networks',
                network: 'api/v1/nodes/$node_id/networks/$network_id'
            },
            stack: {
                stacks: 'api/v1/stacks',
                stack: 'api/v1/stacks/$stack_name',
                services: 'api/v1/stacks/$stack_name/services',
                service: 'api/v1/stacks/$stack_name/services/$service_id',
                tasks: 'api/v1/stacks/$stack_name/services/$service_id/tasks'
            },
            network: {
                network: 'api/v1/networks/$network_id',
                container: 'api/v1/networks/$network_id/container',
                networks: 'api/v1/networks'
            },
            misc: {
                config: 'misc/v1/config'
            }
        };
    }
})();
