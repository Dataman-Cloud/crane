(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceLogCtrl', ServiceLogCtrl);

    /* @ngInject */
    function ServiceLogCtrl(stream, $stateParams, $scope) {
        var self = this;

        self.logs = [];

        activate();

        function activate() {
            listenLog();
        }

        function listenLog() {
            stream = stream.Stream('stack.serviceLog', {
                stack_name: $stateParams.stack_name,
                service_id: $stateParams.service_id
            });
            stream.addHandler('service-logs', function (event) {
                self.logs.push(transformLog(event.data));
                $scope.$apply();
                $('#serviceLog').scrollTop($('#serviceLog')[0].scrollHeight);
            });
            stream.start();

            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }

        function transformLog(log) {
            return log.replace(/(\[DEBUG\]|DEBUG|Debug|\[debug\])/g, "<em class='text-success'>$1</em>")
                .replace(/(\[INFO\]|INFO|Info|\[info\])/g, "<em class='text-info'>$1</em>")
                .replace(/(\[WARN\]|WARM|Warn|\[warn\])/g, "<em class='text-warning'>$1</em>")
                .replace(/(\[ERROR\]|ERROR|Error|\[error\])/g, "<em class='text-danger'>$1</em>")
                .replace(/(\[FATAL\]|FATAL|Fatal|\[fatal\])/g, "<em class='text-danger'>$1</em>")
        }

    }
})();
