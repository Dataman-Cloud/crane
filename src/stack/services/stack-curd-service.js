/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackCurd', stackCurd);


    /* @ngInject */
    function stackCurd(stackBackend, formModal, Notification) {
        //////
        return {
            upServiceScale: upServiceScale
        };
        
        function upServiceScale(ev, stackName, serviceID, curScale) {
            formModal.open('/src/stack/modals/up-scale.html', ev,
                {dataName: 'scale', initData: curScale}).then(function (scale) {
                stackBackend.upServiceScale(stackName, serviceID, scale).then(function (data) {
                    Notification.success('扩缩成功');
                });
            });
        }


    }
})();