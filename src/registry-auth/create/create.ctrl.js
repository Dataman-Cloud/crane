(function () {
    'use strict';
    angular.module('app.registryAuth')
        .controller('RegistryAuthCreateCtrl', RegistryAuthCreateCtrl);

    /* @ngInject */
    function RegistryAuthCreateCtrl() {
        var self = this;

        self.form = {
            Name: '',
            UserName: '',
            PassWord: ''
        };

        self.create = create;

        activate();

        function activate() {
            ///
        }

        function create() {
            console.log(self.form)
        }
    }
})();
