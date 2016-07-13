(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeBackend', nodeBackend);


    /* @ngInject */
    function nodeBackend(gHttp) {
        return {
            listNodes: listNodes,
            getLeaderNode: getLeaderNode
        };

        function listNodes(params, loading) {
            return gHttp.Resource('node.nodes').get({params: params, "loading": loading});
        }

        function getLeaderNode() {
            return gHttp.Resource('node.leader').get();
        }
    }
})();