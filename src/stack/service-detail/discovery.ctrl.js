(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceDiscoveryCtrl', ServiceDiscoveryCtrl);

    /* @ngInject */
    function ServiceDiscoveryCtrl(service, nodes) {
        var self = this;
        self.service = service;
        self.nodes = nodes;
    }
})();
