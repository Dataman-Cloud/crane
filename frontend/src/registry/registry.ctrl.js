(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RegistryCtrl', RegistryCtrl);

    /* @ngInject */
    function RegistryCtrl(registryBackend, registryCurd, $scope) {
        var self = this;

        activate();
        function activate() {
            registryBackend.getNamespace().then(function (data) {
                self.namespace = data.Namespace;
            }, function (resp) {
                if (resp.data && angular.isObject(resp.data) && resp.code && NAMESPACE_NO_FOUND_ERROR_CODE.indexOf(resp.code) != -1) {
                    registryCurd.createNamespace($scope.staticForm);
                }
            });
        }
    }
})();
