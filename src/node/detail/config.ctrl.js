(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeConfigCtrl', NodeConfigCtrl);

    /* @ngInject */
    function NodeConfigCtrl(node) {
        var self = this;

        self.node = node
    }
})();
