(function () {
    'use strict';
    angular.module('glance.node')
        .factory('nodeBackend', nodeBackend);


    /* @ngInject */
    function nodeBackend(gHttp) {
        return {
            listNodes: listNodes
        };

        function listNodes(params, loading) {
            return gHttp.Resource('node.nodes').get({params: params, "loading": loading});
        }
    }
})();