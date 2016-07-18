(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerStatsCtrl', NodeContainerStatsCtrl);

    /* @ngInject */
    function NodeContainerStatsCtrl(stream, $stateParams, $scope) {
        var self = this;
        
        self.stats = []
        activate();
        
        function activate() {
            listenStats();
        }
        
        function listenStats() {
            stream = stream.Stream('node.containerStats', {node_id:$stateParams.node_id, container_id:$stateParams.container_id});
            stream.addHandler('container-stats', function (event) {
                self.stats.push(event.data);
                $scope.$apply();
            });
            stream.start();
            
            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();
