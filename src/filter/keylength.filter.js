/**
 * Created by my9074 on 16/3/11.
 */
(function () {
    'use strict';
    angular.module('app.utils')
        .filter('keylength', keylength);

    /* @ngInject */
    function keylength() {
        //////
        return function (input) {
            if (angular.isObject(input)) {
                return Object.keys(input).length;
            } else {
                return '未配置'
            }
        }
    }
})();