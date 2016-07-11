(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateCtrl', StackCreateCtrl);

    /* @ngInject */
    function StackCreateCtrl($timeout, $scope, $rootScope, $state, Notification, stackBackend, target, $stateParams) {
        var self = this;

        self.target = target;
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

        self.json = angular.toJson($rootScope.STACK_DEFAULT.JsonObj, '\t') || "";

        self.form = {
            name: "",
            json: ""
        };

        self.errorInfo = {
            json: ''
        };

        self.onFileSelect = onFileSelect;
        self.create = create;
        self.update = update;
        self.jsonChange = jsonChange;

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
            self.form.json = angular.fromJson(self.json);
            console.log(self.form)
        }

        function update() {
            ///
        }

        function jsonChange() {
            //clean error Info
            self.errorInfo.json = '';

            jsonValidate()
        }

        function onFileSelect(files) {
            // files is a FileList of File objects. List some properties.
            var file = files[0];

            var reader = new FileReader();
            reader.onload = (function (theFile) {
                return function (e) {
                    self.json = e.target.result;
                    jsonChange();
                    $scope.$digest();
                };
            })(file);

            reader.readAsText(file);
        }

        function jsonValidate() {
            try {
                JSON.parse(self.json)
            } catch (err) {
                self.errorInfo.json = 'JSON 格式有误';
            }
        }
    }
})();