(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListContentCtrl', RepositorieListContentCtrl);


    /* @ngInject */
    function RepositorieListContentCtrl(repositories, type, registryBackend, utils, $rootScope) {
        var self = this;
        
        self.repositories = [];
        self.tags = {};
        
        self.openRepository = openRepository;
        self.closeRepository = closeRepository;
        
        activate()
        
        function activate() {
            angular.forEach(repositories.repositories, function (repository) {
                if ((type === 'my' && utils.startWith(repository, $rootScope.accountId+'/'))||
                    (type !== 'my' && utils.startWith(repository, 'library/'))) {
                    self.repositories.push({name: repository, tags: [], show: false})
                }
            });
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
    }
})();
