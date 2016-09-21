(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceTaskCtrl', ServiceTaskCtrl);

    /* @ngInject */
    function ServiceTaskCtrl(tasks, nodeCurd) {
        var self = this;

        self.tasks = tasks;
        self.nodesMapping = {};

        self.orderByRunning = orderByRunning;

        activate();

        function activate() {
            nodeCurd.getNodesMapping().then(function (mapping) {
                self.nodesMapping = mapping;
            });
        }

        function orderByRunning(task) {
            return task.Status.State !== 'running';
        }

    }
})();
