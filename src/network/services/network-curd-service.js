/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.network')
        .factory('networkCurd', networkCurd);


    /* @ngInject */
    function networkCurd(networkBackend, $state, confirmModal, Notification) {
        //////
        return {
            create: create,
            deleteNetwork: deleteNetwork,
            connectNetwork: connectNetwork,
            disconnectNetwork: disconnectNetwork
        };

        function create(data, form) {
            networkBackend.create(data, form)
                .then(function (data) {
                    $state.go('network.list')
                })
        }

        function deleteNetwork(id) {
            confirmModal.open("是否确认删除该网络？").then(function () {
                networkBackend.deleteNetwork(id)
                    .then(function (data) {
                        Notification.success('删除成功');
                        $state.reload()
                    })
            });
        }

        function connectNetwork(data, id) {

        }

        function disconnectNetwork(id) {

        }

    }
})();