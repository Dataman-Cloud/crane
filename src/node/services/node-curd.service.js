/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeCurd', nodeCurd);


    /* @ngInject */
    function nodeCurd(nodeBackend, $state, confirmModal, Notification, utils, updateLabelsFormModal, formModal) {
        //////
        return {
            deleteVolume: deleteVolume,
            deleteImage: deleteImage,
            removeContainer: removeContainer,
            killContainer: killContainer,
            getNodesMapping: getNodesMapping,
            drainNode: drainNode,
            activeNode: activeNode,
            pauseNode: pauseNode,
            createNetwork: createNetwork,
            updateEndpoint: updateEndpoint,
            updateLabels: updateLabels,
            removeLabels: removeLabels
        };
        
        function updateEndpoint(nodeId, env, endpoint) {
            formModal.open('/src/node/modals/form-nodeIp.html', env, {dataName: 'endpoint', initData: endpoint})
                .then(function (endpoint) {
                nodeBackend.handleNode(nodeId, "endpoint-update", endpoint).then(function (data) {
                    Notification.success('更新主机成功');
                    $state.reload()
                });
            });
        }

        function updateLabels(nodeId, env, labels) {
            var labelList = [];
            angular.forEach(labels, function (value, key) {
                this.push({"key": key, "value": value})
            }, labelList);

            updateLabelsFormModal.open('/src/node/modals/form-labels.html', env, {dataName: 'labels', initData: labelList})
                .then(function (labelList) {
                    var newLabels = {};
                    angular.forEach(labelList, function (label) {
                        this[label.key] = label.value;
                    }, newLabels);
                    nodeBackend.handleNode(nodeId, "label-update", newLabels).then(function (data) {
                        Notification.success('更新主机成功');
                        $state.reload()
                    });
                });
        }

        function removeLabels(nodeId, rmList) {
            nodeBackend.handleNode(nodeId, "label-rm", rmList).then(function (data) {
                Notification.success('更新主机成功');
                $state.reload()
            });
        }
        
        function deleteVolume(id, name) {
            confirmModal.open("是否确认删除该储存卷？").then(function () {
                nodeBackend.deleteVolume(id, name)
                    .then(function (data) {
                        Notification.success('删除成功');
                        $state.reload()
                    })
            });
        }

        function deleteImage(nodeId, imageId) {
            confirmModal.open("是否确认删除该镜像？").then(function () {
                nodeBackend.deleteImage(nodeId, imageId)
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

        function drainNode(nodeId) {
            confirmModal.open("是否确认下线该主机？").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'drain')
                    .then(function (data) {
                        Notification.success('下线成功');
                        $state.reload()
                    })
            });
        }

        function activeNode(nodeId) {
            confirmModal.open("是否确认激活该主机？").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'active')
                    .then(function (data) {
                        Notification.success('激活成功');
                        $state.reload()
                    })
            });
        }

        function pauseNode(nodeId) {
            confirmModal.open("是否确认暂停该主机？").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'pause')
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