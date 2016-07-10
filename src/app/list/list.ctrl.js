(function () {
    'use strict';
    angular.module('glance.app')
        .controller('AppListCtrl', AppListCtrl);


    /* @ngInject */
    function AppListCtrl(apps) {
        var self = this;

        console.log(apps);

        activate();

        function activate() {
            ///
        }
    }
})();
