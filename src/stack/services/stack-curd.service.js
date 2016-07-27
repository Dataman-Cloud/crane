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
            deleteStack: deleteStack,
            stopService: stopService,
            createStack: createStack
        };
        
        function upServiceScale(ev, stackName, serviceID, curScale) {
            formModal.open('/src/stack/modals/up-scale.html', ev,
                {dataName: 'scale', initData: curScale}).then(function (scale) {
                stackBackend.upServiceScale(stackName, serviceID, scale).then(function (data) {
                    Notification.success('修改任务数成功');
                    $state.reload();
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

        function stopService(stackName, serviceID){
            stackBackend.upServiceScale(stackName, serviceID, 0).then(function (data) {
                Notification.success('停止成功');
                $state.reload();
            });
        }

        function createStack(formData, form, groupId){
            stackBackend.createStack(formData, form, groupId)
                .then(function (data) {
                    Notification.success('创建成功');
                    $state.go('stack.detail.service', {stack_name: formData.Namespace})
                })
        }
    }
})();
