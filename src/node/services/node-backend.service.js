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
            getImage: getImage,
            getImageHistory: getImageHistory,
            listContainers: listContainers,
            getContainer: getContainer,
            removeContainer: removeContainer,
            killContainer: killContainer,
            diffContainer: diffContainer
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

        function getContainer(nodeId, containerId) {
            return gHttp.Resource('node.container', {node_id: nodeId, container_id: containerId}).get();
        }

        function removeContainer(nodeId, containerId) {
            var data = {
                method: 'rm'
            };

            return gHttp.Resource('node.container', {node_id: nodeId, container_id: containerId}).delete({data: data});
        }

        function killContainer(nodeId, containerId) {
            var data = {
                method: 'kill'
            };

            return gHttp.Resource('node.container', {node_id: nodeId, container_id: containerId}).delete({data: data});
        }

        function diffContainer(nodeId, containerId) {
            return gHttp.Resource('node.containerDiff', {node_id: nodeId, container_id: containerId}).get();
        }
    }
})();