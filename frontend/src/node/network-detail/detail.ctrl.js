/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeNetworkDetailCtrl', NodeNetworkDetailCtrl);

    /* @ngInject */
    function NodeNetworkDetailCtrl(network) {
        var self = this;

        self.network = network;
    }
})();
