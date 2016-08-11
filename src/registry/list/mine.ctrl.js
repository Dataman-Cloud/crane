(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListMineCtrl', RepositorieListMineCtrl);


    /* @ngInject */
    function RepositorieListMineCtrl(repositories, registryBackend, registryCurd, utils, $rootScope, $stateParams) {
        var self = this;

        self.repositories = repositories;

        self.openRepository = openRepository;
        self.closeRepository = closeRepository;
        self.deleteImage = deleteImage;
        self.publicImage = publicImage;
        self.hideImage = hideImage;
        
        activate()
        
        function activate() {
            angular.forEach(self.repositories, function (repository) {
              repository.tags = [];
              repository.show = false;
              repository.name = repository.Namespace + "/" + repository.Image;
            });
            angular.forEach(self.repositories, function (repository, index) {
                if ($stateParams.open === repository.name) {
                    openRepository(index);
                }
            })
        }
        
        function openRepository(index) {
            registryBackend.listRepositoryTags(self.repositories[index].name).then(function (data) {
                self.repositories[index].show = true;
                self.repositories[index].tags = [];
                angular.forEach(data, function (tag) {
                  self.repositories[index].tags.push({tag: tag.Tag, digest: tag.Digest, size: tag.Size})
                });
            });
        }
        
        function closeRepository(index) {
            self.repositories[index].show = false;
        }
        
        function deleteImage(repository, tag, ev) {
            registryCurd.deleteImage(repository, tag, ev);
        }

        function publicImage(namespace, image) {
            registryCurd.publicImage(namespace, image);
        }

        function hideImage(namespace, image) {
            registryCurd.hideImage(namespace, image);
        }
    }
})();
