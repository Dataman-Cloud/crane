(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RepositorieListCatalogCtrl', RepositorieListCatalogCtrl);

    /* @ngInject */
    function RepositorieListCatalogCtrl(catalogs) {
        var self = this;

        self.apiBase = BACKEND_URL_BASE.defaultBase;
        self.catalogs = catalogs;

    }
})();
