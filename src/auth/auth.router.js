(function () {
    'use strict';
    angular.module('app.auth')
        .config(configure);

    /* @ngInject */
    function configure($urlRouterProvider,
                       $stateProvider) {

        $urlRouterProvider.otherwise('/auth');

        $stateProvider
            .state('auth', {
                url: '/auth',
                template: '<ui-view/>',
                targetState: 'login'
            })
            .state('auth.login', {
                url: '/login?return_to',
                templateUrl: '/src/auth/login/login.html',
                controller: 'LoginCtrl as loginCtrl'
            });
    }
})();