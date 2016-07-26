(function () {
    'use strict';
    angular.module('app.network')
        .controller('NetworkCreateCtrl', NetworkCreateCtrl);

    /* @ngInject */
    function NetworkCreateCtrl(networkCurd, $scope) {
        var self = this;

        self.form = {
            "Name": "",
            "CheckDuplicate": true,
            "Driver": "overlay",
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
                self.form.Options[item.key] = item.value
            });

            networkCurd.create(self.form, $scope.createNetwork)
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

        function listLabel() {
            var labels = self.labels.map(function (item, index) {
                if (item.key) {
                    return item.key
                }
            });

            return labels
        }
    }
})();
