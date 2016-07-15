(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeImageLayerCtrl', NodeImageLayerCtrl);

    /* @ngInject */
    function NodeImageLayerCtrl(layer) {
        var self = this;

        self.layer = layer;
    }
})();
