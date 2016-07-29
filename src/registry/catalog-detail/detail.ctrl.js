(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CatalogDetailCtrl', CatalogDetailCtrl);

    /* @ngInject */
    function CatalogDetailCtrl(catalog) {
        var self = this;

        self.catalog = catalog;
        console.log(catalog)


        self.questions = angular.fromJson(catalog.Questions).questions;
    }
})();
