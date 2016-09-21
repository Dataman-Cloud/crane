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

            Options.prototype._createNetworkOptions = function () {
                var options = chartUtil.createDefaultOptions();
                options.chart.yAxis.axisLabel = '速率';
                options.title.text = $filter('translate')('Network Receive/Send Rate');
                options.chart.yAxis.tickFormat = function (d) {
                    return $filter('netRate')(d);
                };
                options.chart.margin.left = 80;
                return options
            };

            Options.prototype._createCpuOptions = function () {
                var options = chartUtil.createDefaultOptions();
                options.chart.forceY = [0, 1];
                options.chart.yAxis.axisLabel = $filter('translate')('CPU Usage');
                options.title.text = $filter('translate')('CPU Usage');
                return options
            };

            Options.prototype._createMemOptions = function () {
                var options = chartUtil.createDefaultOptions();
                options.chart.forceY = [0, 1];
                options.chart.yAxis.axisLabel = $filter('translate')('Memory Usage');
                options.title.text = $filter('translate')('Memory Usage');
                return options;
            };

            Options.prototype.initNoDataCharts = function () {
                chartUtil.pushData(this.cpuData, "", {
                    x: new Date().getTime(),
                    y: 0
                }, $rootScope.STATS_POINT_NUM);
                chartUtil.pushData(this.memData, "", {
                    x: new Date().getTime(),
                    y: 0
                }, $rootScope.STATS_POINT_NUM);
                chartUtil.pushData(this.networkData, "", {
                    x: new Date().getTime(),
                    y: 0
                }, $rootScope.STATS_POINT_NUM);
            };

            Options.prototype.pushData = function (data, cpuApi, memApi, networkApi) {
                var stats = angular.fromJson(data);
                var containerStat = stats.Stat;
                var serialName = "";
                if (this.serialKey) {
                    serialName = stats[this.serialKey];
                }
                var x = new Date(containerStat.read).getTime();
                this._pushCpuData(serialName, x, containerStat, cpuApi);
                this._pushMemData(serialName, x, containerStat, memApi);
                this._pushNetworkData(serialName, x, stats, networkApi);
            };

            Options.prototype._pushCpuData = function (serialName, x, data, cpuApi) {
                var serialKey = serialName;
                chartUtil.pushData(this.cpuData, serialKey, {
                    x: x,
                    y: this._getCpuUsageRate(data)
                }, $rootScope.STATS_POINT_NUM);
            };

            Options.prototype._pushMemData = function (serialName, x, data, memApi) {
                var serialKey = serialName;
                chartUtil.pushData(this.memData, serialKey, {
                    x: x, y: data.memory_stats.usage / data.memory_stats.limit * 100,
                    total: data.memory_stats.limit, use: data.memory_stats.usage
                }, $rootScope.STATS_POINT_NUM);
            };

            Options.prototype._pushNetworkData = function (serialName, x, data, networkApi) {
                chartUtil.pushData(this.networkData, serialName + '接收', {
                        x: x,
                        y: data.ReceiveRate
                    },
                    $rootScope.STATS_POINT_NUM);
                chartUtil.pushData(this.networkData, serialName + '发送', {
                        x: x,
                        y: data.SendRate
                    },
                    $rootScope.STATS_POINT_NUM);
            };

            Options.prototype._getCpuUsageRate = function (data) {
                var cpuPercent = 0;
                var cpuDelta = data.cpu_stats.cpu_usage.total_usage - data.precpu_stats.cpu_usage.total_usage;
                var systemDelta = data.cpu_stats.system_cpu_usage - data.precpu_stats.system_cpu_usage;

                if (systemDelta > 0 && cpuDelta > 0) {
                    cpuPercent = (cpuDelta / systemDelta) * data.cpu_stats.cpu_usage.percpu_usage.length * 100;
                }
                return cpuPercent
            };

            Options.prototype.flushCharts = function (cpuApi, memApi, networkApi) {
                if (chartUtil.updateForceY(this.cpuOptions.chart, this.cpuData, 0, 1.2, 1, 100)) {
                    cpuApi.refresh();
                } else {
                    cpuApi.update();
                }

                if (chartUtil.updateForceY(this.memOptions.chart, this.memData, 0, 1.2, 1, 100)) {
                    memApi.refresh();
                } else {
                    memApi.update();
                }

                networkApi.update();
            };

            return Options;
        }
    }
})();