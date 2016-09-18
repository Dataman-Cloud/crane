(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackListCtrl', StackListCtrl);


    /* @ngInject */
    function StackListCtrl(stacks, stackCurd) {
        var self = this;
        
        self.stacks = stacks;

        self.openCreateSelect = openCreateSelect;

        activate();

        function activate() {
            formatStackServices(self.stacks);
        }

        function openCreateSelect(ev) {
            stackCurd.createSelect(ev)
        }

        function formatStackServices(stacks) {
            angular.forEach(stacks, function(stack) {
                angular.forEach(stack.Services, function(service) {
                    service.addrs = [];
                    angular.forEach(service.IPs, function(ip) {
                        angular.forEach(service.Ports, function(port) {
                            service.addrs.push({ip: ip, port: port});
                        });
                    });
                });
            });
        }
    }
})();
