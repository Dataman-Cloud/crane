(function () {
    'use strict';
    angular.module('app',
        [
            'ui-notification',
            'ui.router',
            'ngMaterial',
            'ngCookies',
            'ngAnimate',
            'ngSocket',
            'ngSanitize',
            'app.utils',
            'app.stack',
            'app.app',
            'app.node'
        ]);
})();