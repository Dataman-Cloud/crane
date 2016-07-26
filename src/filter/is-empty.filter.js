(function () {
    'use strict';
    angular.module('app.utils')
        .filter('isEmpty', isEmpty);

    /* @ngInject */
    function isEmpty() {
        //////
        return function (value) {
            var empty = false;
            if (!value) {
                empty = true;
            } else if (angular.isArray(value)) {
                if (value.length <= 0) {
                    empty = true;
                }
            } else if (angular.isObject(value)) {
                if (Object.keys(value).length <= 0) {
                    empty = true;
                }
            }
            return empty;
        }
    }
})();