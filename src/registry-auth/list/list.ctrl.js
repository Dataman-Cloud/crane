(function () {
    'use strict';
    angular.module('app.registryAuth')
        .controller('RegistryAuthListCtrl', RegistryAuthListCtrl);


    /* @ngInject */
    function RegistryAuthListCtrl() {
        var self = this;

        self.deleteRegAuth = deleteRegAuth;

        activate();

        function activate() {
            ///
        }

        function deleteRegAuth() {
            //TODO
        }
    }
})();
