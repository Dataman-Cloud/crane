(function () {
    'use strict';
    angular.module('app.node')
        .factory('containerChart', containerChart);


    /* @ngInject */
    function containerChart($filter, $rootScope, chartUtil) {
        
        var optionsCls = createOptionsCls();
        
        return {
           Options: createOptions
        };
        
        function createOptions() {
            return new optionsCls();
        }
        
        function createOptionsCls() {
            function Options() {
                this.memOptions = this._createMemOptions();
                this.cpuOptions = this._createCpuOptions();
                this.networkOptions = this._createNetworkOptions();
                this.memData = [{
                    values: [],
                    key: '内存使用率',
                    area: true
                }];
                this.cpuData = [{
                    values: [],
                    key: 'CPU使用率',
                    area: true
                }];
                this.networkData = [{
                    values: [],
                    key: '网络接收速率',
                    area: true
                },
                {
                    values: [],
                    key: '网络发送速率',
                    area: true
                }];
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
                options.chart.useInteractiveGuideline = false;
                options.chart.tooltip = {
                    contentGenerator: function (e) {
                        var point = e.point;
                        var header = 
                            "<thead>" + 
                            "<tr>" +
                              "<td class='key'><strong>"+$filter('date')(point.x, 'HH:mm:ss')+"</strong></td>" +
                              "<td>&nbsp;</td>" +
                            "</tr>" + 
                          "</thead>";
                        var rows = 
                                "<tr>" +
                                "<td class='key'>" + '内存使用率: ' + "</td>" +
                                "<td class='x-value'>" + d3.format('.02f')(point.y)+'%' + "</td>" +
                                "</tr>";
                        if (point.total) {
                            rows += 
                                "<tr>" +
                                "<td class='key'>" + '内存使用量: ' + "</td>" +
                                "<td class='x-value'>" + $filter('size')(point.use) + "</td>" + 
                                "</tr>" +
                                "<tr>" +
                                "<td class='key'>" + '内存总量: ' + "</td>" +
                                "<td class='x-value'>" + $filter('size')(point.total) + "</td>" +
                                "</tr>";
                        }

                        return "<table>" +
                            header +
                            "<tbody>" + 
                              rows + 
                            "</tbody>" +
                          "</table>";
                    }
                }
                options.title.text = '内存监控'; 
                return options;
            };
            
            Options.prototype.pushData = function(data, cpuApi, memApi, networkApi) {
                data = angular.fromJson(data).Stat;
                var x = new Date(data.read).getTime()
                this._pushCpuData(x, data, cpuApi);
                this._pushMemData(x, data, memApi);
                this._pushNetworkData(x, data, networkApi);
            };
            
            Options.prototype._pushCpuData = function(x, data, cpuApi) {
                chartUtil.pushData(this.cpuData[0], {x:x, y:this._getCpuUsageRate(data)}, $rootScope.CONTAINER_STATS_POINT_NUM);
                if (chartUtil.updateForceY(this.cpuOptions.chart, [this.cpuData[0].values], 0, 1.2, 1, 100)){
                    cpuApi.refresh();
                } else {
                    cpuApi.update();
                }
            }
            
            Options.prototype._pushMemData = function(x, data, memApi) {
                chartUtil.pushData(this.memData[0], {x:x, y:data.memory_stats.usage/data.memory_stats.limit * 100, 
                    total: data.memory_stats.limit, use:data.memory_stats.usage}, $rootScope.CONTAINER_STATS_POINT_NUM);
                if (chartUtil.updateForceY(this.memOptions.chart, [this.memData[0].values], 0, 1.2, 1, 100)){
                    memApi.refresh();
                } else {
                    memApi.update();
                }
            }
            
            Options.prototype._pushNetworkData = function(x, data, networkApi) {
                chartUtil.pushData(this.networkData[0], {x:x, y:Math.sum(data.networks, function (network) {return network.rx_bytes})},
                        $rootScope.CONTAINER_STATS_POINT_NUM);
                chartUtil.pushData(this.networkData[1], {x:x, y:Math.sum(data.networks, function (network) {return network.tx_bytes})},
                        $rootScope.CONTAINER_STATS_POINT_NUM);
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