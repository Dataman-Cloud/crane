(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackServiceCtrl', StackServiceCtrl);

    /* @ngInject */
    function StackServiceCtrl(services) {
        var self = this;

        self.services = services;
    }
})();
