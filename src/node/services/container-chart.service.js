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
                    color: '#2ca02c'
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
                        useInteractiveGuideline: true,
                        xAxis: {
                            axisLabel: '时间',
                            tickFormat: function(d){
                                return $filter('date')(d, 'HH:mm:ss');
                            },
                            showMaxMin: false
                        },
                        yAxis: {
                            axisLabel: '内存使用量',
                            tickFormat: function(d){
                                return d3.format('.02f')(d)+'%';
                            },
                            axisLabelDistance: -10
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
                this._pushData(this.memData[0], x, data.memory_stats.usage/data.memory_stats.limit * 100);
            }
            
            Options.prototype._pushData = function (dataContainer, x, y) {
                dataContainer.values.push({x: x, y: y});
                var curX = x-1000;
                while (dataContainer.values.length !== $rootScope.CONTAINER_STATS_POINT_NUM) {
                    if (dataContainer.values.length > $rootScope.CONTAINER_STATS_POINT_NUM) {
                        dataContainer.values.shift();
                    } else {
                        dataContainer.values.unshift({x: curX, y: 0});
                        curX -= 1000;
                    }
                }
            }
            
            return Options;
        }
    }
})();