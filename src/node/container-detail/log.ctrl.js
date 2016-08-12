(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerLogCtrl', NodeContainerLogCtrl);

    /* @ngInject */
    function NodeContainerLogCtrl(stream, $stateParams, $scope) {
        var self = this;

        self.logs = [];

        activate();

        function activate() {
            listenLog();
        }

        function listenLog() {
            stream = stream.Stream('node.containerLog', {
                node_id: $stateParams.node_id,
                container_id: $stateParams.container_id
            });
            stream.addHandler('container-logs', function (event) {
                self.logs.push(transformLog(event.data));
                $scope.$apply();
                $('#containerLog').scrollTop($('#containerLog')[0].scrollHeight);
            });
            stream.start();

            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();

function transformLog(log) {
    return log.replace(/(\[DEBUG\]|DEBUG|Debug|\[debug\])/g, "<em class='text-success'>$1</em>")
        .replace(/(\[INFO\]|INFO|Info|\[info\])/g, "<em class='text-info'>$1</em>")
        .replace(/(\[WARN\]|WARM|Warn|\[warn\])/g, "<em class='text-warning'>$1</em>")
        .replace(/(\[ERROR\]|ERROR|Error|\[error\])/g, "<em class='text-danger'>$1</em>")
        .replace(/(\[FATAL\]|FATAL|Fatal|\[fatal\])/g, "<em class='text-danger'>$1</em>")
}
