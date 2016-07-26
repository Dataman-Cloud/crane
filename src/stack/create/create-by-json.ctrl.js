(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateByJsonCtrl', StackCreateByJsonCtrl);

    /* @ngInject */
    function StackCreateByJsonCtrl($timeout, $scope, $rootScope, $state, stackBackend, $stateParams) {
        var self = this;

        self.supportReadFile = false;
        self.refreshCodeMirror = false;

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

        self.stack = angular.toJson($rootScope.STACK_SAMPLES.singleService, '\t') || "";

        self.form = {
            Namespace: $stateParams.stack_name || '',
            GroupId: $stateParams.group_id,
            Stack: ""
        };

        self.errorInfo = {
            stack: ''
        };

        self.onFileSelect = onFileSelect;
        self.create = create;
        self.stackChange = stackChange;
        self.example = example;

        activate();

        function activate() {
            self.supportReadFile = !!(window.File && window.FileReader && window.FileList && window.Blob);

            // cload timeout is 10, set long for it;
            var timeoutPromise = $timeout(function () {
                self.refreshCodeMirror = true;
            }, 20, false);

            $scope.$on('$destroy', function () {
                $timeout.cancel(timeoutPromise);
            });

        }

        function create() {
            self.form.Stack = angular.fromJson(self.stack);
            stackBackend.createStack(self.form, $scope.staticForm).then(function (data) {
                $state.go('stack.list');
            });
        }

        function stackChange() {
            //clean error Info
            self.errorInfo.stack = '';

            stackValidate()
        }

        function onFileSelect(files) {
            // files is a FileList of File objects. List some properties.
            var file = files[0];

            var reader = new FileReader();
            reader.onload = (function (theFile) {
                return function (e) {
                    self.stack = e.target.result;
                    stackChange();
                    $scope.$digest();
                };
            })(file);

            reader.readAsText(file);
        }

        function stackValidate() {
            try {
                JSON.parse(self.stack)
            } catch (err) {
                self.errorInfo.stack = 'JSON 格式有误';
            }
        }

        function example() {
            self.stack = angular.toJson($rootScope.STACK_SAMPLES.doubleServices, '\t') || "";
            stackChange();
        }
    }
})();
