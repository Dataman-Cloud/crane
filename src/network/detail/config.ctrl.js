/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .controller('NetworkConfigCtrl', NetworkConfigCtrl);

    /* @ngInject */
    function NetworkConfigCtrl(network) {
        var self = this;

        self.network = network;
    }
})();
