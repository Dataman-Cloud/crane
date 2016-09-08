(function () {
    'use strict';
    angular.module('app.utils')
        .filter('statusZh', statusZh);

    /* @ngInject */
    function statusZh() {
        //////
        return function (statusEn) {
            var en2zh = {
                'second': '秒',
                'seconds': '秒',
                'minute': '分钟',
                'minutes': '分钟',
                'hour': '小时',
                'hours': '小时',
                'day': '天',
                'days': '天',
                'month': '月',
                'months': '月',
                'year': '年',
                'years': '年'
            };
            var one = ["a", "an", "one"];

            var regexPattern = " (\\d+|" + one.join("|") + ") (" + Object.keys(en2zh).join("|") + ")";
            var regexExp = RegExp(regexPattern, 'ig');
            var duration = statusEn.match(regexExp);
            var result = "";
            if ( duration instanceof Array ) {
                duration = duration.pop();
                duration = duration.split(" ");
                if ( duration.length === 3 ) {
                    one.includes(duration[1]) ? result = "1" : result = duration[1];
                }
                result = result + " " + en2zh[duration[2]];
            }

            return result;
        }
    }
})();
