(function () {
    'use strict';
    angular.module('app.node')
        .factory('containerChart', containerChart);


    /* @ngInject */
    function containerChart($filter, $rootScope) {
        
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
                var options = this._createDefaultOptions();
                options.chart.yAxis.axisLabel = '速率';
                options.title.text = '网络监控';
                options.chart.yAxis.tickFormat = function(d){
                        return d;
                    };
                return options
            }
            
            Options.prototype._createCpuOptions = function() {
                var options = this._createDefaultOptions();
                options.chart.yAxis.axisLabel = 'CPU使用率';
                options.title.text = 'CPU监控';
                return options
            };
            
            Options.prototype._createMemOptions = function() {
                var options = this._createDefaultOptions();
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
            
            Options.prototype._createDefaultOptions = function() {
                return {
                    chart: {
                        type: 'lineChart',
                        noData: '暂无数据',
                        height: 450,
                        margin : {
                            top: 20,
                            right: 20,
                            bottom: 40,
                            left: 55
                        },
                        x: function(d){ return d.x; },
                        y: function(d){ return d.y; },
                        useInteractiveGuideline: true,
                        xAxis: {
                            axisLabel: '时间',
                            tickFormat: function(d){
                                return $filter('date')(d, 'HH:mm:ss');
                            },
                            showMaxMin: false
                        },
                        yAxis: {
                            tickFormat: function(d){
                                return d3.format('.02f')(d)+'%';
                            },
                            axisLabelDistance: -10
                        },
                        pointSize: 0.1,
                        forceY: [0],
                        color: [
                                  '#1f77b4',
                                  '#ff7f0e',
                                  '#2ca02c',
                                  '#d62728',
                                  '#9467bd',
                                  '#8c564b',
                                  '#e377c2',
                                  '#7f7f7f',
                                  '#bcbd22',
                                  '#17becf'
                                ],
                    },
                    title: {
                        enable: true
                    }
                }
            };

            Options.prototype.pushData = function(data) {
                data = angular.fromJson(data);
                var x = new Date(data.read).getTime()
                this._pushData(this.memData[0], {x:x, y:data.memory_stats.usage/data.memory_stats.limit * 100, 
                    total: data.memory_stats.limit, use:data.memory_stats.usage});
                this._pushData(this.cpuData[0], {x:x, y:this._getCpuUsageRate(data)});
                this._pushData(this.networkData[0], {x:x, y:Math.sum(data.networks, function (network) {return network.rx_bytes})})
                this._pushData(this.networkData[1], {x:x, y:Math.sum(data.networks, function (network) {return network.tx_bytes})})
            };
            
            Options.prototype._pushData = function (dataContainer, value) {
                dataContainer.values.push(value);
                while (dataContainer.values.length !== $rootScope.CONTAINER_STATS_POINT_NUM) {
                    if (dataContainer.values.length > $rootScope.CONTAINER_STATS_POINT_NUM) {
                        dataContainer.values.shift();
                    } else {
                        dataContainer.values.unshift({x: dataContainer.values[0].x-1000, y: 0});
                    }
                }
            };
            
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