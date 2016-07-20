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
                this.memData = [{
                    values: [],
                    key: '内存使用率',
                    color: '#2ca02c',
                    area: true
                }];
            }
            
            Options.prototype._createMemOptions = function() {
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
                        useInteractiveGuideline: false,
                        xAxis: {
                            axisLabel: '时间',
                            tickFormat: function(d){
                                return $filter('date')(d, 'HH:mm:ss');
                            },
                            showMaxMin: false
                        },
                        yAxis: {
                            axisLabel: '内存使用率',
                            tickFormat: function(d){
                                return d3.format('.02f')(d)+'%';
                            },
                            axisLabelDistance: -10
                        },
                        pointSize: 0.1,
                        forceY: [0],
                        tooltip: {
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
                    },
                    title: {
                        enable: true,
                        text: '内存监控'
                    }
                }
            }

            Options.prototype.pushData = function(data) {
                data = angular.fromJson(data);
                var x = new Date(data.read).getTime()
                this._pushData(this.memData[0], {x:x, y:data.memory_stats.usage/data.memory_stats.limit * 100, 
                    total: data.memory_stats.limit, use:data.memory_stats.usage});
            }
            
            Options.prototype._pushData = function (dataContainer, value) {
                dataContainer.values.push(value);
                while (dataContainer.values.length !== $rootScope.CONTAINER_STATS_POINT_NUM) {
                    if (dataContainer.values.length > $rootScope.CONTAINER_STATS_POINT_NUM) {
                        dataContainer.values.shift();
                    } else {
                        dataContainer.values.unshift({x: dataContainer.values[0].x-1000, y: 0});
                    }
                }
            }
            
            return Options;
        }
    }
})();