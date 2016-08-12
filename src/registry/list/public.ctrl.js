(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListPublicCtrl', RepositorieListPublicCtrl);


    /* @ngInject */
    function RepositorieListPublicCtrl(repositories, registryBackend, registryCurd, utils, $rootScope, $stateParams) {
        var self = this;

       self.repositories = repositories;

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
    }
})();
