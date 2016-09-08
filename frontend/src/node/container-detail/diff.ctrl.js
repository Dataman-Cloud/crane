(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerDiffCtrl', NodeContainerDiffCtrl);

    /* @ngInject */
    function NodeContainerDiffCtrl(diffs) {
        var self = this;
        
        self.diffs = diffs;
    }
})();
