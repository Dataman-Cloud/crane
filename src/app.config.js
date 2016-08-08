/**
 * Created by my9074 on 16/3/21.
 */
(function () {
    'use strict';
    angular.module('app')
        .config(configure);

    /* @ngInject */
    function configure($locationProvider, $interpolateProvider, $translateProvider, NotificationProvider) {
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
    }
})();
