(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerStatsCtrl', NodeContainerStatsCtrl);

    /* @ngInject */
    function NodeContainerStatsCtrl(stream, $stateParams, $scope, statsChart) {
        var self = this;
        
        self.stats = []
        self.chartOptions = statsChart.Options();
        activate();
        
        function activate() {
            listenStats();
        }
        
        function listenStats() {
            stream = stream.Stream('node.containerStats', {node_id:$stateParams.node_id, container_id:$stateParams.container_id});
            stream.addHandler('container-stats', function (event) {
                self.chartOptions.pushData(event.data, self.cpuChartApi, self.memChartApi, self.networkChartApi);
            });
            stream.start();
            
            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();
