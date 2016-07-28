(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RegistryImageCtrl', RegistryImageCtrl);

    /* @ngInject */
    function RegistryImageCtrl(image) {
        var self = this;

        self.image = image;
    }
})();
