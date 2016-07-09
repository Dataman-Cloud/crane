(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('lineChart', lineChart);

    /* @ngInject */
    function lineChart($filter) {
        return {
            buildChartData: buildChartData,
            paint: paint
        };

        function buildChartData(data, timeGetter, dataGetter, options) {
            if (!options) {
                options = {}
            }
            if (!options.long) {
                options.long = 60;
            }
            if (!options.interval) {
                options.interval = 1;
            }
            if (!options.numFixed && options.numFixed !== 0) {
                options.numFixed = 2;
            }
            var chartData = {y: {}, x: []};
            var curTime;
            if (!data) {
                curTime = Math.round((new Date()).getTime()/1000/60);
                data = [];
            } else {
                curTime = timeGetter(data[0]);
            }
            var startTime = curTime - options.long;
            pushXData(chartData, curTime)
            for (var i = 0; i < data.length; i++) {
                var curData = data[i];
                var dataTime = timeGetter(curData);
                if (curTime-options.interval/2 <= dataTime && dataTime < curTime+options.interval/2) {
                    pushYData(chartData, dataGetter(curData));
                } else {
                    calYData(chartData, options.numFixed);
                    curTime -= options.interval;
                    if (curTime <= startTime) {
                        break;
                    } else {
                        pushXData(chartData, curTime);
                    }
                }
            }
            calYData(chartData, options.numFixed);
            supplyData(chartData, curTime, startTime, options.interval);
            reverseChartData(chartData);
            return chartData;
        }
        
        function pushYData(chartData, yData) {
            angular.forEach(yData, function (value, key){
                if (!chartData.y[key]){
                    chartData.y[key] = []
                    for (var i=0; i<chartData.x.length-1; i++) {
                        chartData.y[key].push(undefined);
                    }
                }
                if (!chartData.y[key][chartData.x.length-1]) {
                    chartData.y[key][chartData.x.length-1] = []
                }
                if (value!=undefined){
                    chartData.y[key][chartData.x.length-1].push(value);
                }
            });
        }
        
        function calYData(chartData, numFixed) {
            angular.forEach(chartData.y, function(values, key){
                var curValues = values[chartData.x.length-1];
                if (angular.isArray(curValues)){
                    if (curValues.length >0) {
                        values[chartData.x.length-1] = Number(Math.average(curValues).toFixed(numFixed));
                    } else {
                        values[chartData.x.length-1] = undefined;
                    }
                }
            });
        }
        
        function pushXData(chartData, curTime){
            chartData.x.push(calHourMin(curTime));
            angular.forEach(chartData.y, function (values, key){
                values.push(undefined);
            })
        }
        
        function calHourMin(time) {
            var milliSeconds = time * 60 * 1000;
            var d = new Date();
            d.setTime(milliSeconds);
            return $filter('date')(d, 'HH:mm');
        };
        
        function reverseChartData(chartData) {
            chartData.x.reverse();
            angular.forEach(chartData.y, function(values){
                values.reverse();
            })
        }
        
        function supplyData(chartData, curTime, startTime, interval) {
            curTime -= interval;
            while(curTime > startTime) {
                pushXData(chartData, curTime)
                curTime -= interval;
            }
        }
        
        function paint(chartData, domId, description, seriesNameGetter, max, style) {
            if (!style) {
                style = {
                    lineWidth: 3,
                    axesColor: '#9B9B9B',
                    axiesFontsize: '11px',
                }
            }
            if (!seriesNameGetter) {
                seriesNameGetter = function (key) {return key;};
            }
            if (!max) {
                max = getYMax(chartData.y, seriesNameGetter);
            }
            var option = buildChartOption(chartData, description, seriesNameGetter, max, style)
            var chart = echarts.init(document.getElementById(domId));
            chart.setOption(option);
        }
        
        function getYMax(y, seriesNameGetter) {
            var max = 1;
            angular.forEach(y, function (item, key) {
                if (seriesNameGetter(key)){
                    angular.forEach(item, function (value) {
                        if (value > max) {
                            max = value;
                        }
                    })
                }
            });
            return Math.ceil(max);
        }
        
        function buildChartOption(chartData, description, seriesNameGetter, max, style) {
            var unit = description.unit||''
            var option = {
                grid: {borderWidth: 0},
                xAxis: [{
                    type: 'category',
                    boundaryGap: false,
                    splitLine: {
                        show: false
                    },
                    axisLine: {
                        lineStyle: {
                            color: style.axesColor
                        }
                    },
                    data: chartData.x
                }],
                yAxis: [{
                    type: 'value',
                    axisLabel: {
                        formatter: '{value}'
                    },
                    splitLine: {
                        show: false
                    },
                    axisLine: {
                        lineStyle: {
                            color: style.axesColor
                        }
                    },
                    min: 0,
                    max: max
                }],
                animation: false,
                title: {
                    text: description.title,
                    subtext: description.subtitle + ' (' + unit + ')'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter: function(params){
                        var res = '';
                        for (var i = 0, l = params.length; i < l; i++) {
                            var dataStr;
                            if (params[i].data != undefined) {
                                dataStr = params[i].data + ' ' + unit;
                            } else {
                                dataStr = "无数据";
                            }
                            if (params[i].seriesName !== '') {
                                res += params[i].seriesName + '：' + dataStr + '<br/>';
                            }
                        }
                        return res;
                    }
                },
                series: []
            }
            var keys = Object.keys(chartData.y).sort();
            angular.forEach(keys, function (key) {
                var name  = seriesNameGetter(key);
                if (name) {
                    var values = chartData.y[key];
                    option.series.push(buildSerieStyle(name, values, style));
                }
            })
            if (option.series.length <= 0){
                option.series.push(buildSerieStyle('', [undefined], style))
            }
            return option;
        }
        
        function buildSerieStyle(name, values, style) {
            var serie = {
                name: name,
                type: 'line',
                itemStyle: {
                    normal: {
                        lineStyle: {
                            width: style.lineWidth
                        }
                    }
                },
                symbolSize: '<2|4>',
                data: values
            };
            return serie;
        }
    }

})();