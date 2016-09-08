(function () {
    'use strict';
    angular.module('app.utils')
        .config(i18nCn);

    /* @ngInject */
    function i18nCn($translateProvider) {
        $translateProvider.translations('zh-cn', {
            "LOGOUT": '登出'
        });
    }
})();
