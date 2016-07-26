(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateNavCtrl', StackCreateNavCtrl);


    /* @ngInject */
    function StackCreateNavCtrl($stateParams, userBackend, $rootScope) {
        var self = this;

        self.stackName = $stateParams.stack_name || '';
        self.groups = [];

        self.loadGroups = loadGroups;

        function loadGroups(){
            userBackend.listGroup($rootScope.accountId)
                .then(function(data){
                    self.groups = data
                })
        }
    }
})();
