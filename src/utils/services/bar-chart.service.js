(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('barChart', barChart);

    /* @ngInject */
    function barChart() {
        return {
            buildChartData: buildChartData,
            paint: paint
        };

        function buildChartData(data, xGetter, dataGetter, options) {
            if (!options) {
                options = {}
            }
            if (!options.numFixed && options.numFixed !== 0) {
                options.numFixed = 2;
            }
            var chartData = {y: [], x: []};
            for (var i = 0; i < data.length; i++) {
                var curData = data[i];
                chartData.y.push(dataGetter(curData));
                chartData.x.push(xGetter(curData))
            }
            reverseChartData(chartData);
            return chartData;
        }
        
        function reverseChartData(chartData) {
            chartData.x.reverse();
            chartData.y.reverse();
        }
        
        function paint(chartData, domId, description, seriesName, barWidth) {
            var option = buildChartOption(chartData, description, seriesName, barWidth)
            var chart = echarts.init(document.getElementById(domId));
            chart.setOption(option);
        }
        function buildChartOption(chartData, description, seriesName, barWidth) {
            var unit = description.unit||''
            var option = {
                    title: {
                        text: description.title,
                        subtext: description.subtitle + ' (' + unit + ')'
                    },
                    animation: false,
                    tooltip: {},
                    legend: {
                        data:[seriesName]
                    },
                    xAxis: {
                        data: chartData.x
                    },
                    yAxis: {},
                    series: [{
                        name: seriesName,
                        type: 'bar',
                        data: chartData.y,
                        barWidth: barWidth
                    }],
                };
            return option;
        }
    }

})();