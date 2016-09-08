(function () {
    'use strict';

    angular.module('app.user').factory('userCurd', userCurd);

    /* @ngInject */
    function userCurd(userBackend, formModal, $rootScope, Notification) {

        return {
            registerLicense: registerLicense
        };

        function registerLicense(event) {
            formModal.open('/src/user/modals/register-license.html', event, {dataName: 'serialNum'})
                .then(function (data) {
                    userBackend.registerLicense(data)
                        .then(function (data) {
                            Notification.success('激活成功');
                            $rootScope.licenseValidFlag = true;
                        })
                });
        }
    }
})();