(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerCtrl', NodeContainerCtrl);

    /* @ngInject */
    function NodeContainerCtrl(containers, nodeCurd, $stateParams) {
        var self = this;

        self.containers = containers;

        self.removeContainer = removeContainer;
        self.killContainer = killContainer;

        function removeContainer(containerId){
            nodeCurd.removeContainer($stateParams.node_id, containerId)
        }

        function killContainer(containerId){
            nodeCurd.killContainer($stateParams.node_id, containerId)
        }
    }
})();
