(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('mdSideNav', mdSideNav);

    /* @ngInject */
    function mdSideNav($mdSidenav) {

        return {createSideNav: createSideNav};

        function MdSideNav(navId) {
            this.navId = navId;

            MdSideNav.prototype.buildToggler = function () {
                $mdSidenav(this.navId)
                    .toggle();
            };

            MdSideNav.prototype.isOpenNav = function () {
                return $mdSidenav(this.navId).isOpen();
            };
        }


        function createSideNav(navId) {
            return new MdSideNav(navId)
        }

    }
})();