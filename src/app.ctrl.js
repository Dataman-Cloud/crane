(function () {
    'use strict';
    angular.module('app')
        .controller('RootCtrl', RootCtrl);

    /* @ngInject */
    function RootCtrl($state, $window, mdSideNav, gHttp, $cookies, utils, appCommon) {
        var self = this;

        self.noticeNav = mdSideNav.createSideNav('noticeNav');
        self.goBack = goBack;
        self.togShortMenu = togShortMenu;
        self.isShortMenu = false;

        self.logout = logout;

        activate();

        function activate() {
            ///
            initUser()
        }

        function initUser() {
            var token = $cookies.get('token');
            if (token) {
                gHttp.setToken(token)
            } else {
                utils.redirectLogin(true)
            }
        }

        function logout() {
            appCommon.logout()
                .then(function () {
                    utils.redirectLogin()
                });
        }

        function goBack(state) {
            if (state) {
                $state.go(state);
            } else {
                $window.history.length > 2 ? $window.history.back() : $state.go('dashboard.home');
            }
        }

        function togShortMenu() {
            self.isShortMenu = !self.isShortMenu;
        }

    }
})();