(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateCtrl', StackCreateCtrl);

    /* @ngInject */
    function StackCreateCtrl($timeout, $scope, $state, Notification, stackBackend, target, $stateParams) {
        var self = this;

        var yamlForm = {
            compose: ''
        };
        var patt = new RegExp("^[a-z|_]+$");

        self.target = target;
        self.refreshCodeMirror = false;

        self.editorOptions = {
            mode: 'yaml',
            lineNumbers: true,
            theme: 'midnight',
            matchBrackets: true,
            tabSize: 2,
            'extraKeys': {
                Tab: function (cm) {
                    var spaces = new Array(cm.getOption('indentUnit') + 1).join(' ');
                    cm.replaceSelection(spaces);
                }
            }
        };

        self.form = {
            name: "",
            compose: STACK_DEFAULT.DockerCompose || ""
        };

        self.errorInfo = {
            compose: ''
        };

        self.onFileSelect = onFileSelect;
        self.create = create;
        self.update = update;
        self.onChangeYaml = onChangeYaml;

        activate();

        function activate() {
            generateYaml('compose');

            // cload timeout is 10, set long for it;
            var timeoutPromise = $timeout(function () {
                self.refreshCodeMirror = true;
            }, 20, false);

            $scope.$on('$destroy', function () {
                $timeout.cancel(timeoutPromise);
            });

        }

        function generateYaml(name) {
            var keyList;
            if (!self.form[name]) {
                yamlForm[name] = '';
                self.errorInfo[name] = '';
                return false;
            }

            if (!self.form[name].trim()) {
                yamlForm[name] = '';
                self.errorInfo[name] = '不能为空';
                return false;
            }

            try {
                yamlForm[name] = jsyaml.load(self.form[name]);
            } catch (err) {
                yamlForm[name] = '';
                self.errorInfo[name] = err.message;
                return false;
            }
            if (typeof(yamlForm[name]) === 'string') {
                yamlForm[name] = '';
                self.errorInfo[name] = '必须有多级结构';
                return false;
            }
            keyList = Object.keys(yamlForm[name]);
            for (var i = 0; i < keyList.length; i++) {
                var result = patt.test(keyList[i]);
                if (!result) {
                    yamlForm[name] = '';
                    self.errorInfo[name] = '第一级只能由小写字母和下划线组成';
                    return false;
                }
            }
            self.errorInfo[name] = '';
            return true;
        }

        function create() {
            console.log(self.form)
        }

        function update() {
            ///
        }

        function onChangeYaml(name) {
            generateYaml(name);
        }

        function onFileSelect(files) {
            // files is a FileList of File objects. List some properties.
            var file = files[0];

            var reader = new FileReader();
            reader.onload = (function (theFile) {
                return function (e) {
                    self.form.compose = e.target.result;
                    $scope.$digest();
                };
            })(file);

            reader.readAsText(file);
        }
    }
})();