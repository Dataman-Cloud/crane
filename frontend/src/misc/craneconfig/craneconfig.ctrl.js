(function () {
    'use strict';
    angular.module('app.misc')
        .controller('MiscConfigCtrl', MiscConfigCtrl);


    /* @ngInject */
    function MiscConfigCtrl(craneconfig) {
        var self = this;
        self.craneconfig = craneconfig;

        activate();

        function activate() {
            ///
        }
    }
})();
