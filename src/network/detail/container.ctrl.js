(function () {
    'use strict';
    angular.module('app.network')
        .controller('NetworkContainerCtrl', NetworkContainerCtrl);

    /* @ngInject */
    function NetworkContainerCtrl(network) {
        var self = this;

        self.network = network

    }
})();
