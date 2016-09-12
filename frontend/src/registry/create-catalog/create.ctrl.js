(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CreateCatalog', CreateCatalog);

    /* @ngInject */
    function CreateCatalog(stack, $timeout, $scope, Notification, $rootScope) {
        var self = this;

        self.form = stack;
        self.imageSize = 0;

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
        self.imageUpload = imageUpload;

        activate();

        function activate() {
            self.stack = angular.toJson(stack.Stack, '\t') || "";

            // cload timeout is 10, set long for it;
            var timeoutPromise = $timeout(function () {
                self.refreshCodeMirror = true;
            }, 20, false);

            $scope.$on('$destroy', function () {
                $timeout.cancel(timeoutPromise);
            });
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

        function imageUpload(files) {
            var file = files[0]; //FileList object
            self.imageSize = file.size;
            if(self.imageSize > $rootScope.IMAGE_MAX_SIZE){
                Notification.warning('图片过大，请选择小于 1M 的图片');
            }
            var reader = new FileReader();
            reader.onload = (function (theFile) {
                return function (e) {
                    self.imageData = e.target.result;
                    $scope.$digest();
                };
            })(file);
            reader.readAsDataURL(file);
        }
    }
})();
