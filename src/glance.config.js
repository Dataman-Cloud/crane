/**
 * Created by my9074 on 16/3/21.
 */
(function () {
    'use strict';
    angular.module('glance')
        .config(configure);

    /* @ngInject */
    function configure($locationProvider, $interpolateProvider, $urlRouterProvider) {
        ////
        $urlRouterProvider.otherwise('/layout');

        $locationProvider.html5Mode(true);
        $interpolateProvider.startSymbol('{/');
        $interpolateProvider.endSymbol('/}');
    }
})();
