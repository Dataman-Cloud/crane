(function () {
    'use strict';
    angular.module('app.registry')
        .factory('registryCurd', registryCurd);


    /* @ngInject */
    function registryCurd(registryBackend, confirmModal, $rootScope, utils, $state) {
        //////
        return {
            deleteImage: deleteImage,
            isPublicRepository: isPublicRepository,
            isMyRepository: isMyRepository
        };
        
        function deleteImage(repository, tag, ev) {
            confirmModal.open("是否确认删除镜像？", ev).then(function () {
                if (isPublicRepository(repository)) {
                    $state.go('registry.list.public', {open: repository}, {reload: true});
                } else {
                    $state.go('registry.list.my', {open: repository}, {reload: true});
                }
//                registryBackend.deleteImage(repository, tag)
//                    .then(function (data) {
//                        Notification.success('删除成功');
//                    })
            });
        }
        
        function isPublicRepository(repository) {
            return utils.startWith(repository, 'library/');
        }
        
        function isMyRepository(repository) {
            return utils.startWith(repository, $rootScope.accountId+'/')
        }
    }
})();
