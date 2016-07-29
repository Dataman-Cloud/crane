(function () {
    'use strict';
    angular.module('app.utils')
        .factory('statsChart', statsChart);


    /* @ngInject */
    function statsChart($filter, $rootScope, chartUtil) {
        
        var optionsCls = createOptionsCls();
        
        return {
           Options: createOptions
        };
        
        function createOptions(serialKey) {
            return new optionsCls(serialKey);
        }
        
        function createOptionsCls() {
            function Options(serialKey) {
                this.serialKey = serialKey;
                this.memOptions = this._createMemOptions();
                this.cpuOptions = this._createCpuOptions();
                this.networkOptions = this._createNetworkOptions();
                this.memData = [];
                this.cpuData = [];
                this.networkData = [];
            }
            
            Options.prototype._createNetworkOptions = function() {
                var options = chartUtil.createDefaultOptions();
                options.chart.yAxis.axisLabel = '速率';
                options.title.text = '网络监控';
                options.chart.yAxis.tickFormat = function(d){
                        return d+'b/s';
                    };
                return options
            }
            
            Options.prototype._createCpuOptions = function() {
                var options = chartUtil.createDefaultOptions();
                options.chart.yAxis.axisLabel = 'CPU使用率';
                options.title.text = 'CPU监控';
                return options
            };
            
            Options.prototype._createMemOptions = function() {
                var options = chartUtil.createDefaultOptions();
                options.chart.yAxis.axisLabel = '内存使用率';
                options.title.text = '内存监控'; 
                return options;
            };
            
            Options.prototype.pushData = function(data, cpuApi, memApi, networkApi) {
                var stats = angular.fromJson(data);
                var data = stats.Stat;
                var serialName = "";
                if (this.serialKey) {
                    serialName = stats[this.serialKey];
                }
                var x = new Date(data.read).getTime()
                this._pushCpuData(serialName, x, data, cpuApi);
                this._pushMemData(serialName, x, data, memApi);
                this._pushNetworkData(serialName, x, data, networkApi);
            };
            
            Options.prototype._pushCpuData = function(serialName, x, data, cpuApi) {
                var serialKey = serialName + ' CPU使用率';
                chartUtil.pushData(this.cpuData, serialKey, {x:x, y:this._getCpuUsageRate(data)}, $rootScope.STATS_POINT_NUM);
                if (chartUtil.updateForceY(this.cpuOptions.chart, this.cpuData, 0, 1.2, 1, 100)){
                    cpuApi.refresh();
                } else {
                    cpuApi.update();
                }
            }
            
            Options.prototype._pushMemData = function(serialName, x, data, memApi) {
                var serialKey = serialName + ' 内存使用率';
                chartUtil.pushData(this.memData, serialKey, {x:x, y:data.memory_stats.usage/data.memory_stats.limit * 100, 
                    total: data.memory_stats.limit, use:data.memory_stats.usage}, $rootScope.STATS_POINT_NUM);
                if (chartUtil.updateForceY(this.memOptions.chart, this.memData, 0, 1.2, 1, 100)){
                    memApi.refresh();
                } else {
                    memApi.update();
                }
            }
            
            Options.prototype._pushNetworkData = function(serialName, x, data, networkApi) {
                chartUtil.pushData(this.networkData, serialName+' 网络接收速率', {x:x, y:Math.sum(data.networks, function (network) {return network.rx_bytes})},
                        $rootScope.STATS_POINT_NUM);
                chartUtil.pushData(this.networkData, serialName+' 网络发送速率', {x:x, y:Math.sum(data.networks, function (network) {return network.tx_bytes})},
                        $rootScope.STATS_POINT_NUM);
                networkApi.update();
            }
            
            Options.prototype._getCpuUsageRate = function(data) {
                var cpuPercent=0;
                var cpuDelta = data.cpu_stats.cpu_usage.total_usage - data.precpu_stats.cpu_usage.total_usage;
                var systemDelta = data.cpu_stats.system_cpu_usage - data.precpu_stats.system_cpu_usage;

                if (systemDelta > 0 && cpuDelta > 0) {
                    var cpuPercent = (cpuDelta / systemDelta) * data.cpu_stats.cpu_usage.percpu_usage.length * 100;
                }
                return cpuPercent
            }
            
            return Options;
        }
    }
})();