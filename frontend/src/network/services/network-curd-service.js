/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.network')
        .factory('networkCurd', networkCurd);


    /* @ngInject */
    function networkCurd(networkBackend, $state, confirmModal, Notification, $filter) {
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
                    $state.go('network.list', undefined, {reload: true})
                })
        }

        function deleteNetwork(id) {
            confirmModal.open("Are you sure to delete the network ?").then(function () {
                networkBackend.deleteNetwork(id)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
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