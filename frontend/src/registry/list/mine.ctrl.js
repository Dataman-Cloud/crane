(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListMineCtrl', RepositorieListMineCtrl);


    /* @ngInject */
    function RepositorieListMineCtrl(repositories, registryBackend, registryCurd, utils, $rootScope, $stateParams) {
        var self = this;

        self.repositories = repositories;

        self.publicImage = publicImage;
        self.hideImage = hideImage;
        self.toggleImagePublicity = toggleImagePublicity;

        activate()
        function activate() {
            angular.forEach(self.repositories, function (repository) {
              repository.name = repository.Namespace + "/" + repository.Image;
              repository.tags = [];

              registryBackend.listRepositoryTags(repository.name).then(function (data) {
                  angular.forEach(data, function (tag) {
                    repository.tags.push({tag: tag.Tag, digest: tag.Digest, size: tag.Size})
                  });
              });

            });
            angular.forEach(self.repositories, function (repository, index) {
                if ($stateParams.open === repository.name) {
                    openRepository(index);
                }
            })
        }

        function toggleImagePublicity(repository) {
          if(repository.Publicity == 1){
            publicImage(repository.Namespace, repository.Image)
          }else{
            hideImage(repository.Namespace, repository.Image)
          }
        }

        function publicImage(namespace, image) {
            registryCurd.publicImage(namespace, image);
        }

        function hideImage(namespace, image) {
            registryCurd.hideImage(namespace, image);
        }
    }
})();
