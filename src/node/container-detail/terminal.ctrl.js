(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerTerminalCtrl', NodeContainerTerminalCtrl);

    /* @ngInject */
    function NodeContainerTerminalCtrl($scope, $stateParams, tty) {
        var self = this;
        
        activate();
        
        function activate() {
            var containerTTY = tty.TTY('node.containerTerminal', {node_id:$stateParams.node_id, container_id:$stateParams.container_id});
            $scope.$on('$destroy', function () {
                containerTTY.clear();
            });
        }

    }
})();
