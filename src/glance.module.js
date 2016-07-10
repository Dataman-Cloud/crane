(function () {
    'use strict';
    angular.module('glance',
        [
            'ui-notification',
            'ui.router',
            'ngMaterial',
            'ngCookies',
            'ngAnimate',
            'ngSocket',
            'ngSanitize',
            'glance.utils',
            'glance.layout',
            'glance.app'
        ]);
})();