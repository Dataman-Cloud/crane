/**
 * Created by my9074 on 16/3/11.
 */
(function () {
    'use strict';
    angular.module('app.utils')
        .filter('fixIdWidth', fixed_id_width);

    /* @ngInject */
    function fixed_id_width() {
        //////
        return function (input) {
          return input.substring(0, 12)
        }
    }
})();
