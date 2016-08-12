(function () {
    'use strict';
    angular.module('app.utils')
        .filter('size', size);

    /* @ngInject */
    function size() {
        //////
        return function (rawSize) {
            var units = ['B', 'KB', 'MB', 'GB', 'TB'];
            var unitIndex = 0;
            while(rawSize >= 1024 && unitIndex < units.length-1) {
                rawSize /= 1024;
                unitIndex++;
            }
            return rawSize.toFixed(2) + ' ' + units[unitIndex];
        }
    }
})();
