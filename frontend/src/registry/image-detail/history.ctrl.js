(function () {
    'use strict';
    angular.module('app.registry')
        .controller('RegistryImageHistoryCtrl', RegistryImageHistoryCtrl);

    /* @ngInject */
    function RegistryImageHistoryCtrl($scope) {
        var self = this;
        
        self.historyInfos = [];
        
        activate();
        
        function activate() {
            angular.forEach($scope.registryImageCtrl.image.history, function (historyInfo) {
                var newInfo = {}
                angular.forEach(historyInfo, function (value, key){
                    newInfo[key] = angular.fromJson(value);
                })
                self.historyInfos.push(newInfo);
            })
        }

    }
})();
