(function () {
    'use strict';
    angular.module('app.network')
        .factory('networkBackend', networkBackend);


    /* @ngInject */
    function networkBackend(gHttp) {
        return {
            create: create,
            listNetwork: listNetwork,
            deleteNetwork: deleteNetwork,
            connectNetwork: connectNetwork,
            disconnectNetwork: disconnectNetwork,
            getNetwork: getNetwork
        };

        function create(data, form) {
            return gHttp.Resource('network.networks').post(data, {form: form});
        }

        function listNetwork() {
            return gHttp.Resource('network.networks').get();
        }

        function deleteNetwork(id) {
            return gHttp.Resource('network.network', {network_id: id}).delete();
        }

        function connectNetwork(data, id) {
            data.method = 'connect';

            return gHttp.Resource('network.container', {network_id: id}).patch(data);
        }

        function disconnectNetwork(id) {
            var data = {
                method: 'disconnect'
            };

            return gHttp.Resource('network.container', {network_id: id}).patch(data);
        }

        function getNetwork(id) {
            return gHttp.Resource('network.network', {network_id: id}).get();
        }
    }
})();