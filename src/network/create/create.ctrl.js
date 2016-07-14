(function () {
    'use strict';
    angular.module('app.network')
        .controller('CreateCtrl', CreateCtrl);

    /* @ngInject */
    function CreateCtrl(networkCurd, $scope) {
        var self = this;

        self.form = {
            "Name": "",
            "CheckDuplicate": false,
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

        self.options = [];

        self.create = create;
        self.addOption = addOption;
        self.deleteOption = deleteOption;

        activate();

        function activate() {
            ///
        }

        function create() {
            angular.forEach(self.options, function (item, index) {
                self.form.Options[item.key] = item.value
            });

            networkCurd.create(self.form, $scope.createNetwork)
        }

        function addOption() {
            var option = {
                key: '',
                value: ''
            };

            self.options.push(option)
        }

        function deleteOption(index) {
            self.options.splice(index, 1);
        }
    }
})();
