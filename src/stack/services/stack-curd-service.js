/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackCurd', stackCurd);


    /* @ngInject */
    function stackCurd(stackBackend, formModal, confirmModal, Notification, $state) {
        //////
        return {
            upServiceScale: upServiceScale,
            deleteStack: deleteStack
        };
        
        function upServiceScale(ev, stackName, serviceID, curScale) {
            formModal.open('/src/stack/modals/up-scale.html', ev,
                {dataName: 'scale', initData: curScale}).then(function (scale) {
                stackBackend.upServiceScale(stackName, serviceID, scale).then(function (data) {
                    Notification.success('修改任务数成功');
                });
            });
        }
        
        function deleteStack(ev, stackName) {
            confirmModal.open("是否确认删除编排？", ev).then(function () {
                stackBackend.deleteStack(stackName)
                    .then(function (data) {
                        $state.go('stack.list', undefined, {reload: true});
                    })
            });
        }


    }
})();
