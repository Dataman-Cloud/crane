(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeImageCtrl', NodeImageCtrl);

    /* @ngInject */
    function NodeImageCtrl(images, nodeCurd) {
        var self = this;

        self.images = images;

        self.deleteImage = deleteImage;

        function deleteImage(nodeId, imageId){
            nodeCurd.deleteImage(nodeId, imageId)
        }

    }
})();
