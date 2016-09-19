/**
 * Created by my9074 on 16/3/9.
 */
(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackCurd', stackCurd);


    /* @ngInject */
    function stackCurd(stackBackend, formModal, confirmModal, Notification, $state, createSelectModal, $filter) {
        //////
        return {
            upServiceScale: upServiceScale,
            createSelect: createSelect,
            deleteStack: deleteStack,
            createStack: createStack,
            updateService: updateService
        };

        function upServiceScale(ev, stackName, serviceID, curScale) {
            formModal.open('/src/stack/modals/up-scale.html', ev,
                {dataName: 'scale', initData: curScale}).then(function (scale) {
                stackBackend.upServiceScale(stackName, serviceID, scale).then(function (data) {
                    Notification.success($filter('translate')('Scaled successfully'));
                    $state.reload();
                });
            });
        }

        function createSelect(ev) {
            createSelectModal.open('/src/stack/modals/create-select.html', ev)
        }

        function deleteStack(ev, stackName) {
            confirmModal.open("Modal title for delete stack", ev).then(function () {
                stackBackend.deleteStack(stackName)
                    .then(function (data) {
                        $state.go('stack.list', undefined, {reload: true});
                    })
            });
        }

        function createStack(formData, form, groupId) {
            return stackBackend.createStack(formData, form, groupId)
                .then(function (data) {
                    Notification.success($filter('translate')('Deploying, image pull will cost some time'));
                    $state.go('stack.detail.service', {stack_name: formData.Namespace})
                })
        }

        function updateService(data, form, stackName, serviceID) {
            stackBackend.updateService(data, form, stackName, serviceID).then(function (data) {
                Notification.success($filter('translate')('Updated Successfully'));
                $state.go('stack.serviceDetail.config', {stack_name: stackName, service_id: serviceID}, {reload: true});
            });
        }
    }
})();
