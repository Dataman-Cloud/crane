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
            10000: '服务忙, 请稍后重试',
            10001: '服务忙, 请稍后重试',
            11001: '获取容器列表参数错误',
            11002: '更新容器参数错误',
            11003: '不支持的容器更新操作',
            11004: '不支持的容器删除操作',
            11005: '不支持的删除容器操作',
            11006: '未找到指定容器',
            11007: '指定容器已经处于运行中状态',
            11008: '指定容器需要处于运行中状态',
            11101: '获取镜像列表参数错误',
            11201: '连接到指定网络参数错误',
            11202: '不支持的网络连接操作',
            11203: '创建网络参数错误',
            11204: '查看指定网络详情参数错误',
            11205: '获取网络列表参数错误',
            11206: '不能操作系统内置网络',
            11207: '未找到指定网络',
            11208: '未找到指定容器或网络',
            11301: '更新节点参数错误',
            11302: '不支持的节点操作类型',
            11303: '未定义的节点角色',
            11304: '不支持的节点状态',
            11401: '更新服务参数错误',
            11402: '创建服务参数错误',
            11403: '扩展服务参数错误',
            11404: '获取指定服务任务列表参数错误',
            11501: '创建应用参数错误',
            11601: '创建存储参数错误',
            11701: '连接到指定节点失败',
            11901: '获取配置失败',
            12001: '未定义的用户组',
            12002: '创建用户参数错误',
            12003: '添加用户权限失败',
            12004: '获取指定用户失败',
            12005: '未找到指定用户',
            12006: '登陆参数错误',
            12007: '登陆失败',
            12008: '登出失败',
            12009: '用户组错误',
            12010: '指定用户组内未找到用户',
            12011: '用户ID错误',
            12012: '指定用户组未找到',
            12013: '未定义的用户组ID',
            12014: '未找到指定的用户组ID',
            12015: '获取用户组列表失败',
            12016: '用户没有修改权限',
            12017: '创建用户组参数错误',
            12018: '创建用户组失败',
            12019: '更新用户组参数错误',
            12020: '更新用户组失败',
            12021: '删除用户组失败, 用户组ID错误',
            12022: '删除用户组失败',
            12023: '加入用户组失败, 用户组ID错误',
            12024: '加入用户组失败, 用户ID错误',
            12025: '加入用户组失败',
            12026: '离开用户组失败, 用户组ID错误',
            12027: '离开用户组失败, 用户ID错误',
            12028: '离开用户组失败',
            12029: '添加服务权限失败, 参数错误',
            12030: '添加服务权限失败',
            12031: '移除服务权限失败, 无效的查询词',
            12032: '移除服务权限失败',
            13001: '查询关键词错误',
            14001: '获取镜像仓库配置失败',
            14002: '获取镜像仓库配置失败, 参数解析错误',
            14003: '删除镜像仓库配置失败',
            14004: '镜像公开失败, 参数错误',
            14005: '镜像公开失败, 更新错误',
            14006: '获取镜像仓库应用目录列表失败',
            15001: '获取应用目录失败',
            15002: '获取应用目录列表失败',
            16001: '获取License失败',
            16002: '创建License失败'
        };

        $rootScope.CONTAINER_STATUS_LABELS = {
            running: '运行中',
            paused: '暂停中',
            dead: '已崩溃',
            restarting: '正在重启中',
            created: '初始化中',
            exited: '已退出'
        };

        $rootScope.STATS_POINT_NUM = 180;

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
            auth: {
                login: 'account/v1/login',
                logout: 'account/v1/logout',
                aboutme: 'account/v1/aboutme',
                groups: 'account/v1/accounts/$account_id/groups'
            },
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
                serviceCdUrl: 'api/v1/stacks/$stack_name/services/$service_id/cd_url',
                tasks: 'api/v1/stacks/$stack_name/services/$service_id/tasks',
                serviceLog: 'api/v1/stacks/$stack_name/services/$service_id/logs',
                serviceStats: 'api/v1/stacks/$stack_name/services/$service_id/stats'
            },
            network: {
                network: 'api/v1/networks/$network_id',
                container: 'api/v1/networks/$network_id/container',
                networks: 'api/v1/networks'
            },
            misc: {
                config: 'misc/v1/config'
            },
            registry: {
                publicRepositories: 'registry/v1/repositories/public',
                mineRepositories: 'registry/v1/repositories/mine',
                listTags: 'registry/v1/tag/list/$repository',
                image: 'registry/v1/manifests/$tag/$repository',
                publicity: 'registry/v1/$namespace/$image/publicity',
                catalogs: 'catalog/v1/catalogs',
                catalog: 'catalog/v1/catalogs/$catalog_name'
            },
            layout: {
                search: 'search/v1/luckysearch'
            }
        };
    }
})();
