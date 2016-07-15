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
            deleteVolume: deleteVolume,
            listImages: listImages,
            listContainers: listContainers,
            getImage: getImage,
            getImageHistory: getImageHistory
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

        function listImages(id, params, loading) {
            return gHttp.Resource('node.images', {node_id: id}).get({params: params, "loading": loading});
        }

        function getImage(id, name) {
            return gHttp.Resource('node.image', {node_id: id, image_name: name}).get();
        }

        function getImageHistory(id, name) {
            return gHttp.Resource('node.imageHistory', {node_id: id, image_name: name}).get();
        }
        
        function listContainers(nodeId) {
            return gHttp.Resource('node.containers', {node_id: nodeId}).get();
        }
    }
})();