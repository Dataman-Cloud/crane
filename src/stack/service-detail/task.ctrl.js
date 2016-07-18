(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceTaskCtrl', ServiceTaskCtrl);

    /* @ngInject */
    function ServiceTaskCtrl(tasks, nodeCurd) {
        var self = this;
        
        self.tasks = tasks;
        self.nodesMapping;
        
        activate();
        
        function activate() {
            nodeCurd.getNodesMapping().then(function (mapping) {
                self.nodesMapping = mapping;
            });
        }

    }
})();
