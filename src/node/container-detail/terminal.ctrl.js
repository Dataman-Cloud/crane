(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerTerminalCtrl', NodeContainerTerminalCtrl);

    /* @ngInject */
    function NodeContainerTerminalCtrl($scope, $stateParams, tty) {
        var self = this;
        
        activate();
        
        function activate() {
            tty.TTY('node.containerTerminal', {node_id:$stateParams.node_id, container_id:$stateParams.container_id});
        }

    }
})();
