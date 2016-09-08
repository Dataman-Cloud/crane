(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackDetailCtrl', StackDetailCtrl);

    /* @ngInject */
    function StackDetailCtrl(stack, stackCurd) {
        var self = this;
        
        self.stack = stack;
        self.deleteStack = deleteStack;
        
        function deleteStack(ev) {
            stackCurd.deleteStack(ev, stack.Namespace);
        }
    }
})();
