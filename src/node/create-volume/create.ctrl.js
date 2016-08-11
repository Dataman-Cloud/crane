(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeCreateVolumeCtrl', NodeCreateVolumeCtrl);

    /* @ngInject */
    function NodeCreateVolumeCtrl(nodeBackend, $scope, $state, $stateParams) {
        var self = this;

        self.form = {
            Name: '',
            Driver: '',
            Labels: {}
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

            nodeBackend.createVolume(self.form, $stateParams.node_id, $scope.createVolume)
                .then(function (data) {
                    $state.go('node.detail.volume', {node_id: $stateParams.node_id}, {reload: true})
                })
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