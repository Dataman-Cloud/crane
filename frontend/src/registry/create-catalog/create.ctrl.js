(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CreateCatalog', CreateCatalog);

    /* @ngInject */
    function CreateCatalog(stack) {
        var self = this;

        self.form = stack;

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

        self.errorInfo = {
            stack: ''
        };

        self.stackChange = stackChange;

        activate();

        function activate() {
            self.stack = angular.toJson(stack.Stack, '\t') || "";
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
    }
})();
