(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerLogCtrl', NodeContainerLogCtrl);

    /* @ngInject */
    function NodeContainerLogCtrl(stream, $stateParams, $scope) {
        var self = this;
        
        self.logs = []
        activate();
        
        function activate() {
            listenLog();
        }
        
        function listenLog() {
            stream = stream.Stream('node.containerLog', {node_id:$stateParams.node_id, container_id:$stateParams.container_id});
            stream.addHandler('container-logs', function (event) {
                self.logs.push(event.data);
                $scope.$apply();
                $('#log').scrollTop( $('#log')[0].scrollHeight );
            });
            stream.start();
            
            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();
