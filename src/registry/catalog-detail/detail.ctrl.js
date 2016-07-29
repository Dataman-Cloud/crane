(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CatalogDetailCtrl', CatalogDetailCtrl);

    /* @ngInject */
    function CatalogDetailCtrl(catalog) {
        var self = this;

        self.catalog = catalog;
    }
})();
