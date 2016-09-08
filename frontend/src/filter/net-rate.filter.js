(function () {
    'use strict';
    angular.module('app.utils')
        .filter('netRate', netRate);

    /* @ngInject */
    function netRate() {
        //////
        return function (rawSize) {
            var units = ['b/s', 'Kb/s', 'Mb/s'];
            var unitIndex = 0;
            while(rawSize >= 1024 && unitIndex < units.length-1) {
                rawSize /= 1024;
                unitIndex++;
            }
            return rawSize.toFixed(2) + units[unitIndex];
        }
    }
})();