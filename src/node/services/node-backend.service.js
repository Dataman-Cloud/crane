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

        function createVolume(data, nodeId, form) {
            return gHttp.Resource('node.volumes', {node_id: nodeId}).post(data, {form: form});
        }

        function listVolumes(nodeId) {
            return gHttp.Resource('node.volumes', {node_id: nodeId}).get();
        }

        function deleteVolume(nodeId, name) {
            return gHttp.Resource('node.volume', {node_id: nodeId, volume_name: name}).delete();
        }

        function listImages(nodeId, params, loading) {
            return gHttp.Resource('node.images', {node_id: nodeId}).get({params: params, "loading": loading});
        }

        function getImage(nodeId, imageId) {
            return gHttp.Resource('node.image', {node_id: nodeId, image_id: imageId}).get();
        }

        function getImageHistory(nodeId, imageId) {
            return gHttp.Resource('node.imageHistory', {node_id: nodeId, image_id: imageId}).get();
        }
        
        function listContainers(nodeId) {
            return gHttp.Resource('node.containers', {node_id: nodeId}).get();
        }
    }
})();