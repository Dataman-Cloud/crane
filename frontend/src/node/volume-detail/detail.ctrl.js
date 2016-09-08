/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeVolumeDetailCtrl', NodeVolumeDetailCtrl);

    /* @ngInject */
    function NodeVolumeDetailCtrl(volume) {
        var self = this;

        self.volume = volume;
    }
})();
