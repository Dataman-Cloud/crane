/**
 * Created by my9074 on 16/2/24.
 */
(function () {
    'use strict';

    angular.module('app.utils').factory('chartUtil', chartUtil);

    /* @ngInject */
    function chartUtil($filter) {
        return {
            createDefaultOptions: createDefaultOptions,
            pushData: pushData,
            updateForceY: updateForceY
        };

        function createDefaultOptions() {
            var colors = [
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
            ];
            return {
                chart: {
                    type: 'lineChart',
                    noData: '暂无数据',
                    height: 450,
                    margin: {
                        top: 20,
                        right: 20,
                        bottom: 40,
                        left: 55
                    },
                    x: function (d) {
                        return d.x;
                    },
                    y: function (d) {
                        return d.y;
                    },
                    useInteractiveGuideline: true,
                    showLegend: false,
                    xAxis: {
                        axisLabel: '时间',
                        tickFormat: function (d) {
                            return $filter('date')(d, 'HH:mm:ss');
                        },
                        showMaxMin: false
                    },
                    yAxis: {
                        tickFormat: function (d) {
                            return d3.format('.02f')(d) + '%';
                        },
                        axisLabelDistance: -10
                    },
                    pointSize: 0.1,
                    forceY: [0],
                    color: function (d, i) {
                        if (!colors[i]) {
                            while (true) {
                                var color = '#' + ('00000' + (Math.random() * 0x1000000 << 0).toString(16)).slice(-6);
                                var duplicate = false;
                                for (var j = 0; j < colors.length; j++) {
                                    if (colors[j] == color) {
                                        duplicate = true;
                                        break;
                                    }
                                }
                                if (!duplicate) {
                                    colors[i] = color;
                                    break;
                                }
                            }
                        }
                        return colors[i];
                    }
                },
                title: {
                    enable: true
                }
            }
        }

        function pushData(data, serialKey, value, pointNum, interval, area) {
            if (!interval) {
                interval = 1000;
            }
            if (area === undefined) {
                area = true;
            }
            var i;
            for (i = 0; i < data.length; i++) {
                if (data[i].key === serialKey) {
                    break;
                }
            }
            if (i == data.length) {
                data.push({
                    values: [],
                    key: serialKey,
                    area: area
                });
            }
            data[i].values.push(value);
            while (data[i].values.length !== pointNum) {
                if (data[i].values.length > pointNum) {
                    data[i].values.shift();
                } else {
                    data[i].values.unshift({x: data[i].values[0].x - interval, y: 0});
                }
            }
        }

        function updateForceY(chartOptions, data, min, maxRatio, minMax, maxMax) {
            var newForceY = _buildNewForceY(data, min, maxRatio, minMax, maxMax);
            var flag = false;
            if (!angular.equals(newForceY, chartOptions.forceY)) {
                chartOptions.forceY = newForceY;
                flag = true;
            }
            return flag;
        }

        function _buildNewForceY(data, min, maxRatio, minMax, maxMax) {
            var valueMax = Math.maxPlus(data, function (serialData) {
                return Math.maxPlus(serialData.values, function (value) {
                    return value.y;
                })
            });
            var curMax = valueMax * maxRatio;
            if (maxMax !== undefined && maxMax < curMax) {
                if (maxMax < valueMax) {
                    curMax = valueMax;
                } else {
                    curMax = maxMax;
                }
            }
            if (minMax !== undefined && curMax < minMax) {
                curMax = minMax;
            }
            return [min, Math.ceil(curMax)];
        }

    }
})();