(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceStatsCtrl', ServiceStatsCtrl);

    /* @ngInject */
    function ServiceStatsCtrl(stream, $stateParams, $scope, statsChart) {
        var self = this;
        
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
            
            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();
