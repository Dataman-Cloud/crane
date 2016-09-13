(function () {
    'use strict';
    angular.module('app.registry')
        .factory('registryBackend', registryBackend);


    /* @ngInject */
    function registryBackend(gHttp) {
        return {
            listPublicRepositories: listPublicRepositories,
            listMineRepositories: listMineRepositories,
            listRepositoryTags: listRepositoryTags,
            getImage: getImage,
            listCatalogs: listCatalogs,
            getCatalog: getCatalog,
            createCatalog: createCatalog,
            deleteCatalog: deleteCatalog,
            deleteImage: deleteImage,
            hideImage: hideImage,
            publicImage: publicImage
        };

        function listPublicRepositories() {
            return gHttp.Resource('registry.publicRepositories').get();
        }

        function listMineRepositories() {
            return gHttp.Resource('registry.mineRepositories').get();
        }

        function listRepositoryTags(repository) {
            return gHttp.Resource('registry.listTags', {repository: repository}).get();
        }

        function getImage(repository, tag) {
            return gHttp.Resource('registry.image', {repository: repository, tag: tag}).get();
        }

        function listCatalogs() {
            return gHttp.Resource('registry.catalogs').get();
        }

        function getCatalog(catalogId) {
            return gHttp.Resource('registry.catalog', {catalog_id: catalogId}).get();
        }

        function createCatalog(data, form) {
            return gHttp.Resource('registry.catalogs').post(data, {form: form});
        }

        function deleteCatalog(catalogId) {
            return gHttp.Resource('registry.catalog', {catalog_id: catalogId}).delete();
        }

        function deleteImage(repository, tag) {
            return gHttp.Resource('registry.image', {repository: repository, tag: tag}).delete();
        }

        function hideImage(namespace, image) {
            var data = {"Publicity": 0};
            return gHttp.Resource('registry.publicity', {namespace: namespace, image: image}).patch(data);
        }

        function publicImage(namespace, image) {
            var data = {"Publicity": 1};
            return gHttp.Resource('registry.publicity', {namespace: namespace, image: image}).patch(data);
        }
    }
})();
