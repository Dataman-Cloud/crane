(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListContentCtrl', RepositorieListContentCtrl);


    /* @ngInject */
    function RepositorieListContentCtrl(repositories, type, registryBackend, registryCurd, utils, $rootScope, $stateParams) {
        var self = this;
        
        self.repositories = [];
        
        self.openRepository = openRepository;
        self.closeRepository = closeRepository;
        self.deleteImage = deleteImage;
        
        activate()
        
        function activate() {
            angular.forEach(repositories.repositories, function (repository) {
                if ((type === 'my' && registryCurd.isMyRepository(repository))||
                    (type !== 'my' && registryCurd.isPublicRepository(repository))) {
                    self.repositories.push({name: repository, tags: [], show: false});
                }
            });
            angular.forEach(self.repositories, function (repository, index) {
                if ($stateParams.open === repository.name) {
                    openRepository(index);
                }
            })
        }
        
        function openRepository(index) {
            registryBackend.listRepositoryTags(self.repositories[index].name).then(function (data) {
                self.repositories[index].tags = data.tags;
                self.repositories[index].show = true;
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
