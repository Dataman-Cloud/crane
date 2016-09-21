/**
 * Created by my9074 on 16/3/21.
 */
(function () {
    'use strict';
    angular.module('app')
        .config(configure);

    /* @ngInject */
    function configure($locationProvider, $interpolateProvider, $translateProvider, NotificationProvider, $breadcrumbProvider) {
        ////
        $locationProvider.html5Mode(true);

        $interpolateProvider.startSymbol('{/');
        $interpolateProvider.endSymbol('/}');

        $translateProvider.useSanitizeValueStrategy(null);

        NotificationProvider.setOptions({
            delay: 5000 * 3,
            replaceMessage: true
        });

        $breadcrumbProvider.setOptions({
            template: '<div id="breadcrumb-main"><span ng-repeat="step in steps" ng-switch="$last || !!step.abstract">' +
            '<a ng-switch-when="false" href="{/step.ncyBreadcrumbLink/}">{/step.ncyBreadcrumbLabel/}</a>' +
            '<i class="fa fa-angle-right" aria-hidden="true" ng-switch-when="false"></i>' +
            '<span ng-switch-when="true" >{/step.ncyBreadcrumbLabel/}</span></div>'
        });
    }
})();
