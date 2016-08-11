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

        NotificationProvider.setOptions({
            delay: 3000,
            positionX: 'right',
            positionY: 'top',
            replaceMessage: true,
            startTop: 20,
            startRight: 260
        });

        $translateProvider.preferredLanguage('zh-cn');
        $translateProvider.useSanitizeValueStrategy(null);

        $breadcrumbProvider.setOptions({
            template: '<div id="breadcrumb-main"><span ng-repeat="step in steps" ng-switch="$last || !!step.abstract">' +
            '<a ng-switch-when="false" href="{/step.ncyBreadcrumbLink/}">{/step.ncyBreadcrumbLabel/}</a>' +
            '<i class="fa fa-angle-right" aria-hidden="true"></i>' +
            '<span ng-switch-when="true" >{{step.ncyBreadcrumbLabel}}</span></div>'
        });
    }
})();
