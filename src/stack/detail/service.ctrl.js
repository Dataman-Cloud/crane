(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackServiceCtrl', StackServiceCtrl);

    /* @ngInject */
    function StackServiceCtrl($state, services, stackCurd) {
        var self = this;

        self.services = services;
        self.upServiceScale = stackCurd.upServiceScale;
        self.stopService = stackCurd.stopService;
        self.updateService = updateService;

        function updateService(stack_name, service_id) {
            $state.go('stack.serviceUpdate', {'stack_name': stack_name, 'service_id': service_id});
        }
    }
})();
