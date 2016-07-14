(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackServiceCtrl', StackServiceCtrl);

    /* @ngInject */
    function StackServiceCtrl(services, stackCurd) {
        var self = this;

        self.services = services;
        self.upServiceScale = stackCurd.upServiceScale;
    }
})();
