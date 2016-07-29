(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceLogCtrl', ServiceLogCtrl);

    /* @ngInject */
    function ServiceLogCtrl(stream, $stateParams, $scope) {
        var self = this;
        
        self.logs = []
        activate();
        
        function activate() {
            listenLog();
        }
        
        function listenLog() {
            stream = stream.Stream('stack.serviceLog', {stack_name:$stateParams.stack_name, service_id:$stateParams.service_id});
            stream.addHandler('service-logs', function (event) {
                self.logs.push(event.data);
                $scope.$apply();
                $('#log').scrollTop( $('#log')[0].scrollHeight);
            });
            stream.start();
            
            $scope.$on('$destroy', function () {
                stream.stop();
            });
        }
    }
})();
