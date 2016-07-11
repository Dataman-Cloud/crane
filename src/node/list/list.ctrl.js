(function () {
    'use strict';
    angular.module('app.node')
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
