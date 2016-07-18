(function () {
    'use strict';
    angular.module('app.misc')
        .controller('MiscConfigCtrl', MiscConfigCtrl);


    /* @ngInject */
    function MiscConfigCtrl(rolexconfig) {
        var self = this;
        self.rolexconfig = rolexconfig;

        activate();

        function activate() {
            ///
        }
    }
})();
