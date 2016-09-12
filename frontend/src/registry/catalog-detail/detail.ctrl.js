(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CatalogDetailCtrl', CatalogDetailCtrl);

    /* @ngInject */
    function CatalogDetailCtrl(catalog, userBackend, $rootScope, stackCurd, $scope) {
        var self = this;

        self.catalog = catalog;
        self.stack = self.catalog.Bundle || "";
        self.description = self.catalog.Description || "";

        self.stackName = '';
        self.groups = [];

        self.form = {
            Namespace: "",
            Stack: ""
        };

        self.editorOptions = {
            theme: 'midnight',
            lineNumbers: true,
            indentWithTabs: true,
            matchBrackets: true,
            mode: 'Javascript',
            tabSize: 2,
            extraKeys: {
                Tab: function (cm) {
                    var spaces = new Array(cm.getOption('indentUnit') + 1).join(' ');
                    cm.replaceSelection(spaces);
                }
            }
        };

        self.errorInfo = {
            stack: ''
        };

        self.loadGroups = loadGroups;
        self.stackChange = stackChange;
        self.create = create;

        function loadGroups() {
            userBackend.listGroup($rootScope.accountId)
                .then(function (data) {
                    self.groups = data
                })
        }

        function stackChange() {
            //clean error Info
            self.errorInfo.stack = '';

            stackValidate()
        }

        function stackValidate() {
            try {
                JSON.parse(self.stack)
            } catch (err) {
                self.errorInfo.stack = 'JSON 格式有误';
            }
        }

        function create() {
            self.form.Stack = angular.fromJson(self.stack);
            stackCurd.createStack(self.form, $scope.staticForm, self.selectGroupId)
        }

    }
})();
