(function () {
    'use strict';
    angular.module('app.registryAuth')
        .factory('registryAuthCurd', registryAuthCurd);


    /* @ngInject */
    function registryAuthCurd(registryAuthBackend, $state, Notification, confirmModal) {
        return {
            createRegAuth: createRegAuth,
            deleteRegAuth: deleteRegAuth
        };

        function createRegAuth(data, form) {
            registryAuthBackend.createRegAuth(data, form)
                .then(function (data) {
                    Notification.success('创建成功');
                    $state.go('registryAuth.list', undefined, {reload: true})
                })
        }

        function deleteRegAuth(name) {
            confirmModal.open("是否确认删除该认证？").then(function () {
                registryAuthBackend.deleteRegAuth(name)
                    .then(function (data) {
                        Notification.success('删除成功');
                        $state.reload()
                    })
            });
        }
    }
})();
