(function () {
    'use strict';
    angular.module('app.registryAuth')
        .factory('registryAuthCurd', registryAuthCurd);


    /* @ngInject */
    function registryAuthCurd(registryAuthBackend, $state, Notification, confirmModal, $filter) {
        return {
            createRegAuth: createRegAuth,
            deleteRegAuth: deleteRegAuth
        };

        function createRegAuth(data, form) {
            registryAuthBackend.createRegAuth(data, form)
                .then(function (data) {
                    Notification.success($filter('translate')('Created successfully'));
                    $state.go('registryAuth.list', undefined, {reload: true})
                })
        }

        function deleteRegAuth(name) {
            confirmModal.open("Are you sure to delete this auth pair?").then(function () {
                registryAuthBackend.deleteRegAuth(name)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
                        $state.reload()
                    })
            });
        }
    }
})();
