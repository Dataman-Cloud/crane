(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RegistryImageCtrl', RegistryImageCtrl);

    /* @ngInject */
    function RegistryImageCtrl(image, $stateParams, registryCurd) {
        var self = this;

        self.image = image;
        self.deleteImage = deleteImage;
        
        function deleteImage(ev) {
            registryCurd.deleteImage($stateParams.repository, $stateParams.tag, ev);
        }
    }
})();
