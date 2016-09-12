(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceStatsCtrl', ServiceStatsCtrl);

    /* @ngInject */
    function ServiceStatsCtrl(stream, $stateParams, $scope, statsChart, $interval) {
        var self = this;
        var stopTime;
        
        self.stats = [];
        self.chartOptions = statsChart.Options('TaskName');
        activate();
        
        function activate() {
            listenStats();
        }
        
        function listenStats() {
            stream = stream.Stream('stack.serviceStats', {stack_name:$stateParams.stack_name, service_id:$stateParams.service_id});
            stream.addHandler('service-stats', function (event) {
                self.chartOptions.pushData(event.data, self.cpuChartApi, self.memChartApi, self.networkChartApi);
            });
            stream.start();

            stopTime = $interval(flushCharts, 5000);
            
            $scope.$on('$destroy', function () {
                stream.stop();
                $interval.cancel(stopTime);

            });
        }

        function flushCharts() {
            self.chartOptions.flushCharts(self.cpuChartApi, self.memChartApi, self.networkChartApi);
        }
    }
})();
