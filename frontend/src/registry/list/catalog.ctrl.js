(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListCatalogCtrl', RepositorieListCatalogCtrl);

    /* @ngInject */
    function RepositorieListCatalogCtrl(catalogs, registryCurd) {
        var self = this;

        self.catalogs = catalogs;

        self.deleteCatalog = deleteCatalog;

        function deleteCatalog(id) {
            registryCurd.deleteCatalog(id)
        }

    }
})();
