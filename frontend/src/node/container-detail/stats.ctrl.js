(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerStatsCtrl', NodeContainerStatsCtrl);

    /* @ngInject */
    function NodeContainerStatsCtrl(stream, $stateParams, $scope, statsChart, $interval) {
        var self = this;
        var stopTime;
        
        self.stats = [];
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

            stopTime = $interval(flushCharts, 5000);
            
            $scope.$on('$destroy', function () {
                stream.stop();
                $interval.cancel(stopTime)
            });
        }

        function flushCharts() {
            self.chartOptions.flushCharts(self.cpuChartApi, self.memChartApi, self.networkChartApi);
        }
    }
})();
