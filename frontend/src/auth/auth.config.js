(function () {
    'use strict';
    angular.module('app.auth')
        .config(configure);

    /* ngInject */
    function configure($locationProvider, $interpolateProvider, $translateProvider) {

        $locationProvider.html5Mode(true);
        $interpolateProvider.startSymbol('{/');
        $interpolateProvider.endSymbol('/}');

        var language = window.localStorage.getItem('language') || 'en';
        $translateProvider.preferredLanguage(language);
        $translateProvider.useSanitizeValueStrategy(null);
    }
})();