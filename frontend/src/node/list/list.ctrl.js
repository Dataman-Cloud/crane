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
        self.deleteNode = deleteNode;
        self.updateEndpoint = updateEndpoint;
        self.updateLabels = updateLabels;
        self.addWorkerNode = addWorkerNode;

        activate();

        function activate() {
            ///
        }

        function drainNode(nodeId) {
            nodeCurd.drainNode(nodeId)
        }

        function activeNode(nodeId) {
            nodeCurd.activeNode(nodeId)
        }

        function pauseNode(nodeId) {
            nodeCurd.pauseNode(nodeId)
        }

        function deleteNode(nodeId) {
            nodeCurd.deleteNode(nodeId)
        }

        function updateEndpoint(nodeId, env, labels) {
            var endpoint = "";
            if (labels && labels.hasOwnProperty(NODE_ENDPOINT_LABEL)) {
                endpoint = labels[NODE_ENDPOINT_LABEL]
            }

            nodeCurd.updateEndpoint(nodeId, env, endpoint)
        }

        function updateLabels(nodeId, env, labels) {
            labels = labels || {};
            nodeCurd.updateLabels(nodeId, env, labels)
        }

        function addWorkerNode(env){
            nodeCurd.addWorkerNode(env)
        }
    }
})();
