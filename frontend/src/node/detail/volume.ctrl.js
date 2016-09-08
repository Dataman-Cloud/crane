(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeVolumeCtrl', NodeVolumeCtrl);

    /* @ngInject */
    function NodeVolumeCtrl(volumes, nodeCurd, $stateParams) {
        var self = this;

        self.volumes = volumes;

        self.deleteVolume = deleteVolume;

        function deleteVolume(volumeName) {
            nodeCurd.deleteVolume($stateParams.node_id, volumeName)
        }
    }
})();
