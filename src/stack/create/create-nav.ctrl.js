(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateNavCtrl', StackCreateNavCtrl);


    /* @ngInject */
    function StackCreateNavCtrl($stateParams) {
        var self = this;

        self.stackName = $stateParams.stack_name || '';
    }
})();
