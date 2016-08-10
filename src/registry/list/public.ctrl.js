(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListPublicCtrl', RepositorieListPublicCtrl);


    /* @ngInject */
    function RepositorieListPublicCtrl(repositories, registryBackend, registryCurd, utils, $rootScope, $stateParams) {
        var self = this;


       self.repositories = [];

        self.openRepository = openRepository;
        self.closeRepository = closeRepository;
        self.deleteImage = deleteImage;
        
        activate()
        
        function activate() {
            angular.forEach(repositories, function (repository) {
                self.repositories.push({name: repository.Namespace + "/" + repository.Image,
                                       pullCount: repository.PullCount,
                                       pushCount: repository.PushCount,
                                       tags: [],
                                       show: false});
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
                  self.repositories[index].tags.push({tag: tag.Tag, digest: tag.Digest})
                });
            });
        }
        
        function closeRepository(index) {
            self.repositories[index].show = false;
        }
        
        function deleteImage(repository, tag, ev) {
            registryCurd.deleteImage(repository, tag, ev);
        }
    }
})();
