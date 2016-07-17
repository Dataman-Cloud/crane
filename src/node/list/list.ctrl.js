(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeListCtrl', NodeListCtrl);


    /* @ngInject */
    function NodeListCtrl(nodes, nodeCurd) {
        var self = this;

        self.nodes = nodes;

        self.drainNode = drainNode;
        self.activeNode = activeNode;
        self.pauseNode = pauseNode;

        activate();

        function activate() {
            ///
        }

        function drainNode(nodeId, node) {
            var data = handleNodeFormat(node);
            nodeCurd.drainNode(nodeId, data)
        }

        function activeNode(nodeId, node) {
            var data = handleNodeFormat(node);
            nodeCurd.activeNode(nodeId, data)
        }

        function pauseNode(nodeId, node) {
            var data = handleNodeFormat(node);
            nodeCurd.pauseNode(nodeId, data)
        }

        function handleNodeFormat(node) {
            var data = {
                Spec: node.Spec || '',
                Version: node.Version || ''
            };

            return data;
        }
    }
})();
