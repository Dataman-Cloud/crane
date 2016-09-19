(function () {
    'use strict';
    angular.module('app.registry')
        .factory('registryCurd', registryCurd);


    /* @ngInject */
    function registryCurd(registryBackend, confirmModal, $rootScope, utils, $state, Notification, $filter) {
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
            confirmModal.open("Are you sure to remove the image ?", ev).then(function () {
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
            confirmModal.open("Are you sure to delete the stack template ?", ev).then(function () {
                registryBackend.deleteCatalog(catalogId)
                    .then(function (data) {
                        Notification.success($filter('translate')('Successfully deleted'));
                        $state.go('registry.list.catalogs', null, {reload: true});
                    })
            });
        }

        function updateCatalog(catalogId, data) {
            registryBackend.updateCatalog(catalogId, data)
                .then(function (data) {
                    Notification.success($filter('translate')('Updated Successfully'));
                    $state.go('registry.list.catalogs', null, {reload: true});
                })
        }
    }
})();
