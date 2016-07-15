/**
 * Created by my9074 on 16/3/11.
 */
(function () {
    'use strict';
    angular.module('app.utils')
        .filter('spliceStr', spliceStr);

    /* @ngInject */
    function spliceStr() {
        //////
        return function (input, separator, filterKey) {
            input = input || '';
            // conditional based on optional argument
            var index = input.indexOf(separator);
            if (filterKey === 'value') {
                if (index != -1) {
                    input = input.slice(index + 1)
                } else {
                    input = ""
                }
            } else if (filterKey === 'key') {
                if (index != -1) {
                    input = input.slice(0, index)
                }
            }

            return input;
        }
    }
})();