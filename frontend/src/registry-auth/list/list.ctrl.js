(function () {
    'use strict';
    angular.module('app.registryAuth')
        .controller('RegistryAuthListCtrl', RegistryAuthListCtrl);


    /* @ngInject */
    function RegistryAuthListCtrl(reAuths, registryAuthCurd) {
        var self = this;

        self.reAuths = reAuths;

        self.deleteRegAuth = deleteRegAuth;

        activate();

        function activate() {
            ///
        }

        function deleteRegAuth(name) {
            registryAuthCurd.deleteRegAuth(name)
        }
    }
})();
