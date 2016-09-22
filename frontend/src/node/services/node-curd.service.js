/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeCurd', nodeCurd);


    /* @ngInject */
    function nodeCurd(nodeBackend, $state, confirmModal, Notification, utils, updateLabelsFormModal, formModal, $filter) {
        //////
        return {
            deleteVolume: deleteVolume,
            deleteImage: deleteImage,
            removeContainer: removeContainer,
            killContainer: killContainer,
            getNodesMapping: getNodesMapping,
            drainNode: drainNode,
            activeNode: activeNode,
            deleteNode: deleteNode,
            pauseNode: pauseNode,
            createNetwork: createNetwork,
            updateEndpoint: updateEndpoint,
            updateLabels: updateLabels,
            removeLabels: removeLabels,
            addWorkerNode: addWorkerNode
        };

        function updateEndpoint(nodeId, env, endpoint) {
            formModal.open('/src/node/modals/form-nodeIp.html', env, {dataName: 'endpoint', initData: endpoint})
                .then(function (endpoint) {
                    nodeBackend.handleNode(nodeId, "endpoint-update", endpoint).then(function (data) {
                        Notification.success($filter('translate')('Host updated successfully'));
                        $state.reload()
                    });
                });
        }

        function updateLabels(nodeId, env, labels) {
            var labelList = [];
            angular.forEach(labels, function (value, key) {
                this.push({"key": key, "value": value})
            }, labelList);

            updateLabelsFormModal.open('/src/node/modals/form-labels.html', env, {
                dataName: 'labels',
                initData: labelList
            })
                .then(function (labelList) {
                    var newLabels = {};
                    angular.forEach(labelList, function (label) {
                        this[label.key] = label.value;
                    }, newLabels);
                    nodeBackend.handleNode(nodeId, "label-update", newLabels).then(function (data) {
                        Notification.success($filter('translate')('Host updated successfully'));
                        $state.reload()
                    });
                });
        }

        function removeLabels(nodeId, rmList) {
            nodeBackend.handleNode(nodeId, "label-rm", rmList).then(function (data) {
                Notification.success($filter('translate')('Host updated successfully'));
                $state.reload()
            });
        }

        function deleteVolume(id, name) {
            confirmModal.open("Are you sure to delete the storage volume ?").then(function () {
                nodeBackend.deleteVolume(id, name)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
                        $state.reload()
                    })
            });
        }

        function deleteImage(nodeId, imageId) {
            confirmModal.open("Are you sure to delete the image ?").then(function () {
                nodeBackend.deleteImage(nodeId, imageId)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
                        $state.reload()
                    })
            });
        }

        function removeContainer(nodeId, containerId) {
            confirmModal.open("Are you sure to remove the container ?").then(function () {
                nodeBackend.removeContainer(nodeId, containerId)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
                        $state.reload()
                    })
            });
        }

        function killContainer(nodeId, containerId) {
            confirmModal.open("Are you sure to kill the container ?").then(function () {
                nodeBackend.killContainer(nodeId, containerId)
                    .then(function (data) {
                        Notification.success($filter('translate')('Killed successfully'));
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
            confirmModal.open("Are you sure offline the host ?").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'drain')
                    .then(function (data) {
                        Notification.success($filter('translate')('Offline successfully'));
                        $state.reload()
                    })
            });
        }

        function activeNode(nodeId) {
            confirmModal.open("Are you sure to activate the host ?").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'active')
                    .then(function (data) {
                        Notification.success($filter('translate')('Activated successfully'));
                        $state.reload()
                    })
            });
        }

        function deleteNode(nodeId) {
            confirmModal.open("Are you sure to delete this host ?").then(function () {
                nodeBackend.deleteNode(nodeId)
                    .then(function (data) {
                        Notification.success($filter('translate')('Host successfully deleted'));
                        $state.reload()
                    })
            });
        }

        function pauseNode(nodeId) {
            confirmModal.open("Are you sure to pause the node?").then(function () {
                nodeBackend.handleNode(nodeId, "availability", 'pause')
                    .then(function (data) {
                        Notification.success($filter('translate')('Paused successfully'));
                        $state.reload()
                    })
            });
        }

        function createNetwork(data, nodeId, form) {
            nodeBackend.createNetwork(data, nodeId, form)
                .then(function (data) {
                    Notification.success($filter('translate')('Created successfully'));
                    $state.go('node.networkDetail', {node_id: nodeId, network_id: data.Id}, {reload: true})
                })
        }

        function addWorkerNode(env) {
            formModal.open('/src/node/modals/create-node.html', env, {dataName: 'endpoint'})
                .then(function (endpoint) {
                    Notification.success($filter('translate')('Add node pending'));
                    var data = {
                        Role: "worker",
                        Endpoint: endpoint
                    };

                    nodeBackend.addWorkerNode(data)
                        .then(function (data) {
                            Notification.success($filter('translate')('Created successfully'));
                            $state.reload()
                        })
                });
        }

    }
})();
