(function () {
    'use strict';
    angular.module('app.utils')
        .filter('statusZh', statusZh);

    /* @ngInject */
    function statusZh($filter) {
        //////
        return function (statusEn) {
            var en2zh = ['second', 'seconds', 'minute', 'minutes', 'hour', 'hours', 'day', 'days', 'week', 'weeks', 'month', 'months', 'year', 'years'];
            var one = ["a", "an", "one"];

            var regexPattern = " (\\d+|" + one.join("|") + ") (" + en2zh.join("|") + ")";
            var regexExp = RegExp(regexPattern, 'ig');
            var duration = statusEn.match(regexExp);
            var result = "";
            if ( duration instanceof Array ) {
                duration = duration.pop();
                duration = duration.split(" ");
                if ( duration.length === 3 ) {
                    one.includes(duration[1]) ? result = "1" : result = duration[1];
                }
                result = result + " " + $filter('translate')(duration[2]);
            }
            return result;
        }
    }
})();
