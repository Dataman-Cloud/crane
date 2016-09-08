(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeDetailCtrl', NodeDetailCtrl);

    /* @ngInject */
    function NodeDetailCtrl(node) {
        var self = this;

        self.node = node
    }
})();