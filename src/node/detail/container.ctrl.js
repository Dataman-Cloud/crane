(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerCtrl', NodeContainerCtrl);

    /* @ngInject */
    function NodeContainerCtrl(containers) {
        var self = this;

        self.containers = containers;
    }
})();
