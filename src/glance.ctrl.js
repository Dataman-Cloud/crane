(function () {
    'use strict';
    angular.module('glance')
        .controller('RootCtrl', RootCtrl);

    /* @ngInject */
    function RootCtrl($state, $window, mdSideNav) {
        var self = this;

        self.noticeNav = mdSideNav.createSideNav('noticeNav');
        self.goBack = goBack;
        self.togShortMenu = togShortMenu;
        self.isShortMenu = false;

        activate();

        function activate() {
            ///
        }

        function goBack(state) {
            if(state){
                $state.go(state);
            }else{
                $window.history.length > 2 ? $window.history.back() : $state.go('dashboard.home');
            }
        }

        function togShortMenu() {
            self.isShortMenu = !self.isShortMenu;
        }

    }
})();