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
            deleteImage: deleteImage
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

        function getCatalog(catalogName) {
            return gHttp.Resource('registry.catalog', {catalog_name: catalogName}).get();
        }
        
        function deleteImage(repository, tag) {
            return gHttp.Resource('registry.image', {repository: repository, tag: tag}).delete();
        }
    }
})();
