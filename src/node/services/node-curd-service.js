/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeCurd', nodeCurd);


    /* @ngInject */
    function nodeCurd(nodeBackend, $state, confirmModal, Notification) {
        //////
        return {
            deleteVolume: deleteVolume,
            removeContainer: removeContainer,
            killContainer: killContainer
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

    }
})();