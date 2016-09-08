(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeCreateNetworkCtrl', NodeCreateNetworkCtrl);

    /* @ngInject */
    function NodeCreateNetworkCtrl($scope, $state, $stateParams, nodeCurd) {
        var self = this;

        self.form = {
            "Name": "",
            "CheckDuplicate": true,
            "Driver": "bridge",
            "EnableIPv6": false,
            "IPAM": {
                "Driver": "default",
                "Config": []
            },
            "Internal": true,
            "Options": {},
            "Labels": {}
        };

        self.labels = [];

        self.create = create;
        self.addLabel = addLabel;
        self.deleteLabel = deleteLabel;
        self.listLabel = listLabel;

        activate();

        function activate() {
            ///
        }

        function create() {
            angular.forEach(self.labels, function (item, index) {
                self.form.Labels[item.key] = item.value
            });

            nodeCurd.createNetwork(self.form, $stateParams.node_id, $scope.createNetwork)
        }

        function addLabel() {
            var label = {
                key: '',
                value: ''
            };

            self.labels.push(label)
        }

        function deleteLabel(index) {
            self.labels.splice(index, 1);
        }

        function listLabel(curIndex) {
            var labels = self.labels.map(function (item, index) {
                if (item.key && curIndex !== index) {
                    return item.key
                }
            });

            return labels
        }
    }
})();