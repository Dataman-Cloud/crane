(function () {
    'use strict';
    angular.module('app.network')
        .controller('NetworkDetailCtrl', NetworkDetailCtrl);

    /* @ngInject */
    function NetworkDetailCtrl(network) {
        var self = this;

        self.network = network

    }
})();