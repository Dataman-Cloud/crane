(function () {
    'use strict';
    angular.module('app.network')
        .controller('NetworkListCtrl', NetworkListCtrl);


    /* @ngInject */
    function NetworkListCtrl(networks, networkCurd) {
        var self = this;
        self.networks = networks;

        self.deleteNetwork = deleteNetwork;
        self.connectNetwork = connectNetwork;
        self.disconnectNetwork = disconnectNetwork;

        activate();

        function activate() {
            ///
        }

        function deleteNetwork(id) {
            networkCurd.deleteNetwork(id)
        }

        function connectNetwork(id) {

        }

        function disconnectNetwork(id) {

        }
    }
})();
