(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeNetworkCtrl', NodeNetworkCtrl);

    /* @ngInject */
    function NodeNetworkCtrl(networks, networkCurd) {
        var self = this;

        self.networks = networks;
        self.deleteNetwork = deleteNetwork;

        function deleteNetwork(id) {
            networkCurd.deleteNetwork(id)
        }
    }
})();
