(function () {
    'use strict';
    angular.module('app.registry')
        .factory('registryCurd', registryCurd);


    /* @ngInject */
    function registryCurd(registryBackend, confirmModal, $rootScope, utils, $state, Notification) {
        //////
        return {
            deleteImage: deleteImage,
            isPublicRepository: isPublicRepository,
            hideImage: hideImage,
            publicImage: publicImage,
            isMyRepository: isMyRepository,
            createCatalog: createCatalog,
            deleteCatalog: deleteCatalog,
            updateCatalog: updateCatalog
        };

        function deleteImage(repository, tag, ev) {
            confirmModal.open("是否确认删除镜像？", ev).then(function () {
                if (isPublicRepository(repository)) {
                    $state.go('registry.list.public', {open: repository}, {reload: true});
                } else {
                    $state.go('registry.list.mine', {open: repository}, {reload: true});
                }
            });
        }

        function isPublicRepository(repository) {
            return utils.startWith(repository, 'library/');
        }

        function isMyRepository(repository) {
            return utils.startWith(repository, $rootScope.accountId + '/')
        }

        function publicImage(namespace, image) {
            registryBackend.publicImage(namespace, image)
        }

        function hideImage(namespace, image) {
            registryBackend.hideImage(namespace, image)
        }

        function createCatalog(data, form) {
            registryBackend.createCatalog(data, form)
                .then(function (data) {
                    $state.go('registry.list.catalogs', null, {reload: true});
                })
        }

        function deleteCatalog(catalogId, ev) {
            confirmModal.open("是否确认删除该项目模板？", ev).then(function () {
                registryBackend.deleteCatalog(catalogId)
                    .then(function (data) {
                        Notification.success('删除成功');
                        $state.go('registry.list.catalogs', null, {reload: true});
                    })
            });
        }

        function updateCatalog(catalogId, data) {
            registryBackend.updateCatalog(catalogId, data)
                .then(function (data) {
                    Notification.success('更新成功');
                    $state.go('registry.list.catalogs', null, {reload: true});
                })
        }
    }
})();
