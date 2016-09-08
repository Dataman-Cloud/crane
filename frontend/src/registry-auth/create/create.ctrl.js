(function () {
    'use strict';
    angular.module('app.registryAuth')
        .controller('RegistryAuthCreateCtrl', RegistryAuthCreateCtrl);

    /* @ngInject */
    function RegistryAuthCreateCtrl(registryAuthCurd) {
        var self = this;

        self.form = {
            Name: '',
            Username: '',
            Password: ''
        };

        self.create = create;

        activate();

        function activate() {
            ///
        }

        function create() {
            registryAuthCurd.createRegAuth(self.form)
        }
    }
})();
