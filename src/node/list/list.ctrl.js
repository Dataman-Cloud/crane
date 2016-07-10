(function () {
    'use strict';
    angular.module('glance.node')
        .controller('NodeListCtrl', NodeListCtrl);


    /* @ngInject */
    function NodeListCtrl(nodes) {
        var self = this;

        self.nodes = nodes;

        activate();

        function activate() {
            ///
        }
    }
})();
