(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackServiceCtrl', StackServiceCtrl);

    /* @ngInject */
    function StackServiceCtrl($state, services, stackCurd) {
        var self = this;

        self.services = services;
        self.upServiceScale = stackCurd.upServiceScale;
        self.updateService = updateService;

        function updateService(stackName, serviceId) {
            $state.go('stack.serviceUpdate', {'stack_name': stackName, 'service_id': serviceId});
        }
    }
})();
