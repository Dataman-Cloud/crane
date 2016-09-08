(function () {
    'use strict';
    angular.module('app.utils')
        .config(i18nEn);

    /* @ngInject */
    function i18nEn($translateProvider) {
        $translateProvider.translations('en-us', {
            "LOGOUT": 'logout'
        });
    }
})();
