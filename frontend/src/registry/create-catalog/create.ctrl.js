(function () {
    'use strict';
    angular.module('app.registry')
        .controller('CreateUpdateCatalog', CreateUpdateCatalog);

    /* @ngInject */
    function CreateUpdateCatalog(stack, $timeout, $scope, Notification, $rootScope, registryCurd, target) {
        var self = this;

        self.target = target;
        self.imageSize = 0;

        if (target === 'create') {
            self.form = {
                Name: '',
                Bundle: angular.toJson(stack.Stack, '\t') || ""
            };
        } else {
            self.form = {
                Name: stack.Name || '',
                Bundle: stack.Bundle || '',
                Description: stack.Description || ''
            };

            self.image = stack.IconData || ''
        }


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
        self.create = create;
        self.update = update;

        activate();

        function activate() {
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
                JSON.parse(self.form.Bundle)
            } catch (err) {
                self.errorInfo.stack = 'JSON 格式有误';
            }
        }

        function imageUpload(files) {
            var file = files[0]; //FileList object
            self.imageSize = file.size;
            if (self.imageSize > $rootScope.IMAGE_MAX_SIZE) {
                Notification.warning('图片过大，请选择小于 1M 的图片');
            }
            var reader = new FileReader();
            reader.onload = (function (theFile) {
                return function (e) {
                    self.image = e.target.result;
                    $scope.$digest();
                };
            })(file);
            reader.readAsDataURL(file)
        }

        function create() {
            registryCurd.createCatalog(self.form, $scope.staticForm)
        }

        function update() {
            //TODO
        }
    }
})();
