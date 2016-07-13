(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackDetailCtrl', StackDetailCtrl);

    /* @ngInject */
    function StackDetailCtrl(stack) {
        var self = this;
        
        self.stack = stack;
    }
})();