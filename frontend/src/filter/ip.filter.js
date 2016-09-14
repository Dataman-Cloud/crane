/**
 * Created by wtzhou.
 */
(function () {
    'use strict';
    angular.module('app.utils')
        .filter('ip', ip);

    /* @ngInject */
    function ip() {
        //////
        return function (input) {
            input = input || '';

            //http://www.regular-expressions.info/examples.html
            var r = /\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b/;
            var t = input.match(r);
            return t && t.length ? t[0] : ""
        }
    }
})();
