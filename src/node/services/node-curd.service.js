/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeCurd', nodeCurd);


    /* @ngInject */
    function nodeCurd(nodeBackend, $state, confirmModal, Notification, utils) {
        //////
        return {
            deleteVolume: deleteVolume,
            removeContainer: removeContainer,
            killContainer: killContainer,
            getNodesMapping: getNodesMapping,
            drainNode: drainNode,
            activeNode: activeNode,
            pauseNode: pauseNode,
            createNetwork: createNetwork
        };

        function deleteVolume(id, name) {
            confirmModal.open("是否确认删除该储存卷？").then(function () {
                nodeBackend.deleteVolume(id, name)
                    .then(function (data) {
                        Notification.success('删除成功');
                        $state.reload()
                    })
            });
        }

        function removeContainer(nodeId, containerId) {
            confirmModal.open("是否确认移除该容器？").then(function () {
                nodeBackend.removeContainer(nodeId, containerId)
                    .then(function (data) {
                        Notification.success('移除成功');
                        $state.reload()
                    })
            });
        }

        function killContainer(nodeId, containerId) {
            confirmModal.open("是否确认杀死该容器？").then(function () {
                nodeBackend.killContainer(nodeId, containerId)
                    .then(function (data) {
                        Notification.success('杀死成功');
                        $state.reload()
                    })
            });
        }

        function getNodesMapping() {
            return nodeBackend.listNodes().then(function (nodes) {
                return utils.convert2Mapping(nodes);
            })
        }

        function drainNode(nodeId, data) {
            confirmModal.open("是否确认下线该主机？").then(function () {
                nodeBackend.handleNode(nodeId, data, 'drain')
                    .then(function (data) {
                        Notification.success('下线成功');
                        $state.reload()
                    })
            });
        }

        function activeNode(nodeId, data) {
            confirmModal.open("是否确认激活该主机？").then(function () {
                nodeBackend.handleNode(nodeId, data, 'active')
                    .then(function (data) {
                        Notification.success('激活成功');
                        $state.reload()
                    })
            });
        }

        function pauseNode(nodeId, data) {
            confirmModal.open("是否确认暂停该主机？").then(function () {
                nodeBackend.handleNode(nodeId, data, 'pause')
                    .then(function (data) {
                        Notification.success('暂停成功');
                        $state.reload()
                    })
            });
        }

        function createNetwork(data, nodeId, form) {
            nodeBackend.createNetwork(data, nodeId, form)
                .then(function (data) {
                    Notification.success('创建成功');
                    $state.go('node.networkDetail', {node_id: nodeId, network_id: data.Id}, {reload: true})
                })
        }

    }
})();