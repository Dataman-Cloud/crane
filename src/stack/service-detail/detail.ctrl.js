(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceDetailCtrl', ServiceDetailCtrl);

    /* @ngInject */
    function ServiceDetailCtrl(service) {
        var self = this;
        self.service = service;
    }
})();
