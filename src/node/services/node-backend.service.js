(function () {
    'use strict';
    angular.module('app.node')
        .factory('nodeBackend', nodeBackend);


    /* @ngInject */
    function nodeBackend(gHttp) {
        return {
            listNodes: listNodes,
            getLeaderNode: getLeaderNode,
            createVolume: createVolume,
            listVolumes: listVolumes,
            deleteVolume: deleteVolume
        };

        function listNodes(params, loading) {
            return gHttp.Resource('node.nodes').get({params: params, "loading": loading});
        }

        function getLeaderNode() {
            return gHttp.Resource('node.leader').get();
        }

        function createVolume(data, id, form) {
            return gHttp.Resource('node.volumes', {node_id: id}).post(data, {form: form});
        }

        function listVolumes(id) {
            return gHttp.Resource('node.volumes', {node_id: id}).get();
        }

        function deleteVolume(id, name) {
            return gHttp.Resource('node.volume', {node_id: id, volume_name: name}).delete();
        }
    }
})();