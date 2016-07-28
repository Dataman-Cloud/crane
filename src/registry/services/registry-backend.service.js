(function () {
    'use strict';
    angular.module('app.registry')
        .factory('registryBackend', registryBackend);


    /* @ngInject */
    function registryBackend(gHttp) {
        return {
            listRepositories: listRepositories,
            listRepositoryTags: listRepositoryTags,
            getImage: getImage
        };
        
        function listRepositories() {
            return gHttp.Resource('registry.repositories').get();
        }
        
        function listRepositoryTags(repository) {
            return gHttp.Resource('registry.listTags', {repository: repository}).get();
        }
        
        function getImage(repository, tag) {
            return gHttp.Resource('registry.image', {repository: repository, tag: tag}).get();
        }
    }
})();
