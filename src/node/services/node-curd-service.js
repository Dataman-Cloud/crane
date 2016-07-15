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
            deleteVolume: deleteVolume
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

    }
})();