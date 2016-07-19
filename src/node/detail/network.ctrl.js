(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeNetworkCtrl', NodeNetworkCtrl);

    /* @ngInject */
    function NodeNetworkCtrl(networks) {
        var self = this;

        self.networks = networks;
    }
})();
