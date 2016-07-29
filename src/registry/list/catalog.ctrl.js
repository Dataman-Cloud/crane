(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListCatalogCtrl', RepositorieListCatalogCtrl);

    /* @ngInject */
    function RepositorieListCatalogCtrl(catalogs) {
        var self = this;

        self.catalogs = catalogs;
    }
})();
